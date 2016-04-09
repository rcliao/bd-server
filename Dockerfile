FROM golang:onbuild

EXPOSE 9000

RUN mkdir /usr/src/bd-server

COPY main.go /usr/src/bd-server

WORKDIR /usr/src/bd-server

RUN go build -o main .

CMD ["./main"]
