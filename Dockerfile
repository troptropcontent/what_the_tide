FROM golang:1.22

ADD go.mod go.sum /

RUN go mod download

WORKDIR /app

COPY . .

ENV GO_ENV=production

# Download SASS
RUN curl -sL -o /app/bin/dart-sass-1.78.0-linux-x64.tar.gz https://github.com/sass/dart-sass/releases/download/1.78.0/dart-sass-1.78.0-linux-x64.tar.gz
# Unzip sass
RUN tar -xvf /app/bin/dart-sass-1.78.0-linux-x64.tar.gz -C /app/bin/
RUN rm /app/bin/dart-sass-1.78.0-linux-x64.tar.gz
# Compile SCSS files into CSS files
RUN /app/bin/dart-sass/sass --no-source-map /app/internal/assets/stylesheets/index.scss /app/public/css/index.css
# Delete SASS binary and scss files as those are useless now
RUN rm -rf /app/bin/dart-sass
RUN rm -rf /app/internal/stylesheets

RUN go build -o /app/web /app/cmd/web/main.go

EXPOSE 3001
CMD ["/app/web"]