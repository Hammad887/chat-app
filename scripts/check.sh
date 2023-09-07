#!/bin/bash


set -o errexit
set -o pipefail

golangci-lint run --color always --print-resources-usage -c .golangci.yml -v