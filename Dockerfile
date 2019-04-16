FROM andreychuk/go-wkhtmltopdf:v0.1.0

WORKDIR /src

COPY ./ ./

RUN mkdir reports

RUN go mod download

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /fop_docs .

CMD ["./fop_docs"]