#!/bin/bash

# Set the color variable
green='\033[0;33m'
magenta='\033[0;35m'


# Clear the color after that
clear='\033[0m'

go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/kisielk/errcheck@latest

pkgs=$(go list ./...)

printf "${magenta}Running go vet${clear}\n"
go vet $pkgs

printf "${magenta}Running staticcheck${clear}\n"
staticcheck $pkgs

printf "${magenta}Running errcheck${clear}\n"
errcheck $pkgs

printf "${magenta}Running go golangci-lint${clear}\n"
golangci-lint run