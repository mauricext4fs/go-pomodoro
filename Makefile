## Inspired by Trevor Sawler https://github.com/tsawler

BINARY_NAME="Go Pomodoro.app"
APP_NAME="Go Pomodoro"
APP_ID="ch.mauricext4fs.gopomodoro"
VERSION=8.0.0
BUILD_NO=8

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -appID ${APP_ID}
	@## Removing the following line will crash the app when sound is enabled
	cp notification.wav Go\ Pomodoro.app/Contents/Resources/

build_win:
	rm -rf Go\ Pomodoro.exe
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -appID ${APP_ID}

package_win:
	rm -rf package_w11x86
	mkdir -p package_w11x86
	cp notification.wav package_w11x86/
	cp Go\ Pomodoro.exe package_w11x86/
	zip -r GoPomodoro.zip package_w11x86/*

bundle:
	@echo "Bundling ressource into bundled.go"
	go generate
	@#fyne bundle -o bundled.go icon.png
	@#fyne bundle -o bundled.go -append icon_systray.png
	@#fyne bundle -o bundled.go -append notification.wav

## run: builds and runs the application
run:
	env DB_PATH="./sql.db" go run -v .

release:
	@echo "Create package for release"
	fyne release -appID ${APP_ID} -os darwin -profile ${PROFILE} -icon icon.png -category productivity

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@rm -rf package/*
	@echo "Cleaned!"


## test: runs all tests
test:
	go test -v ./...
