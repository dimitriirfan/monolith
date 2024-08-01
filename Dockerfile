FROM nginx:1.21.6-alpine
COPY index.html /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf