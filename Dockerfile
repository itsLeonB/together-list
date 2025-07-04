# --- Build Stage ---
FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -buildvcs=false -ldflags='-w -s' \
    -o /together-list ./cmd/whatsapp/main.go

# --- Release Stage with Chrome ---
FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    ca-certificates fonts-liberation gnupg libasound2 libatk-bridge2.0-0 \
    libatk1.0-0 libc6 libgbm1 libgtk-3-0 libnss3 libu2f-udev libx11-xcb1 \
    libxcomposite1 libxdamage1 libxrandr2 libxss1 libxtst6 lsb-release \
    tzdata wget xdg-utils \
 && wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | gpg --dearmor > /usr/share/keyrings/google.gpg \
 && echo "deb [arch=amd64 signed-by=/usr/share/keyrings/google.gpg] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list \
 && apt-get update && apt-get install -y google-chrome-stable \
 && apt-get clean && rm -rf /var/lib/apt/lists/* \
 && useradd -m -u 10001 appuser

COPY --from=build-stage /together-list /together-list

USER appuser

WORKDIR /

ENV GOOGLE_CHROME_SHIM=/usr/bin/google-chrome

ENTRYPOINT ["/together-list"]
