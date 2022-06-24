FROM golang as builder
WORKDIR /go/src/
COPY ./redirector.go .
RUN go build *.go
FROM gcr.io/distroless/base-debian11 as release
COPY --from=builder /go/src/redirector /redirector
# RUN chmod +x /redirector
CMD [ "/redirector" ]