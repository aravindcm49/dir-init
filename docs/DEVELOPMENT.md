# Development

## Project Structure

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
└── README.md             # Project overview
```

## Running Tests

```bash
go test ./...
```

## Building

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
