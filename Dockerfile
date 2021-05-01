FROM golang:latest


WORKDIR /go/src/app
COPY . .

# RUN go get -d -v github.com/adam-lavrik/go-imath.git
# RUN go build ./...

CMD ["factorizer"]