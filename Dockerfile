FROM golang:1.22

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build

CMD [ "go", "run", "main.go" ]