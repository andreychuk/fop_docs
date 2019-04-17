FROM golang:1.12-alpine3.9 as builder

WORKDIR /src

RUN apk add --no-cache ca-certificates git

COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /fop_docs .

FROM andreychuk/go-wkhtmltopdf:v0.1.0 AS final

WORKDIR /src

RUN mkdir reports
RUN mkdir templates

COPY --from=builder /src/templates/* /src/templates/
COPY --from=builder /fop_docs /fop_docs

CMD ["/fop_docs"]