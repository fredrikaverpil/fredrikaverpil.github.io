.PHONY: serve build clean all test ci

all: clean build serve

serve:
	go tool hugo server -D

test:
	cd ./content/blog/2025-12-28-gos-secret-weapon/examples && go test -v ./...

build:
	go tool hugo --minify --environment production
	# legacy mkdocs-created RSS feed
	cp public/blog/index.xml public/feed_rss_created.xml
	bunx pagefind --site public

clean:
	rm -rf public resources

ci: test build
