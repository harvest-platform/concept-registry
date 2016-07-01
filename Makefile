PROG_NAME := "harvest-domain-service"
GIT_VERSION := $(shell git log -1 --pretty=format:"%h (%ci)" .)

setup: glide tls compiledaemon

micro-cmd:
	@go get -u github.com/micro/micro

glide:
	@if command -v glide &> /dev/null; then \
		echo >&2 'Installing library dependences'; \
		glide install; \
	else \
		echo >&2 'Glide required: https://glide.sh'; \
		exit 1; \
	fi

gen:
	@go generate ./types.go

tls:
	@if [ ! -a cert.pem ]; then \
		echo >&2 'Creating self-signed TLS certs.'; \
		go run $(GOROOT)/src/crypto/tls/generate_cert.go --host localhost; \
	fi

compiledaemon:
	@if command -v CompileDaemon &> /dev/null; then \
		echo >&2 'Getting CompileDaemon for auto-reload.'; \
		go get github.com/githubnemo/CompileDaemon; \
	fi

watch:
	CompileDaemon \
		-build="make build" \
		-command="$(PROG_NAME)" \
		-graceful-kill=true \
		-exclude-dir=.git \
		-exclude-dir=vendor \
		-exclude-dir=client \
		-color=true

build:
	go build -ldflags "-X \"main.buildVersion=$(GIT_VERSION)\"" \
		-o $(GOPATH)/bin/$(PROG_NAME) .

dist-build:
	mkdir -p dist

	# Enable CGO for linking to the SQLite package.
	gox -output="./dist/{{.OS}}-{{.Arch}}/$(PROG_NAME)" \
		-cgo -ldflags "-X \"main.buildVersion=$(GIT_VERSION)\"" \
		-os "windows linux darwin" \
		-arch "amd64" . > /dev/null

dist-zip:
	cd dist && zip $(PROG_NAME)-darwin-amd64.zip darwin-amd64/*
	cd dist && zip $(PROG_NAME)-linux-amd64.zip linux-amd64/*
	cd dist && zip $(PROG_NAME)-windows-amd64.zip windows-amd64/*

dist: dist-build dist-zip

install:
	glide install

test:
	go test -cover $(glide novendor)

doc:
	godoc -http=:6060

.PHONY: test build gen dist-build dist
