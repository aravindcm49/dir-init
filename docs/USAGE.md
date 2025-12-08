# Usage Guide

## Interactive Mode (Default)

By default, `dir-init` runs in interactive mode, guiding you through creating directories with a beautiful TUI:

```bash
# Simply run dir-init (interactive mode is default)
dir-init

# Or explicitly enable interactive mode
dir-init --interactive
# or
dir-init -i
```

### Interactive Flow

1. **Select Frontend**: Choose from React, Vue, Angular, Next.js, Svelte, etc. (or add custom)
2. **Select Backend**: Choose from Node.js, Python, Go, Java, etc. (or add custom)
3. **Select Category**: Choose from food, animals, pop, silly, dev, or all
4. **Select Suffix Type**: Alphabetic, Numeric, Mixed, or Timestamp
5. **Enter Count**: How many directories to create (1-10)

### Example Output

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

### Interactive Mode Flags
- `--no-interactive`: Skip interactive mode and show help
- `--interactive` or `-i`: Explicitly enable interactive mode (overrides `--no-interactive`)

---

## Non-Interactive Mode: Generate Names Only

Use the `generate` command to generate names without creating directories:

```bash
# Generate from any category
dir-init generate -c tech
# Output: ruby-v0nk

# Generate a food-related name
dir-init generate -c food
# Output: wings-j0kr
```

### Generate Multiple Names
```bash
# Generate 5 tech names (does not create directories)
dir-init generate -c tech -n 5
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
dir-init generate -c silly -s alpha -l 6
# Output: potato-abcdef

# Numeric suffix only
dir-init generate -c animals -s numeric -l 3
# Output: penguin-123

# Mixed suffix (default)
dir-init generate -c pop -s mixed -l 4
# Output: ninja-1a2b

# Timestamp suffix
dir-init generate -c dev -s timestamp
# Output: api-1714567890
```

### List Categories
```bash
dir-init categories
```

### Show Examples
```bash
# Show examples for all categories
dir-init examples

# Show examples for specific category
dir-init examples -c food -n 5
```

### JSON Output
```bash
# Generate names in JSON format
dir-init generate -c all -n 3 -o json
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
dir-init generate -c tech -S 12345
dir-init generate -c tech -S 12345  # Same result
```

---

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
