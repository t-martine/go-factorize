FROM golang:latest

WORKDIR /go/src/app
COPY . .

#RUN go get -d -v ./...
RUN go build ./...

CMD ["factorizer"]