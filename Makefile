# Building GetNighthawk

jekyll=bundle exec jekyll

site:
	cd docs; $(jekyll) serve --drafts --livereload --config _config.yml

setup:
	docker pull envoyproxy/nighthawk-dev; cd cmd; go mod tidy;
run:
	go run cmd/main.go

