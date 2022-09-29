FROM alpine:3.16

LABEL maintainer="swquinc"

RUN apk update && \
	apk upgrade && \
	apk add mariadb su-exec go
#	go get github.com/go-sql-driver/mysql

COPY ./conf /app

ENTRYPOINT ["go", "run", "-modfile=/app/go.mod", "/app/setup.go"]
CMD ["mariadbd", "--defaults-extra-file=/app/mariadb.conf", "--datadir=/app/data"]
EXPOSE 3306
