FROM golang:1.21

WORKDIR /app

COPY . ./

RUN go build -o ./bin cmd/server/main.go

EXPOSE 8001

CMD [ "./bin/main"]