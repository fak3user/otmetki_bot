FROM okteto/golang:1.21 AS builder

RUN go version

COPY ./ /otmetki
WORKDIR /otmetki

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 go build -o ./app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /otmetki/app .

EXPOSE 8080

CMD [ "./app" ]
