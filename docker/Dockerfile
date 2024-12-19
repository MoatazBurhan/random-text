FROM golang:1.19

WORKDIR /app

COPY ./dict/words ./
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /random-text

RUN groupadd -r burhan && useradd -r -g burhan moataz
USER moataz

EXPOSE 8080

CMD ["/random-text"]