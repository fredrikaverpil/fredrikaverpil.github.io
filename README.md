# fredrikaverpil.github.io

This is the source code of [my personal blog](https://fredrikaverpil.github.io), orginally based on [poole/hyde](https://github.com/poole/hyde).

## Run website locally

Full instructions [here](https://help.github.com/articles/setting-up-your-github-pages-site-locally-with-jekyll/#step-1-create-a-local-repository-for-your-jekyll-site).

### Install bundler

```bash
gem install bundler  # requires Ruby 2.x.x.
```

### Set up Gemfile and install

```bash
cd <GITHUB_PAGES_REPO>
echo "source 'https://rubygems.org'" > Gemfile
echo "gem 'github-pages', group: :jekyll_plugins" >> Gemfile
apt-get install ruby ruby-dev
bundle config set --local path 'vendor/bundle'
bundle install --path vendor/bundle
```

### Serve website

```bash
bundle exec jekyll serve
```

### Visit website

- Visit website at http://localhost:4000
- Cancel serving with ctrl+c
