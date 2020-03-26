FROM golang:1.13
RUN mkdir /app
WORKDIR /app
COPY ./Server/go.mod /app
RUN go mod download
COPY ./Server /app
RUN go build -o main /app

ENTRYPOINT [ "/app/main" ]
