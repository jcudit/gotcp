PROJECT=`basename $(PWD)`
GIT_USER=`git config --get user.name`
GIT_USER=jcudit
GO_PROJECTS="/go/src/github.com/$(GIT_USER)"

build-container:
	docker build --rm -t $(PROJECT) .

test: build-container
	docker run -it --rm \
		--name $(PROJECT) \
		--privileged \
		--device /dev/net/tun:/dev/net/tun \
		-v $(PWD):$(GO_PROJECTS)/$(PROJECT) \
		-w $(GO_PROJECTS)/$(PROJECT) $(PROJECT) \
		go test -v ../$(PROJECT)...

.PHONY: build test
