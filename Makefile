server:
	air
sass:
	bin/dart-sass/sass --no-source-map --watch internal/assets/stylesheets/index.scss public/css/index.css
worker:
	go run cmd/worker/main.go