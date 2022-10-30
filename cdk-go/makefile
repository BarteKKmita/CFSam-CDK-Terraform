.PHONY: build
.PHONY: zip
.PHONY: deploy

#linux
build:
	cd lambda/cmd; pwd; GOOS=linux GOARCH=amd64 go build -tags cloud -ldflags="-s -w" -o ./dist/main handler/main.go

zip: build
	zip lambda/cmd/dist/main.zip lambda/cmd/dist/main

deploy: zip
	cd cdk; cdk bootstrap; cdk deploy --require-approval never
