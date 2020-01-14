SHELL := /bin/zsh
ENV=local
include .env.${ENV}

.PHONY: import
import:
	@mysqldef -u ${USER} -P${PORT} -p${PASSWORD} --host ${HOST} ${DB} < schema.sql

.PHONY: init
init:
	@mysql -u ${USER} -p${PASSWORD} -P${PORT} --host ${HOST} -e'create database if not exists executor;'