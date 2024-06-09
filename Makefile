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

test_with_coverage:
	mkdir build -p
	go test ./... -coverprofile=build/profile.cov ./...

coverage: test_with_coverage
	go tool cover -o build/coverage.html -html build/profile.cov
	go tool cover -o build/coverage.md -func build/profile.cov
	sed -i 's/^\(total\)/# \1/g' build/coverage.md
	sed -i '/^#/!s/\(.\+\)/- \1/g' build/coverage.md
	tail -n 1 build/coverage.md > build/coverage.md.tmp
	head -n -1 build/coverage.md >> build/coverage.md.tmp
	mv build/coverage.md.tmp build/coverage.md

coverage_only:
	go tool cover -o build/coverage.html -html build/profile.cov
	go tool cover -o build/coverage.md -func build/profile.cov
	sed -i 's/^\(total\)/# \1/g' build/coverage.md
	sed -i '/^#/!s/\(.\+\)/- \1/g' build/coverage.md
	tail -n 1 build/coverage.md > build/coverage.md.tmp
	head -n -1 build/coverage.md >> build/coverage.md.tmp
	mv build/coverage.md.tmp build/coverage.md

api_doc:
	go run generate_doc.go

.PHONY: generate_linux generate_windows generate_mac generate_all clean docker_image podman_image test
