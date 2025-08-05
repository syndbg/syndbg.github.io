package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	postsDir := "content/posts"

	files, err := os.ReadDir(postsDir)
	if err != nil {
		fmt.Printf("Error reading posts directory: %v\n", err)
		return
	}

	fmt.Printf("🔄 Renaming %d posts with date prefixes...\n\n", len(files))

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if !strings.HasSuffix(filename, ".md") {
			continue
		}

		filePath := filepath.Join(postsDir, filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("❌ Error reading file %s: %v\n", filename, err)
			continue
		}

		// Extract date from frontmatter
		date, err := extractDateFromPost(string(content))
		if err != nil {
			fmt.Printf("❌ Error extracting date from %s: %v\n", filename, err)
			continue
		}

		// Create new filename with date prefix
		newFilename := createDatePrefixedFilename(date, filename)
		newFilePath := filepath.Join(postsDir, newFilename)

		// Skip if filename is already date-prefixed
		if strings.HasPrefix(filename, date.Format("2006-01-02")) {
			fmt.Printf("⏭️  Skipped %s (already has date prefix)\n", filename)
			continue
		}

		// Rename the file
		err = os.Rename(filePath, newFilePath)
		if err != nil {
			fmt.Printf("❌ Error renaming %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Renamed: %s → %s\n", filename, newFilename)
	}

	fmt.Printf("\n✅ Post renaming complete!\n")
}

func extractDateFromPost(content string) (time.Time, error) {
	// Extract frontmatter
	re := regexp.MustCompile(`(?s)^---\n(.*?)\n---`)
	matches := re.FindStringSubmatch(content)

	if len(matches) < 2 {
		return time.Time{}, fmt.Errorf("could not find frontmatter")
	}

	frontMatter := matches[1]

	// Find date line
	lines := strings.Split(frontMatter, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "date:") {
			dateStr := strings.TrimSpace(strings.TrimPrefix(line, "date:"))

			// Try parsing multiple date formats
			layouts := []string{
				"2006-01-02T15:04:05Z",
				"2006-01-02T15:04:05-07:00",
				"2006-01-02 15:04:05",
				"2006-01-02",
			}

			for _, layout := range layouts {
				t, err := time.Parse(layout, dateStr)
				if err == nil {
					return t, nil
				}
			}

			return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
		}
	}

	return time.Time{}, fmt.Errorf("no date found in frontmatter")
}

func createDatePrefixedFilename(date time.Time, originalFilename string) string {
	// Remove .md extension
	baseName := strings.TrimSuffix(originalFilename, ".md")

	// Create date prefix in YYYY-MM-DD format
	datePrefix := date.Format("2006-01-02")

	// If the basename is just a year (like "2015"), make it more descriptive
	if regexp.MustCompile(`^\d{4}$`).MatchString(baseName) {
		baseName = baseName + "-retrospective"
	}

	// Create slug-friendly filename
	slug := strings.ToLower(baseName)
	slug = regexp.MustCompile(`[^a-z0-9-]+`).ReplaceAllString(slug, "-")
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	return fmt.Sprintf("%s-%s.md", datePrefix, slug)
}
