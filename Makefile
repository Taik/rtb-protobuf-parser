NAME=go-adx-parser
OS=linux darwin
VERSION=0.0.2

.PHONY: build install release

deps:
	go get github.com/progrium/gh-release/...
	go get github.com/mitchellh/gox/...
	go get github.com/tools/godep
	go get || true

build:
	mkdir -p build/
	gox -os="$(OS)" -arch="amd64" -output="build/{{.OS}}/$(NAME)"

install: build
	install build/$(shell uname -s)/$(NAME) /usr/local/bin

release:
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_x86_64.tgz -C build/linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_x86_64.tgz -C build/darwin $(NAME)
	gh-release checksums sha256
	gh-release create taik/$(NAME) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD) v$(VERSION); true
