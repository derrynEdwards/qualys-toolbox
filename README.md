# Qualys-ToolBox

Qualys-ToolBox is a command-line tool for interacting with Qualys services 
via their API. This tool provides a convenient way to perform various security-related
tasks, such as vulnerability scanning and compliance assessments, directly from
the command line.

## Prerequisites

Before you can use Qualys-ToolBox, you'll need to have the following prerequisites:  
- Qualys API credentials (API username and password or API token)

## Installation

You have two options for installation:

### Option 1: Binary Release

You can download the latest binary release for your platform from the
[Releases](https://github.com/derrynEdwards/qualys-toolbox/releases) page. Once downloaded, 
make it executable and move it to a directory in your system's PATH.

```shell
# Example for Linux
chmod +x qualys-toolbox
sudo mv qualys-toolbox /usr/local/bin/qualys-toolbox
```

### Option 2: Build from Source

To build qualys-toolbox from soruce, you can clone this repository and build it manually:  

```shell
git clone https://github.com/derrynEdwards/qualys-toolbox.git
cd qualys-toolbox
go build
```

## Usage

```shell
./qualys-toolbox

# or

qualys-toolbox
```

## Configuration

## Contributing

Contributions are welcome! If you'd like to contribute to this project, pelase follow the guidelines
in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](https://chat.openai.com/c/LICENSE).