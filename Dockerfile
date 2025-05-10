FROM golang:1.22

ADD . /app

WORKDIR /app

RUN go build -o /knight_board

CMD [ "/knight_board" ]