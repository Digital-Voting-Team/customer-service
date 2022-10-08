FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/customer-service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/customer-service /go/src/customer-service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/customer-service /usr/local/bin/customer-service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["customer-service"]
