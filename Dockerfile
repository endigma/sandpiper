FROM golang:1.15.8 as build-env

WORKDIR /go/src/sandpiper
ADD . /go/src/sandpiper

RUN go get -d -v ./...

RUN go build -o /go/bin/sandpiper

FROM gcr.io/distroless/base
COPY ./assets/ /assets
COPY --from=build-env /go/bin/sandpiper /

CMD ["/sandpiper", "/assets/config.json"]