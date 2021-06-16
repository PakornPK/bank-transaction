FROM golang:1.16-buster as build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLE=0

RUN mkdir -p /workspace
WORKDIR /workspace
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -o .build/bin -ldflags "-w -s" ./cmd/ 

FROM gcr.io/distroless/base-debian10

COPY --from=build /workspace/.build/* /

EXPOSE 3000
ENTRYPOINT [ "/bin" ]
