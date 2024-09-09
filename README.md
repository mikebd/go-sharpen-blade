# go-sharpen-blade

Sharpening my Golang skills with coding challenges and skunkworks.

Some packages are only executed as unit tests, others can be run as commands using the
`-c <command>` argument.  The `-l` argument will list the available commands.

## Work in Progress

* This repository is currently transitioning from a single `main.go` to multiple `cmd` packages

## Vanilla Go

* This repository aims to minimize third-party library dependencies, allowing the author the most opportunity to
  idiomatically leverage the Standard Library.
* Reusable code is extracted to [github.com/mikebd/go-util](https://github.com/mikebd/go-util)

## Running in Docker

* `docker build -t go-sharpen-blade .` - Build the Docker image
* `docker run go-sharpen-blade` - Run without arguments<br/>
  OR <br/>
  `docker run go-sharpen-blade /app/go-sharpen-blade <args>` - Run with the given arguments

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Arguments

| Argument           | Description               |
|--------------------|---------------------------|
| `-c <command>`     | Command to run            |
| `-d <directory>`   | Working directory         |
| `-h` \| `--help`   | Print the help message    |
| `-l`               | List commands             |
| `-lt`              | Log with timestamps (UTC) |
| `-o11y-prometheus` | Enable Prometheus metrics |

## Prometheus Metrics

When run with `-o11y-prometheus -c <command>`, prometheus metrics are exposed on `:2112/metrics`.
The program will wait for 20 seconds before exiting, so that the metrics can be scraped.

`prometheus.yml`:
```yaml
scrape_configs:
  - job_name: sharpen_blade
    scrape_interval: 10s
    static_configs:
      - targets:
        - localhost:2112
```