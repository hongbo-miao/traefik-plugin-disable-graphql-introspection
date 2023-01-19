export GO111MODULE=on

default: test

test:
		go test -v -cover ./...
