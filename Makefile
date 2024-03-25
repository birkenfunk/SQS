APP_NAME=weather-proxy-api
BACKEND_DIR=src/backend/

linux: switchToBackend
	cd ${BACKEND_DIR} && go build -o build/linux/$(APP_NAME)-linux main.go

windows: switchToBackend
	cd ${BACKEND_DIR} && GOOS=windows GOARCH=amd64 go build -o build/windows/$(APP_NAME)-windows.exe main.go

mac: switchToBackend
	cd ${BACKEND_DIR} && GOOS=darwin GOARCH=amd64 go build -o build/mac/$(APP_NAME)-mac main.go

docker_image: linux
	cd ${BACKEND_DIR} && docker build -f ./docker/Dockerfile -t codeberg.org/birkenfunk/sqs-external-api .

podman_image: linux
	cd ${BACKEND_DIR} && podman build -f ./docker/Dockerfile -t codeberg.org/birkenfunk/sqs-external-api .

generate_all: linux windows mac

test: switchToBackend
	cd ${BACKEND_DIR} && go test ./...

clean: switchToBackend
	cd ${BACKEND_DIR} && rm -rf build

.PHONY: generate_linux generate_windows generate_mac generate_all clean docker_image podman_image test switchToBackend
