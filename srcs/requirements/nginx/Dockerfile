FROM alpine:3.16

LABEL maintainer="swquinc"

WORKDIR /app

RUN	apk update && \
	apk upgrade && \
	apk add openrc openssl3 nginx

COPY ./conf /app
RUN mkdir -p /run/openrc/ && \
    mkdir -p /run/nginx && \
    mkdir -p /var/www/html && \
    cp /app/nginx.conf  /etc/nginx/nginx.conf && \
    cp /app/default.conf /etc/nginx/http.d/default.conf && \
    mv /var/lib/nginx/html/index.html /var/www/html && \
    openssl3 req -x509 -noenc -newkey rsa:4096 -keyout /app/nginx.key -out /app/nginx.crt \
    -sha256 -days 365 -subj "/C=RU/ST=Moscow/L=Moscow/O=42 ecole/OU=21 school/CN=swquinc" && \
    touch /run/nginx/nginx.pid && \
    touch /run/openrc/softlevel && \
    openrc

#	rc-service nginx start && \
#	nginx -s stop

ENTRYPOINT ["nginx", "-g", "daemon off;"]

EXPOSE 443
