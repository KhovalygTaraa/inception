FROM alpine:3.16

LABEL maintainer="swquinc"

WORKDIR /app

RUN echo "http://dl-cdn.alpinelinux.org/alpine/v3.16/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/v3.16/community" >> /etc/apk/repositories && \
    apk update && \
	apk upgrade && \
	apk add mariadb=10.6.9-r0 && \
    apk add mariadb-client=10.6.9-r0 && \
    apk add go=1.18.6-r0 && \
    apk add su-exec=0.2-r1

COPY ./conf /app

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["mariadbd", "--defaults-extra-file=/app/mariadb.conf", "--datadir=/app/data"]
EXPOSE 3306
