# Config Management

`dir-init` uses a YAML configuration file at `~/.dir-init/config.yaml` to store your custom frontends, backends, tech stacks, frameworks, and category words.

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
# Example: dir-init config add techstack fejs "Frontend JS"

# Add a framework
dir-init config add framework <techstack> <code> <description>
# Example: dir-init config add framework fejs react "React Framework"

# Add a word to a category
dir-init config add word <category> <word>
# Example: dir-init config add word food taco
```

## Remove Items

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

## Config Structure

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
- `config add framework <techstack> <code> <description>`: Add a framework
- `config add word <category> <word>`: Add a word to a category

**Remove Subcommands:**
- `config remove techstack <code>`: Remove a tech stack
- `config remove framework <techstack> <code>`: Remove a framework
- `config remove word <category> <word>`: Remove a word from a category
