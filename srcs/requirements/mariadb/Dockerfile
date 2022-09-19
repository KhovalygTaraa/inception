FROM alpine:3.16

LABEL maintainer="swquinc"

COPY ./conf /app

RUN apk update && \
	apk upgrade && \
	apk add mariadb su-exec

ENTRYPOINT ["/app/setup.sh"]
CMD ["mariadbd", "--default-file=/app/mariadb.conf"]
EXPOSE 3306
