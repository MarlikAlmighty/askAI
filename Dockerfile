FROM golang:1.14-alpine3.14 AS builder

WORKDIR /go/src/kickHisAss

COPY . .

RUN go build -o /go/src/kickHisAss/bot /go/src/kickHisAss/cmd/main.go

FROM alpine:3.14

COPY --from=builder /go/src/kickHisAss/bot /

ENV AI_TOKEN="XXX"
ENV BOT_TOKEN="XXXX:XXXX"
ENV CHANNEL="-100XXX"

EXPOSE 3000
CMD ["/bot"]
