FROM golang:1.12-alpine as builder

WORKDIR /go/src/

COPY . .

# RUN GOOS=linux go build -o main -ldflags="-s -w"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# CMD ["./main"]

FROM scratch

WORKDIR /go/src/

COPY --from=builder /go/src/ .

CMD [ "./main" ]