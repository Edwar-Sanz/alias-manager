# am — Alias Manager

A CLI to add, remove, and list your shell aliases from the terminal.

## Requirements

- bash or zsh
- Go 1.25+ (to build from source)

## Installation

### Default installation
```bash
# step 1: install am
curl -fsSL https://raw.githubusercontent.com/Edwar-Sanz/alias-manager/main/install.sh | bash
# step 2: reload shell
source ~/.zshrc # or ~/.bashrc depending on your shell
```

### Or if you want to build from source:

```bash
git clone https://github.com/Edwar-Sanz/alias-manager
cd alias-manager
go build -o am .
```


## Usage

### Commands
- `a | add` for adding
- `r | rm | remove` for removing
- `l | ls | list` for listing

```bash
# add alias
am a <alias> <command>
# add with category
am a <alias> <command> <category>
# add with category and description
am a <alias> <command> <category> <description>
# remove alias
am r <alias>
# list all aliases
am l
# list by category
am l -c <category>
# list category names only
am l -C
```


## Alias file

Aliases are stored at `~/.config/am/.amf` in standard shell format, readable and editable by hand.


