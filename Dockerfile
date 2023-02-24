FROM registry.c4i.vn/c4i/golang:1.19-buster AS builder
LABEL stage=builder

ARG SVC

WORKDIR /go/src/go-template
COPY . .

RUN make go-build \
    && mv bin/$SVC /exe

FROM scratch
COPY --from=builder /exe /
ENTRYPOINT ["./exe", "start"]