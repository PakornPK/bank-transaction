FROM golang:1.16 AS build
WORKDIR /go/src
COPY . .

RUN go get -d -v ./...

RUN go build -a -o bin github.com/jimmiepr/bank-transaction/cmd

FROM scratch AS runtime
COPY --from=build /go/src/bin ./
EXPOSE 3000/tcp
ENTRYPOINT ["./bin"]
