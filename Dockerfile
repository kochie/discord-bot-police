FROM golang AS golang

WORKDIR /go/src/policebot

RUN apt update && apt install -y \
    tzdata \
    zip \
    ca-certificates \
    ffmpeg \
    build-essential \
    git \
    pkgconf

# Download, build and install Opus 1.1.2
# Keep this for the moment, I've got a suspicion that the vendored version will be the wrong platform.
# Will need to symlink with vendor
#RUN wget https://archive.mozilla.org/pub/opus/opus-1.1.2.tar.gz \
#    && tar -xzvf opus-1.1.2.tar.gz \
#    && cd opus-1.1.2 \
#    && ./configure \
#    && make \
#    && make install \
#    && ldconfig \
#    && cd .. \
#    && rm -rf opus-1.1.2 opus-1.1.2.tar.gz

# Verify installation
# RUN pkg-config --modversion opus

COPY . .

RUN go build -ldflags="-w -s" -o policebot src/main.go
RUN chmod +x ./policebot
CMD ["./policebot"]
