FROM golang:1.16-alpine3.13 AS base

ENV APP_USER app
ENV APP_HOME /subtitleDownloader

WORKDIR $APP_HOME
COPY go.mod .

RUN go mod download
RUN go install -v ./...

COPY src src

RUN cd src \
&& go build -o /bin/subtitle-downloader .

FROM scratch AS download-exe
COPY --from=base /bin/subtitle-downloader /