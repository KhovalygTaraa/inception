FROM alpine:3.16

LABEL maintainer="swquinc"

RUN apk update &&\
	apk upgrade &&\
    apk add go

WORKDIR $APP_DIR

COPY ./conf /app

ENTRYPOINT ["go", "run", "setup.go"]

EXPOSE 9999