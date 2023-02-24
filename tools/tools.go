package tools

// This file may incorporate tools that may be *both* used as CLI and as lib
// Keep in mind that these global tools change the go.mod/go.sum dependency tree
// Other tooling may be installed as *static binary* directly within the Dockerfile

import (
	_ "github.com/swaggo/swag/cmd/swag"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
