// This script has been vibe coded using Claude Haiku 4.5
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
)

func loadEnvFile(envPath string) (string, error) {
	// Load .env file using godotenv
	if err := godotenv.Load(envPath); err != nil {
		return "", err
	}

	// Get OBSIDIAN_PATH from environment variables
	obsidianPath := os.Getenv("OBSIDIAN_PATH")
	if obsidianPath == "" {
		return "", fmt.Errorf("OBSIDIAN_PATH not found in .env file")
	}

	return obsidianPath, nil
}

func readFileAndExtractDate(filePath string) (string, string, []byte, error) {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", nil, err
	}

	// Compute hash of file content
	hash := fmt.Sprintf("%x", sha256.Sum256(content))

	// Parse frontmatter from content
	var matter map[string]interface{}
	_, err = frontmatter.Parse(strings.NewReader(string(content)), &matter)
	if err != nil {
		return hash, "NOT FOUND", content, nil
	}

	// Extract date from frontmatter
	date := "NOT FOUND"
	if dateValue, exists := matter["date"]; exists {
		switch v := dateValue.(type) {
		case time.Time:
			date = v.Format("2006-01-02T15:04:05")
		case string:
			date = v
		default:
			date = fmt.Sprintf("%v", v)
		}
	}

	return hash, date, content, nil
}

func extractYearFromDate(date string) string {
	// Extract year from date string (e.g., "2026-01-11T12:00:00" -> "2026")
	if date == "NOT FOUND" {
		return ""
	}
	year, _, found := strings.Cut(date, "-")
	if !found {
		return ""
	}
	return year
}

func filesAreIdentical(srcHash string, dstPath string) (bool, error) {
	// Read destination file
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		return false, err
	}

	// Compute hash of destination file
	dstHash := fmt.Sprintf("%x", sha256.Sum256(dstContent))

	// Compare hashes
	return srcHash == dstHash, nil
}

func copyFile(dst string, content []byte, srcHash string) (string, error) {
	// Check if destination file already exists
	if _, err := os.Stat(dst); err == nil {
		// File exists, check if they're identical using hash
		identical, err := filesAreIdentical(srcHash, dst)
		if err != nil {
			return "ERROR", err
		}
		if identical {
			return "SKIPPED (identical)", nil
		}
		// Files are different, will proceed to copy (overwrite)
	}

	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return "ERROR", err
	}

	// Write content to destination file
	if err := os.WriteFile(dst, content, 0644); err != nil {
		return "ERROR", err
	}

	return "SUCCESS", nil
}

type FileRecord struct {
	RelativePath    string
	Date            string
	DestinationPath string
	Status          string
}

func walkBlogFolder(blogPath string, contentBlogPath string, obsidianPath string) ([]FileRecord, error) {
	var records []FileRecord

	err := filepath.Walk(blogPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process markdown files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			// Read file once and extract both date and hash
			fileHash, date, content, err := readFileAndExtractDate(path)
			if err != nil {
				return err
			}

			year := extractYearFromDate(date)

			// Calculate relative path from obsidian Blog folder
			relPath, err := filepath.Rel(obsidianPath, path)
			if err != nil {
				return err
			}

			record := FileRecord{
				RelativePath: relPath,
				Date:         date,
			}

			// Only copy if date was found
			if year != "" {
				destPath := filepath.Join(contentBlogPath, year, info.Name())
				status, err := copyFile(destPath, content, fileHash)
				if err != nil {
					record.Status = "ERROR: " + err.Error()
				} else if status == "SUCCESS" {
					record.DestinationPath = destPath
					record.Status = status
				} else {
					record.Status = status
				}
			} else {
				record.Status = "SKIPPED (no date)"
			}

			records = append(records, record)
		}

		return nil
	})

	return records, err
}

func main() {
	// Define command-line flags
	envPath := flag.String("env", ".env", "Path to the .env file")
	flag.Parse()

	// Load environment variables from .env file
	obsidianPath, err := loadEnvFile(*envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Navigate to Blog folder
	blogPath := filepath.Join(obsidianPath, "Blog")

	// Check if Blog folder exists
	info, err := os.Stat(blogPath)
	if err != nil {
		log.Fatalf("Error accessing Blog folder: %v", err)
	}
	if !info.IsDir() {
		log.Fatal("Blog path is not a directory")
	}

	// Get the content/blog path relative to the script location
	// Script is in <workspace>/scripts, so go up one level to get to workspace
	scriptDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	contentBlogPath := filepath.Join(scriptDir, "..", "content", "blog")

	fmt.Println("=========================================================")
	fmt.Printf("Blog folder: %s\n", blogPath)
	fmt.Printf("Destination: %s\n", contentBlogPath)
	fmt.Println()

	// Walk through Blog folder and process markdown files
	records, err := walkBlogFolder(blogPath, contentBlogPath, obsidianPath)
	if err != nil {
		log.Fatalf("Error walking Blog folder: %v", err)
	}

	// Print summary table
	t := table.NewWriter()
	t.AppendHeader(table.Row{"File Path", "Date", "Destination Path", "Status"})

	successCount := 0
	skippedCount := 0
	for _, record := range records {
		t.AppendRow(table.Row{record.RelativePath, record.Date, record.DestinationPath, record.Status})

		if record.Status == "SUCCESS" {
			successCount++
		} else if record.Status == "SKIPPED (identical)" {
			skippedCount++
		}
	}

	t.SetStyle(table.StyleDefault)
	fmt.Println(t.Render())

	fmt.Printf("Total files processed: %d, Successfully copied: %d, Skipped (identical): %d\n", len(records), successCount, skippedCount)
}
