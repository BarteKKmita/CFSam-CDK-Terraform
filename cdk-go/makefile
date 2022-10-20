.PHONY: build
.PHONY: zip
.PHONY: deploy


build:
	cd cmd; GOOS=linux GOARCH=amd64 go build -tags cloud -ldflags="-s -w" -o ./dist/main handler/main.go

zip: build
	zip cmd/dist/main.zip cmd/dist/main

deploy: zip
	cd cdk/; cdk bootstrap; cdk deploy --require-approval never
