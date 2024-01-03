FROM golang:1.21

WORKDIR /usr/src/app

# System settings
RUN ln -sf /usr/share/zoneinfo/Asia/Taipei /etc/localtime

# Install the dockerize
ENV DOCKERIZE_VERSION v0.7.0

RUN apt-get update \
    && apt-get install -y wget \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apt-get autoremove -yqq --purge wget && rm -rf /var/lib/apt/lists/*

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go mod tidy
RUN go build -v -o /usr/local/bin/app .

CMD dockerize -wait tcp://db:5432 -timeout 60s && app