FROM golang:1.18.3-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o music-shop ./app/cmd/main.go

CMD ["./music-shop"]
