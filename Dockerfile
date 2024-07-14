FROM golang:alpine as golang

WORKDIR /go/src/policebot
COPY . .
RUN apk --no-cache add \
    tzdata \
    zip \
    ca-certificates \
    opus-tools \
    ffmpeg \
    build-base \
    git \
    pkgconf \
    opus-dev
RUN go build -ldflags="-w -s" -o policebot
# RUN chmod +x policebot


#FROM alpine:latest as alpine
#RUN apk --no-cache add tzdata zip ca-certificates opus-tools
#WORKDIR /usr/share/zoneinfo
## -0 means no compression.  Needed because go's
## tz loader doesn't handle compressed data.
#RUN zip -r -0 /zoneinfo.zip .



#FROM scratch
## the test program:
#COPY --from=golang /go/src/policebot/policebot /policebot
#COPY --from=golang /go/src/policebot/assets /assets
## the timezone data:
#ENV ZONEINFO /zoneinfo.zip
#COPY --from=alpine /zoneinfo.zip /
## the tls certificates:
#COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#ENV DISCORD_TOKEN = ""
RUN chmod +x ./policebot
CMD ["./policebot"]
