//go:build tools
// +build tools

package docs_codegen

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/joho/godotenv/cmd/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/pressly/goose/v3/cmd/goose"
)
