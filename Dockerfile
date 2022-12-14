FROM golang:1.19

WORKDIR /app
COPY go.mod .

RUN go mod download

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]

EXPOSE 3000