GO_PROJECT_PACKAGES := `go list ./...`
GO_PROJECT_LINT_PACKAGES := $(GO_PROJECT_PACKAGES)
GO_PROJECT_ERRCHECK_PACKAGES := `go list ./... | grep -v github.com/theaufish-git/dndml/tools`
GO_PROJECT_STATICCHECK_PACKAGES := $(GO_PROJECT_ERRCHECK_PACKAGES)

PB_OPTIONS := --experimental_allow_proto3_optional
PB_GO_INCLUDE := -I pkg \
	-I vendor
PB_GO_COMPILE := --go_opt=paths=source_relative \
	--go_out=pkg \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=pkg
PB_GO_SRC := pkg/dndml pkg/dndml/enums

PB_GO_TAG_INCLUDE := $(PB_GO_INCLUDE) \
	-I .
PB_GO_TAG_SRC := 

PB_PY_INCLUDE := -I pkg \
	-I vendor
PB_PY_COMPILE := --python_out=pkg
PB_PY_SRC := pkg/dndml pkg/dndml/enums
