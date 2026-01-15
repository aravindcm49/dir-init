# dir-init

A fun Go CLI tool that generates funny, randomized folder names with customizable categories and alphanumeric suffixes. Features an interactive TUI mode for guided directory creation.

## Quick Start

```bash
# Install
go install github.com/aravindcm49/dir-init@latest

# Run (interactive mode)
dir-init
```

## Features

- **Interactive TUI Mode** — Step-by-step guided directory creation
- **6 Fun Categories** — Tech, Food, Animals, Pop Culture, Silly, Developer
- **Flexible Suffixes** — Alpha, numeric, mixed, or timestamp
- **YAML Config** — Customize tech stacks and category words
- **Multiple Outputs** — Plain text, JSON, colored terminal

## Example

```
========
dir-init
========
Step 1/4: Enter Nickname >> myapp
Step 2/4: Select Tech Stack >> node
Step 3/4: Select Category >> food
Step 4/4: Select Suffix Type >> mixed

myapp-node-pizza-a1b2 created!
```

**Format**: `{nickname}-{techstack}-{categoryword}-{suffix}`

## Shell Helper (Recommended)

Add this function to your `~/.zshrc` for the best experience — it automatically `cd`s into the newly created directory:

```bash
di() {
    # Remember current directory listing
    local before=$(ls -1d */ 2>/dev/null | sort)
    
    # Run dir-init normally (preserves colors and TUI)
    command dir-init "$@"
    
    # Get new directory listing
    local after=$(ls -1d */ 2>/dev/null | sort)
    
    # Find new directories (ones in after but not in before)
    local new_dirs=$(comm -13 <(echo "$before") <(echo "$after"))
    
    # Count new directories
    local count=$(echo "$new_dirs" | grep -v '^$' | wc -l | tr -d ' ')
    
    # If exactly one new directory, cd into it
    if [ "$count" = "1" ]; then
        local dir=$(echo "$new_dirs" | head -n1 | tr -d '/')
        if [ -n "$dir" ] && [ -d "$dir" ]; then
            cd "$dir"
        fi
    fi
}
```

Then use `di` instead of `dir-init` to create and enter directories in one go.

## Documentation

- [Installation](docs/INSTALLATION.md) — Prerequisites, GOPATH setup, installation methods
- [Usage Guide](docs/USAGE.md) — Interactive mode, generate command, flags reference
- [Config Management](docs/CONFIG.md) — Customize frontends, backends, and words
- [Categories](docs/CATEGORIES.md) — All 6 category descriptions with examples
- [Development](docs/DEVELOPMENT.md) — Project structure, building, contributing

## License

MIT License
