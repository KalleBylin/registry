package validators

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	// Regular expressions for validating repository URLs
	// These regex patterns ensure the URL is in the format of a valid GitHub or GitLab repository
	// For example:	// - GitHub: https://github.com/user/repo
	githubURLRegex = regexp.MustCompile(`^https?://(www\.)?github\.com/[\w.-]+/[\w.-]+/?$`)
	gitlabURLRegex = regexp.MustCompile(`^https?://(www\.)?gitlab\.com/[\w.-]+/[\w.-]+/?$`)
)

// IsValidRepositoryURL checks if the given URL is valid for the specified repository source
func IsValidRepositoryURL(source RepositorySource, url string) bool {
	switch source {
	case SourceGitHub:
		return githubURLRegex.MatchString(url)
	case SourceGitLab:
		return gitlabURLRegex.MatchString(url)
	}
	return false
}

// HasNoSpaces checks if a string contains no spaces
func HasNoSpaces(s string) bool {
	return !strings.Contains(s, " ")
}

// IsValidURL checks if a URL is in valid format
func IsValidURL(rawURL string) bool {
	// Parse the URL
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	// Check if scheme is present (http or https)
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" || u.Hostname() == "localhost" {
		return false
	}
	return true
}

// IsValidSubfolderPath checks if a subfolder path is valid
func IsValidSubfolderPath(path string) bool {
	// Empty path is valid (subfolder is optional)
	if path == "" {
		return true
	}

	// Must not start with / (must be relative)
	if strings.HasPrefix(path, "/") {
		return false
	}

	// Must not end with / (clean path format)
	if strings.HasSuffix(path, "/") {
		return false
	}

	// Check for valid path characters (alphanumeric, dash, underscore, dot, forward slash)
	validPathRegex := regexp.MustCompile(`^[a-zA-Z0-9\-_./]+$`)
	if !validPathRegex.MatchString(path) {
		return false
	}

	// Check that path segments are valid
	segments := strings.Split(path, "/")
	for _, segment := range segments {
		// Disallow empty segments ("//"), current dir ("."), and parent dir ("..")
		if segment == "" || segment == "." || segment == ".." {
			return false
		}
	}

	return true
}
