FROM golang:1.10 AS build
WORKDIR /go/src
COPY cmd internal ./

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o bin github.com/jimmiepr/bank-transaction/cmd

FROM scratch AS runtime
COPY --from=build /go/src/bin ./
EXPOSE 3000/tcp
ENTRYPOINT ["./bin"]
