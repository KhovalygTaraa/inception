FROM alpine:3.16

LABEL maintainer="swquinc"

COPY ./conf /app

RUN apk update && \
	apk upgrade && \
	apk add su-exec mariadb-openrc mariadb mariadb-client && \
	mkdir -p /run/openrc/ && \
	touch /run/openrc/softlevel

#ENTRYPOINT ["mariadb", ";"]
CMD ["sleep", "infinity"]


