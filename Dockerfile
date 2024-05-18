FROM golang:1.22.3-alpine3.19 AS builder

WORKDIR /app

COPY . . 

ENV GOARCH=amd64
ENV GOOS=linux 

RUN go build -o main .

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

ENV DB_HOST=host.docker.internal
ENV DB_PORT=5432
ENV DB_USER=dionfananie
ENV DB_PASSWORD=yoloyolo
ENV DB_NAME=eniqilodb
ENV DB_PARAMS=sslmode=disable
ENV JWT_SECRET=secretly
ENV BCRYPT_SALT=8

CMD ["/app/main"]