FROM node:23-alpine

WORKDIR /app

COPY package*.json ./

RUN npm ci --only=production

COPY . .

RUN chown -R node:node /app

USER node

CMD ["npm", "start"]
