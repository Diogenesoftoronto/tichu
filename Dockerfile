# Install the base requirements for app.
# This stage is to support development

FROM golang:1.19.4-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /tichu

EXPOSE 8080

ENTRYPOINT [ "/tichu" ]
