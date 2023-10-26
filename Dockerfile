FROM golang:1.20

ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./bin/app ./cmd/household-account-book

CMD [ "./bin/app" ]
