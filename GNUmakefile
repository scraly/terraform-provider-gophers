TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=terraform.local
# HOSTNAME=registry.terraform.io
NAMESPACE=local
# NAMESPACE=hashicorp
NAME=gophers
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
#OS_ARCH=darwin_amd64
OS_ARCH=darwin_arm64

default: install

build:
	go build -o ${BINARY}

release:
	goreleaser release --clean --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4  
	
# default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
