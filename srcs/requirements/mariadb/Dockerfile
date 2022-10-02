FROM alpine:3.16

LABEL maintainer="swquinc"

RUN echo "http://dl-cdn.alpinelinux.org/alpine/v3.16/main" >> /etc/apk/repositories &&\
    echo "http://dl-cdn.alpinelinux.org/alpine/v3.16/community" >> /etc/apk/repositories &&\
    apk update &&\
	apk upgrade &&\
	apk add mariadb &&\
    apk add mariadb-client &&\
    apk add go &&\
    apk add su-exec

ENV MARIADB_CONFIG=/app/mariadb.conf\
    DATA_DIR=/app/data\
    APP_DIR=/app

WORKDIR $APP_DIR

COPY ./conf $APP_DIR

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["mariadbd", "--defaults-extra-file=/app/mariadb.conf", "--datadir=/app/data"]
EXPOSE 3306
