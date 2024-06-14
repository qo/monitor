#!/bin/sh
source ./export_env.sh
go test -v ./internal/db
