FROM golang:latest

WORKDIR /go/src/WebService
COPY . .

RUN go get ./...
RUN go install -v ./...
RUN go build

CMD ./WebService