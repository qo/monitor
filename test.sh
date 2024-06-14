#!/bin/sh

# Экспортиовать переменные среды
source ./export_env.sh

# Запустить тесты
go test -v ./internal/db
