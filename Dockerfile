FROM golang:1.20-alpine as builder
WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# ======================================================

FROM golang:1.20-alpine as final

ARG ENVIRONMENT
ARG AUTH_SERVER
ARG SEC_PREFIX

WORKDIR /go/src/app

COPY --from=builder /go/bin/api.default.marincor /go/bin/
COPY --from=builder /go/src/app/private.pem .
COPY --from=builder /go/src/app/config.yaml .
COPY --from=builder /go/src/app/marincor.json .

EXPOSE 9090

ENV CREDENTIALS=/go/src/app/marincor.json
ENV AUTH_SERVER=$AUTH_SERVER
ENV SEC_PREFIX=$SEC_PREFIX
ENV ENVIRONMENT=$ENVIRONMENT

RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD api.default.marincor
