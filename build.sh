#!/bin/bash

# builds a fyne macos app, signs it, creates a .dmg, signs that, and then requests notarization.
# ... stapling the notarization should wait until the request is approved which can take a minute ...
#
# Pre-reqs:
#   Install the full xcode package
#   Install the fyne command line tool
#   Create a certificate request, send to apple developer site, and import the resulting certificate
#   Create an application password in appleid.apple.com tied to the app's name in the keychain. Put it in the system keychain.

# The name of the application
APP_NAME="Go Pomodoro"
# The bundle / app ID for the app. Make sure you use a domain that you own!
APP_ID="ch.mauricext4fs.gopomodoro"
APP_KEY_ID=""
# This is the CN from the code signing cert. You can get this information from the keychain.
CERT=""
# This is your login for the apple developer site
EMAIL=""
# Apple Team ID
TEAM_ID=""
# This is the name for an application password stored in the keychain
PASSWORD=""
# A working dir where temp files are created, and final disk image will be placed.
PKG_DIR="package"
# The entrypoint to compile:
MAIN="."
# Set to "TRUE" if the app version should use the git tag:
TAG="FALSE"
### End of settings ###

# Load previous variables from local .env file
. .env

# should permit a package built on MacOS versions >= 10.15 to run on 10.14 and later, maybe.
export CGO_CFLAGS="-mmacosx-version-min=10.14"
export CGO_LDFLAGS="-mmacosx-version-min=10.14"

set -x
# don't overwrite something that already exists when we compile our bin.
BIN=$(dirname "${MAIN}" | awk -F/ '{print $NF}')
MAIN_DIR=$(dirname "${MAIN}")
[ -f "${MAIN_DIR}/${BIN}" ] && {
	echo "refusing to continue, ${MAIN_DIR}/${BIN} already exists"
	kill 0
}

# Cleanup old packages if needed
[ -d "${PKG_DIR}" ] && {
	mkdir -p "${PKG_DIR}/old"
	mv "${PKG_DIR}/"*.dmg "${PKG_DIR}/old/"
}
[ -d "${PKG_DIR}/${APP_NAME}" ] && rm -fr "${PKG_DIR}/${APP_NAME}"
mkdir -p "${PKG_DIR}/${APP_NAME}"

go build -ldflags "-s -w" -o "${MAIN_DIR}/${BIN}" "${MAIN}" || exit

fyne package -sourceDir "${MAIN_DIR}" -name "${APP_NAME}" -os darwin -appID "${APP_ID}" && mv "${APP_NAME}.app" "${PKG_DIR}/${APP_NAME}/" || exit
rm -f "${MAIN_DIR}/${BIN}"

[ "${TAG}" == "TRUE" ] && sed -i'' -e 's/.string.1\.0.\/string./\<string>'$(git describe --tags --always --long)'\<\/string>/g' "${PKG_DIR}/${APP_NAME}/${APP_NAME}.app/Contents/Info.plist"

pushd "${PKG_DIR}"
pushd "${APP_NAME}"

# remove any existing extended attributes from new app:
xattr -rc "${APP_NAME}.app"

# Add wav manually to the package as it will crash otherwise
cp ../../notification.wav Go\ Pomodoro.app/Contents/Resources

# Add a link for /Applications in the Directory, more covenient when becomes a .dmg:
ln -s /Applications Applications

# sign the app:
codesign --force --options runtime --deep --sign "${CERT}" -i "${APP_ID}" "${APP_NAME}.app" || exit

popd
# Create a disk image
hdiutil create -srcfolder "${APP_NAME}" "${APP_NAME}.dmg" || exit
# sign the .dmg
codesign --sign "${CERT}" -i "${APP_ID}" "${APP_NAME}.dmg" || exit

# Request anotization from apple
xcrun notarytool submit --team-id "${TEAM_ID}" --key-id "${APP_KEY_ID}" --apple-id "${EMAIL}" --password "${PASSWORD}" --wait "${APP_NAME}".dmg
rm -fr "${APP_NAME}/"

set +x

# wait until approved and then staple the dmg
xcrun stapler staple -v "${APP_NAME}.dmg"
