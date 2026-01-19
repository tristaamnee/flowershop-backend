FROM golang:1.25.5-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /out/main ./
ENTRYPOINT ["/out/main"]

CMD ./out/main