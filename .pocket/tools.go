package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/fredrikaverpil/pocket/pk"
	"github.com/fredrikaverpil/pocket/pk/download"
	"github.com/fredrikaverpil/pocket/pk/platform"
	"github.com/fredrikaverpil/pocket/pk/repopath"
)

const (
	htmltestName = "htmltest"
	// renovate: datasource=github-releases depName=wjdp/htmltest
	htmltestVersion = "0.17.0"
)

// InstallHTMLTest ensures htmltest is available.
var InstallHTMLTest = &pk.Task{
	Name:   "install:htmltest",
	Usage:  "install htmltest",
	Body:   installHTMLTest(),
	Hidden: true,
	Global: true,
}

func installHTMLTest() pk.Runnable {
	binDir := repopath.FromToolsDir(htmltestName, htmltestVersion, "bin")
	binaryName := platform.BinaryName(htmltestName)
	binaryPath := filepath.Join(binDir, binaryName)

	url := fmt.Sprintf(
		"https://github.com/wjdp/htmltest/releases/download/v%s/htmltest_%s_%s_%s.%s",
		htmltestVersion, htmltestVersion, htmltestOS(), platform.HostArch(), htmltestArchiveFormat(),
	)

	return download.Download(url,
		download.WithDestDir(binDir),
		download.WithFormat(htmltestArchiveFormat()),
		download.WithExtract(download.WithExtractFile(binaryName)),
		download.WithSymlink(),
		download.WithSkipIfExists(binaryPath),
	)
}

func htmltestOS() string {
	switch runtime.GOOS {
	case "darwin":
		return "macos"
	default:
		return runtime.GOOS
	}
}

func htmltestArchiveFormat() string {
	if runtime.GOOS == "windows" {
		return "zip"
	}
	return "tar.gz"
}
