.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/addAuthor addAuthor/addAuthor.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/updateAuthor updateAuthor/updateAuthor.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/deleteAuthor deleteAuthor/deleteAuthor.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/getAuthor getAuthor/getAuthor.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/getAuthorName getAuthorName/getAuthorName.go
clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
