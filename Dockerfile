FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go build -o app .

###
FROM alpine:latest
COPY --from=builder /go/src/app/app /usr/local/bin/
EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/app" ]