FROM golang:1.23

COPY ./ /src
WORKDIR /src
RUN go mod download
RUN go build -o main main.go

EXPOSE 8080
CMD ["./main"]