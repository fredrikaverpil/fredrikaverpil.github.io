.PHONY: serve build clean

serve:
	hugo server -D

build:
	hugo --minify
	npx pagefind --site public

clean:
	rm -rf public resources