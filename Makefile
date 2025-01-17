# Load env from file
ifneq (,$(wildcard ./.env))
	include ./.env
	export
endif

run:
	LIBRARY_PATH=${LIBRARY_PATH} go run .

release:
	docker build --platform linux/amd64 -t lidarr-artist-webhook:latest .
	docker tag lidarr-artist-webhook ghcr.io/nathanaelcunningham/lidarr-artist-webhook:latest
	docker push ghcr.io/nathanaelcunningham/lidarr-artist-webhook:latest
