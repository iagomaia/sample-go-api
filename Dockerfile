FROM golang:1.19-alpine as build
RUN apk add --no-cache make

WORKDIR $GOPATH/cps-go-api
COPY . .

RUN mkdir -p build build/static
RUN go build -mod=mod -v -o ./build ./...
RUN cp ./docs/swagger.yaml ./build/static/swagger.yaml

FROM alpine:latest

RUN apk update
RUN apk upgrade
RUN apk add --no-cache bash
RUN apk add --upgrade --no-cache coreutils

WORKDIR /

COPY --from=build /go/cps-go-api/build /usr/local/bin

ARG PORT=3000
ENV SV_PORT=${PORT}

EXPOSE ${SV_PORT}

CMD ["api"]