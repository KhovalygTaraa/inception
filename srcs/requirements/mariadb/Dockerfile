FROM alpine:3.16

LABEL maintainer="swquinc"

RUN apk update && \
	apk upgrade && \
	apk add mariadb su-exec go && \
	go install github.com/go-sql-driver/mysql

COPY ./conf /app

ENTRYPOINT ["go", "run", "/app/setup.go"]
CMD ["mariadbd", "--defaults-extra-file=/app/mariadb.conf", "--datadir=/app/data"]
EXPOSE 3306
