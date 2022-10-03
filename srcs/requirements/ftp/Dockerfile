FROM alpine:3.16

LABEL maintainer="swquinc"

RUN apk update &&\
	apk upgrade &&\
    apk add vsftpd &&\
    apk add go bash

WORKDIR /app

COPY ./conf /app

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["vsftpd"]
EXPOSE 20 21 20005 20006 20007
