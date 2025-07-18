package curl

import (
	"fmt"
	"strings"
)

// Builder constructs curl command arguments
type Builder struct {
	method  string
	url     string
	args    []string
	headers []string
}

// NewBuilder creates a new curl command builder
func NewBuilder(method, url string) *Builder {
	b := &Builder{
		method:  method,
		url:     url,
		args:    []string{},
		headers: []string{},
	}

	// Set method if not GET or HEAD
	if method != "GET" && method != "HEAD" {
		b.args = append(b.args, "--request", method)
	}

	return b
}

// AddFlag adds a simple flag to curl command
func (b *Builder) AddFlag(flag string) *Builder {
	b.args = append(b.args, flag)
	return b
}

// AddHeader adds a header
func (b *Builder) AddHeader(key, value string) *Builder {
	b.headers = append(b.headers, fmt.Sprintf("%s: %s", key, value))
	return b
}

// AddRawHeader adds a raw header string
func (b *Builder) AddRawHeader(header string) *Builder {
	b.headers = append(b.headers, header)
	return b
}

// AddData adds data for POST/PUT requests
func (b *Builder) AddData(data string) *Builder {
	b.args = append(b.args, "--data", data)
	return b
}

// AddDataFile adds data from file
func (b *Builder) AddDataFile(filename string) *Builder {
	b.args = append(b.args, "--data", filename)
	return b
}

// AddOutput sets output file
func (b *Builder) AddOutput(filename string) *Builder {
	b.args = append(b.args, "--output", filename)
	return b
}

// AddBasicAuth adds basic authentication
func (b *Builder) AddBasicAuth(userPass string) *Builder {
	b.args = append(b.args, "--user", userPass)
	return b
}

// Build returns the complete curl command arguments
func (b *Builder) Build() []string {
	args := []string{}

	// Add method and other flags first
	args = append(args, b.args...)

	// Add headers
	for _, h := range b.headers {
		args = append(args, "--header", h)
	}

	// Add URL last
	args = append(args, b.url)

	return args
}

// BuildCommand returns the complete curl command as a string for display
func (b *Builder) BuildCommand() string {
	args := b.Build()
	quotedArgs := make([]string, len(args))

	for i, arg := range args {
		// Quote arguments that contain spaces or special characters
		if strings.ContainsAny(arg, " \t\n\"'") {
			quotedArgs[i] = fmt.Sprintf("'%s'", strings.ReplaceAll(arg, "'", "'\"'\"'"))
		} else {
			quotedArgs[i] = arg
		}
	}

	return "curl " + strings.Join(quotedArgs, " ")
}
