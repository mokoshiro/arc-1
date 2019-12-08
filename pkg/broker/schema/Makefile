SHELL := /bin/zsh
ENV=local
include .env.${ENV}

.PHONY: import
import:
	@mysqldef -u ${USER} -P${PORT} -p${PASSWORD} --host ${HOST} ${DB} < schema.sql
