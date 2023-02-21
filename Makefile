-include .common.mk

codegen: protobuf generate

generate: generate-go

lint: lint-go

protobuf: protobuf-go protobuf-python

tool-chain: tool-chain-go	

vendor: vendor-go

## go specific directives
generate-go:
	go generate ./...

lint-go:
	echo $(GO_PROJECT_PACKAGES) | xargs go fmt
	golint $(GO_PROJECT_LINT_PACKAGES)
	staticcheck $(GO_PROJECT_STATICCHECK_PACKAGES)
	errcheck -ignoretests -ignoregenerated -asserts -exclude .errcheck_exclude $(GO_PROJECT_ERRCHECK_PACKAGES)

tool-chain-go:
	go install \
		github.com/kisielk/errcheck \
		github.com/maxbrunsfeld/counterfeiter/v6 \
		github.com/rleszilm/tag-version \
		github.com/srikrsna/protoc-gen-gotag \
		github.com/spf13/cobra-cli \
		golang.org/x/lint/golint \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		honnef.co/go/tools/cmd/staticcheck

vendor-go:
	go mod vendor

#### protobuf
protobuf-go: vendor-go protobuf-go-compile protobuf-go-tag

protobuf-go-compile:
	$(foreach SRC, $(PB_GO_SRC), $(shell protoc $(PB_OPTIONS) $(PB_GO_INCLUDE) $(PB_GO_COMPILE) `ls $(SRC)/*.proto`))

protobuf-go-tag:
	$(foreach SRC, $(PB_GO_TAG_SRC), $(shell protoc $(PB_OPTIONS) -I . $(PB_GO_TAG_INCLUDE) --gotag_opt=paths=source_relative,output_path=$(SRC) --gotag_out=:. `ls $(SRC)/*.proto`))


## python specific directives
#### protobuf
protobuf-python: protobuf-python-compile

protobuf-python-compile:
	$(foreach SRC, $(PB_PY_SRC), $(shell protoc $(PB_OPTIONS) $(PB_PY_INCLUDE) $(PB_PY_COMPILE) `ls $(SRC)/*.proto`))

