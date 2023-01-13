#!/bin/bash
clc -sS -e accelhint_test.go
go mod tidy
go fmt .
staticcheck .
go vet .
golangci-lint run
git st
