package tools

import (
	// Included so we can vendor the verions of the tools used.
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/kisielk/errcheck"
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
	_ "github.com/rleszilm/tag-version"
	_ "github.com/spf13/cobra-cli"
	_ "github.com/srikrsna/protoc-gen-gotag"
	_ "golang.org/x/lint/golint"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
