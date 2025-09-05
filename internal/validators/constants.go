package validators

import "errors"

// Error messages for validation
var (
	// Repository validation errors
	ErrInvalidRepositoryURL = errors.New("invalid repository URL")

	// Package validation errors
	ErrPackageNameHasSpaces = errors.New("package name cannot contain spaces")

	// Remote validation errors
	ErrInvalidRemoteURL = errors.New("invalid remote URL")

	// Registry validation errors
	ErrUnsupportedRegistryBaseURL   = errors.New("unsupported registry base URL")
	ErrMismatchedRegistryTypeAndURL = errors.New("registry type and base URL do not match")

	// Argument validation errors
	ErrNamedArgumentNameRequired     = errors.New("named argument name is required")
	ErrInvalidNamedArgumentName      = errors.New("invalid named argument name format")
	ErrArgumentValueStartsWithName   = errors.New("argument value cannot start with the argument name")
	ErrArgumentDefaultStartsWithName = errors.New("argument default cannot start with the argument name")
)

// RepositorySource represents valid repository sources
type RepositorySource string

const (
	SourceGitHub RepositorySource = "github"
	SourceGitLab RepositorySource = "gitlab"
)
