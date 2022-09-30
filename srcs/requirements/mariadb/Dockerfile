FROM alpine:3.16

LABEL maintainer="swquinc"

WORKDIR /app

RUN apk update && \
	apk upgrade && \
	apk add mariadb mariadb-client su-exec go
#    go mod init mariadb && \
#    go get github.com/go-sql-driver/mysql

COPY ./conf /app

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["mariadbd", "--defaults-extra-file=/app/mariadb.conf", "--datadir=/app/data"]
EXPOSE 3306
