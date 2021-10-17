export tag=v1.0
root:
	export ROOT=github.com/myangvgo/cn-practice

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t myangvgo/httpserver:${tag} .

push: release
	echo "pushing myangvgo/httpserver"
	docker push myangvgo/httpserver:v1.0
