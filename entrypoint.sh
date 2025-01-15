#!/bin/sh
# Set PUID and PGID if they are provided
if [ -n "$PUID" ] && [ -n "$PGID" ]; then
    echo "Creating user with UID: $PUID and GID: $PGID"
    addgroup -g "$PGID" appgroup
    adduser -u "$PUID" -G appgroup -D appuser
    chown appuser:appgroup /webhook
    exec su-exec appuser /webhook
else
    echo "Running as default user"
    exec /webhook
fi
