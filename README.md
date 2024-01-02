# go-sharpen-blade

Sharpening my Golang skills with coding challenges and skunkworks.

Some packages are only executed as unit tests, others can be run as commands using the
`-c <command>` argument.  The `-l` argument will list the available commands.

## Running in Docker

* `docker build -t go-sharpen-blade .` - Build the Docker image
* `docker run go-sharpen-blade` - Run without arguments<br/>
  OR <br/>
  `docker run go-sharpen-blade /app/go-sharpen-blade <args>` - Run with the given arguments

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Arguments

| Argument         | Description               |
|------------------|---------------------------|
| `-c <command>`   | Command to run            |
| `-d <directory>` | Working directory         |
| `-h` \| `--help` | Print the help message    |
| `-l`             | List commands             |
| `-lt`            | Log with timestamps (UTC) |