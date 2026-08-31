#!/usr/bin/env bash

set -e

VPS_USER="usman"
VPS_HOST="cp.telzz.com"
SSH_KEY="$HOME/.ssh/usman-contabo"

DOMAIN="api.erosyncng.com"

NGINX_CONTAINER="nginx"

docker run --rm \
    -v ~/nginx/certbot/letsencrypt:/etc/letsencrypt \
    -v ~/nginx/certbot/www:/var/www/certbot \
    certbot/certbot:latest certonly --webroot -w /var/www/certbot --non-interactive --agree-tos  --email admin@erosyncng.com  -d api.erosyncng.com



server {
    listen 80;
    server_name api.erosyncng.com;

    location /.well-known/acme-challenge/ {
        root /var/www/certbot;
    }

    location / {
        return 301 https://$host$request_uri;
    }
}

server {
    listen 443 ssl;
    server_name api.erosyncng.com;

    ssl_certificate /etc/letsencrypt/live/api.erosyncng.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.erosyncng.com/privkey.pem;

    location / {
        proxy_pass http://erosync-api:8000;

        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}