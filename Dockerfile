FROM golang:1.16 as build

COPY . /go/src/github.com/danshu93/go-skeleton

WORKDIR /go/src/github.com/danshu93/go-skeleton

RUN go get ./...
RUN go test ./...

RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o /app/app /go/src/github.com/danshu93/go-skeleton/cmd/app/app.go

FROM alpine as app

COPY --from=build /app/app /app/app

CMD /app/app
