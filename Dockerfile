FROM golang:1.16 AS build
WORKDIR /go/src
COPY . .

RUN go get -d -v ./...

RUN go build -o /out/bin ./cmd/bank-transaction.go

FROM scratch AS runtime
COPY --from=build /out/bin ./
EXPOSE 3000/tcp
ENTRYPOINT ["./bin"]
