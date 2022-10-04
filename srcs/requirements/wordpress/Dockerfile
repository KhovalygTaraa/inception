FROM alpine:3.16

LABEL maintainer="swquinc"

RUN apk update &&\
	apk upgrade &&\
    apk add go &&\
    apk add su-exec &&\
    apk add php8 &&\
    apk add php8-fpm &&\
    apk add php8-mysqli &&\
    apk add php8-curl &&\
    apk add php8-exif &&\
    apk add php8-phar &&\
    apk add php8-fileinfo &&\
    apk add php8-mbstring &&\
    apk add php8-xml &&\
    apk add php8-xmlreader &&\
    apk add php8-simplexml &&\
    apk add php8-zip &&\
    apk add php8-dom &&\
    apk add php8-ctype &&\
    apk add php8-pecl-redis &&\
    apk add php8-pecl-imagick &&\
    apk add pcre

ENV WP_CONFIG="/app/wp-config.php"\
    WORDPRESS_PATH="/app/data/wordpress"\
    PHP_FPM_GLOBAL_CONFIG="/app/php-fpm.conf"\
    PHP_FPM_WWW_CONFIG="/app/www.conf"\
    APP_DIR="/app"\
    DATA_DIR="/app/data"\
    WORDPRESS_LINK="https://wordpress.org/wordpress-6.0.tar.gz"\
    WP_CLI_LINK="https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar"

WORKDIR $APP_DIR

COPY ./conf $APP_DIR

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["php-fpm8"]
EXPOSE 9000
