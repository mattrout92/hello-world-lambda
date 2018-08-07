build:

	env GOOS=linux go build -ldflags="-s -w" -o bin/hello-world-lambda main.go