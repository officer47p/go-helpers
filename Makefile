build: 
	@go build 

test:
	@go test -cover -v -count=1 -p=1 ./...