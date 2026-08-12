tailwind:
	tailwindcss -i ./static/css/app.css -o ./static/css/app.min.css --minify --watch

templ:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ./cmd/app"

dev:
	make -j2 tailwind templ