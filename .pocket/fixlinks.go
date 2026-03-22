package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fredrikaverpil/pocket/pk/repopath"
)

// linkPattern matches markdown links with ../ relative paths.
var linkPattern = regexp.MustCompile(`\]\(\.\./([^)]+)\)`)

// fixBrokenLinks scans content/blog/ markdown files for relative ../
// links and rewrites them to absolute /blog/... paths using the built
// public/ directory as the source of truth for valid URLs.
func fixBrokenLinks() error {
	urlMap, err := buildURLMap()
	if err != nil {
		return fmt.Errorf("build URL map: %w", err)
	}

	contentDir := repopath.FromGitRoot("content", "blog")
	entries, err := os.ReadDir(contentDir)
	if err != nil {
		return fmt.Errorf("read content dir: %w", err)
	}

	var totalFixed int
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		filePath := filepath.Join(contentDir, entry.Name())
		fixed, err := fixLinksInFile(filePath, urlMap)
		if err != nil {
			return fmt.Errorf("fix links in %s: %w", entry.Name(), err)
		}
		totalFixed += fixed
	}

	fmt.Printf("Fixed %d links across content/blog/\n", totalFixed)
	return nil
}

// fixLinksInFile rewrites relative ../ links in a single markdown file.
// Returns the number of links fixed.
func fixLinksInFile(filePath string, urlMap map[string]string) (int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("read file: %w", err)
	}

	content := string(data)
	var fixed int

	result := linkPattern.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the relative path from ](../some/path/).
		sub := linkPattern.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		relPath := strings.TrimSuffix(sub[1], "/")

		absURL, ok := resolveLink(relPath, urlMap)
		if !ok {
			return match
		}

		fixed++
		fileName := filepath.Base(filePath)
		fmt.Printf("  %s: ../%s/ -> %s\n", fileName, relPath, absURL)
		return "](" + absURL + ")"
	})

	if fixed == 0 {
		return 0, nil
	}

	if err := os.WriteFile(filePath, []byte(result), 0o644); err != nil {
		return 0, fmt.Errorf("write file: %w", err)
	}
	return fixed, nil
}

// resolveLink maps a relative link path to an absolute /blog/... URL.
func resolveLink(relPath string, urlMap map[string]string) (string, bool) {
	// Pattern: YYYY-MM-DD-slug (filename-style link).
	if url, ok := urlMap[relPath]; ok {
		return url, true
	}

	// Pattern: YYYY/MM/DD/slug (old URL-style link).
	// Convert to YYYY-MM-DD prefix and search by date + slug overlap.
	parts := strings.SplitN(relPath, "/", 4)
	if len(parts) == 4 {
		datePrefix := parts[0] + "-" + parts[1] + "-" + parts[2] + "-"
		slug := parts[3]
		for stem, url := range urlMap {
			if !strings.HasPrefix(stem, datePrefix) {
				continue
			}
			// Match if the link slug is a substring of the filename slug or vice versa.
			fileSuffix := strings.TrimPrefix(stem, datePrefix)
			if strings.Contains(fileSuffix, slug) || strings.Contains(slug, fileSuffix) {
				return url, true
			}
		}
	}

	return "", false
}

// buildURLMap creates a mapping from content filename stems to their
// generated absolute URLs by cross-referencing content/blog/ files
// with public/blog/ output.
func buildURLMap() (map[string]string, error) {
	publicDir := repopath.FromGitRoot("public", "blog")
	contentDir := repopath.FromGitRoot("content", "blog")

	// Collect all content file stems grouped by date.
	// e.g. "2016-07-25" -> ["2016-07-25-developing-with-qt-py", "2016-07-25-dealing-with-maya-2017-and-pyside2"]
	contentEntries, err := os.ReadDir(contentDir)
	if err != nil {
		return nil, fmt.Errorf("read content dir: %w", err)
	}
	stemsByDate := make(map[string][]string)
	for _, entry := range contentEntries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		stem := strings.TrimSuffix(entry.Name(), ".md")
		if len(stem) < 10 {
			continue
		}
		date := stem[:10] // YYYY-MM-DD
		stemsByDate[date] = append(stemsByDate[date], stem)
	}

	// Walk public/blog/YYYY/MM/DD/slug/ directories and match to content stems.
	urlMap := make(map[string]string)
	err = filepath.WalkDir(publicDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() != "index.html" {
			return nil
		}

		// Extract URL from path: public/blog/YYYY/MM/DD/slug/index.html -> /blog/YYYY/MM/DD/slug/
		rel, err := filepath.Rel(repopath.FromGitRoot(), path)
		if err != nil {
			return nil
		}
		dir := filepath.Dir(rel) // public/blog/YYYY/MM/DD/slug
		url := "/" + strings.TrimPrefix(dir, "public/") + "/"

		// Extract date from URL path.
		urlParts := strings.Split(strings.Trim(url, "/"), "/")
		if len(urlParts) < 5 {
			return nil
		}
		date := urlParts[1] + "-" + urlParts[2] + "-" + urlParts[3] // YYYY-MM-DD

		stems, ok := stemsByDate[date]
		if !ok {
			return nil
		}

		// URL slug is everything after /blog/YYYY/MM/DD/.
		urlSlug := strings.Join(urlParts[4:], "/")

		// Find the best matching content stem for this URL.
		bestStem := matchStem(stems, date, urlSlug)
		if bestStem != "" {
			urlMap[bestStem] = url
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk public dir: %w", err)
	}

	return urlMap, nil
}

// matchStem finds the content file stem that best matches a URL slug.
func matchStem(stems []string, date, urlSlug string) string {
	if len(stems) == 1 {
		return stems[0]
	}

	// Normalize the URL slug for comparison (replace / with -).
	normalizedURL := strings.ReplaceAll(urlSlug, "/", "-")

	var bestStem string
	var bestScore int
	for _, stem := range stems {
		fileSuffix := strings.TrimPrefix(stem, date+"-")
		// Score by length of common prefix between filename slug and URL slug.
		score := commonPrefixLen(fileSuffix, normalizedURL)
		if score > bestScore {
			bestScore = score
			bestStem = stem
		}
	}
	return bestStem
}

// commonPrefixLen returns the length of the common prefix between two strings.
func commonPrefixLen(a, b string) int {
	limit := len(a)
	if len(b) < limit {
		limit = len(b)
	}
	for i := range limit {
		if a[i] != b[i] {
			return i
		}
	}
	return limit
}
