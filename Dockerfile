FROM golang:1.14-alpine3.12 AS builder

WORKDIR /go/src/kickHisAss

COPY . .

RUN go build -o /go/src/kickHisAss/kick /go/src/kickHisAss/cmd/main.go

FROM alpine:3.12

COPY --from=builder /go/src/kickHisAss/kick /

ENV WEB_HOOK="https://domen.com/"
ENV BOT_TOKEN="XXXX:XXXX"
ENV PORT="3000"

EXPOSE 3000
CMD ["/kick"]
