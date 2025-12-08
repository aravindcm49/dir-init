# dir-init

A fun Go CLI tool that generates funny, randomized folder names with customizable categories and alphanumeric suffixes. Features an interactive TUI mode for guided directory creation and a powerful config system for customizing your experience.

## Features

- **Interactive Mode (Default)**:
  - Beautiful TUI-based interactive flow
  - Step-by-step guidance: Frontend → Backend → Category → Suffix → Count
  - Automatically creates directories with enhanced format: `{frontend}-{backend}-{category}-{suffix}`
  - Example: `rct-node-pizza-a1b2` (React + Node.js + Food category)

- **6 Categories of Funny Names**:
  - Tech & Programming
  - Food & Cooking
  - Animals & Nature
  - Pop Culture
  - Silly & Absurd
  - Developer-related

- **Config System**:
  - YAML-based configuration at `~/.dir-init/config.yaml`
  - Customize frontends, backends, tech stacks, frameworks, and category words
  - Add/remove items via CLI commands
  - Validate and edit config easily

- **Flexible Generation Options**:
  - Multiple suffix types (alpha, numeric, mixed, timestamp)
  - Configurable suffix length (1-8 characters)
  - Generate multiple names at once
  - Random seed support for reproducible results
  - Non-interactive `generate` command for name generation only

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

### Interactive Mode (Default)

By default, `dir-init` runs in interactive mode, guiding you through creating directories with a beautiful TUI:

```bash
# Simply run dir-init (interactive mode is default)
dir-init

# Or explicitly enable interactive mode
dir-init --interactive
# or
dir-init -i
```

**Interactive Flow:**
1. **Select Frontend**: Choose from React, Vue, Angular, Next.js, Svelte, etc. (or add custom)
2. **Select Backend**: Choose from Node.js, Python, Go, Java, etc. (or add custom)
3. **Select Category**: Choose from food, animals, pop, silly, dev, or all
4. **Select Suffix Type**: Alphabetic, Numeric, Mixed, or Timestamp
5. **Enter Count**: How many directories to create (1-10)

**Example Output:**
```
========
dir-init
========
Step 1/4: Select Frontend >> rct
Step 2/4: Select Backend >> node
Step 3/4: Select Category >> food
Step 4/4: Select Suffix Type >> mixed
How many directories to create? <default 1, enter number to change> 2

rct-node-pizza-a1b2 created!
rct-node-burger-x9y3 created!
```

**Interactive Mode Flags:**
- `--no-interactive`: Skip interactive mode and show help
- `--interactive` or `-i`: Explicitly enable interactive mode (overrides `--no-interactive`)

### Non-Interactive Mode: Generate Names Only

Use the `generate` command to generate names without creating directories:

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
# Generate 5 tech names (does not create directories)
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

## Config Management

`dir-init` uses a YAML configuration file at `~/.dir-init/config.yaml` to store your custom frontends, backends, tech stacks, frameworks, and category words.

### Initialize Config

```bash
# Create config file with default values
dir-init config init
```

### View Config

```bash
# Show config file path
dir-init config path

# Show all loaded collections
dir-init config show

# Validate config syntax
dir-init config validate
```

### Edit Config

```bash
# Open config in your default editor ($EDITOR or vim)
dir-init config edit
```

### Add Items

```bash
# Add a tech stack
dir-init config add techstack <code> <description>
# Example: dir-init config add techstack fejs "Frontend JS"

# Add a framework
dir-init config add framework <techstack> <code> <description>
# Example: dir-init config add framework fejs react "React Framework"

# Add a word to a category
dir-init config add word <category> <word>
# Example: dir-init config add word food taco
```

### Remove Items

```bash
# Remove a tech stack
dir-init config remove techstack <code>
# Example: dir-init config remove techstack fejs

# Remove a framework
dir-init config remove framework <techstack> <code>
# Example: dir-init config remove framework fejs react

# Remove a word from a category
dir-init config remove word <category> <word>
# Example: dir-init config remove word food taco
```

### Config Structure

The config file structure:

