.PHONY: build packr vendor

build: vendor packr
	@echo "Building relax..."
	@go build -o relax ./cmd
	@echo "Done. Enjoy!"

packrBin := $(shell command -v packr 2> /dev/null)
packr:
ifndef packrBin
	@echo "Installing packr"
	@go get -u github.com/gobuffalo/packr/...
endif
	@echo "Bundling tracks..."
	@packr

DEP := $(shell command -v dep 2> /dev/null)
vendor:
ifndef DEP
	@echo "Installing dep..."
	@go get -u github.com/golang/dep/cmd/dep
endif
	@echo "Installing/updating depenencies..."
	@dep ensure --vendor-only
