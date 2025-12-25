.PHONY: serve build clean all

all: clean build serve

serve:
	go tool hugo server -D

build:
	go tool hugo --minify --environment production
	# legacy mkdocs-created RSS feed
	cp public/blog/index.xml public/feed_rss_created.xml
	bunx pagefind --site public

clean:
	rm -rf public resources