```yaml
# dir-init Custom Collections
frontends:
  - code: rct
    description: React
  - code: vue
    description: Vue.js
  # ... more frontends

backends:
  - code: node
    description: Node.js
  - code: py
    description: Python
  # ... more backends

categories:
  food:
    - pizza
    - burger
    - taco
    # ... more words
  animals:
    - penguin
    - koala
    # ... more words
  # ... more categories
```

## Categories

### Tech & Programming
Technology and programming related words including languages, frameworks, tools, and concepts.

**Examples**: 
- Generate command: `code-1234`, `debug-abc2`, `api-v2beta`
- Interactive mode: `rct-node-code-a1b2`, `vue-py-debug-x9y3`

### Food & Cooking
Food, cooking, and beverage related words for delicious folder names.

**Examples**: 
- Generate command: `pizza-fresh`, `burger-delicious`, `taco-hot`
- Interactive mode: `rct-node-pizza-a1b2`, `vue-py-burger-x9y3`

### Animals & Nature
Animals and nature related words for cute and wild folder names.

**Examples**: 
- Generate command: `penguin-cute`, `koala-gentle`, `dolphin-smart`
- Interactive mode: `rct-node-penguin-a1b2`, `vue-py-koala-x9y3`

### Pop Culture
Pop culture, fantasy, and creative arts related words.

**Examples**: 
- Generate command: `ninja-epic`, `wizard-magical`, `knight-brave`
- Interactive mode: `rct-node-ninja-a1b2`, `vue-py-wizard-x9y3`

### Silly & Absurd
Silly, funny, and absurd words for humorous folder names.

**Examples**: 
- Generate command: `potato-silly`, `banana-goofy`, `unicorn-odd`
- Interactive mode: `rct-node-potato-a1b2`, `vue-py-banana-x9y3`

### Developer-related
Development tools, programming languages, and devops related words.

**Examples**: 
- Generate command: `github-v1`, `docker-prod`, `react-ui`
- Interactive mode: `rct-node-github-a1b2`, `vue-py-docker-x9y3`

## Command Reference

### Root Command

**Default Behavior**: Running `dir-init` without arguments starts interactive mode, which creates directories.

**Flags:**
- `--no-interactive`: Skip interactive mode and show help instead
- `--interactive, -i`: Explicitly enable interactive mode (overrides `--no-interactive`)

### `generate`
Generate funny folder names (does not create directories, only outputs names).

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

### `config`
Manage custom word collections and configuration.

**Subcommands:**
- `config init`: Initialize config file with default values
- `config path`: Show config file path
- `config show`: Display all loaded collections
- `config validate`: Validate config file syntax
- `config edit`: Open config in default editor

**Add Subcommands:**
- `config add techstack <code> <description>`: Add a tech stack
- `config add framework <techstack> <code> <description>`: Add a framework
- `config add word <category> <word>`: Add a word to a category

**Remove Subcommands:**
- `config remove techstack <code>`: Remove a tech stack
- `config remove framework <techstack> <code>`: Remove a framework
- `config remove word <category> <word>`: Remove a word from a category

## Development

### Project Structure
```
dir-init/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command and flags
│   ├── generate.go        # Generate command (non-interactive)
│   ├── categories.go      # Categories command
│   ├── examples.go        # Examples command
│   ├── config.go          # Config management commands
│   ├── interactive.go     # Interactive TUI mode
│   ├── interactive_helpers.go  # Interactive mode helpers
│   └── tui/               # Terminal UI components
│       └── models/
│           └── selector.go  # TUI selector model
├── internal/
│   ├── config/            # Configuration management
│   │   ├── loader.go      # Config loading and saving
│   │   ├── save_helpers.go  # Helper functions for saving
│   │   └── types.go       # Config type definitions
│   ├── generator/         # Name generation logic
│   │   └── generator.go  # Generator implementation
│   └── utils/            # Utility functions
│       └── filesystem.go  # Filesystem utilities
├── main.go               # Application entry point
├── go.mod                # Go module
├── go.sum                # Go module checksums
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