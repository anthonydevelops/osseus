# Build all 
build: agent

# Clean all
clean: agent

# Build agent
agent:
	@echo "=> building agent"
	cd cmd/agent && go build -i -v
	cd plugins/generator && go build -i -v
	cd plugins/restapi && go build -i -v

# Clean agent
clean-agent:
	@echo "=> cleaning agent"
	cd cmd/agent && make clean
	cd plugins/generator && make clean
	cd plugins/restapi && make clean

# Code generation
generate: generate-proto

# Get generator tools
get-proto-generators:
	go install ./vendor/github.com/gogo/protobuf/protoc-gen-gogo

# Generate proto models
generate-proto: get-proto-generators
	@echo "=> generating proto"
	go generate plugins/generator
	go generate plugins/restapi

# Get client dependency manager tool
get-npm: 
	@echo "=> installing nodejs & npm"
	sudo apt install -y nodejs
	sudo apt install -y npm

# Install client dependencies
client-install: get-npm
	@echo "=> installing client dependencies"
	cd client/ && npm install

# Get dependency manager tool
get-dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep version

# Install the project's dependencies
dep-install: get-dep
	dep ensure

# Update the locked versions of all dependencies
dep-update: get-dep
	dep ensure -update

# Check state of dependencies
dep-check: get-dep
	@echo "=> checking dependencies"
	dep check

.PHONY: build clean \
		get-dep dep-install dep-update \
		agent clean-agent \