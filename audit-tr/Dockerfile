FROM daystram/go-builder:1.16 as builder
WORKDIR /build
COPY audit-tr audit-tr
COPY audit-be audit-be
COPY proto proto
COPY Makefile .
RUN make
WORKDIR /build/audit-tr
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o app .

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/audit-tr/app ./
ENTRYPOINT ["/app"]
