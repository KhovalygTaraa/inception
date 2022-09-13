FROM alpine:3.16

LABEL maintainer="swquinc"

COPY ./conf /app

RUN apk update && \
	apk upgrade && \
	apk add su-exec openrc php8 && \
	mkdir -p /run/openrc/ && \
	touch /run/openrc/softlevel && \
	openrc

CMD ["sh", "/app/script.sh"]

EXPOSE 9000
