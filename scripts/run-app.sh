#!/bin/sh
# This script is used to run the local version of the application.

export ZINC_FIRST_ADMIN_USER="admin"
export ZINC_FIRST_ADMIN_PASSWORD="password"

go run main.go