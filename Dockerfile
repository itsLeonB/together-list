# Use slim image with Chromium and required libs
FROM node:23-slim

WORKDIR /app

# Install Chromium + dependencies
RUN apt-get update && apt-get install -y \
  chromium \
  libatk-bridge2.0-0 \
  libatk1.0-0 \
  libcups2 \
  libdbus-1-3 \
  libgdk-pixbuf2.0-0 \
  libnspr4 \
  libnss3 \
  libx11-xcb1 \
  libxcomposite1 \
  libxdamage1 \
  libxrandr2 \
  fonts-liberation \
  libasound2 \
  libappindicator3-1 \
  xdg-utils \
  ca-certificates \
  wget \
  && apt-get clean && rm -rf /var/lib/apt/lists/*

# Copy app
COPY package*.json ./
RUN npm ci --only=production
COPY . .

# Fix permissions and switch user
RUN chown -R node:node /app
USER node

CMD ["npm", "start"]
