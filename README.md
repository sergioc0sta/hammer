# Hammer
A powerful command-line tool written in Go for performing load tests on web services. This tool allows you to test the performance and reliability of any HTTP endpoint by sending concurrent requests and generating detailed reports.

## Features
- 🚀 **Concurrent Request Execution** - Configure the number of simultaneous requests
- 📊 **Detailed Reporting** - Get comprehensive statistics about your load tests
- 🐳 **Docker Support** - Run the tool in a containerized environment
- ⚡ **Fast and Lightweight** - Built with Go for optimal performance
- 📈 **HTTP Status Distribution** - Detailed breakdown of response status codes

## Installation

### Using Docker (Recommended)
```bash
docker pull hammer
```

### From Source

```bash
git clone https://github.com/sergioc0sta/hammer
cd hammer
go build -o hammer cmd/main.go
```
## Usage

### Docker

```bash
docker run --name hammer hammer --url=https://httpbin.org/get --requests=1000 --concurrency=10
```

### Binary

```bash
./hammer --url=http://google.com --requests=1000 --concurrency=10
```
## Command-Line Options

| Flag | Description | Required | Example |
|------|-------------|----------|---------|
| `--url` | URL of the service to test | Yes | `--url=http://example.com` |
| `--requests` | Total number of requests to make | Yes | `--requests=1000` |
| `--concurrency` | Number of concurrent requests | Yes | `--concurrency=10` |

## Example Output
```yaml
=== Hammer Report ===
Total time: 5.234s
Total requests: 1000
Successful requests (200): 950

Status code distribution:
200: 950 (95.0%)
404: 30 (3.0%)
500: 20 (2.0%)
```
## Examples

### Basic load test
Test a website with 100 requests, 10 at a time:

```bash
docker run --name hammer hammer --url=https://example.com --requests=100 --concurrency=10
```

### Heavy load test
Simulate 10,000 requests with high concurrency:

```bash
docker run --name hammer hammer --url=https://api.example.com/health --requests=10000 --concurrency=100
```
### Single request test
Test basic connectivity:

```bash
docker run --name hammer hammer --url=https://example.com --requests=1 --concurrency=1
```

## Building the Docker Image
```bash
docker build -t hammer .
```
