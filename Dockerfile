FROM golang:alpine3.16

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /demo-go-app

EXPOSE 8080

CMD [ "/demo-go-app" ]
