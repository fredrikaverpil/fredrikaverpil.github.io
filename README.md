# fredrikaverpil.github.io

## WIP Log

```bash
brew install hugo
hugo new site hugo
```

```bash
cd hugo
hugo new posts/my-first-post.md
```

```bash
# test
cd hugo
git clone https://github.com/theNewDynamic/gohugo-theme-ananke.git themes/ananke
echo theme = \"ananke\" >> config.toml
```

Now update `.gitignore` with the following, so git won't track generated files:

```
# Hugo
hugo/archetypes
hugo/data
hugo/layouts
hugo/public
hugo/resources
hugo/static
```

```bash
cd hugo

# run server
hugo server -D

# build static pages
hugo

# build static pages (include drafts)
hugo -D
```

## Hugo installation

- Instructions: https://gohugo.io/installation

```bash
brew install hugo
```