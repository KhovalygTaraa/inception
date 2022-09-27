FROM alpine:3.16

LABEL maintainer="swquinc"

COPY ./conf /app

RUN apk update && \
	apk upgrade && \
	apk add mariadb su-exec go

ENTRYPOINT ["go", "run", "/app/setup.go"]
CMD ["mariadbd", "--defaults-extra-file=/app/mariadb.conf", "--datadir=/app/data"]
EXPOSE 3306
