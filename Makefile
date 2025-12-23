.PHONY: serve build clean all

all: clean build serve

serve:
	go tool hugo server -D

build:
	go tool hugo --minify --environment production
	npx pagefind --site public

clean:
	rm -rf public resources