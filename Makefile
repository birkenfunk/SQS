APP_NAME=weather-proxy-api

linux:
	go build -o build/linux/$(APP_NAME)-linux main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o build/windows/$(APP_NAME)-windows.exe main.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o build/mac/$(APP_NAME)-mac main.go

docker_image: linux
	docker build -f ./docker/Dockerfile -t codeberg.org/birkenfunk/sqs-external-api .

podman_image: linux
	podman build -f ./docker/Dockerfile -t codeberg.org/birkenfunk/sqs-external-api .

generate_all: linux windows mac

test:
	go test ./...

clean:
	rm -rf build

.PHONY: generate_linux generate_windows generate_mac generate_all clean docker_image podman_image test
