package opencode2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client is an HTTP client for the OpenCode2 V2 API.
type Client struct {
	baseURL string
	http    *http.Client
}

// NewClient creates a new client for the OpenCode2 V2 API.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		http:    &http.Client{},
	}
}

// NewClientWithHTTP allows customizing the HTTP client.
func NewClientWithHTTP(baseURL string, hc *http.Client) *Client {
	return &Client{baseURL: strings.TrimRight(baseURL, "/"), http: hc}
}

// V2AgentList - List agents
// HTTP: GET /api/agent
func (c *Client) V2AgentList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/agent"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2AgentGet - Get agent
// HTTP: GET /api/agent/{agentID}
func (c *Client) V2AgentGet(ctx context.Context, agentID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/agent/%s", url.PathEscape(agentID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2CommandList - List commands
// HTTP: GET /api/command
func (c *Client) V2CommandList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/command"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ConfigGet - Get configuration
// HTTP: GET /api/config
func (c *Client) V2ConfigGet(ctx context.Context, query map[string]string) ([]interface{}, error) {
	path := "/api/config"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result []interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2CredentialUpdate - Update credential
// HTTP: PATCH /api/credential/{credentialID}
func (c *Client) V2CredentialUpdate(ctx context.Context, credentialID string, query map[string]string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/credential/%s", url.PathEscape(credentialID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "PATCH", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2CredentialRemove - Remove credential
// HTTP: DELETE /api/credential/{credentialID}
func (c *Client) V2CredentialRemove(ctx context.Context, credentialID string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/credential/%s", url.PathEscape(credentialID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2DebugLocationList - List loaded locations
// HTTP: GET /api/debug/location
func (c *Client) V2DebugLocationList(ctx context.Context) ([]interface{}, error) {
	path := "/api/debug/location"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result []interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2DebugLocationEvict - Evict a loaded location
// HTTP: DELETE /api/debug/location
func (c *Client) V2DebugLocationEvict(ctx context.Context, query map[string]string) (json.RawMessage, error) {
	path := "/api/debug/location"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2EventSubscribe - Subscribe to events
// HTTP: GET /api/event
func (c *Client) V2EventSubscribe(ctx context.Context) (json.RawMessage, error) {
	path := "/api/event"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2ExperimentalIntegrationWellknownAdd - Add wellknown integration
// HTTP: POST /api/experimental/integration/wellknown
func (c *Client) V2ExperimentalIntegrationWellknownAdd(ctx context.Context, query map[string]string, body interface{}) (json.RawMessage, error) {
	path := "/api/experimental/integration/wellknown"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2ExperimentalMigrationV1Status - Get V1 migration status
// HTTP: GET /api/experimental/migration/v1
func (c *Client) V2ExperimentalMigrationV1Status(ctx context.Context) (interface{}, error) {
	path := "/api/experimental/migration/v1"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return 0, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionLog - Read the session log
// HTTP: GET /api/experimental/session/{sessionID}/log
func (c *Client) V2SessionLog(ctx context.Context, sessionID string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/experimental/session/%s/log", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["after"]; ok {
		q.Set("after", v)
	}
	if v, ok := query["follow"]; ok {
		q.Set("follow", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2FormRequestList - List pending form requests
// HTTP: GET /api/form/request
func (c *Client) V2FormRequestList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/form/request"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2FsFind - Find files
// HTTP: GET /api/fs/find
func (c *Client) V2FsFind(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/fs/find"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	if v, ok := query["query"]; ok {
		q.Set("query", v)
	}
	if v, ok := query["type"]; ok {
		q.Set("type", v)
	}
	if v, ok := query["limit"]; ok {
		q.Set("limit", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2FsList - List directory
// HTTP: GET /api/fs/list
func (c *Client) V2FsList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/fs/list"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	if v, ok := query["path"]; ok {
		q.Set("path", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2FsRead - Read file
// HTTP: GET /api/fs/read/*
func (c *Client) V2FsRead(ctx context.Context, query map[string]string) (json.RawMessage, error) {
	path := "/api/fs/read/*"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2GenerateText - Generate text
// HTTP: POST /api/generate
func (c *Client) V2GenerateText(ctx context.Context, query map[string]string, body interface{}) (GenerateTextResponse, error) {
	path := "/api/generate"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return GenerateTextResponse{}, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return GenerateTextResponse{}, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return GenerateTextResponse{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return GenerateTextResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return GenerateTextResponse{}, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return GenerateTextResponse{}, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result GenerateTextResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return GenerateTextResponse{}, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2HealthGet - Check server health
// HTTP: GET /api/health
func (c *Client) V2HealthGet(ctx context.Context) (ServiceHealth, error) {
	path := "/api/health"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return ServiceHealth{}, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return ServiceHealth{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return ServiceHealth{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ServiceHealth{}, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ServiceHealth{}, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result ServiceHealth
	if err := json.Unmarshal(data, &result); err != nil {
		return ServiceHealth{}, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2IntegrationList - List integrations
// HTTP: GET /api/integration
func (c *Client) V2IntegrationList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/integration"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2IntegrationGet - Get integration
// HTTP: GET /api/integration/{integrationID}
func (c *Client) V2IntegrationGet(ctx context.Context, integrationID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/integration/%s", url.PathEscape(integrationID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2IntegrationCommandConnect - Begin command connection
// HTTP: POST /api/integration/{integrationID}/connect/command
func (c *Client) V2IntegrationCommandConnect(ctx context.Context, integrationID string, query map[string]string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/command", url.PathEscape(integrationID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2IntegrationCommandStatus - Get command attempt status
// HTTP: GET /api/integration/{integrationID}/connect/command/{attemptID}
func (c *Client) V2IntegrationCommandStatus(ctx context.Context, integrationID string, attemptID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/command/%s", url.PathEscape(integrationID), url.PathEscape(attemptID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2IntegrationCommandCancel - Cancel command connection
// HTTP: DELETE /api/integration/{integrationID}/connect/command/{attemptID}
func (c *Client) V2IntegrationCommandCancel(ctx context.Context, integrationID string, attemptID string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/command/%s", url.PathEscape(integrationID), url.PathEscape(attemptID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2IntegrationConnectKey - Connect with key
// HTTP: POST /api/integration/{integrationID}/connect/key
func (c *Client) V2IntegrationConnectKey(ctx context.Context, integrationID string, query map[string]string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/key", url.PathEscape(integrationID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2IntegrationOauthConnect - Begin OAuth connection
// HTTP: POST /api/integration/{integrationID}/connect/oauth
func (c *Client) V2IntegrationOauthConnect(ctx context.Context, integrationID string, query map[string]string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/oauth", url.PathEscape(integrationID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2IntegrationOauthStatus - Get OAuth attempt status
// HTTP: GET /api/integration/{integrationID}/connect/oauth/{attemptID}
func (c *Client) V2IntegrationOauthStatus(ctx context.Context, integrationID string, attemptID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/oauth/%s", url.PathEscape(integrationID), url.PathEscape(attemptID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2IntegrationOauthCancel - Cancel OAuth connection
// HTTP: DELETE /api/integration/{integrationID}/connect/oauth/{attemptID}
func (c *Client) V2IntegrationOauthCancel(ctx context.Context, integrationID string, attemptID string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/oauth/%s", url.PathEscape(integrationID), url.PathEscape(attemptID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2IntegrationOauthComplete - Complete OAuth connection
// HTTP: POST /api/integration/{integrationID}/connect/oauth/{attemptID}/complete
func (c *Client) V2IntegrationOauthComplete(ctx context.Context, integrationID string, attemptID string, query map[string]string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/integration/%s/connect/oauth/%s/complete", url.PathEscape(integrationID), url.PathEscape(attemptID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2LocationGet - Get location
// HTTP: GET /api/location
func (c *Client) V2LocationGet(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/api/location"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return 0, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2McpList - List MCP servers
// HTTP: GET /api/mcp
func (c *Client) V2McpList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/mcp"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2McpResourceCatalog - List MCP resources
// HTTP: GET /api/mcp/resource
func (c *Client) V2McpResourceCatalog(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/mcp/resource"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2McpAdd - Add MCP server
// HTTP: PUT /api/mcp/{server}
func (c *Client) V2McpAdd(ctx context.Context, server string, query map[string]string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/mcp/%s", url.PathEscape(server))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2McpRemove - Remove MCP server
// HTTP: DELETE /api/mcp/{server}
func (c *Client) V2McpRemove(ctx context.Context, server string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/mcp/%s", url.PathEscape(server))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2McpConnect - Connect MCP server
// HTTP: POST /api/mcp/{server}/connect
func (c *Client) V2McpConnect(ctx context.Context, server string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/mcp/%s/connect", url.PathEscape(server))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2McpDisconnect - Disconnect MCP server
// HTTP: POST /api/mcp/{server}/disconnect
func (c *Client) V2McpDisconnect(ctx context.Context, server string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/mcp/%s/disconnect", url.PathEscape(server))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2ModelList - List models
// HTTP: GET /api/model
func (c *Client) V2ModelList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/model"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ModelDefault - Get default model
// HTTP: GET /api/model/default
func (c *Client) V2ModelDefault(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/model/default"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PermissionRequestList - List pending permission requests
// HTTP: GET /api/permission/request
func (c *Client) V2PermissionRequestList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/permission/request"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PermissionSavedList - List saved permissions
// HTTP: GET /api/permission/saved
func (c *Client) V2PermissionSavedList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/permission/saved"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["projectID"]; ok {
		q.Set("projectID", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PermissionSavedRemove - Remove saved permission
// HTTP: DELETE /api/permission/saved/{id}
func (c *Client) V2PermissionSavedRemove(ctx context.Context, id string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/permission/saved/%s", url.PathEscape(id))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2PluginList - List plugins
// HTTP: GET /api/plugin
func (c *Client) V2PluginList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/plugin"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ProjectList - List projects
// HTTP: GET /api/project
func (c *Client) V2ProjectList(ctx context.Context) ([]Project, error) {
	path := "/api/project"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result []Project
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ProjectCurrent - Get current project
// HTTP: GET /api/project/current
func (c *Client) V2ProjectCurrent(ctx context.Context, query map[string]string) (interface{}, error) {
	path := "/api/project/current"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return 0, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ProviderList - List providers
// HTTP: GET /api/provider
func (c *Client) V2ProviderList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/provider"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ProviderGet - Get provider
// HTTP: GET /api/provider/{providerID}
func (c *Client) V2ProviderGet(ctx context.Context, providerID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/provider/%s", url.PathEscape(providerID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PtyList - List PTY sessions
// HTTP: GET /api/pty
func (c *Client) V2PtyList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/pty"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PtyCreate - Create PTY session
// HTTP: POST /api/pty
func (c *Client) V2PtyCreate(ctx context.Context, query map[string]string, body interface{}) (map[string]interface{}, error) {
	path := "/api/pty"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PtyGet - Get PTY session
// HTTP: GET /api/pty/{ptyID}
func (c *Client) V2PtyGet(ctx context.Context, ptyID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/pty/%s", url.PathEscape(ptyID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PtyUpdate - Update PTY session
// HTTP: PUT /api/pty/{ptyID}
func (c *Client) V2PtyUpdate(ctx context.Context, ptyID string, query map[string]string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/pty/%s", url.PathEscape(ptyID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PtyRemove - Remove PTY session
// HTTP: DELETE /api/pty/{ptyID}
func (c *Client) V2PtyRemove(ctx context.Context, ptyID string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/pty/%s", url.PathEscape(ptyID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2PtyConnect - Connect to PTY session
// HTTP: GET /api/pty/{ptyID}/connect
func (c *Client) V2PtyConnect(ctx context.Context, ptyID string, query map[string]string) (bool, error) {
	path := fmt.Sprintf("/api/pty/%s/connect", url.PathEscape(ptyID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return false, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location[directory]"]; ok {
		q.Set("location[directory]", v)
	}
	if v, ok := query["location[workspace]"]; ok {
		q.Set("location[workspace]", v)
	}
	if v, ok := query["cursor"]; ok {
		q.Set("cursor", v)
	}
	if v, ok := query["ticket"]; ok {
		q.Set("ticket", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return false, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return false, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result bool
	if err := json.Unmarshal(data, &result); err != nil {
		return false, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2PtyConnectToken - Create PTY WebSocket token
// HTTP: POST /api/pty/{ptyID}/connect-token
func (c *Client) V2PtyConnectToken(ctx context.Context, ptyID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/pty/%s/connect-token", url.PathEscape(ptyID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ReferenceList - List references
// HTTP: GET /api/reference
func (c *Client) V2ReferenceList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/reference"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ServerGet - Get server information
// HTTP: GET /api/server
func (c *Client) V2ServerGet(ctx context.Context) (map[string]interface{}, error) {
	path := "/api/server"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2HealthStop - Stop the managed server
// HTTP: POST /api/service/stop
func (c *Client) V2HealthStop(ctx context.Context, body interface{}) (ServiceStopResponse, error) {
	path := "/api/service/stop"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return ServiceStopResponse{}, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return ServiceStopResponse{}, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return ServiceStopResponse{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return ServiceStopResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ServiceStopResponse{}, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ServiceStopResponse{}, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result ServiceStopResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return ServiceStopResponse{}, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionList - List sessions
// HTTP: GET /api/session
func (c *Client) V2SessionList(ctx context.Context, query map[string]string) (SessionsResponse, error) {
	path := "/api/session"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return SessionsResponse{}, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["workspace"]; ok {
		q.Set("workspace", v)
	}
	if v, ok := query["limit"]; ok {
		q.Set("limit", v)
	}
	if v, ok := query["order"]; ok {
		q.Set("order", v)
	}
	if v, ok := query["search"]; ok {
		q.Set("search", v)
	}
	if v, ok := query["parentID"]; ok {
		q.Set("parentID", v)
	}
	if v, ok := query["directory"]; ok {
		q.Set("directory", v)
	}
	if v, ok := query["project"]; ok {
		q.Set("project", v)
	}
	if v, ok := query["subpath"]; ok {
		q.Set("subpath", v)
	}
	if v, ok := query["cursor"]; ok {
		q.Set("cursor", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return SessionsResponse{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return SessionsResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return SessionsResponse{}, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return SessionsResponse{}, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result SessionsResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return SessionsResponse{}, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionCreate - Create session
// HTTP: POST /api/session
func (c *Client) V2SessionCreate(ctx context.Context, body interface{}) (map[string]interface{}, error) {
	path := "/api/session"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionActive - List active sessions
// HTTP: GET /api/session/active
func (c *Client) V2SessionActive(ctx context.Context) (map[string]interface{}, error) {
	path := "/api/session/active"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionImport - Import session
// HTTP: POST /api/session/import
func (c *Client) V2SessionImport(ctx context.Context, body interface{}) (map[string]interface{}, error) {
	path := "/api/session/import"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionGet - Get session
// HTTP: GET /api/session/{sessionID}
func (c *Client) V2SessionGet(ctx context.Context, sessionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionRemove - Delete session
// HTTP: DELETE /api/session/{sessionID}
func (c *Client) V2SessionRemove(ctx context.Context, sessionID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionSwitchAgent - Switch session agent
// HTTP: POST /api/session/{sessionID}/agent
func (c *Client) V2SessionSwitchAgent(ctx context.Context, sessionID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/agent", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionBackground - Background blocking session tools
// HTTP: POST /api/session/{sessionID}/background
func (c *Client) V2SessionBackground(ctx context.Context, sessionID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/background", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionCommand - Run command
// HTTP: POST /api/session/{sessionID}/command
func (c *Client) V2SessionCommand(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/command", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionCompact - Compact session
// HTTP: POST /api/session/{sessionID}/compact
func (c *Client) V2SessionCompact(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/compact", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionContext - Get session context
// HTTP: GET /api/session/{sessionID}/context
func (c *Client) V2SessionContext(ctx context.Context, sessionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/context", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionExport - Export session
// HTTP: GET /api/session/{sessionID}/export
func (c *Client) V2SessionExport(ctx context.Context, sessionID string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/export", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["sanitize"]; ok {
		q.Set("sanitize", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionFork - Fork session
// HTTP: POST /api/session/{sessionID}/fork
func (c *Client) V2SessionFork(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/fork", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionFormList - List session forms
// HTTP: GET /api/session/{sessionID}/form
func (c *Client) V2SessionFormList(ctx context.Context, sessionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/form", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionFormCreate - Create session form
// HTTP: POST /api/session/{sessionID}/form
func (c *Client) V2SessionFormCreate(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/form", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionFormGet - Get session form
// HTTP: GET /api/session/{sessionID}/form/{formID}
func (c *Client) V2SessionFormGet(ctx context.Context, sessionID string, formID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/form/%s", url.PathEscape(sessionID), url.PathEscape(formID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionFormCancel - Cancel form
// HTTP: POST /api/session/{sessionID}/form/{formID}/cancel
func (c *Client) V2SessionFormCancel(ctx context.Context, sessionID string, formID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/form/%s/cancel", url.PathEscape(sessionID), url.PathEscape(formID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionFormReply - Reply to form
// HTTP: POST /api/session/{sessionID}/form/{formID}/reply
func (c *Client) V2SessionFormReply(ctx context.Context, sessionID string, formID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/form/%s/reply", url.PathEscape(sessionID), url.PathEscape(formID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionFormState - Get form state
// HTTP: GET /api/session/{sessionID}/form/{formID}/state
func (c *Client) V2SessionFormState(ctx context.Context, sessionID string, formID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/form/%s/state", url.PathEscape(sessionID), url.PathEscape(formID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionGenerate - Generate text from session context
// HTTP: POST /api/session/{sessionID}/generate
func (c *Client) V2SessionGenerate(ctx context.Context, sessionID string, body interface{}) (SessionGenerateResponse, error) {
	path := fmt.Sprintf("/api/session/%s/generate", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return SessionGenerateResponse{}, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return SessionGenerateResponse{}, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return SessionGenerateResponse{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return SessionGenerateResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return SessionGenerateResponse{}, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return SessionGenerateResponse{}, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result SessionGenerateResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return SessionGenerateResponse{}, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionInboxList - List session inbox
// HTTP: GET /api/session/{sessionID}/inbox
func (c *Client) V2SessionInboxList(ctx context.Context, sessionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/inbox", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionInboxCancel - Cancel inbox input
// HTTP: DELETE /api/session/{sessionID}/inbox/{inboxID}
func (c *Client) V2SessionInboxCancel(ctx context.Context, sessionID string, inboxID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/inbox/%s", url.PathEscape(sessionID), url.PathEscape(inboxID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionInboxQueue - Queue steered item
// HTTP: POST /api/session/{sessionID}/inbox/{inboxID}/queue
func (c *Client) V2SessionInboxQueue(ctx context.Context, sessionID string, inboxID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/inbox/%s/queue", url.PathEscape(sessionID), url.PathEscape(inboxID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionInboxSteer - Steer queued item
// HTTP: POST /api/session/{sessionID}/inbox/{inboxID}/steer
func (c *Client) V2SessionInboxSteer(ctx context.Context, sessionID string, inboxID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/inbox/%s/steer", url.PathEscape(sessionID), url.PathEscape(inboxID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionInstructionsEntryList - List instruction entries
// HTTP: GET /api/session/{sessionID}/instructions/entries
func (c *Client) V2SessionInstructionsEntryList(ctx context.Context, sessionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/instructions/entries", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionInstructionsEntryPut - Put instruction entry
// HTTP: PUT /api/session/{sessionID}/instructions/entries/{key}
func (c *Client) V2SessionInstructionsEntryPut(ctx context.Context, sessionID string, key string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/instructions/entries/%s", url.PathEscape(sessionID), url.PathEscape(key))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionInstructionsEntryRemove - Remove instruction entry
// HTTP: DELETE /api/session/{sessionID}/instructions/entries/{key}
func (c *Client) V2SessionInstructionsEntryRemove(ctx context.Context, sessionID string, key string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/instructions/entries/%s", url.PathEscape(sessionID), url.PathEscape(key))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionInterrupt - Interrupt session execution
// HTTP: POST /api/session/{sessionID}/interrupt
func (c *Client) V2SessionInterrupt(ctx context.Context, sessionID string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/interrupt", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["continue"]; ok {
		q.Set("continue", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2MessageList - Get session messages
// HTTP: GET /api/session/{sessionID}/message
func (c *Client) V2MessageList(ctx context.Context, sessionID string, query map[string]string) (SessionMessagesResponse, error) {
	path := fmt.Sprintf("/api/session/%s/message", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return SessionMessagesResponse{}, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["limit"]; ok {
		q.Set("limit", v)
	}
	if v, ok := query["order"]; ok {
		q.Set("order", v)
	}
	if v, ok := query["cursor"]; ok {
		q.Set("cursor", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return SessionMessagesResponse{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return SessionMessagesResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return SessionMessagesResponse{}, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return SessionMessagesResponse{}, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result SessionMessagesResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return SessionMessagesResponse{}, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionMessage - Get session message
// HTTP: GET /api/session/{sessionID}/message/{messageID}
func (c *Client) V2SessionMessage(ctx context.Context, sessionID string, messageID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/message/%s", url.PathEscape(sessionID), url.PathEscape(messageID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionSwitchModel - Switch session model
// HTTP: POST /api/session/{sessionID}/model
func (c *Client) V2SessionSwitchModel(ctx context.Context, sessionID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/model", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionMove - Move session
// HTTP: POST /api/session/{sessionID}/move
func (c *Client) V2SessionMove(ctx context.Context, sessionID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/move", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionPermissionCreate - Create permission request
// HTTP: POST /api/session/{sessionID}/permission
func (c *Client) V2SessionPermissionCreate(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/permission", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionPermissionList - List session permission requests
// HTTP: GET /api/session/{sessionID}/permission
func (c *Client) V2SessionPermissionList(ctx context.Context, sessionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/permission", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionPermissionGet - Get permission request
// HTTP: GET /api/session/{sessionID}/permission/{requestID}
func (c *Client) V2SessionPermissionGet(ctx context.Context, sessionID string, requestID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/permission/%s", url.PathEscape(sessionID), url.PathEscape(requestID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionPermissionReply - Reply to pending permission request
// HTTP: POST /api/session/{sessionID}/permission/{requestID}/reply
func (c *Client) V2SessionPermissionReply(ctx context.Context, sessionID string, requestID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/permission/%s/reply", url.PathEscape(sessionID), url.PathEscape(requestID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionPrompt - Send message
// HTTP: POST /api/session/{sessionID}/prompt
func (c *Client) V2SessionPrompt(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/prompt", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionRename - Rename session
// HTTP: POST /api/session/{sessionID}/rename
func (c *Client) V2SessionRename(ctx context.Context, sessionID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/rename", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionRevertClear - Clear staged revert
// HTTP: POST /api/session/{sessionID}/revert/clear
func (c *Client) V2SessionRevertClear(ctx context.Context, sessionID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/revert/clear", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionRevertCommit - Commit staged revert
// HTTP: POST /api/session/{sessionID}/revert/commit
func (c *Client) V2SessionRevertCommit(ctx context.Context, sessionID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/revert/commit", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionRevertStage - Stage session revert
// HTTP: POST /api/session/{sessionID}/revert/stage
func (c *Client) V2SessionRevertStage(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/revert/stage", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionShell - Run shell command
// HTTP: POST /api/session/{sessionID}/shell
func (c *Client) V2SessionShell(ctx context.Context, sessionID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/shell", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionSkill - Activate skill
// HTTP: POST /api/session/{sessionID}/skill
func (c *Client) V2SessionSkill(ctx context.Context, sessionID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/skill", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2SessionSynthetic - Add synthetic message
// HTTP: POST /api/session/{sessionID}/synthetic
func (c *Client) V2SessionSynthetic(ctx context.Context, sessionID string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/session/%s/synthetic", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SessionWait - Wait for session
// HTTP: POST /api/session/{sessionID}/wait
func (c *Client) V2SessionWait(ctx context.Context, sessionID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/session/%s/wait", url.PathEscape(sessionID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2ShellList - List running shell commands
// HTTP: GET /api/shell
func (c *Client) V2ShellList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/shell"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ShellCreate - Run shell command
// HTTP: POST /api/shell
func (c *Client) V2ShellCreate(ctx context.Context, query map[string]string, body interface{}) (map[string]interface{}, error) {
	path := "/api/shell"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ShellGet - Get shell command
// HTTP: GET /api/shell/{id}
func (c *Client) V2ShellGet(ctx context.Context, id string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/shell/%s", url.PathEscape(id))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ShellRemove - Remove shell command
// HTTP: DELETE /api/shell/{id}
func (c *Client) V2ShellRemove(ctx context.Context, id string, query map[string]string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/shell/%s", url.PathEscape(id))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2ShellOutput - Read shell output
// HTTP: GET /api/shell/{id}/output
func (c *Client) V2ShellOutput(ctx context.Context, id string, query map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/shell/%s/output", url.PathEscape(id))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	if v, ok := query["cursor"]; ok {
		q.Set("cursor", v)
	}
	if v, ok := query["limit"]; ok {
		q.Set("limit", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2ShellTimeout - Update shell timeout
// HTTP: PATCH /api/shell/{id}/timeout
func (c *Client) V2ShellTimeout(ctx context.Context, id string, query map[string]string, body interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/api/shell/%s/timeout", url.PathEscape(id))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "PATCH", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2SkillList - List skills
// HTTP: GET /api/skill
func (c *Client) V2SkillList(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/skill"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2VcsGet - VCS info
// HTTP: GET /api/vcs
func (c *Client) V2VcsGet(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/vcs"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2VcsDiff - VCS diff
// HTTP: GET /api/vcs/diff
func (c *Client) V2VcsDiff(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/vcs/diff"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	if v, ok := query["mode"]; ok {
		q.Set("mode", v)
	}
	if v, ok := query["context"]; ok {
		q.Set("context", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2VcsStatus - VCS status
// HTTP: GET /api/vcs/status
func (c *Client) V2VcsStatus(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/vcs/status"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2WebsearchQuery - Search the web
// HTTP: POST /api/websearch
func (c *Client) V2WebsearchQuery(ctx context.Context, query map[string]string, body interface{}) (map[string]interface{}, error) {
	path := "/api/websearch"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2WebsearchProviders - List web search providers
// HTTP: GET /api/websearch/provider
func (c *Client) V2WebsearchProviders(ctx context.Context, query map[string]string) (map[string]interface{}, error) {
	path := "/api/websearch/provider"
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	q := u.Query()
	if v, ok := query["location"]; ok {
		q.Set("location", v)
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2WorktreeList - List worktrees
// HTTP: GET /api/worktree/{projectID}
func (c *Client) V2WorktreeList(ctx context.Context, projectID string) (interface{}, error) {
	path := fmt.Sprintf("/api/worktree/%s", url.PathEscape(projectID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return 0, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2WorktreeCreate - Create worktree
// HTTP: POST /api/worktree/{projectID}
func (c *Client) V2WorktreeCreate(ctx context.Context, projectID string, body interface{}) (interface{}, error) {
	path := fmt.Sprintf("/api/worktree/%s", url.PathEscape(projectID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return 0, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return 0, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}
	return result, nil
}

// V2WorktreeRemove - Remove worktree
// HTTP: DELETE /api/worktree/{projectID}
func (c *Client) V2WorktreeRemove(ctx context.Context, projectID string, body interface{}) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/worktree/%s", url.PathEscape(projectID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	bodyReader := io.Reader(nil)
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = strings.NewReader(string(b))
	}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}

// V2WorktreeRefresh - Refresh worktrees
// HTTP: POST /api/worktree/{projectID}/refresh
func (c *Client) V2WorktreeRefresh(ctx context.Context, projectID string) (json.RawMessage, error) {
	path := fmt.Sprintf("/api/worktree/%s/refresh", url.PathEscape(projectID))
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}

	return data, nil
}
