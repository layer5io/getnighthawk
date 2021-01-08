# Building GetNighthawk project site

jekyll=bundle exec jekyll

site:
	$(jekyll) serve --drafts --livereload --config _config.yml

docker:
	docker run --name getnighthawk --rm -p 4000:4000 -v `pwd`:"/srv/jekyll" jekyll/jekyll:4.0.0 bash -c "bundle install; jekyll serve --drafts --livereload --config _config.yml"