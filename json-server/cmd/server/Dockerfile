FROM golang:1.14.0 as gobuilder
WORKDIR /go-app
COPY . /go-app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o json-server

FROM scratch
COPY --from=gobuilder /go-app/json-server .
ENTRYPOINT ["/json-server"]