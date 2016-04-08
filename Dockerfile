FROM golang:onbuild

EXPOSE 9000

COPY main.go .

CMD go-wrapper run main.go
