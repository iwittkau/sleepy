FROM golang:1.12

ARG GOPROXY

ADD . /build

WORKDIR /build

RUN CGO_ENABLED=0 go build -o sleepy .

FROM scratch

COPY --from=0 /build/sleepy /sleepy

ENTRYPOINT [ "/sleepy" ]