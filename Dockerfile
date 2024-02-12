FROM golang:latest

ENV MYAPP_PORT=3000

WORKDIR /app

COPY . .

Run go build -o main main.go


EXPOSE $MYAPP_PORT

CMD ["/app/main"]
