package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// OpenCodeService expune către frontend verificarea statusului
// și setup-ul serviciului OpenCode2.
type OpenCodeService struct{}

// OpenCodeStatus descrie starea serviciului OpenCode2.
type OpenCodeStatus struct {
	Ready bool   `json:"ready"`
	URL   string `json:"url"`
	Error string `json:"error"`
}

// runCommand execută o comandă și întoarce output-ul standard.
func runCommand(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("%s: %w", strings.TrimSpace(string(out)), err)
	}
	return strings.TrimSpace(string(out)), nil
}

// IsReady verifică dacă serviciul OpenCode2 rulează.
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

// Setup instalează CLI-ul OpenCode2 dacă lipsește și pornește serviciul.
func (s *OpenCodeService) Setup() OpenCodeStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if _, err := exec.LookPath("opencode2"); err != nil {
		cmd := exec.CommandContext(ctx, "npm", "install", "-g", "@opencode-ai/cli@next")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return OpenCodeStatus{
				Ready: false,
				Error: fmt.Sprintf("instalare eșuată: %s", strings.TrimSpace(string(out))),
			}
		}
	}

	out, err := runCommand(ctx, "opencode2", "service", "start")
	if err != nil {
		return OpenCodeStatus{Ready: false, Error: err.Error()}
	}

	out = strings.TrimSpace(out)
	if out == "" {
		return OpenCodeStatus{Ready: false, Error: "serviciul a pornit fără să întoarcă un URL"}
	}

	// Re-verificăm că serviciul răspunde.
	return s.IsReady()
}
