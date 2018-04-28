FROM golang:1.10

WORKDIR /go/src/github.com/suddutt1/noop/
COPY *.go .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT /go/bin/noop

# Document that the service listens on port 8080.
EXPOSE 8080
