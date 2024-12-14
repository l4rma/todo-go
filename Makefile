BINARY_NAME=bootstrap
APP=./cmd/task-api/main.go

hello:
	echo "Hello world!"

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

tidy:
	go mod tidy

build: tidy
	GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o ${BINARY_NAME} ${APP}
	@#zip myLambda.zip ${BINARY_NAME}
	@#chmod 755 myLambda.zip

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm bootstrap
	@#rm myLambda.zip
