FROM nginx:alpine
VOLUME /var/cache/nginx
COPY ./Client/dist /usr/share/nginx/html
COPY ./Client/config/nginx.conf /etc/nginx/conf.d/default.conf