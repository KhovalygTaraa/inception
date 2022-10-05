FROM alpine:3.16

LABEL maintainer="swquinc"

RUN apk update &&\
	apk upgrade &&\
    apk add prometheus &&\
    apk add go bash

WORKDIR /app

COPY conf /app

ENTRYPOINT ["go", "run", "setup.go"]
CMD ["prometheus", "--config.file=/app/prometheus.yml"]
EXPOSE 9090