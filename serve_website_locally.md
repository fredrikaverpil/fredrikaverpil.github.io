---
layout: null
title: null
---

# Full instructions at https://help.github.com/articles/setting-up-your-github-pages-site-locally-with-jekyll/#step-1-create-a-local-repository-for-your-jekyll-site

# Install bundler (requires Ruby 2.x.x)
gem install bundler

# Set up Gemfile and install
cd GITHUB_PAGES_REPO
echo "source 'https://rubygems.org'" > Gemfile
echo "gem 'github-pages', group: :jekyll_plugins" >> Gemfile
bundle install

# Serve website
bundle exec jekyll serve

# Visit website at http://localhost:4000
# Cancel serving with ctrl+c
