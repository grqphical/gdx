package analysis

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
)

type InputConfig struct {
	Name       string
	Deadzone   float32
	Keybinding string
}

type GodotProjectFile struct {
	ApplicationName string
	InputConfigs    []InputConfig
}

type Config map[string]map[string]string

func ParseGodotProjectFile(contents []byte) (*GodotProjectFile, error) {
	result := GodotProjectFile{}

	config := make(Config)
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	currentSection := "DEFAULT"
	config[currentSection] = make(map[string]string)

	var lastKey string

	for scanner.Scan() {
		line := scanner.Text()

		// Strip comments
		if idx := strings.Index(line, ";"); idx != -1 {
			line = line[:idx]
		}
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			currentSection = strings.TrimSpace(trimmed[1 : len(trimmed)-1])
			if _, exists := config[currentSection]; !exists {
				config[currentSection] = make(map[string]string)
			}
			lastKey = ""
		} else if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
			// Continuation line for the last key
			if lastKey != "" {
				config[currentSection][lastKey] += "\n" + strings.TrimSpace(line)
			}
		} else if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			config[currentSection][key] = val
			lastKey = key
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if application, ok := config["application"]; ok {
		result.ApplicationName = application["config/name"]
	} else {
		return nil, errors.New("no application field in config")
	}

	if input, ok := config["input"]; ok {
		for key := range input {
			result.InputConfigs = append(result.InputConfigs, InputConfig{
				Name: key,
			})
		}
	}

	return &result, nil
}
