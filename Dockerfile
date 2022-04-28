FROM golang:latest

WORKDIR /

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o cmd/eeveentory/eeveentory

EXPOSE 8888

CMD [ "cmd/eeveentory/eeveentory" ]