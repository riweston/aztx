## README

This tool is a helper for azure-cli that leverages fzf for a nice interface to switch between subscription contexts.

Additionally, it can also be used to quickly switch back to a previous subscription context using the `aztx -` command in a similar way to `cd -` in bash.

### Demo

[![asciicast](https://asciinema.org/a/Rk36acdIGN9K6w5WO5Rx74NwA.svg)](https://asciinema.org/a/Rk36acdIGN9K6w5WO5Rx74NwA)

### Prerequisites

> [!NOTE]
> This tool is built on top of the azure-cli and fzf and requires them to be installed and configured.
> If you use the Brew or Scoop package managers, these pre-requisites should be handled during the installation.

- go >=1.16.6
- azure-cli >= 2.22.1
- fzf >= 0.20.0

### Installation Options

#### [Brew](https://brew.sh/) (Mac/Linux)

```sh
$ brew tap riweston/aztx
$ brew install aztx
```

#### [Scoop](https://scoop.sh/) (Windows)

```sh
$ scoop bucket add riweston https://github.com/riweston/scoop-bucket.git
$ scoop update
$ scoop install riweston/aztx
```

#### Download a prebuilt binary (Linux/Mac/Windows)

Download the latest release from the [releases page](https://github.com/riweston/aztx/releases) and add it to your PATH.

#### Install from source (Linux/Mac/Windows)

```sh
$ go install github.com/riweston/aztx
```

## Usage

```sh
# Run the tool in interactive mode
$ aztx
# Switch back to the previous subscription context
$ aztx -
Switched to "Test Subscription 1" (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
# Switch back again to the previous subscription context
$ aztx -
Switched to "Test Subscription 2" (yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy)
```

### Show your support

Give a ⭐️ if this project helped you!
