.PHONY: serve build clean

serve:
	hugo server -D

build:
	hugo --minify

clean:
	rm -rf public resources