package analysis

import (
	"bufio"
	"bytes"
	"fmt"
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

type IniData map[string]map[string]string

func parseIniFile(contents []byte) (IniData, error) {
	data := make(IniData)
	scanner := bufio.NewScanner(bytes.NewReader(contents))

	// Default section for entries before the first named section
	currentSection := "default"
	data[currentSection] = make(map[string]string)

	lineNum := 0
	var inMultilineValue bool
	var currentKey string
	var multilineValue strings.Builder

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if line == "" && !inMultilineValue {
			continue
		}

		if strings.HasPrefix(line, ";") && !inMultilineValue {
			continue
		}

		if inMultilineValue {
			// Check if this line ends the multiline value
			if strings.HasSuffix(line, "}") {
				multilineValue.WriteString(line[:len(line)-1])
				data[currentSection][currentKey] = multilineValue.String()
				inMultilineValue = false
				multilineValue.Reset()
			} else {
				// Add this line to the multiline value
				multilineValue.WriteString(line)
				multilineValue.WriteString("\n")
			}
			continue
		}

		// Check for section
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// Extract section name without brackets
			currentSection = line[1 : len(line)-1]
			// Initialize the section if it doesn't exist
			if _, exists := data[currentSection]; !exists {
				data[currentSection] = make(map[string]string)
			}
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format at line %d: %s", lineNum, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Check if this is the start of a multiline value
		if strings.HasPrefix(value, "{") {
			if strings.HasSuffix(value, "}") {
				// Single line with curly braces, just strip them
				data[currentSection][key] = value[1 : len(value)-1]
			} else {
				// Start of multiline value
				inMultilineValue = true
				currentKey = key
				multilineValue.WriteString(value[1:])
				multilineValue.WriteString("\n")
			}
		} else {
			// Regular key-value pair
			data[currentSection][key] = value
		}
	}

	// Check if we ended in the middle of a multiline value
	if inMultilineValue {
		return nil, fmt.Errorf("unclosed multiline value starting at key %s", currentKey)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func ParseGodotProjectFile(contents []byte) (*GodotProjectFile, error) {
	var projectData GodotProjectFile

	iniData, err := parseIniFile(contents)
	if err != nil {
		return nil, err
	}

	projectData.ApplicationName = iniData["application"]["config/name"]

	for key := range iniData["input"] {
		projectData.InputConfigs = append(projectData.InputConfigs, InputConfig{
			Name: key,
		})
	}

	return &projectData, nil
}
