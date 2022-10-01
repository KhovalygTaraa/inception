FROM alpine:3.16

LABEL maintainer="swquinc"

WORKDIR /app

RUN apk update && \
	apk upgrade && \
    apk add go=1.18.6-r0 && \
    apk add su-exec=0.2-r1 && \
    apk add php8=8.0.22-r0 && \
    apk add php8-fpm=8.0.22-r0 && \
    apk add php8-mysqli=8.0.22-r0 && \
    apk add php8-curl=8.0.22-r0 && \
    apk add php8-exif=8.0.22-r0 && \
    apk add php8-fileinfo=8.0.22-r0 && \
    apk add php8-mbstring=8.0.22-r0 && \
    apk add php8-xml=8.0.22-r0 && \
    apk add php8-xmlreader=8.0.22-r0 && \
    apk add php8-simplexml=8.0.22-r0 && \
    apk add php8-zip=8.0.22-r0 && \
    apk add php8-dom=8.0.22-r0 && \
    apk add php8-pecl-redis=5.3.7-r0 && \
    apk add php8-pecl-imagick=3.7.0-r0 && \
    apk add pcre=8.45-r2 && \
    wget -O wordpress.tar.gz https://wordpress.org/wordpress-6.0.2.tar.gz && \
    tar -xzf wordpress.tar.gz

COPY ./conf /app

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["php8-fpm"]
EXPOSE 9000
