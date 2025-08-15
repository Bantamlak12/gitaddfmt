# gitaddfmt

A Go package that simplifies your development workflow by automatically formatting Go code before adding files to git.

## What it does

This package provides a `go-git` binary that runs `go fmt ./...` before executing git commands. When you use `go git add <files>`, it will:

1. Format all Go code in your project using `go fmt ./...`
2. Add the specified files to git staging area
3. Provide clear feedback about the formatting and staging process

## Installation

### Quick Install

Run the installer to set up `go-git` and configure your shell:

```bash
go run ./install
```

This will:
- Install the `go-git` binary to your `$GOPATH/bin`
- Add a `go` function to your shell configuration (`.bashrc`, `.zshrc`, or PowerShell profile)
- Allow you to use `go git` syntax instead of `go-git`

### Manual Installation

If you prefer to install manually:

```bash
# Install the go-git binary
go install ./go-git

# Add the go function to your shell config (see shell configuration below)
```

## Usage

After installation, you can use the `go git` command just like regular git commands:

```bash
# Format code and add specific files
go git add main.go utils.go

# Format code and add all files
go git add .

# Other git commands work normally
go git commit -m "feat: add new feature"
go git status
go git log
```

## Shell Configuration

The installer automatically adds a `go` function to your shell configuration:

### Bash/Zsh
```bash
go() {
    if [ "$1" = "git" ]; then
        shift
        go-git "$@"
    else
        command go "$@"
    fi
}
```

### PowerShell
```powershell
function go {
    param($sub, $rest)
    if ($sub -eq 'git') {go-git @rest}
    else { & go.exe $sub @rest }
}
```

This allows you to use `go git` syntax while preserving all other `go` command functionality.

## How it works

The `go-git` binary intercepts git commands and:

1. **For `git add` commands**: Runs `go fmt ./...` first to format all Go code in the project
2. **For other git commands**: Passes them through to the regular git command
3. **Provides feedback**: Shows status messages about formatting and staging

## Requirements

- Go 1.24.5 or later
- Git installed and configured
- Bash, Zsh, or PowerShell (for shell integration)

## Project Structure

```
.
├── go-git/          # Main go-git binary
│   └── main.go
├── install/         # Installer script
│   └── main.go
├── go.mod           # Go module definition
└── README.md        # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test the installation and usage
5. Submit a pull request

## License

This project is open source. Feel free to use and modify as needed.

## Troubleshooting

### Command not found: go-git
Make sure the binary is installed and `$GOPATH/bin` is in your PATH:
```bash
go install ./go-git
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

### go git command not working
Restart your terminal or source your shell configuration:
```bash
source ~/.bashrc  # or ~/.zshrc
```

### Formatting fails
Ensure you're in a Go project directory with valid Go code and that `go fmt ./...` works manually. 