# Config Management

`dir-init` uses a YAML configuration file at `~/.dir-init/config-v2.yaml` to store your custom tech stacks and category words. Nicknames are entered during interactive mode and not saved to config.

## Initialize Config

```bash
# Create config file with default values
dir-init config init
```

## View Config

```bash
# Show config file path
dir-init config path

# Show all loaded collections
dir-init config show

# Validate config syntax
dir-init config validate
```

## Edit Config

```bash
# Open config in your default editor ($EDITOR or vim)
dir-init config edit
```

## Add Items

```bash
# Add a tech stack
dir-init config add techstack <code> <description>
# Example: dir-init config add techstack react "React"

# Add a word to a category
dir-init config add word <category> <word>
# Example: dir-init config add word food taco
```

## Remove Items

```bash
# Remove a tech stack
dir-init config remove techstack <code>
# Example: dir-init config remove techstack react

# Remove a word from a category
dir-init config remove word <category> <word>
# Example: dir-init config remove word food taco
```

## Config Structure

The config file structure:

```yaml
# dir-init Custom Collections
tech-stacks:
  - code: react
    description: React
  - code: node
    description: Node.js
  - code: python
    description: Python
  # ... more tech stacks

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

## Command Reference

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
- `config add word <category> <word>`: Add a word to a category

**Remove Subcommands:**
- `config remove techstack <code>`: Remove a tech stack
- `config remove word <category> <word>`: Remove a word from a category
