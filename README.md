# console_go - Simple and Colorful Logger for Go

`console_go` is a lightweight logging package for Go that supports colored logs and multiple log levels. It eliminates the need for explicit initialization and allows for easy configuration via a JSON file.

## Features

- Color-coded log levels
- Configurable log output
- Toggle logging levels on and off
- Retrieve file and line number of the caller

## Installation

To install the package, use the following command:

```shell
go get github.com/navidpirsajed/console_go
```

## Usage

### Basic Usage

Simply import the package and start logging:

```go
package main

import (
	"github.com/navidpirsajed/console_go"
)

func main() {
	console_go.Log("This is a log message")
	console_go.Info("This is an info message")
	console_go.Debug("This is a debug message")
	console_go.Error("This is an error message")
	console_go.Warn("This is a warning message")
}
```

### Configuration

You can load configuration from a JSON file to set the logging preferences:

```go
package main

import (
	"github.com/navidpirsajed/console_go"
	"log"
)

func main() {
	if err := console_go.LoadConfig("config.json"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	console_go.Log("This is a log message")
	console_go.Info("This is an info message")
	console_go.Debug("This is a debug message")
	console_go.Error("This is an error message")
	console_go.Warn("This is a warning message")
}
```

#### JSON Configuration File

Here is an example of a JSON configuration file (`config.json`):

```json
{
  "enable_log": true,
  "enable_info": true,
  "enable_debug": true,
  "enable_error": true,
  "enable_warn": true,
  "output": "/path/to/log/file.log"
}
```

In the configuration file:

- `enable_log`, `enable_info`, `enable_debug`, `enable_error`, `enable_warn` are boolean flags to enable or disable respective log levels.
- `output` is the file path where the logs will be written. If omitted, logs will be printed to the standard output.

### Custom Output

By default, logs are written to `os.Stdout`. You can change the output by calling `LoadConfig` and setting the `output` field.
