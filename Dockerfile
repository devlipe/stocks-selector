FROM golang:1.18

WORKDIR /go/app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /main  -v

CMD ["/main"]