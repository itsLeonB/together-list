FROM node:23-alpine

WORKDIR /app

# Copy dependency declarations and install
COPY package*.json ./
RUN npm ci --only=production

# Copy app source code
COPY . .

# Copy and set permissions on entrypoint script
COPY bin/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh && \
	chown -R node:node /app

# Run as non-root user
# USER node

# Start with exec form to comply with SonarQube rule S7019
CMD ["/usr/local/bin/docker-entrypoint.sh"]
