package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fredrikaverpil/pocket/pk"
	"github.com/fredrikaverpil/pocket/pk/repopath"
	"github.com/fredrikaverpil/pocket/pk/run"
	"github.com/fredrikaverpil/pocket/tools/pagefind"
)

// Serve starts the Hugo development server with drafts enabled.
var Serve = &pk.Task{
	Name:  "serve",
	Usage: "start local dev server with drafts (LiveReload)",
	Do: func(ctx context.Context) error {
		return run.Exec(ctx, "go", "tool", "hugo", "server", "-D")
	},
}

// Build runs the production build: Hugo + legacy RSS copy + Pagefind indexing.
var Build = &pk.Task{
	Name:  "build",
	Usage: "production build (Hugo + Pagefind indexing)",
	Body: pk.Serial(
		pagefind.Install,
		pk.Do(func(ctx context.Context) error {
			if err := run.Exec(ctx, "go", "tool", "hugo", "--minify", "--environment", "production"); err != nil {
				return err
			}
			// Legacy mkdocs-created RSS feed.
			rssSrc := repopath.FromGitRoot("public", "blog", "index.xml")
			rssDst := repopath.FromGitRoot("public", "feed_rss_created.xml")
			src, err := os.ReadFile(rssSrc)
			if err != nil {
				return fmt.Errorf("read RSS feed: %w", err)
			}
			if err := os.WriteFile(rssDst, src, 0o644); err != nil {
				return fmt.Errorf("write legacy RSS feed: %w", err)
			}
			return run.Exec(ctx, repopath.FromBinDir(pagefind.Name), "--site", "public")
		}),
	),
}

// CheckLinks builds the site and checks for broken internal links.
var CheckLinks = &pk.Task{
	Name:  "check-links",
	Usage: "build site and check for broken links",
	Body: pk.Serial(
		InstallHTMLTest,
		Build,
		pk.Do(func(ctx context.Context) error {
			return run.Exec(ctx, repopath.FromBinDir(htmltestName), "-c", repopath.FromGitRoot(".htmltest.yml"))
		}),
	),
}

// Clean removes build artifacts.
var Clean = &pk.Task{
	Name:  "clean",
	Usage: "remove public/ and resources/",
	Do: func(ctx context.Context) error {
		for _, dir := range []string{"public", "resources"} {
			if err := os.RemoveAll(repopath.FromGitRoot(dir)); err != nil {
				return err
			}
		}
		return nil
	},
}
