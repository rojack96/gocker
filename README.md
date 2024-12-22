# Gocker

**Gocker** is a CLI generator to create and execute Docker commands in Go.

## Installation

To install Gocker, use `go get`:

```bash
go get github.com/rojack96/gocker
```

### Basic Example

Here is an example of how to use Gocker to execute a Docker Compose command:

```go
package main

import (
    "fmt"
    "github.com/rojack96/gocker"
)

func main() {
    compose := gocker.Compose()

    result := compose.FileName("compose.yml").Up().Build().ServiceNames()

    fmt.Println(result.GetCommand())
    // output: "docker compose --file compose.yml up --build hello-world"

    result.Exec(false)
}
```

### Supported Commands

Gocker supports Docker commands for version 26.1.3.

Currently, all `compose` commands are implemented, and support for other Docker commands is coming soon.

### Create a Command

```go
// Create compose instance
compose := gocker.Compose()

result := compose.FileName("compose.yml"). //(optional) Docker file
                    Up(). // up command
                    Build(). // --build command
                    ServiceNames() // service name to start, if empty, all services will be started

fmt.Println(result.GetCommand())
// output: "docker compose --file compose.yml up --build hello-world"
```

### Execute a Command

```go
// Create compose instance
compose := gocker.Compose()

result := compose.FileName("compose.yml"). // Docker file, if empty it will not be included
                    Up(). // up command
                    Build(). // --build command
                    ServiceNames() // service name to start, if empty, all services will be started
result.Exec(false) // if true, the command will be executed with administrator privileges
```

## Contributing

If you want to contribute to Gocker, feel free to do so! You can create a pull request or open an issue to discuss changes.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.