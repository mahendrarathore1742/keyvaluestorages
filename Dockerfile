FROM golang:1.20-alpine as builder

WORKDIR /keyvaluestore
COPY . ./

RUN go build -o main ./cmd/main.go


FROM alpine:buster-slim

WORKDIR /app
COPY --from=builder /keyvaluestore/main .

CMD [ "./main" ]