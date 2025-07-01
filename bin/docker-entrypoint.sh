#!/bin/sh
set -e

# Ensure auth folder exist
mkdir -p /app/.wwebjs_auth

# Set correct permissions for the node user
chown -R node:node /app/.wwebjs_auth

# Run the app as node
exec npm start
