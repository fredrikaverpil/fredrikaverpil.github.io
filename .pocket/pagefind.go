package main

import (
	"fmt"
	"path/filepath"

	"github.com/fredrikaverpil/pocket/pk"
	"github.com/fredrikaverpil/pocket/pk/download"
	"github.com/fredrikaverpil/pocket/pk/platform"
	"github.com/fredrikaverpil/pocket/pk/repopath"
)

const (
	pagefindName = "pagefind"
	// renovate: datasource=github-releases depName=CloudCannon/pagefind
	pagefindVersion = "1.4.0"
)

// pagefindInstall ensures pagefind is available.
var pagefindInstall = &pk.Task{
	Name:   "install:pagefind",
	Usage:  "install pagefind",
	Body:   installPagefind(),
	Hidden: true,
	Global: true,
}

func installPagefind() pk.Runnable {
	binDir := repopath.FromToolsDir(pagefindName, pagefindVersion, "bin")
	binaryName := platform.BinaryName(pagefindName)
	binaryPath := filepath.Join(binDir, binaryName)

	url := fmt.Sprintf(
		"https://github.com/CloudCannon/pagefind/releases/download/v%s/pagefind-v%s-%s.tar.gz",
		pagefindVersion, pagefindVersion, pagefindPlatformTarget(),
	)

	return download.Download(url,
		download.WithDestDir(binDir),
		download.WithFormat("tar.gz"),
		download.WithExtract(download.WithExtractFile(binaryName)),
		download.WithSymlink(),
		download.WithSkipIfExists(binaryPath),
	)
}

func pagefindPlatformTarget() string {
	hostOS := platform.HostOS()
	hostArch := platform.HostArch()

	archName := platform.ArchToX8664(hostArch)

	switch hostOS {
	case platform.Darwin:
		return archName + "-apple-darwin"
	case platform.Linux:
		return archName + "-unknown-linux-musl"
	case platform.Windows:
		return archName + "-pc-windows-msvc"
	default:
		return fmt.Sprintf("%s-%s", archName, hostOS)
	}
}
