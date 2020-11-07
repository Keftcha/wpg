FROM golang:1.15 as builder

WORKDIR /usr/src/wpg
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/wpg


FROM alpine

WORKDIR /bin/wpg

COPY --from=builder /bin/wpg .

COPY ./pages pages/

CMD ["./wpg"]
