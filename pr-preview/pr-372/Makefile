# Building GetNighthawk project site

jekyll=bundle exec jekyll
jekyll_production=JEKYLL_ENV=production $(jekyll)
jekyll_preview=JEKYLL_ENV=preview $(jekyll)

setup:
	bundle install

site:
	$(jekyll) serve --drafts --livereload --config _config.yml

build:
	@if [ -n "$(SITE_URL)" ]; then \
		runtime_config="$$(mktemp)"; \
		printf 'url: "%s"\n' "$(SITE_URL)" > "$$runtime_config"; \
		$(jekyll_production) build --config _config.yml,"$$runtime_config" $(if $(BASEURL),--baseurl "$(BASEURL)"); \
		rm -f "$$runtime_config"; \
	else \
		$(jekyll_production) build --config _config.yml $(if $(BASEURL),--baseurl "$(BASEURL)"); \
	fi

build-preview:
	@if [ -n "$(SITE_URL)" ]; then \
		runtime_config="$$(mktemp)"; \
		printf 'url: "%s"\n' "$(SITE_URL)" > "$$runtime_config"; \
		$(jekyll_preview) build --config _config.yml,"$$runtime_config" --baseurl "$(BASEURL)"; \
		rm -f "$$runtime_config"; \
	else \
		$(jekyll_preview) build --config _config.yml --baseurl "$(BASEURL)"; \
	fi

docker:
	docker run --name getnighthawk --rm -p 4000:4000 -v `pwd`:"/srv/jekyll" jekyll/jekyll:4.0.0 bash -c "bundle install; jekyll serve --drafts --livereload --config _config.yml"
