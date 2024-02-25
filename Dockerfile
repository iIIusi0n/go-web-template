FROM golang:1.19

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go build api-server/cmd/api-server

CMD ["./api-server"]
