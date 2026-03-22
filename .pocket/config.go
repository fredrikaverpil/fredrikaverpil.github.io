package main

import (
	"github.com/fredrikaverpil/pocket/pk"
	"github.com/fredrikaverpil/pocket/tasks/github"
	"github.com/fredrikaverpil/pocket/tasks/golang"
)

var Config = &pk.Config{
	Auto: pk.Parallel(

		// TODO: enable
		// pk.WithOptions(
		// 	markdown.Tasks(),
		// 	pk.WithSkipPath("public"),
		// ),

		pk.WithOptions(
			golang.Tasks(),
			pk.WithDetect(golang.Detect()),
			pk.WithSkipPath(`^\.$`, "public"),
			pk.WithSkipTask(golang.Lint, "content/blog"),   // TODO: enable?
			pk.WithSkipTask(golang.Format, "content/blog"), // TODO: enable?
			pk.WithSkipTask(golang.Fix, "content/blog")),   // TODO: enable?

		pk.WithOptions(
			github.Tasks(),
			pk.WithFlags(github.WorkflowFlags{
				ReleasePleaseWorkflow: new(false),
				StaleWorkflow:         new(false),
				Platforms:             []github.Platform{github.Ubuntu},
				ExternalWorkflows:     []string{"gh-pages.yml"},
			}),
		),
	),
	Manual: []pk.Runnable{
		Serve,
		Build,
		Clean,
		CheckLinks,
	},
}
