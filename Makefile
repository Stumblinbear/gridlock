# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY := gridlockd
VERSION ?= 0.0.0
PLATFORMS := windows linux darwin
os = $(word 1, $@)

PLUGINS := $(subst plugins/,,$(shell find plugins/* -type d))

all: test_all build_all


.PHONY: get
get:
	cd src && $(GOGET) -u
	for p in $(PLUGINS); do \
		cd plugins/$$p ; \
		$(GOGET) -u ; \
		cd ../../ ; \
	done


.PHONY: build_src
build_src:
	cd src/cmd/gridlockd && $(GOBUILD) -o ../../../build/$(BINARY) -v -ldflags "-X main.version=$(VERSION)"

# Build things in the plugins folder
.PHONY: build_plugins
build_plugins:
	for p in $(PLUGINS); do \
		cd plugins/$$p ; \
		$(GOBUILD) -buildmode=plugin -o ../../build/plugins/$$p.so -v; \
		cd ../../ ; \
	done

.PHONY: build
build: build_src build_plugins


test_src:
	cd src && $(GOTEST) -v ./...

# Test things in the plugins folder
test_plugins:
	for p in $(PLUGINS); do \
		cd plugins/$$p ; \
		$(GOTEST) -v ./... ; \
		cd ../../ ; \
	done

test: test_src test_plugins


# Go clean src folder
.PHONY: clean_src
clean_src:
	cd src && $(GOCLEAN)

# Clear the build/ folder
.PHONY: clean_build
clean_build:
	rm -rf build/*

# Clear the releases/ folder
.PHONY: clean_release
clean_release:
	rm -rf release/*

# Go clean the plugin
.PHONY: clean_plugins
clean_plugins:
	for p in $(PLUGINS); do \
		cd plugins/$$p ; \
		$(GOCLEAN) ; \
		cd ../../ ; \
	done

# Clear everything
.PHONY: clean
clean: clean_build clean_release clean_src clean_plugins


run: clean build
	./build/$(BINARY)


.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	cd src && GOOS=$(os) GOARCH=amd64 go build -o ../release/$(BINARY)-v$(VERSION)-$(os)-amd64 -ldflags "-X main.version=$(VERSION)"

.PHONY: release
release: windows linux darwin