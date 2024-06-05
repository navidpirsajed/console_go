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
	console "github.com/navidpirsajed/console_go"
)

func main() {
	console.Log("This is a log message")
	console.Info("%d/%d/%d", 1,3,2024)
	console.Debug("This is a debug message")
	console.Error("This is an error message")
	console.Warn("This is a warning message")
}
```

### Configuration

You can load configuration from a JSON file to set the logging preferences. Here is an example of a `config.json`:

```json
{
  "enable_log": true,
  "enable_info": true,
  "enable_debug": true,
  "enable_error": true,
  "enable_warn": true
}
```

In the configuration file:

- `enable_log`, `enable_info`, `enable_debug`, `enable_error`, `enable_warn` are boolean flags to enable or disable respective log levels.
