FROM golang:latest as golang
WORKDIR /go/src/policebot
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o policebot


FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -r -0 /zoneinfo.zip .


FROM scratch
# the test program:
COPY --from=golang /go/src/policebot/policebot /policebot
# the timezone data:
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
# the tls certificates:
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV DISCORD_TOKEN = ""

CMD ["/policebot"]
