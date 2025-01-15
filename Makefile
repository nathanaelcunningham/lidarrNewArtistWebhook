# Load env from file
ifneq (,$(wildcard ./.env))
	include ./.env
	export
endif

run:
	LIBRARY_PATH=${LIBRARY_PATH} go run .
