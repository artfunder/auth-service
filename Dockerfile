FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server

# Final Image
FROM scratch

WORKDIR /bin

COPY --from=builder /app/server .
COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/

CMD [ "/bin/server" ]