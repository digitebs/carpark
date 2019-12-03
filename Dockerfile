FROM golang:1.13.4

EXPOSE 10000

WORKDIR /go/src/github.com/iody/carpark
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["carpark"]