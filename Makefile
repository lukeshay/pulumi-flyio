PROJECT_NAME := Pulumi Fly.io Resource Provider

PACK             := flyio
PACKDIR          := sdk
PACKDIR_SUPERFLY := ${PACKDIR}/superfly
PROJECT          := github.com/lukeshay/pulumi-flyio
NODE_MODULE_NAME := pulumi-flyio
NUGET_PKG_NAME   := Pulumi.Flyio

PROVIDER        := pulumi-resource-${PACK}
VERSION         ?= $(shell pulumictl get version)
PROVIDER_PATH   := provider
VERSION_PATH    := ${PROVIDER_PATH}/pkg/provider.Version

GOPATH			:= $(shell go env GOPATH)

WORKING_DIR     := $(shell pwd)
EXAMPLES_DIR    := ${WORKING_DIR}/examples/yaml
TESTPARALLELISM := 4

OS    := $(shell uname)
SHELL := /bin/bash

ensure::
	cd provider && go mod tidy
	cd tests && go mod tidy

gen::
	curl 'https://docs.machines.dev/spec/openapi3.json' -o tmp/fly-openapi3.tmp.json --create-dirs
	jq '.paths."/apps/{app_name}/secrets".get += {"parameters": [{"name":"app_name","in":"path","description":"Fly App Name","required":true,"schema":{"type":"string"}}]}' tmp/fly-openapi3.tmp.json > tmp/fly-openapi3.json
	mkdir -p provider/pkg/flyio
	rm -f $(WORKING_DIR)/bin/$(PROVIDER)
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest --config oapi.yaml tmp/fly-openapi3.json
	sed -i 's/float32/float64/g' provider/pkg/flyio/flyio.gen.go
	cd provider && go generate -ldflags "-X $(PROJECT)/$(VERSION_PATH)=$(VERSION)" $(PROJECT)/$(PROVIDER_PATH)/cmd/$(PROVIDER)

provider:: gen
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

provider_debug::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -gcflags="all=-N -l" -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

test_provider::
	cd tests && go test -short -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM} ./...

dotnet_sdk:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
dotnet_sdk::
	rm -rf sdk/dotnet
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language dotnet
	cd ${PACKDIR}/dotnet/&& \
		dotnet build /p:Version=${DOTNET_VERSION} && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' ./pulumi-plugin.json && \
		rm ./pulumi-plugin.json.bak

java_sdk::
	rm -rf sdk/java
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language java
	# cd ${PACKDIR}/java/&& \
	# 	sed -i.bak 's/$${VERSION}/$(VERSION)/g' ./pulumi-plugin.json && \
	# 	rm ./pulumi-plugin.json.bak

go_sdk:: $(WORKING_DIR)/bin/$(PROVIDER)
	rm -rf sdk/go
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language go
	cd ${PACKDIR}/go/ && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' ./flyio/pulumi-plugin.json && \
		rm ./flyio/pulumi-plugin.json.bak

nodejs_sdk:: VERSION := $(shell pulumictl get version --language javascript)
nodejs_sdk::
	rm -rf sdk/nodejs
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language nodejs
	cd ${PACKDIR}/nodejs/ && \
		bun install && \
		bun run tsc && \
		cp ../../README.md ../../LICENSE package.json bun.lockb bin/ && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' bin/package.json && \
		rm ./bin/package.json.bak
nodejs_superfly_sdk:: VERSION := $(shell pulumictl get version --language javascript)
nodejs_superfly_sdk::
	rm -rf sdk/nodejs-superfly
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language nodejs --out ${PACKDIR_SUPERFLY}
	cd ${PACKDIR_SUPERFLY}/nodejs/ && \
		bun install && \
		bun run tsc && \
		cp ../../../README.md ../../../LICENSE package.json bun.lockb bin/ && \
		sed -i.bak -e 's/$${VERSION}/$(VERSION)/g' -e 's/pulumi-flyio/pulumi-superfly/g' bin/package.json && \
		rm ./bin/package.json.bak

python_sdk:: PYPI_VERSION := $(shell pulumictl get version --language python)
python_sdk::
	rm -rf sdk/python
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language python
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' ./bin/pulumi_flyio/pulumi-plugin.json && \
		rm ./bin/setup.py.bak ./bin/pulumi_flyio/pulumi-plugin.json.bak && \
		cd ./bin && python3 setup.py build sdist

python_superfly_sdk:: PYPI_VERSION := $(shell pulumictl get version --language python)
python_superfly_sdk::
	rm -rf sdk/python
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language python --out ${PACKDIR_SUPERFLY}
	cp README.md ${PACKDIR_SUPERFLY}/python/
	cd ${PACKDIR_SUPERFLY}/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' -e "s/name='pulumi_flyio'/name='pulumi_superfly'/g" ./bin/setup.py && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' ./bin/pulumi_flyio/pulumi-plugin.json && \
		rm ./bin/setup.py.bak ./bin/pulumi_flyio/pulumi-plugin.json.bak && \
		cd ./bin && python3 setup.py build sdist

gen_examples: gen_go_example \
		gen_nodejs_example \
		gen_python_example \
		gen_java_example \
		gen_dotnet_example

gen_%_example:
	rm -rf ${WORKING_DIR}/examples/$*
	pulumi convert \
		--cwd ${WORKING_DIR}/examples/yaml \
		--logtostderr \
		--generate-only \
		--non-interactive \
		--language $* \
		--out ${WORKING_DIR}/examples/$*

define pulumi_login
    export PULUMI_CONFIG_PASSPHRASE=asdfqwerty1234; \
    pulumi login --local;
endef

up::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack init dev || true && \
	pulumi stack select dev && \
	pulumi config set name dev && \
	pulumi up -y

down::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack select dev && \
	pulumi destroy -y && \
	pulumi stack rm dev -y

devcontainer::
	git submodule update --init --recursive .devcontainer
	git submodule update --remote --merge .devcontainer
	cp -f .devcontainer/devcontainer.json .devcontainer.json

.PHONY: build

build:: provider dotnet_sdk go_sdk nodejs_sdk python_sdk java_sdk python_superfly_sdk nodejs_superfly_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

lint::
	for DIR in "provider" "sdk" "tests" ; do \
		pushd $$DIR && golangci-lint run -c ../.golangci.yml --timeout 10m && popd ; \
	done

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

GO_TEST 	 := go test -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM}

test_all:: test_provider
	cd tests/sdk/nodejs && $(GO_TEST) ./...
	cd tests/sdk/python && $(GO_TEST) ./...
	cd tests/sdk/dotnet && $(GO_TEST) ./...
	cd tests/sdk/go && $(GO_TEST) ./...

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/$(NUGET_PKG_NAME).*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

prepare::
	git tag --delete ${VERSION} || true
	git tag ${VERSION}
prepare:: build gen_examples
prepare::
	git tag --delete ${VERSION}

push::
	git tag ${VERSION}
	git push
	git push origin ${VERSION}
	
