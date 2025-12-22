# Make sure to specify the same Go version as the one in the go.mod file.
# For example, golang:1.22.1-alpine.
FROM golang:1.25.5-alpine

WORKDIR /app

ARG bin_to_build

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN go build -o svr cmd/${bin_to_build}/*.go

EXPOSE 8080
CMD [ "./svr" ]
