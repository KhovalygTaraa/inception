FROM alpine:3.16

LABEL maintainer="swquinc"

RUN apk update &&\
	apk upgrade &&\
    apk add redis &&\
    apk add go bash

WORKDIR /app

COPY conf /app

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["redis-server", "/app/redis.conf"]
EXPOSE 6379