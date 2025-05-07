package lsp

import (
	"gdx/analysis"
	"log"
	"os"
	"path/filepath"
)

type Registration struct {
	Id     string `json:"id"`
	Method string `json:"method"`
}

type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

func HandleInitialized(logger *log.Logger, state *ServerState) error {
	workspacePath := state.WorkspacePath
	projectFilePath := filepath.Join(workspacePath, "project.godot")

	data, err := os.ReadFile(projectFilePath)
	if err != nil {
		return err
	}

	projectConfig, err := analysis.ParseGodotProjectFile(data)
	if err != nil {
		return err
	}

	state.ProjectConfig = *projectConfig

	logger.Printf("loaded Godot project: %s\n", state.ProjectConfig.ApplicationName)

	return nil
}
