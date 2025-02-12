# aztx - Azure Tenant Context Switcher

`aztx` is a command-line tool designed to streamline the management of Azure tenant and subscription contexts. It provides an intuitive fuzzy-finder interface for switching between Azure subscriptions and tenants, making it easier to work with multiple Azure environments.

## Features

- 🔍 Fuzzy search interface for finding subscriptions and tenants
- ⚡ Quick context switching between subscriptions
- 🔄 Easy switching to previous context (similar to `cd -`)
- 🎯 Tenant-first selection mode
- 🔧 Configurable logging levels

### Demo

[![asciicast](https://asciinema.org/a/Rk36acdIGN9K6w5WO5Rx74NwA.svg)](https://asciinema.org/a/Rk36acdIGN9K6w5WO5Rx74NwA)

## Prerequisites

> [!NOTE]
> This tool is built on top of the azure-cli and fzf and requires them to be installed and configured.
> If you use the Brew or Scoop package managers, these pre-requisites will be handled during installation.

- go >=1.16.6
- azure-cli >= 2.22.1
- fzf >= 0.20.0

## Installation

### [Brew](https://brew.sh/) (Mac/Linux)

```sh
brew tap riweston/aztx
brew install aztx
```

### [Scoop](https://scoop.sh/) (Windows)

```sh
scoop bucket add riweston https://github.com/riweston/scoop-bucket.git
scoop update
scoop install riweston/aztx
```

### Download Prebuilt Binary

Download the latest release from the [releases page](https://github.com/riweston/aztx/releases) and add it to your PATH.

### Install from Source

```sh
go install github.com/riweston/aztx
```

## Usage

### Basic Subscription Switching

```sh
# Launch interactive subscription selector
aztx

# Switch to previous subscription context
aztx -
```

### Tenant-First Selection

```sh
# Select tenant before choosing subscription
aztx --by-tenant
```

## Configuration

Configuration is stored in `~/.aztx.yml`. The following options are available:

```yaml
# Log level: debug, info, warn, error
log-level: info
```

You can also set configuration via environment variables:
- `AZTX_LOG_LEVEL`: Set logging level
- `AZTX_BY_TENANT`: Enable tenant-first selection mode

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Show your support

Give a ⭐️ if this project helped you!
