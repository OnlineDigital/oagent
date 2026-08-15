package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// OpenCodeService exposes status checking and OpenCode2 service setup to the frontend.
type OpenCodeService struct{}

// OpenCodeStatus describes the state of the OpenCode2 service.
type OpenCodeStatus struct {
	Ready bool   `json:"ready"`
	URL   string `json:"url"`
	Error string `json:"error"`
}

// runCommand executes a command and returns its standard output.
func runCommand(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("%s: %w", strings.TrimSpace(string(out)), err)
	}
	return strings.TrimSpace(string(out)), nil
}

// IsReady checks whether the OpenCode2 service is running.
func (s *OpenCodeService) IsReady() OpenCodeStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	out, err := runCommand(ctx, "opencode2", "service", "status")
	if err != nil {
		return OpenCodeStatus{Ready: false, Error: err.Error()}
	}

	out = strings.TrimSpace(out)
	if out == "" || out == "stopped" {
		return OpenCodeStatus{Ready: false}
	}

	return OpenCodeStatus{Ready: true, URL: out}
}

// Setup installs the OpenCode2 CLI if missing and starts the service.
func (s *OpenCodeService) Setup() OpenCodeStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if _, err := exec.LookPath("opencode2"); err != nil {
		cmd := exec.CommandContext(ctx, "npm", "install", "-g", "@opencode-ai/cli@next")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return OpenCodeStatus{
				Ready: false,
				Error: fmt.Sprintf("installation failed: %s", strings.TrimSpace(string(out))),
			}
		}
	}

	out, err := runCommand(ctx, "opencode2", "service", "start")
	if err != nil {
		return OpenCodeStatus{Ready: false, Error: err.Error()}
	}

	out = strings.TrimSpace(out)
	if out == "" {
		return OpenCodeStatus{Ready: false, Error: "service started without returning a URL"}
	}

	// Re-check that the service is responding.
	return s.IsReady()
}
