PROJECT=`basename $(PWD)`
# FIXME
# GIT_USER=`git config --get user.name`
GIT_USER=jcudit
GO_PROJECTS="/go/src/github.com/$(GIT_USER)"

all: build test

build:
	go build -v
	docker build --rm -t $(PROJECT) .
	docker-compose build

test: build build-tunnel
	docker run --rm \
		--name $(PROJECT) \
		--volumes-from gotcp-tun \
		-v $(PWD):$(GO_PROJECTS)/$(PROJECT) \
		-w $(GO_PROJECTS)/$(PROJECT) $(PROJECT) go test -v ../$(PROJECT)...
	docker-compose down

build-tunnel:
	docker-compose up -d

clean:
	docker-compose down
	docker stop $(PROJECT) && docker rm -v $(PROJECT)

watch:
	fswatch -r -o *.go  | xargs -n1 -I{} -L1 make
