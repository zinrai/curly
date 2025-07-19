# curly

A curl wrapper that simplifies common HTTP operations with an command-line interface.

## Features

- Support for common HTTP methods (GET, POST)
- JSON and form data posting
- Response headers inspection
- Basic authentication support
- File upload/download capabilities
- Show curl command for learning
- Verbose mode for debugging
- **Pass-through curl options after `--`**

## Requirements

curl must be installed and available in PATH.

## Installation

```bash
$ go install github.com/zinrai/curly@latest
```

## Usage

### POST Requests

Send JSON data:

```bash
$ curly post https://api.example.com/users --json '{"name":"John","age":30}'
```

Send form data:

```bash
$ curly post https://api.example.com/form --data "name=John&age=30"
```

Send data from file:

```bash
$ curly post https://api.example.com/upload --file @data.json
```

### GET Requests

Simple GET:

```bash
$ curly get https://example.com
```

Save response to file:

```bash
$ curly get https://example.com/image.jpg --output image.jpg
```

### View Headers Only

```bash
$ curly headers https://example.com
```

Follow redirects:

```bash
$ curly headers https://bit.ly/example --follow
```

### Basic Authentication

```bash
$ curly basic-auth https://api.example.com --user username:password
```

### Global Options

- `--show-command`: Show the curl command being executed
- `--verbose`: Enable curl verbose output (shows request/response details)
- `--silent`: Silent mode (hide progress)
- `--header`: Add custom headers (can be used multiple times)
- `--user-agent`: Set User-Agent string
- `--follow`: Follow redirects

### Examples with Global Options

Show the curl command being executed:

```bash
$ curly get https://api.example.com --show-command
```

Enable curl verbose output:

```bash
$ curly get https://api.example.com --verbose
```

Add custom headers:

```bash
$ curly post https://api.example.com \
  --header "Authorization: Bearer token123" \
  --header "X-Custom-Header: value" \
  --json '{"data":"value"}'
```

Set User-Agent:

```bash
$ curly get https://example.com --user-agent "MyBot/1.0"
```

### Advanced Usage with Pass-through Options

You can pass any curl option after `--`:

Include response headers in output:

```bash
$ curly get https://example.com -- --include
```

Save response headers to file:

```bash
$ curly get https://example.com -- --dump-header headers.txt
```

Trace HTTP communication:

```bash
$ curly get https://example.com -- --trace trace.log
```

ASCII trace

```bash
$ curly get https://example.com -- --trace-ascii trace.txt
```

Combine multiple curl options:

```bash
$ curly get https://example.com -- --include --location --max-time 30 --compressed
```

Use with POST:

```bash
$ curly post https://api.example.com --json '{"test":true}' -- --include --trace-ascii trace.txt
```

Pass-through with basic auth:

```bash
$ curly basic-auth https://api.example.com --user admin:pass -- --include --dump-header auth-headers.txt
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
