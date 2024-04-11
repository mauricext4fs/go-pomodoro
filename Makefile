## Source from Trevor Sawler https://github.com/tsawler

BINARY_NAME="Go Pomodoro.app"
APP_NAME="Go Pomodoro"
VERSION=3.0.0
BUILD_NO=3

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -appID ch.mauricext4fs.gopomodoro -release

## run: builds and runs the application
run:
	env DB_PATH="./sql.db" go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...
