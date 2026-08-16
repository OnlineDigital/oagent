package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"changeme/opencode2"
)

// OpenCodeService exposes status checking and OpenCode2 service setup to the frontend.
type OpenCodeService struct{}

// OpenCodeStatus describes the state of the OpenCode2 service.
type OpenCodeStatus struct {
	Ready bool   `json:"ready"`
	URL   string `json:"url"`
	Error string `json:"error"`
}

// HarnessInfo describes one selectable harness or device.
type HarnessInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Local       bool   `json:"local"`
	Online      bool   `json:"online"`
}

// McpInfo describes one configured MCP server.
type McpInfo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Active bool   `json:"active"`
}

// PluginInfo describes one plugin and whether it is currently enabled.
type PluginInfo struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`
}

// SkillInfo describes one loaded skill.
type SkillInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Active      bool   `json:"active"`
}

// SessionInfo is a frontend-friendly projection of an OpenCode2 session.
type SessionInfo struct {
	ID        string `json:"id"`
	ParentID  string `json:"parentId,omitempty"`
	ProjectID string `json:"projectId"`
	Agent     string `json:"agent,omitempty"`
	Title     string `json:"title"`
	Directory string `json:"directory,omitempty"`
	Subpath   string `json:"subpath,omitempty"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Active    bool   `json:"active"`
}

// ProjectInfo is a frontend-friendly projection of an OpenCode2 project.
type ProjectInfo struct {
	ID        string `json:"id"`
	Canonical string `json:"canonical"`
	Name      string `json:"name,omitempty"`
	VCS       string `json:"vcs,omitempty"`
	UpdatedAt int64  `json:"updatedAt"`
}

// ConversationMessage is a normalized transcript entry for one session.
type ConversationMessage struct {
	ID        string     `json:"id"`
	Role      string     `json:"role"` // user, assistant, system, synthetic, shell, skill
	Agent     string     `json:"agent,omitempty"`
	Model     string     `json:"model,omitempty"`
	Text      string     `json:"text,omitempty"`
	Reasoning string     `json:"reasoning,omitempty"`
	Tools     []ToolCall `json:"tools,omitempty"`
	CreatedAt int64      `json:"createdAt"`
	Completed int64      `json:"completedAt,omitempty"`
}

