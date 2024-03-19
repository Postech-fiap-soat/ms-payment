FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o payment ./cmd/server

EXPOSE 8001

CMD [ "./payment"]