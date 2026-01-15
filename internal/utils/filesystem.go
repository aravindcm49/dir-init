package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// IsValidDirectoryName checks if a directory name is safe for filesystem use
func IsValidDirectoryName(name string) bool {
	if name == "" {
		return false
	}

	// Reserved names on Windows
	reservedNames := []string{
		"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	}

	// Check for reserved names (case insensitive)
	upperName := strings.ToUpper(name)
	for _, reserved := range reservedNames {
		if upperName == reserved {
			return false
		}
	}

	// Check for invalid characters
	invalidChars := regexp.MustCompile(`[<>:"/\\|?*]`)
	if invalidChars.MatchString(name) {
		return false
	}

	// Check for leading/trailing spaces or dots
	if strings.HasPrefix(name, " ") || strings.HasSuffix(name, " ") ||
		strings.HasPrefix(name, ".") || strings.HasSuffix(name, ".") {
		return false
	}

	// Check for control characters
	for _, r := range name {
		if unicode.IsControl(r) {
			return false
		}
	}

	// Check maximum length (255 characters)
	if len(name) > 255 {
		return false
	}

	return true
}

// SanitizeDirectoryName cleans a directory name to make it filesystem-safe
func SanitizeDirectoryName(name string) string {
	// Replace invalid characters with underscores
	invalidChars := regexp.MustCompile(`[<>:"/\\|?*]`)
	sanitized := invalidChars.ReplaceAllString(name, "_")

	// Remove leading/trailing spaces and dots
	sanitized = strings.TrimSpace(sanitized)
	sanitized = strings.TrimLeft(sanitized, ".")
	sanitized = strings.TrimRight(sanitized, ".")

	// Replace multiple consecutive underscores with single underscore
	multipleUnderscores := regexp.MustCompile(`_{2,}`)
	sanitized = multipleUnderscores.ReplaceAllString(sanitized, "_")

	// Remove control characters
	var builder strings.Builder
	for _, r := range sanitized {
		if !unicode.IsControl(r) {
			builder.WriteRune(r)
		}
	}
	sanitized = builder.String()

	// Limit length
	if len(sanitized) > 255 {
		sanitized = sanitized[:255]
	}

	// Ensure it's not empty
	if sanitized == "" {
		sanitized = "unnamed"
	}

	return sanitized
}

// CreateDirectory safely creates a directory with validation
func CreateDirectory(path string) error {
	if !IsValidDirectoryName(path) {
		return fmt.Errorf("invalid directory name: %s", path)
	}

	// Ensure parent directory exists
	parent := filepath.Dir(path)
	if parent != "." {
		err := os.MkdirAll(parent, 0755)
		if err != nil {
			return fmt.Errorf("failed to create parent directory: %v", err)
		}
	}

	// Create the directory
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	return nil
}

// DirectoryExists checks if a directory already exists
func DirectoryExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ValidateDirectoryCount checks if the requested number of directories is reasonable
func ValidateDirectoryCount(count int) error {
	if count < 1 {
		return fmt.Errorf("count must be at least 1")
	}
	if count > 20 {
		return fmt.Errorf("count cannot exceed 20")
	}
	return nil
}