// ToolCall is a single tool invocation embedded in an assistant message.
type ToolCall struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Input     string `json:"input,omitempty"`
	Output    string `json:"output,omitempty"`
	Diff      string `json:"diff,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	Completed int64  `json:"completedAt,omitempty"`
}

// SubagentInfo is a child session that was spawned by a parent session.
type SubagentInfo struct {
	ID       string                `json:"id"`
	ParentID string                `json:"parentId"`
	Agent    string                `json:"agent,omitempty"`
	Title    string                `json:"title"`
	Status   string                `json:"status"`
	Messages []ConversationMessage `json:"messages"`
}

// runCommand executes a command and returns its standard output.
func runCommand(ctx context.Context, name string, args ...string) (string, error) {
	log.Printf("[opencode2] exec %s %s", name, strings.Join(args, " "))
	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("[opencode2] exec error: %v (stdout=%s)", err, strings.TrimSpace(string(out)))
		return string(out), fmt.Errorf("%s: %w", strings.TrimSpace(string(out)), err)
	}
	log.Printf("[opencode2] exec ok: %s", strings.TrimSpace(string(out)))
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

// Harnesses returns the selectable harnesses/devices, with the local device first.
func (s *OpenCodeService) Harnesses() ([]HarnessInfo, error) {
	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		hostname = "Local device"
	}

	status := s.IsReady()
	return []HarnessInfo{
		{
			ID:          "local",
			Name:        "opencode2",
			Description: hostname,
			Local:       true,
			Online:      status.Ready,
		},
	}, nil
}

// McpServers returns configured MCP servers and their connection status.
func (s *OpenCodeService) McpServers() ([]McpInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := s.serviceClient(ctx)
	if err != nil {
		return nil, err
	}

	raw, err := client.V2McpList(ctx, nil)
	if err != nil {
		return nil, err
	}

	data, _ := raw["data"].([]interface{})
	servers := make([]McpInfo, 0, len(data))
	for _, item := range data {
		b, err := json.Marshal(item)
		if err != nil {
			continue
		}
		var mcp struct {
			Name   string      `json:"name"`
			Status interface{} `json:"status"`
		}
		if err := json.Unmarshal(b, &mcp); err != nil {
			continue
		}

		status := "unknown"
		if statusMap, ok := mcp.Status.(map[string]interface{}); ok {
			status = asString(statusMap["status"])
		}

		servers = append(servers, McpInfo{
			Name:   mcp.Name,
			Status: status,
			Active: status == "connected",
		})
	}

	return servers, nil
}

// ToggleMcp connects or disconnects an MCP server.
func (s *OpenCodeService) ToggleMcp(server string, enable bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := s.serviceClient(ctx)
	if err != nil {
		return err
	}

	if enable {
		_, err = client.V2McpConnect(ctx, server, nil)
	} else {
		_, err = client.V2McpDisconnect(ctx, server, nil)
	}
	return err
}

// Plugins returns configured plugins and whether each one is currently enabled.
func (s *OpenCodeService) Plugins() ([]PluginInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := s.serviceClient(ctx)
	if err != nil {
		return nil, err
	}

	raw, err := client.V2PluginList(ctx, nil)
	if err != nil {
		return nil, err
	}

	active := map[string]bool{}
	if data, ok := raw["data"].([]interface{}); ok {
		for _, item := range data {
			if m, ok := item.(map[string]interface{}); ok {
				if id := asString(m["id"]); id != "" {
					active[id] = true
				}
			}
		}
	}

	configured := s.configuredPlugins()
	seen := map[string]bool{}
	plugins := make([]PluginInfo, 0, len(configured)+len(active))

	for _, id := range configured {
		seen[id] = true
		plugins = append(plugins, PluginInfo{ID: id, Active: active[id]})
	}
	for id := range active {
		if seen[id] {
			continue
		}
		plugins = append(plugins, PluginInfo{ID: id, Active: true})
	}

	return plugins, nil
}

// TogglePlugin enables or disables a plugin in the user config.
func (s *OpenCodeService) TogglePlugin(id string, enable bool) error {
	return updatePluginConfig(id, enable)
}

// Skills returns skills loaded by OpenCode2.
func (s *OpenCodeService) Skills() ([]SkillInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := s.serviceClient(ctx)
	if err != nil {
		return nil, err
	}

	raw, err := client.V2SkillList(ctx, nil)
	if err != nil {
		return nil, err
	}

	denied := s.deniedSkills()

	data, _ := raw["data"].([]interface{})
	skills := make([]SkillInfo, 0, len(data))
	for _, item := range data {
		b, err := json.Marshal(item)
		if err != nil {
			continue
		}
		var skill SkillInfo
		if err := json.Unmarshal(b, &skill); err != nil {
			continue
		}
		skill.Active = !denied[skill.ID]
		skills = append(skills, skill)
	}

	return skills, nil
}

// ToggleSkill enables or disables a skill via permission rules.
func (s *OpenCodeService) ToggleSkill(id string, enable bool) error {
	return updateSkillPermission(id, enable)
}

// deniedSkills returns skill IDs denied by the active OpenCode config.
func (s *OpenCodeService) deniedSkills() map[string]bool {
	denied := map[string]bool{}
	for _, config := range readOpenCodeConfigs() {
		permissions := config.Permissions
		if agentRaw, ok := config.Raw["agents"].(map[string]interface{}); ok {
			for _, agentValue := range agentRaw {
				agentMap, ok := agentValue.(map[string]interface{})
				if !ok {
					continue
				}
				if list, ok := agentMap["permissions"].([]interface{}); ok {
					permissions = append(permissions, list...)
				}
			}
		}

		for _, rule := range permissions {
			if !isSkillRule(rule, "deny") {
				continue
			}
			resource := ruleResource(rule)
			if resource == "" {
				continue
			}
			if resource == "*" {
				for _, id := range config.SkillIDs {
					denied[id] = true
				}
				continue
			}
			for _, id := range config.SkillIDs {
				if matchWildcard(resource, id) {
					denied[id] = true
				}
			}
		}
	}
	return denied
}

// openCodeConfig is a flattened view of one OpenCode config document.
type openCodeConfig struct {
	Path        string
	Raw         map[string]interface{}
	Permissions []interface{}
	SkillIDs    []string
}

func readOpenCodeConfigs() []openCodeConfig {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	paths := []string{
		filepath.Join(home, ".config", "opencode", "opencode.json"),
		filepath.Join(home, ".opencode", "opencode.json"),
	}

	var configs []openCodeConfig
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var raw map[string]interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			continue
		}

		config := openCodeConfig{Path: path, Raw: raw}
		if list, ok := raw["permissions"].([]interface{}); ok {
			config.Permissions = list
		}
		if entries, ok := raw["skills"].([]interface{}); ok {
			for _, entry := range entries {
				if id, ok := entry.(string); ok && id != "" {
					config.SkillIDs = append(config.SkillIDs, id)
				}
			}
		}
		configs = append(configs, config)
	}

	return configs
}

// updateSkillPermission enables or disables a skill by editing permission rules.
func updateSkillPermission(id string, enable bool) error {
	configs := readOpenCodeConfigs()
	if len(configs) == 0 {
		return fmt.Errorf("no OpenCode config found")
	}

	// Apply the change to the highest-priority config that declares skills or permissions.
	for i := len(configs) - 1; i >= 0; i-- {
		config := configs[i]
		if _, hasPermissions := config.Raw["permissions"]; !hasPermissions {
			if len(config.SkillIDs) == 0 {
				continue
			}
		}

		permissions, _ := config.Raw["permissions"].([]interface{})
		if permissions == nil {
			permissions = []interface{}{}
		}

		updated := make([]interface{}, 0, len(permissions)+1)
		for _, rule := range permissions {
			if isSkillRule(rule, "deny") && ruleResource(rule) == id {
				continue
			}
			updated = append(updated, rule)
		}

		if !enable {
			updated = append(updated, map[string]interface{}{
				"action":   "skill",
				"resource": id,
				"effect":   "deny",
			})
		}

		config.Raw["permissions"] = updated
		out, err := json.MarshalIndent(config.Raw, "", "  ")
		if err != nil {
			return err
		}
		return os.WriteFile(config.Path, out, 0o644)
	}

	return fmt.Errorf("skill %q not found in config", id)
}

func isSkillRule(rule interface{}, effect string) bool {
	ruleMap, ok := rule.(map[string]interface{})
	if !ok {
		return false
	}
	action := asString(ruleMap["action"])
	ruleEffect := asString(ruleMap["effect"])
	return action == "skill" && ruleEffect == effect
}

func ruleResource(rule interface{}) string {
	if ruleMap, ok := rule.(map[string]interface{}); ok {
		return asString(ruleMap["resource"])
	}
	return ""
}

func matchWildcard(pattern, value string) bool {
	if !strings.Contains(pattern, "*") {
		return pattern == value
	}
	prefix, suffix, _ := strings.Cut(pattern, "*")
	return strings.HasPrefix(value, prefix) && strings.HasSuffix(value, suffix)
}

// configuredPlugins reads plugin IDs from the user's OpenCode config files.
func (s *OpenCodeService) configuredPlugins() []string {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	paths := []string{
		filepath.Join(home, ".config", "opencode", "opencode.json"),
		filepath.Join(home, ".opencode", "opencode.json"),
	}

	var out []string
	seen := map[string]bool{}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var config struct {
			Plugins []interface{} `json:"plugins"`
		}
		if err := json.Unmarshal(data, &config); err != nil {
			continue
		}
		for _, entry := range config.Plugins {
			switch v := entry.(type) {
			case string:
				if v != "" && !seen[v] {
					seen[v] = true
					out = append(out, v)
				}
			case map[string]interface{}:
				if pkg := asString(v["package"]); pkg != "" && !seen[pkg] {
					seen[pkg] = true
					out = append(out, pkg)
				}
			}
		}
	}

	return out
}

// updatePluginConfig toggles a plugin in the user config files.
func updatePluginConfig(id string, enable bool) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	paths := []string{
		filepath.Join(home, ".config", "opencode", "opencode.json"),
		filepath.Join(home, ".opencode", "opencode.json"),
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var config map[string]interface{}
		if err := json.Unmarshal(data, &config); err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}

		plugins, ok := config["plugins"].([]interface{})
		if !ok {
			plugins = []interface{}{}
		}

		var updated []interface{}
		changed := false
		for _, entry := range plugins {
			var packageID string
			switch v := entry.(type) {
			case string:
				packageID = v
			case map[string]interface{}:
				packageID = asString(v["package"])
			}

			if packageID == id {
				changed = true
				if enable {
					updated = append(updated, entry)
				}
				continue
			}
			updated = append(updated, entry)
		}

		if !changed {
			continue
		}

		config["plugins"] = updated
		out, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, out, 0o644); err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("plugin %q not found in config", id)
}

// serviceClient builds an authenticated client for the running OpenCode2 service.
func (s *OpenCodeService) serviceClient(ctx context.Context) (*opencode2.Client, error) {
	log.Printf("[opencode2] building service client")
	out, err := runCommand(ctx, "opencode2", "service", "status")
	if err != nil {
		return nil, fmt.Errorf("service status: %w", err)
	}

	baseURL := strings.TrimSpace(out)
	if baseURL == "" || baseURL == "stopped" {
		return nil, errors.New("opencode2 service is not running")
	}

	password, err := runCommand(ctx, "opencode2", "service", "get", "password")
	if err != nil {
		return nil, fmt.Errorf("read service password: %w", err)
	}

	httpClient := &http.Client{
		Transport: &authRoundTripper{
			username: "opencode",
			password: strings.TrimSpace(password),
			next:     http.DefaultTransport,
		},
	}

	log.Printf("[opencode2] service client ready (baseURL=%s)", baseURL)
	return opencode2.NewClientWithHTTP(baseURL, httpClient), nil
}

// authRoundTripper injects the service basic-auth credentials into every request.
type authRoundTripper struct {
	username string
	password string
	next     http.RoundTripper
}

func (rt *authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.SetBasicAuth(rt.username, rt.password)
	log.Printf("[opencode2] HTTP %s %s", req.Method, req.URL.String())
	resp, err := rt.next.RoundTrip(clone)
	if err != nil {
		log.Printf("[opencode2] HTTP error: %v", err)
		return resp, err
	}
	log.Printf("[opencode2] HTTP %d %s", resp.StatusCode, req.URL.String())
	return resp, err
}

// Projects returns the projects known to OpenCode2.
func (s *OpenCodeService) Projects() ([]ProjectInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := s.serviceClient(ctx)
	if err != nil {
		return nil, err
	}

	raw, err := client.V2ProjectList(ctx)
	if err != nil {
		log.Printf("[opencode2] Projects failed: %v", err)
		return nil, err
	}

	projects := make([]ProjectInfo, 0, len(raw))
	for _, item := range raw {
		projects = append(projects, ProjectInfo{
			ID:        item.Id,
			Canonical: item.Canonical,
			Name:      item.Name,
			VCS:       asString(item.Vcs),
			UpdatedAt: asInt64(item.Time),
		})
	}

	log.Printf("[opencode2] Projects returned %d projects", len(projects))
	return projects, nil
}

// Sessions returns recent OpenCode2 sessions, newest first.
func (s *OpenCodeService) Sessions(limit int) ([]SessionInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := s.serviceClient(ctx)
	if err != nil {
		return nil, err
	}

	if limit <= 0 {
		limit = 50
	}

	response, err := client.V2SessionList(ctx, map[string]string{
		"limit": fmt.Sprintf("%d", limit),
		"order": "desc",
	})
	if err != nil {
		log.Printf("[opencode2] Sessions list failed: %v", err)
		return nil, err
	}

	active, err := client.V2SessionActive(ctx)
	if err != nil {
		// Active sessions are an enhancement, not a hard requirement.
		log.Printf("[opencode2] Sessions active lookup failed: %v", err)
		active = nil
	}

	activeIDs := map[string]bool{}
	if data, ok := active["data"].(map[string]interface{}); ok {
		for id := range data {
			activeIDs[id] = true
		}
	}

	sessions := make([]SessionInfo, 0, len(response.Data))
	for _, item := range response.Data {
		raw, err := json.Marshal(item)
		if err != nil {
			log.Printf("[opencode2] Sessions marshal item failed: %v", err)
			continue
		}

		var info opencode2.Session_Info
		if err := json.Unmarshal(raw, &info); err != nil {
			log.Printf("[opencode2] Sessions unmarshal item failed: %v", err)
			continue
		}

		id := asString(info.Id)
		sessions = append(sessions, SessionInfo{
			ID:        id,
			ParentID:  asString(info.ParentID),
			ProjectID: info.ProjectID,
			Agent:     info.Agent,
			Title:     info.Title,
			Directory: asString(info.Location),
			Subpath:   info.Subpath,
			CreatedAt: mapInt64(info.Time, "created"),
			UpdatedAt: mapInt64(info.Time, "updated"),
			Active:    activeIDs[id],
		})
	}

	log.Printf("[opencode2] Sessions returned %d sessions", len(sessions))
	return sessions, nil
}

// Conversation returns a normalized transcript for a session.
func (s *OpenCodeService) Conversation(sessionID string) ([]ConversationMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	log.Printf("[opencode2] Conversation requested for session %s", sessionID)
	client, err := s.serviceClient(ctx)
	if err != nil {
		return nil, err
	}

	response, err := client.V2MessageList(ctx, sessionID, map[string]string{
		"limit": "200",
		"order": "asc",
	})
	if err != nil {
		log.Printf("[opencode2] Conversation list failed for %s: %v", sessionID, err)
		return nil, err
	}

	messages := make([]ConversationMessage, 0, len(response.Data))
	for _, raw := range response.Data {
		message := normalizeMessage(raw)
		if message.ID == "" && message.Text == "" && len(message.Tools) == 0 && message.Reasoning == "" {
			continue
		}
		messages = append(messages, message)
	}

	log.Printf("[opencode2] Conversation returned %d messages for %s", len(messages), sessionID)
	return messages, nil
}

// Subagents returns child sessions spawned by a parent session.
func (s *OpenCodeService) Subagents(sessionID string) ([]SubagentInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Printf("[opencode2] Subagents requested for parent %s", sessionID)
	client, err := s.serviceClient(ctx)
	if err != nil {
		return nil, err
	}

	children, err := client.V2SessionList(ctx, map[string]string{
		"parentID": sessionID,
		"limit":    "100",
		"order":    "desc",
	})
	if err != nil {
		log.Printf("[opencode2] Subagents list failed for %s: %v", sessionID, err)
		return nil, err
	}

	active, err := client.V2SessionActive(ctx)
	if err != nil {
		active = nil
	}
	activeIDs := map[string]bool{}
	if data, ok := active["data"].(map[string]interface{}); ok {
		for id := range data {
			activeIDs[id] = true
		}
	}

	subagents := make([]SubagentInfo, 0, len(children.Data))
	for _, raw := range children.Data {
		b, err := json.Marshal(raw)
		if err != nil {
			continue
		}
		var info opencode2.Session_Info
		if err := json.Unmarshal(b, &info); err != nil {
			continue
		}

		id := asString(info.Id)
		status := "done"
		if activeIDs[id] {
			status = "active"
		}

		messages, err := s.Conversation(id)
		if err != nil {
			log.Printf("[opencode2] Subagents child messages failed for %s: %v", id, err)
			messages = nil
		}

		subagents = append(subagents, SubagentInfo{
			ID:       id,
			ParentID: sessionID,
			Agent:    info.Agent,
			Title:    info.Title,
			Status:   status,
			Messages: messages,
		})
	}

	log.Printf("[opencode2] Subagents returned %d children for %s", len(subagents), sessionID)
	return subagents, nil
}

// normalizeMessage converts a raw message map into a ConversationMessage.
func normalizeMessage(raw interface{}) ConversationMessage {
	b, err := json.Marshal(raw)
	if err != nil {
		log.Printf("[opencode2] normalizeMessage marshal failed: %v", err)
		return ConversationMessage{}
	}

	var envelope struct {
		ID      interface{}            `json:"id"`
		Type    string                 `json:"type"`
		Agent   string                 `json:"agent"`
		Model   interface{}            `json:"model"`
		Text    string                 `json:"text"`
		Time    map[string]interface{} `json:"time"`
		Content []interface{}          `json:"content"`
	}
	if err := json.Unmarshal(b, &envelope); err != nil {
		log.Printf("[opencode2] normalizeMessage unmarshal failed: %v", err)
		return ConversationMessage{}
	}

	message := ConversationMessage{
		ID:        asString(envelope.ID),
		Role:      envelope.Type,
		Agent:     envelope.Agent,
		Model:     asString(envelope.Model),
		Text:      envelope.Text,
		CreatedAt: mapInt64(envelope.Time, "created"),
		Completed: mapInt64(envelope.Time, "completed"),
	}

	// Older migrated sessions flatten tool calls into synthetic text messages
	// using the "Called the <Tool> tool with the following input: ..." shape.
	// Parse those into real tool cards instead of showing them as plain text.
	if len(envelope.Content) == 0 && envelope.Text != "" {
		if tools, leftover := parseLegacySyntheticTools(envelope.Text, asString(envelope.ID)); len(tools) > 0 {
			message.Tools = append(message.Tools, tools...)
			message.Text = strings.TrimSpace(leftover)
		}
	}

	for _, part := range envelope.Content {
		partJSON, err := json.Marshal(part)
		if err != nil {
			continue
		}
		var meta struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(partJSON, &meta); err != nil {
			continue
		}

		switch meta.Type {
		case "reasoning":
			var r struct {
				Text string `json:"text"`
			}
			if err := json.Unmarshal(partJSON, &r); err == nil && r.Text != "" {
				if message.Reasoning != "" {
					message.Reasoning += "\n"
				}
				message.Reasoning += r.Text
			}
		case "text":
			var t struct {
				Text string `json:"text"`
			}
			if err := json.Unmarshal(partJSON, &t); err == nil && t.Text != "" {
				if message.Text != "" {
					message.Text += "\n"
				}
				message.Text += t.Text
			}
		case "tool":
			if tool := normalizeTool(part); tool.ID != "" {
				message.Tools = append(message.Tools, tool)
			}
		}
	}

	return message
}

// normalizeTool converts a raw tool part into a ToolCall.
func normalizeTool(raw interface{}) ToolCall {
	b, err := json.Marshal(raw)
	if err != nil {
		return ToolCall{}
	}

	var part struct {
		ID    string                 `json:"id"`
		Name  string                 `json:"name"`
		State map[string]interface{} `json:"state"`
		Time  map[string]interface{} `json:"time"`
	}
	if err := json.Unmarshal(b, &part); err != nil {
		return ToolCall{}
	}

	tool := ToolCall{
		ID:        part.ID,
		Name:      part.Name,
		Status:    asString(part.State["status"]),
		CreatedAt: mapInt64(part.Time, "created"),
		Completed: mapInt64(part.Time, "completed"),
	}

	if input, ok := part.State["input"]; ok {
		tool.Input = stringifyForDisplay(input)
	}
	if output, ok := part.State["content"]; ok {
		tool.Output = stringifyForDisplay(output)
	}
	if meta, ok := part.State["metadata"].(map[string]interface{}); ok {
		if diff, ok := meta["diff"].(string); ok {
			tool.Diff = diff
		}
	}

	return tool
}

// stringifyForDisplay renders arbitrary JSON values as a stable, readable string.
func stringifyForDisplay(value interface{}) string {
	if value == nil {
		return ""
	}
	if s, ok := value.(string); ok {
		return s
	}
	b, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	return string(b)
}

// parseLegacySyntheticTools extracts tool invocations from flattened synthetic
// messages that follow the "Called the <Name> tool with the following input: {...}"
// format. It returns the extracted tools and any text left over between them.
func parseLegacySyntheticTools(text string, messageID string) ([]ToolCall, string) {
	const prefix = "Called the "
	const suffix = " tool with the following input: "

	lines := strings.Split(text, "\n")
	tools := make([]ToolCall, 0)
	var leftover strings.Builder

	for i := 0; i < len(lines); {
		line := lines[i]
		marker := strings.Index(line, suffix)
		if !strings.HasPrefix(line, prefix) || marker == -1 {
			leftover.WriteString(line)
			leftover.WriteString("\n")
			i++
			continue
		}

		name := strings.TrimSpace(line[len(prefix):marker])
		input := strings.TrimSpace(line[marker+len(suffix):])
		i++

		var outputLines []string
		for i < len(lines) {
			next := lines[i]
			if strings.HasPrefix(next, prefix) && strings.Contains(next, suffix) {
				break
			}
			outputLines = append(outputLines, next)
			i++
		}

		tools = append(tools, ToolCall{
			ID:     fmt.Sprintf("%s:legacy-%d", messageID, len(tools)+1),
			Name:   name,
			Status: "completed",
			Input:  input,
			Output: strings.TrimSpace(strings.Join(outputLines, "\n")),
		})
	}

	return tools, strings.TrimSpace(leftover.String())
}

// asString returns the string representation of a value, handling nested fields.
func asString(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	case map[string]interface{}:
		if dir, ok := v["directory"].(string); ok {
			return dir
		}
	case *opencode2.Location_Info:
		return v.Directory
	case opencode2.Location_Info:
		return v.Directory
	default:
		if b, err := json.Marshal(value); err == nil {
			return string(b)
		}
	}
	return ""
}

// asInt64 returns the time value as milliseconds since epoch when possible.
func asInt64(value interface{}) int64 {
	switch v := value.(type) {
	case nil:
		return 0
	case float64:
		return int64(v)
	case int64:
		return v
	case int:
		return int64(v)
	case map[string]interface{}:
		if updated, ok := v["updated"].(float64); ok {
			return int64(updated)
		}
	case *opencode2.Project_Time:
		if v != nil {
			return toInt64(v.Updated)
		}
	case opencode2.Project_Time:
		return toInt64(v.Updated)
	}
	return 0
}

func mapInt64(value map[string]interface{}, key string) int64 {
	if value == nil {
		return 0
	}
	if raw, ok := value[key].(float64); ok {
		return int64(raw)
	}
	return 0
}

func toInt64(value interface{}) int64 {
	switch v := value.(type) {
	case float64:
		return int64(v)
	case int64:
		return v
	case int:
		return int64(v)
	}
	return 0
}
