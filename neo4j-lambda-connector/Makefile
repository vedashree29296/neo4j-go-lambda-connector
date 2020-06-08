.PHONY: clean build deploy

build:
	/bin/bash create_layer.sh
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/connect example/connect.go

clean:
	rm -rf ../layer
	rm -rf ./.serverless
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy