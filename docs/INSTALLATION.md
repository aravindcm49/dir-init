# Installation

## Prerequisites
- Go 1.19 or later installed on your system

## GOPATH Setup (if not already configured)

Go automatically sets up GOPATH, but you should verify it's in your PATH:

```bash
# Check if GOPATH is set
go env GOPATH

# Check if GOPATH/bin is in your PATH
echo $PATH | grep -o "$(go env GOPATH)/bin"

# If GOPATH/bin is not in PATH, add it to your shell profile:
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
# or for bash:
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bash_profile

# Apply the changes
source ~/.zshrc  # or source ~/.bash_profile
```

## Installation Methods

### Method 1: Install from GitHub (Recommended)
```bash
go install github.com/aravindcm49/dir-init@latest
```

### Method 2: Build from Source
```bash
git clone https://github.com/aravindcm49/dir-init.git
cd dir-init
go build -o dir-init main.go
./dir-init --help
```

### Method 3: Build and Install Locally
```bash
git clone https://github.com/aravindcm49/dir-init.git
cd dir-init
go install .
```

## Verification

After installation, verify the command works:
```bash
dir-init --help
```

If you get "command not found", restart your terminal or run `source ~/.zshrc` (or `source ~/.bash_profile`) to ensure PATH changes take effect.
