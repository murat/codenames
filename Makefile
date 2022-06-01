build: clean
	GOARCH=amd64 GOOS=linux go build -o ./bin/server cmd/server/main.go

clean:
	go clean
	-rm ./bin/codenames

test:
	go test ./...

coverage:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

lint:
	golangci-lint run ./... -c ./.golangci.yml

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
	   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	   protos/codenames.proto