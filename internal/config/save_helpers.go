package config

import "fmt"

// SaveFrontend adds a frontend to the config
func SaveFrontend(code, description string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Check if already exists
	for _, fe := range config.Frontends {
		if fe.Code == code {
			return fmt.Errorf("frontend '%s' already exists", code)
		}
	}

	// Add new frontend
	config.Frontends = append(config.Frontends, Frontend{
		Code:        code,
		Description: description,
	})

	return SaveConfig(config)
}

// SaveBackend adds a backend to the config
func SaveBackend(code, description string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Check if already exists
	for _, be := range config.Backends {
		if be.Code == code {
			return fmt.Errorf("backend '%s' already exists", code)
		}
	}

	// Add new backend
	config.Backends = append(config.Backends, Backend{
		Code:        code,
		Description: description,
	})

	return SaveConfig(config)
}
