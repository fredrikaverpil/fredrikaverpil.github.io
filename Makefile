.PHONY: serve build clean all

all: clean build serve

serve:
	hugo server -D

build:
	hugo --minify
	npx pagefind --site public

clean:
	rm -rf public resources