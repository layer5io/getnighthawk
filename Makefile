# Building GetNighthawk

jekyll=bundle exec jekyll

site:
	cd docs; $(jekyll) serve --drafts --livereload --config _config.yml

site-setup:
	$(MAKE) -C docs setup

site-build:
	$(MAKE) -C docs build BASEURL="$(BASEURL)" SITE_URL="$(SITE_URL)"

site-build-preview:
	$(MAKE) -C docs build-preview BASEURL="$(BASEURL)" SITE_URL="$(SITE_URL)"

setup:
	docker pull envoyproxy/nighthawk-dev; cd cmd; go mod tidy;
run:
	go run cmd/main.go

