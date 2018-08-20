CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-X github.com/kiririmode/exif-namer.revision=$(CURRENT_REVISION)"
ifdef update
	u=-u
endif

deps:
	go get ${u} github.com/golang/dep/cmd/dep
	dep ensure

devel-deps:
	go get ${u} github.com/golang/lint/golint \
                github.com/haya14busa/goverage \
                github.com/Songmu/goxz/cmd/goxz \
                github.com/motemen/gobump/cmd/gobump

test: deps
	go test ./...

lint: devel-deps
	go vet ./...
	go list ./... | xargs golint -set_exit_status

build: deps
	go build -ldflags=${BUILD_LDFLAGS}

cover: devel-deps
	goverage -v -race -covermode=atomic ./...

crossbuild: devel-deps
	goxz -pv=v$(shell gobump show -r) -build-ldflags=$(BUILD_LDFLAGS) -d=./dist/v$(shell gobump show -r) .

release:
	scripts/releng

.PHONY: deps devel-deps test lint build cover crossbuild
