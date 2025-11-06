# dir-init

A fun Go CLI tool that generates funny, randomized folder names with customizable categories and alphanumeric suffixes.

## Features

- **6 Categories of Funny Names**:
  - Tech & Programming
  - Food & Cooking
  - Animals & Nature
  - Pop Culture
  - Silly & Absurd
  - Developer-related

- **Flexible Generation Options**:
  - Multiple suffix types (alpha, numeric, mixed, timestamp)
  - Configurable suffix length (1-8 characters)
  - Generate multiple names at once
  - Random seed support for reproducible results

- **Multiple Output Formats**:
  - Plain text (default)
  - JSON output
  - Colored terminal output

## Installation

### Prerequisites
- Go 1.19 or later installed on your system

### GOPATH Setup (if not already configured)

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

### Installation Methods

#### Method 1: Install from GitHub (Recommended)
```bash
go install github.com/aravindcm49/dir-init@latest
```

#### Method 2: Build from Source
```bash
git clone https://github.com/aravindcm49/dir-init.git
cd dir-init
go build -o dir-init main.go
./dir-init --help
```

#### Method 3: Build and Install Locally
```bash
git clone https://github.com/aravindcm49/dir-init.git
cd dir-init
go install .
```

### Verification
After installation, verify the command works:
```bash
dir-init --help
```

If you get "command not found", restart your terminal or run `source ~/.zshrc` (or `source ~/.bash_profile`) to ensure PATH changes take effect.

## Usage

### Generate a Single Name
```bash
# Generate from any category
./dir-init generate -c tech
# Output: ruby-v0nk

# Generate a food-related name
./dir-init generate -c food
# Output: wings-j0kr
```

### Generate Multiple Names
```bash
# Generate 5 tech names
./dir-init generate -c tech -n 5
# Output:
# Generated folder names:
# 1. code-1234
# 2. debug-abc2
# 3. api-x42
# 4. service-test
# 5. binary-v1
```

### Different Suffix Types
```bash
# Alpha suffix only
./dir-init generate -c silly -s alpha -l 6
# Output: potato-abcdef

# Numeric suffix only
./dir-init generate -c animals -s numeric -l 3
# Output: penguin-123

# Mixed suffix (default)
./dir-init generate -c pop -s mixed -l 4
# Output: ninja-1a2b

# Timestamp suffix
./dir-init generate -c dev -s timestamp
# Output: api-1714567890
```

### List Categories
```bash
./dir-init categories
```

### Show Examples
```bash
# Show examples for all categories
./dir-init examples

# Show examples for specific category
./dir-init examples -c food -n 5
```

### JSON Output
```bash
# Generate names in JSON format
./dir-init generate -c all -n 3 -o json
# Output:
# {
#   "count": 3,
#   "names": [
#     "pizza-smpc",
#     "debug-7x8c",
#     "koala-v0nk"
#   ]
# }
```

### Reproducible Results
```bash
# Use a specific seed for reproducible results
./dir-init generate -c tech -S 12345
./dir-init generate -c tech -S 12345  # Same result
```

## Categories

### Tech & Programming
Technology and programming related words including languages, frameworks, tools, and concepts.

**Examples**: `code-1234`, `debug-abc2`, `api-v2beta`

### Food & Cooking
Food, cooking, and beverage related words for delicious folder names.

**Examples**: `pizza-fresh`, `burger-delicious`, `taco-hot`

### Animals & Nature
Animals and nature related words for cute and wild folder names.

**Examples**: `penguin-cute`, `koala-gentle`, `dolphin-smart`

### Pop Culture
Pop culture, fantasy, and creative arts related words.

**Examples**: `ninja-epic`, `wizard-magical`, `knight-brave`

### Silly & Absurd
Silly, funny, and absurd words for humorous folder names.

**Examples**: `potato-silly`, `banana-goofy`, `unicorn-odd`

### Developer-related
Development tools, programming languages, and devops related words.

**Examples**: `github-v1`, `docker-prod`, `react-ui`

## Command Reference

### `generate`
Generate funny folder names.

**Flags:**
- `-c, --category`: Category to use (tech, food, animals, pop, silly, dev, all)
- `-s, --suffix`: Suffix type (alpha, numeric, mixed, timestamp)
- `-l, --length`: Suffix length (1-8)
- `-n, --count`: Number of names to generate
- `-S, --seed`: Random seed for reproducible results
- `-o, --output`: Output format (text, json)

### `categories`
List all available categories with descriptions and word counts.

### `examples`
Show example folder names for demonstration.

**Flags:**
- `-c, --category`: Show examples for specific category
- `-n, --count`: Number of examples to show

## Development

### Project Structure
```
dir-init/
├── cmd/                    # CLI commands
│   ├── main.go            # Main entry point
│   ├── root.go            # Root command
│   ├── generate.go        # Generate command
│   ├── categories.go      # Categories command
│   └── examples.go        # Examples command
├── internal/
│   ├── categories/        # Category word lists
│   │   ├── tech.go
│   │   ├── food.go
│   │   ├── animals.go
│   │   ├── pop.go
│   │   ├── silly.go
│   │   └── dev.go
│   └── generator/         # Name generation logic
│       └── generator.go
├── main.go               # Application entry point
├── go.mod                # Go module
└── README.md             # This file
```

### Running Tests
```bash
go test ./...
```

### Building
```bash
go build -o dir-init main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add your funny words to the appropriate category files
4. Test your changes
5. Submit a pull request

## License

MIT License - feel free to use this tool for any purpose!