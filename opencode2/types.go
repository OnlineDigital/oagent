package opencode2

type Agent_Info struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Model       interface{} `json:"model,omitempty"`
	Request     interface{} `json:"request"`
	System      string      `json:"system,omitempty"`
	Description string      `json:"description,omitempty"`
	Mode        string      `json:"mode"`
	Hidden      bool        `json:"hidden"`
	Color       interface{} `json:"color,omitempty"`
	Steps       interface{} `json:"steps,omitempty"`
	Permissions interface{} `json:"permissions"`
}

type AgentNotFoundError struct {
	Tag     string `json:"_tag"`
	AgentID string `json:"agentID"`
	Message string `json:"message"`
}

type Command_Info struct {
	Name        string      `json:"name"`
	Template    string      `json:"template"`
	Description string      `json:"description,omitempty"`
	Agent       string      `json:"agent,omitempty"`
	Model       interface{} `json:"model,omitempty"`
	Subtask     bool        `json:"subtask,omitempty"`
}

type CommandEvaluationError struct {
	Tag     string `json:"_tag"`
	Command string `json:"command"`
	Message string `json:"message"`
}

type CommandNotFoundError struct {
	Tag     string `json:"_tag"`
	Command string `json:"command"`
	Message string `json:"message"`
}

type Config_Agent struct {
	Model       interface{}            `json:"model,omitempty"`
	Request     map[string]interface{} `json:"request,omitempty"`
	System      string                 `json:"system,omitempty"`
	Description string                 `json:"description,omitempty"`
	Mode        string                 `json:"mode,omitempty"`
	Hidden      bool                   `json:"hidden,omitempty"`
	Color       interface{}            `json:"color,omitempty"`
	Steps       interface{}            `json:"steps,omitempty"`
	Disabled    bool                   `json:"disabled,omitempty"`
	Permissions interface{}            `json:"permissions,omitempty"`
}

type Config_AgentsDirectory struct {
	F_type string `json:"type"`
	Path   string `json:"path"`
}

type Config_ClaudeDirectory struct {
	F_type string `json:"type"`
	Path   string `json:"path"`
}

type Config_Command struct {
	Template    string      `json:"template"`
	Description string      `json:"description,omitempty"`
	Agent       string      `json:"agent,omitempty"`
	Model       interface{} `json:"model,omitempty"`
	Subtask     bool        `json:"subtask,omitempty"`
}

type Config_Directory struct {
	F_type string `json:"type"`
	Path   string `json:"path"`
}

type Config_Document struct {
	F_type string      `json:"type"`
	Path   string      `json:"path,omitempty"`
	Info   interface{} `json:"info"`
}

type Config_Formatter_Entry struct {
	Disabled    bool                   `json:"disabled,omitempty"`
	Command     []string               `json:"command,omitempty"`
	Environment map[string]interface{} `json:"environment,omitempty"`
	Extensions  []string               `json:"extensions,omitempty"`
}

type Config_Info struct {
	Schema       string                 `json:"$schema,omitempty"`
	Shell        string                 `json:"shell,omitempty"`
	Model        interface{}            `json:"model,omitempty"`
	DefaultAgent string                 `json:"default_agent,omitempty"`
	Autoupdate   interface{}            `json:"autoupdate,omitempty"`
	Share        string                 `json:"share,omitempty"`
	Enterprise   map[string]interface{} `json:"enterprise,omitempty"`
	Username     string                 `json:"username,omitempty"`
	Permissions  interface{}            `json:"permissions,omitempty"`
	Agents       map[string]interface{} `json:"agents,omitempty"`
	Snapshots    bool                   `json:"snapshots,omitempty"`
	Watcher      map[string]interface{} `json:"watcher,omitempty"`
	Formatter    interface{}            `json:"formatter,omitempty"`
	Lsp          interface{}            `json:"lsp,omitempty"`
	Media        map[string]interface{} `json:"media,omitempty"`
	ToolOutput   map[string]interface{} `json:"tool_output,omitempty"`
	Mcp          map[string]interface{} `json:"mcp,omitempty"`
	Compaction   map[string]interface{} `json:"compaction,omitempty"`
	Skills       []string               `json:"skills,omitempty"`
	Commands     map[string]interface{} `json:"commands,omitempty"`
	Instructions []string               `json:"instructions,omitempty"`
	References   map[string]interface{} `json:"references,omitempty"`
	Websearch    interface{}            `json:"websearch,omitempty"`
	Plugins      []interface{}          `json:"plugins,omitempty"`
	Warming      interface{}            `json:"warming,omitempty"`
	Providers    map[string]interface{} `json:"providers,omitempty"`
	Experimental map[string]interface{} `json:"experimental,omitempty"`
}

type Config_LSP_Server struct {
	Command        []string               `json:"command"`
	Extensions     []string               `json:"extensions,omitempty"`
	Disabled       bool                   `json:"disabled,omitempty"`
	Env            map[string]interface{} `json:"env,omitempty"`
	Initialization map[string]interface{} `json:"initialization,omitempty"`
}

type Config_Model struct {
	ModelID       string                   `json:"modelID,omitempty"`
	Family        string                   `json:"family,omitempty"`
	Name          string                   `json:"name,omitempty"`
	Compatibility interface{}              `json:"compatibility,omitempty"`
	F_package     string                   `json:"package,omitempty"`
	Settings      map[string]interface{}   `json:"settings,omitempty"`
	Headers       map[string]interface{}   `json:"headers,omitempty"`
	Body          map[string]interface{}   `json:"body,omitempty"`
	Capabilities  interface{}              `json:"capabilities,omitempty"`
	Variants      []map[string]interface{} `json:"variants,omitempty"`
	Cost          interface{}              `json:"cost,omitempty"`
	Disabled      bool                     `json:"disabled,omitempty"`
	Limit         map[string]interface{}   `json:"limit,omitempty"`
}

type Config_Model_Cost struct {
	Tier   map[string]interface{} `json:"tier,omitempty"`
	Input  interface{}            `json:"input"`
	Output interface{}            `json:"output"`
	Cache  map[string]interface{} `json:"cache,omitempty"`
}

type Config_Plugin_Entry struct {
	F_package string                 `json:"package"`
	Options   map[string]interface{} `json:"options,omitempty"`
}

type Config_Provider struct {
	Name      string                 `json:"name,omitempty"`
	Env       []string               `json:"env,omitempty"`
	F_package string                 `json:"package,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	Headers   map[string]interface{} `json:"headers,omitempty"`
	Body      map[string]interface{} `json:"body,omitempty"`
	Models    map[string]interface{} `json:"models,omitempty"`
}

type Config_Reference_Git struct {
	Repository  string `json:"repository"`
	Branch      string `json:"branch,omitempty"`
	Description string `json:"description,omitempty"`
	Hidden      bool   `json:"hidden,omitempty"`
}

type Config_Reference_Local struct {
	Path        string `json:"path"`
	Description string `json:"description,omitempty"`
	Hidden      bool   `json:"hidden,omitempty"`
}

type Config_Warming struct {
	Prompt   string `json:"prompt,omitempty"`
	Interval string `json:"interval,omitempty"`
	Duration string `json:"duration,omitempty"`
}

type ConfigWebSearch_Info struct {
	Provider interface{} `json:"provider"`
}

type ConflictError struct {
	Tag      string      `json:"_tag"`
	Message  string      `json:"message"`
	Resource interface{} `json:"resource,omitempty"`
}

type Connection_CredentialInfo struct {
	F_type string `json:"type"`
	Id     string `json:"id"`
	Label  string `json:"label"`
}

type Connection_EnvInfo struct {
	F_type string `json:"type"`
	Name   string `json:"name"`
}

type EventLog_Synced struct {
	F_type      string      `json:"type"`
	AggregateID string      `json:"aggregateID"`
	Seq         interface{} `json:"seq,omitempty"`
}

type FileDiff_Info struct {
	File      string      `json:"file"`
	Patch     string      `json:"patch"`
	Additions interface{} `json:"additions"`
	Deletions interface{} `json:"deletions"`
	Status    string      `json:"status"`
}

type FileSystem_Entry struct {
	Path   string `json:"path"`
	F_type string `json:"type"`
}

type ForbiddenError struct {
	Tag     string `json:"_tag"`
	Message string `json:"message"`
}

type Form_Answer map[string]interface{}

type Form_Answer1 map[string]interface{}

type Form_BooleanField struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	F_default   bool          `json:"default,omitempty"`
}

type Form_BooleanField1 struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	F_default   bool          `json:"default,omitempty"`
}

type Form_CreatePayload struct {
	Id       interface{} `json:"id,omitempty"`
	Title    string      `json:"title"`
	Metadata interface{} `json:"metadata,omitempty"`
	Fields   interface{} `json:"fields"`
}

type Form_ExternalField struct {
	Key         string `json:"key"`
	F_type      string `json:"type"`
	Url         string `json:"url"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Form_Info struct {
	Id        interface{} `json:"id"`
	SessionID string      `json:"sessionID"`
	Title     string      `json:"title"`
	Metadata  interface{} `json:"metadata,omitempty"`
	Fields    interface{} `json:"fields"`
}

type Form_Info1 struct {
	Id        interface{} `json:"id"`
	SessionID string      `json:"sessionID"`
	Title     string      `json:"title"`
	Metadata  interface{} `json:"metadata,omitempty"`
	Fields    interface{} `json:"fields"`
}

type Form_IntegerField struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Minimum     interface{}   `json:"minimum,omitempty"`
	Maximum     interface{}   `json:"maximum,omitempty"`
	F_default   interface{}   `json:"default,omitempty"`
}

type Form_IntegerField1 struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Minimum     interface{}   `json:"minimum,omitempty"`
	Maximum     interface{}   `json:"maximum,omitempty"`
	F_default   interface{}   `json:"default,omitempty"`
}

type Form_Metadata map[string]interface{}

type Form_Metadata1 map[string]interface{}

type Form_MultiselectField struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Options     []interface{} `json:"options"`
	MinItems    interface{}   `json:"minItems,omitempty"`
	MaxItems    interface{}   `json:"maxItems,omitempty"`
	Custom      bool          `json:"custom,omitempty"`
	F_default   []string      `json:"default,omitempty"`
}

type Form_MultiselectField1 struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Options     []interface{} `json:"options"`
	MinItems    interface{}   `json:"minItems,omitempty"`
	MaxItems    interface{}   `json:"maxItems,omitempty"`
	Custom      bool          `json:"custom,omitempty"`
	F_default   []string      `json:"default,omitempty"`
}

type Form_NumberField struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Minimum     interface{}   `json:"minimum,omitempty"`
	Maximum     interface{}   `json:"maximum,omitempty"`
	F_default   interface{}   `json:"default,omitempty"`
}

type Form_NumberField1 struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Minimum     interface{}   `json:"minimum,omitempty"`
	Maximum     interface{}   `json:"maximum,omitempty"`
	F_default   interface{}   `json:"default,omitempty"`
}

type Form_Option struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
}

type Form_Reply struct {
	Answer interface{} `json:"answer"`
}

type Form_StringField struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Format      string        `json:"format,omitempty"`
	MinLength   interface{}   `json:"minLength,omitempty"`
	MaxLength   interface{}   `json:"maxLength,omitempty"`
	Pattern     string        `json:"pattern,omitempty"`
	Placeholder string        `json:"placeholder,omitempty"`
	F_default   string        `json:"default,omitempty"`
	Options     []interface{} `json:"options,omitempty"`
	Custom      bool          `json:"custom,omitempty"`
}

type Form_StringField1 struct {
	Key         string        `json:"key"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	When        []interface{} `json:"when,omitempty"`
	F_type      string        `json:"type"`
	Format      string        `json:"format,omitempty"`
	MinLength   interface{}   `json:"minLength,omitempty"`
	MaxLength   interface{}   `json:"maxLength,omitempty"`
	Pattern     string        `json:"pattern,omitempty"`
	Placeholder string        `json:"placeholder,omitempty"`
	F_default   string        `json:"default,omitempty"`
	Options     []interface{} `json:"options,omitempty"`
	Custom      bool          `json:"custom,omitempty"`
}

type Form_When struct {
	Key   string      `json:"key"`
	Op    string      `json:"op"`
	Value interface{} `json:"value"`
}

type Form_When1 struct {
	Key   string      `json:"key"`
	Op    string      `json:"op"`
	Value interface{} `json:"value"`
}

type FormAlreadySettledError struct {
	Tag     string `json:"_tag"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

type FormInvalidAnswerError struct {
	Tag     string `json:"_tag"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

type FormNotFoundError struct {
	Tag     string `json:"_tag"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

type GenerateTextResponse struct {
	Data map[string]interface{} `json:"data"`
}

type InstructionEntry_Info struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

type InstructionEntry_Key struct {
}

type InstructionEntryValueTooLargeError struct {
	Tag         string `json:"_tag"`
	ActualBytes int    `json:"actualBytes"`
	MaxBytes    int    `json:"maxBytes"`
	Message     string `json:"message"`
}

type Integration_Attempt struct {
	AttemptID    string                 `json:"attemptID"`
	Url          string                 `json:"url"`
	Instructions string                 `json:"instructions"`
	Mode         string                 `json:"mode"`
	Time         map[string]interface{} `json:"time"`
}

type Integration_CommandAttempt struct {
	AttemptID string                 `json:"attemptID"`
	Time      map[string]interface{} `json:"time"`
}

type Integration_CommandMethod struct {
	Id      string   `json:"id"`
	F_type  string   `json:"type"`
	Label   string   `json:"label"`
	Command []string `json:"command"`
}

type Integration_EnvMethod struct {
	F_type string   `json:"type"`
	Names  []string `json:"names"`
}

type Integration_Info struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Methods     []interface{} `json:"methods"`
	Connections []interface{} `json:"connections"`
}

type Integration_KeyMethod struct {
	F_type string      `json:"type"`
	Label  string      `json:"label,omitempty"`
	Form   interface{} `json:"form,omitempty"`
}

type Integration_OAuthMethod struct {
	Id     string      `json:"id"`
	F_type string      `json:"type"`
	Label  string      `json:"label"`
	Form   interface{} `json:"form,omitempty"`
}

type InvalidCursorError struct {
	Tag     string `json:"_tag"`
	Message string `json:"message"`
}

type InvalidRequestError struct {
	Tag     string      `json:"_tag"`
	Message string      `json:"message"`
	Kind    interface{} `json:"kind,omitempty"`
	Field   interface{} `json:"field,omitempty"`
}

type InvalidRequestError1 struct {
	Tag     string      `json:"_tag"`
	Message string      `json:"message"`
	Kind    interface{} `json:"kind,omitempty"`
	Field   interface{} `json:"field,omitempty"`
}

type Location_Info struct {
	Directory   string                 `json:"directory"`
	WorkspaceID interface{}            `json:"workspaceID,omitempty"`
	Project     map[string]interface{} `json:"project"`
}

type Location_Ref struct {
	Directory   string      `json:"directory"`
	WorkspaceID interface{} `json:"workspaceID,omitempty"`
}

type Mcp_LocalConfig struct {
	F_type      string                 `json:"type"`
	Command     []string               `json:"command"`
	Cwd         string                 `json:"cwd,omitempty"`
	Environment map[string]interface{} `json:"environment,omitempty"`
	Disabled    bool                   `json:"disabled,omitempty"`
	Codemode    bool                   `json:"codemode,omitempty"`
	Timeout     map[string]interface{} `json:"timeout,omitempty"`
}

type Mcp_OAuthConfig struct {
	ClientId     string      `json:"client_id,omitempty"`
	ClientSecret string      `json:"client_secret,omitempty"`
	Scope        string      `json:"scope,omitempty"`
	CallbackPort interface{} `json:"callback_port,omitempty"`
	RedirectUri  string      `json:"redirect_uri,omitempty"`
}

type Mcp_RemoteConfig struct {
	F_type   string                 `json:"type"`
	Url      string                 `json:"url"`
	Headers  map[string]interface{} `json:"headers,omitempty"`
	Oauth    interface{}            `json:"oauth,omitempty"`
	Disabled bool                   `json:"disabled,omitempty"`
	Codemode bool                   `json:"codemode,omitempty"`
	Timeout  map[string]interface{} `json:"timeout,omitempty"`
}

type Mcp_Resource struct {
	Server      string `json:"server"`
	Name        string `json:"name"`
	Uri         string `json:"uri"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
}

type Mcp_ResourceCatalog struct {
	Resources []interface{} `json:"resources"`
	Templates []interface{} `json:"templates"`
}

type Mcp_ResourceTemplate struct {
	Server      string `json:"server"`
	Name        string `json:"name"`
	UriTemplate string `json:"uriTemplate"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
}

type Mcp_Server struct {
	Name          string      `json:"name"`
	Status        interface{} `json:"status"`
	IntegrationID string      `json:"integrationID,omitempty"`
}

type Mcp_Status_Connected struct {
	Status string `json:"status"`
}

type Mcp_Status_Disabled struct {
	Status string `json:"status"`
}

type Mcp_Status_Failed struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type Mcp_Status_NeedsAuth struct {
	Status string `json:"status"`
}

type Mcp_Status_Pending struct {
	Status string `json:"status"`
}

type McpServerNotFoundError struct {
	Tag     string `json:"_tag"`
	Server  string `json:"server"`
	Message string `json:"message"`
}

type MessageNotFoundError struct {
	Tag       string `json:"_tag"`
	SessionID string `json:"sessionID"`
	MessageID string `json:"messageID"`
	Message   string `json:"message"`
}

type Model_Capabilities struct {
	Tools  bool     `json:"tools"`
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

type Model_Compatibility struct {
	ReasoningField      interface{} `json:"reasoningField,omitempty"`
	MaxTokensField      interface{} `json:"maxTokensField,omitempty"`
	RequireFinishReason bool        `json:"requireFinishReason,omitempty"`
}

type Model_Cost struct {
	Tier   map[string]interface{} `json:"tier,omitempty"`
	Input  interface{}            `json:"input"`
	Output interface{}            `json:"output"`
	Cache  map[string]interface{} `json:"cache"`
}

type Model_Info struct {
	Id            string                 `json:"id"`
	ModelID       string                 `json:"modelID"`
	ProviderID    string                 `json:"providerID"`
	Family        string                 `json:"family,omitempty"`
	Name          string                 `json:"name"`
	Compatibility interface{}            `json:"compatibility,omitempty"`
	F_package     string                 `json:"package,omitempty"`
	Settings      map[string]interface{} `json:"settings,omitempty"`
	Headers       map[string]interface{} `json:"headers,omitempty"`
	Body          map[string]interface{} `json:"body,omitempty"`
	Capabilities  interface{}            `json:"capabilities"`
	Variants      []interface{}          `json:"variants"`
	Time          map[string]interface{} `json:"time"`
	Cost          []interface{}          `json:"cost"`
	Status        string                 `json:"status"`
	Enabled       bool                   `json:"enabled"`
	Limit         map[string]interface{} `json:"limit"`
}

type Model_Ref struct {
	Id         string `json:"id"`
	ProviderID string `json:"providerID"`
	Variant    string `json:"variant,omitempty"`
}

type Model_Variant struct {
	Id       string                 `json:"id"`
	Settings map[string]interface{} `json:"settings,omitempty"`
	Headers  map[string]interface{} `json:"headers,omitempty"`
	Body     map[string]interface{} `json:"body,omitempty"`
}

type Permission_Request struct {
	Id        interface{}            `json:"id"`
	SessionID interface{}            `json:"sessionID"`
	Action    string                 `json:"action"`
	Resources []string               `json:"resources"`
	Save      []string               `json:"save,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Source    interface{}            `json:"source,omitempty"`
}

type Permission_Rule struct {
	Action   string      `json:"action"`
	Resource string      `json:"resource"`
	Effect   interface{} `json:"effect"`
}

type PermissionNotFoundError struct {
	Tag       string `json:"_tag"`
	RequestID string `json:"requestID"`
	Message   string `json:"message"`
}

type PermissionSaved_Info struct {
	Id        string `json:"id"`
	ProjectID string `json:"projectID"`
	Action    string `json:"action"`
	Resource  string `json:"resource"`
}

type Plugin_Info struct {
	Id string `json:"id"`
}

type Project struct {
	Id        string      `json:"id"`
	Canonical string      `json:"canonical"`
	Vcs       interface{} `json:"vcs,omitempty"`
	Name      string      `json:"name,omitempty"`
	Icon      interface{} `json:"icon,omitempty"`
	Commands  interface{} `json:"commands,omitempty"`
	Time      interface{} `json:"time"`
	Sandboxes []string    `json:"sandboxes"`
}

type Project_Commands struct {
	Start string `json:"start,omitempty"`
}

type Project_Current struct {
	Id        string `json:"id"`
	Directory string `json:"directory"`
	Canonical string `json:"canonical"`
}

type Project_Icon struct {
	Url      string `json:"url,omitempty"`
	Override string `json:"override,omitempty"`
	Color    string `json:"color,omitempty"`
}

type Project_Time struct {
	Created     interface{} `json:"created"`
	Updated     interface{} `json:"updated"`
	Initialized interface{} `json:"initialized,omitempty"`
}

type Prompt_AgentAttachment struct {
	Name    string      `json:"name"`
	Mention interface{} `json:"mention,omitempty"`
}

type Prompt_Base64 struct {
}

type Prompt_FileAttachment struct {
	Data        interface{} `json:"data"`
	Mime        string      `json:"mime"`
	Source      interface{} `json:"source"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Mention     interface{} `json:"mention,omitempty"`
}

type Prompt_Mention struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Text  string  `json:"text"`
}

type Prompt_SkillAttachment struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Text    string      `json:"text"`
	Mention interface{} `json:"mention,omitempty"`
}

type PromptInput_FileAttachment struct {
	Uri         string      `json:"uri"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Mention     interface{} `json:"mention,omitempty"`
}

type PromptInput_SkillAttachment struct {
	Id      string      `json:"id"`
	Mention interface{} `json:"mention,omitempty"`
}

type Provider_Info struct {
	Id            string                 `json:"id"`
	IntegrationID string                 `json:"integrationID,omitempty"`
	Name          string                 `json:"name"`
	Activation    string                 `json:"activation"`
	F_package     string                 `json:"package"`
	Settings      map[string]interface{} `json:"settings,omitempty"`
	Headers       map[string]interface{} `json:"headers,omitempty"`
	Body          map[string]interface{} `json:"body,omitempty"`
}

type Provider_Request struct {
	Settings interface{}            `json:"settings"`
	Headers  map[string]interface{} `json:"headers"`
	Body     map[string]interface{} `json:"body"`
}

type Provider_Settings map[string]interface{}

type ProviderNotFoundError struct {
	Tag        string `json:"_tag"`
	ProviderID string `json:"providerID"`
	Message    string `json:"message"`
}

type Pty struct {
	Id       interface{} `json:"id"`
	Title    string      `json:"title"`
	Command  string      `json:"command"`
	Args     []string    `json:"args"`
	Cwd      string      `json:"cwd"`
	Status   string      `json:"status"`
	Pid      interface{} `json:"pid"`
	ExitCode interface{} `json:"exitCode,omitempty"`
}

type PtyNotFoundError struct {
	Tag     string `json:"_tag"`
	PtyID   string `json:"ptyID"`
	Message string `json:"message"`
}

type PtyTicket_ConnectToken struct {
	Ticket    string      `json:"ticket"`
	ExpiresIn interface{} `json:"expires_in"`
}

type Reference_GitSource struct {
	F_type      string `json:"type"`
	Repository  string `json:"repository"`
	Branch      string `json:"branch,omitempty"`
	Description string `json:"description,omitempty"`
	Hidden      bool   `json:"hidden,omitempty"`
}

type Reference_Info struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Description string      `json:"description,omitempty"`
	Hidden      bool        `json:"hidden,omitempty"`
	Source      interface{} `json:"source"`
}

type Reference_LocalSource struct {
	F_type      string `json:"type"`
	Path        string `json:"path"`
	Description string `json:"description,omitempty"`
	Hidden      bool   `json:"hidden,omitempty"`
}

type ServiceHealth struct {
	Healthy bool        `json:"healthy"`
	Version string      `json:"version"`
	Pid     interface{} `json:"pid"`
}

type ServiceStopRequest struct {
	InstanceID string `json:"instanceID"`
}

type ServiceStopResponse struct {
	Accepted bool `json:"accepted"`
}

type ServiceUnavailableError struct {
	Tag     string      `json:"_tag"`
	Message string      `json:"message"`
	Service interface{} `json:"service,omitempty"`
}

type Session_Inbox_Compaction struct {
	Id          interface{} `json:"id"`
	SessionID   interface{} `json:"sessionID"`
	TimeCreated float64     `json:"timeCreated"`
	F_type      string      `json:"type"`
	Payload     interface{} `json:"payload"`
	Delivery    interface{} `json:"delivery"`
}

type Session_Inbox_Move struct {
	Id          interface{} `json:"id"`
	SessionID   interface{} `json:"sessionID"`
	TimeCreated float64     `json:"timeCreated"`
	F_type      string      `json:"type"`
	Payload     interface{} `json:"payload"`
	Delivery    interface{} `json:"delivery"`
}

type Session_Inbox_MovePayload struct {
	Location  interface{} `json:"location"`
	ProjectID string      `json:"projectID"`
	Subpath   string      `json:"subpath,omitempty"`
}

type Session_Inbox_Synthetic struct {
	Id          interface{} `json:"id"`
	SessionID   interface{} `json:"sessionID"`
	TimeCreated float64     `json:"timeCreated"`
	F_type      string      `json:"type"`
	Payload     interface{} `json:"payload"`
	Delivery    interface{} `json:"delivery"`
}

type Session_Inbox_SyntheticPayload struct {
	Text        string                 `json:"text"`
	Description string                 `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type Session_Inbox_SyntheticPayload1 struct {
	Text        string                 `json:"text"`
	Description string                 `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type Session_Inbox_User struct {
	Id          interface{} `json:"id"`
	SessionID   interface{} `json:"sessionID"`
	TimeCreated float64     `json:"timeCreated"`
	F_type      string      `json:"type"`
	Payload     interface{} `json:"payload"`
	Delivery    interface{} `json:"delivery"`
}

type Session_Inbox_UserPayload struct {
	Text     string                 `json:"text"`
	Files    []interface{}          `json:"files,omitempty"`
	Agents   []interface{}          `json:"agents,omitempty"`
	Skills   []interface{}          `json:"skills,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Session_Inbox_UserPayload1 struct {
	Text     string                 `json:"text"`
	Files    []interface{}          `json:"files,omitempty"`
	Agents   []interface{}          `json:"agents,omitempty"`
	Skills   []interface{}          `json:"skills,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Session_Info struct {
	Id        interface{}            `json:"id"`
	ParentID  interface{}            `json:"parentID,omitempty"`
	Fork      map[string]interface{} `json:"fork,omitempty"`
	ProjectID string                 `json:"projectID"`
	Agent     string                 `json:"agent,omitempty"`
	Model     interface{}            `json:"model,omitempty"`
	Cost      interface{}            `json:"cost"`
	Tokens    interface{}            `json:"tokens"`
	Time      map[string]interface{} `json:"time"`
	Title     string                 `json:"title,omitempty"`
	Location  interface{}            `json:"location"`
	Subpath   string                 `json:"subpath,omitempty"`
	Revert    interface{}            `json:"revert,omitempty"`
}

type Session_Message_AgentSelected struct {
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	F_type   string                 `json:"type"`
	Agent    string                 `json:"agent"`
	Previous string                 `json:"previous,omitempty"`
}

type Session_Message_Assistant struct {
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	F_type   string                 `json:"type"`
	Agent    string                 `json:"agent"`
	Model    interface{}            `json:"model"`
	Content  []interface{}          `json:"content"`
	Snapshot map[string]interface{} `json:"snapshot,omitempty"`
	Finish   string                 `json:"finish,omitempty"`
	Cost     interface{}            `json:"cost,omitempty"`
	Tokens   interface{}            `json:"tokens,omitempty"`
	Error    interface{}            `json:"error,omitempty"`
	Retry    interface{}            `json:"retry,omitempty"`
}

type Session_Message_Assistant_Reasoning struct {
	F_type string                 `json:"type"`
	Text   string                 `json:"text"`
	State  interface{}            `json:"state,omitempty"`
	Time   map[string]interface{} `json:"time,omitempty"`
}

type Session_Message_Assistant_Retry struct {
	Attempt interface{} `json:"attempt"`
	At      float64     `json:"at"`
	Error   interface{} `json:"error"`
}

type Session_Message_Assistant_Text struct {
	F_type string      `json:"type"`
	Text   string      `json:"text"`
	State  interface{} `json:"state,omitempty"`
}

type Session_Message_Assistant_Tool struct {
	F_type              string                 `json:"type"`
	Id                  string                 `json:"id"`
	Name                string                 `json:"name"`
	Executed            bool                   `json:"executed,omitempty"`
	ProviderState       interface{}            `json:"providerState,omitempty"`
	ProviderResultState interface{}            `json:"providerResultState,omitempty"`
	State               interface{}            `json:"state"`
	Time                map[string]interface{} `json:"time"`
}

type Session_Message_Compaction_Completed struct {
	F_type   string                 `json:"type"`
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	Status   string                 `json:"status"`
	Reason   string                 `json:"reason"`
	Summary  string                 `json:"summary"`
	Recent   string                 `json:"recent"`
}

type Session_Message_Compaction_Failed struct {
	F_type   string                 `json:"type"`
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	Status   string                 `json:"status"`
	Reason   string                 `json:"reason"`
	Error    interface{}            `json:"error"`
}

type Session_Message_Compaction_Running struct {
	F_type   string                 `json:"type"`
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	Status   string                 `json:"status"`
	Reason   string                 `json:"reason"`
	Summary  string                 `json:"summary"`
	Recent   string                 `json:"recent"`
}

type Session_Message_LocationSwitched struct {
	Id        interface{}            `json:"id"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Time      map[string]interface{} `json:"time"`
	F_type    string                 `json:"type"`
	Location  interface{}            `json:"location"`
	ProjectID string                 `json:"projectID,omitempty"`
	Subpath   string                 `json:"subpath,omitempty"`
	Previous  map[string]interface{} `json:"previous,omitempty"`
}

type Session_Message_ModelSelected struct {
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	F_type   string                 `json:"type"`
	Model    interface{}            `json:"model"`
	Previous interface{}            `json:"previous,omitempty"`
}

type Session_Message_ProviderState map[string]interface{}

type Session_Message_ProviderState4 map[string]interface{}

type Session_Message_ProviderState5 map[string]interface{}

type Session_Message_ProviderState6 map[string]interface{}

type Session_Message_ProviderState7 map[string]interface{}

type Session_Message_ProviderState8 map[string]interface{}

type Session_Message_ProviderState9 map[string]interface{}

type Session_Message_Shell struct {
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	F_type   string                 `json:"type"`
	ShellID  interface{}            `json:"shellID"`
	Command  string                 `json:"command"`
	Status   string                 `json:"status"`
	Exit     interface{}            `json:"exit,omitempty"`
	Output   map[string]interface{} `json:"output,omitempty"`
}

type Session_Message_Skill struct {
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	F_type   string                 `json:"type"`
	Skill    string                 `json:"skill"`
	Name     string                 `json:"name"`
	Text     string                 `json:"text"`
}

type Session_Message_Synthetic struct {
	Id          interface{}            `json:"id"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Time        map[string]interface{} `json:"time"`
	Text        string                 `json:"text"`
	Description string                 `json:"description,omitempty"`
	F_type      string                 `json:"type"`
}

type Session_Message_System struct {
	Id          interface{}            `json:"id"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Time        map[string]interface{} `json:"time"`
	F_type      string                 `json:"type"`
	Text        string                 `json:"text"`
	Description string                 `json:"description,omitempty"`
}

type Session_Message_ToolState_Completed struct {
	Status   string                 `json:"status"`
	Input    map[string]interface{} `json:"input"`
	Content  []interface{}          `json:"content"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Session_Message_ToolState_Error struct {
	Status   string                 `json:"status"`
	Input    map[string]interface{} `json:"input"`
	Error    interface{}            `json:"error"`
	Content  []interface{}          `json:"content,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Session_Message_ToolState_Running struct {
	Status   string                 `json:"status"`
	Input    map[string]interface{} `json:"input"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Session_Message_ToolState_Streaming struct {
	Status string `json:"status"`
	Input  string `json:"input"`
}

type Session_Message_User struct {
	Id       interface{}            `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Time     map[string]interface{} `json:"time"`
	Text     string                 `json:"text"`
	Files    []interface{}          `json:"files,omitempty"`
	Agents   []interface{}          `json:"agents,omitempty"`
	Skills   []interface{}          `json:"skills,omitempty"`
	F_type   string                 `json:"type"`
}

type Session_Revert struct {
	MessageID interface{}   `json:"messageID"`
	PartID    string        `json:"partID,omitempty"`
	Snapshot  string        `json:"snapshot,omitempty"`
	Files     []interface{} `json:"files,omitempty"`
}

type Session_StructuredError struct {
	F_type  string      `json:"type"`
	Message string      `json:"message"`
	Status  interface{} `json:"status,omitempty"`
}

type SessionActive struct {
	F_type string `json:"type"`
}

type SessionBusyError struct {
	Tag       string `json:"_tag"`
	SessionID string `json:"sessionID"`
	Message   string `json:"message"`
}

type SessionGenerateResponse struct {
	Data map[string]interface{} `json:"data"`
}

type SessionMessagesResponse struct {
	Data   []interface{}          `json:"data"`
	Cursor map[string]interface{} `json:"cursor"`
}

type SessionNotFoundError struct {
	Tag       string `json:"_tag"`
	SessionID string `json:"sessionID"`
	Message   string `json:"message"`
}

type SessionTransfer_Data struct {
	Info     interface{}   `json:"info"`
	Messages []interface{} `json:"messages"`
}

type SessionsResponse struct {
	Data   []interface{}          `json:"data"`
	Cursor map[string]interface{} `json:"cursor"`
}

type Shell_Info struct {
	Id       interface{}            `json:"id"`
	Status   string                 `json:"status"`
	Command  string                 `json:"command"`
	Cwd      string                 `json:"cwd"`
	Shell    string                 `json:"shell"`
	File     string                 `json:"file"`
	Pid      interface{}            `json:"pid,omitempty"`
	Exit     float64                `json:"exit,omitempty"`
	Metadata map[string]interface{} `json:"metadata"`
	Time     map[string]interface{} `json:"time"`
}

type Shell_Info1 struct {
	Id       interface{}            `json:"id"`
	Status   string                 `json:"status"`
	Command  string                 `json:"command"`
	Cwd      string                 `json:"cwd"`
	Shell    string                 `json:"shell"`
	File     string                 `json:"file"`
	Pid      interface{}            `json:"pid,omitempty"`
	Exit     float64                `json:"exit,omitempty"`
	Metadata map[string]interface{} `json:"metadata"`
	Time     map[string]interface{} `json:"time"`
}

type ShellNotFoundError struct {
	Tag     string `json:"_tag"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

type Skill_Info struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Slash       bool   `json:"slash,omitempty"`
	Autoinvoke  bool   `json:"autoinvoke,omitempty"`
	Location    string `json:"location"`
	Content     string `json:"content"`
}

type SkillNotFoundError struct {
	Tag     string `json:"_tag"`
	Skill   string `json:"skill"`
	Message string `json:"message"`
}

type TokenUsage_Info struct {
	Input     float64                `json:"input"`
	Output    float64                `json:"output"`
	Reasoning float64                `json:"reasoning"`
	Cache     map[string]interface{} `json:"cache"`
}

type Tool_FileContent struct {
	F_type string      `json:"type"`
	Uri    string      `json:"uri"`
	Mime   string      `json:"mime"`
	Name   interface{} `json:"name,omitempty"`
}

type Tool_FileContent1 struct {
	F_type string      `json:"type"`
	Uri    string      `json:"uri"`
	Mime   string      `json:"mime"`
	Name   interface{} `json:"name,omitempty"`
}

type Tool_TextContent struct {
	F_type string `json:"type"`
	Text   string `json:"text"`
}

type UnauthorizedError struct {
	Tag     string `json:"_tag"`
	Message string `json:"message"`
}

type UnknownError struct {
	Tag     string      `json:"_tag"`
	Message string      `json:"message"`
	Ref     interface{} `json:"ref,omitempty"`
}

type V2Event_server_connected struct {
	Id       interface{} `json:"id"`
	Metadata interface{} `json:"metadata,omitempty"`
	Location interface{} `json:"location,omitempty"`
	F_type   string      `json:"type"`
	Data     interface{} `json:"data"`
}

type Vcs_Branch struct {
	Current   string `json:"current,omitempty"`
	F_default string `json:"default,omitempty"`
}

type Vcs_FileStatus struct {
	File      string      `json:"file"`
	Additions interface{} `json:"additions"`
	Deletions interface{} `json:"deletions"`
	Status    string      `json:"status"`
}

type Vcs_Info struct {
	Branch interface{} `json:"branch"`
}

type WebSearch_Provider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type WebSearch_Response struct {
	ProviderID string        `json:"providerID"`
	Results    []interface{} `json:"results"`
}

type WebSearch_Result struct {
	Url     string                 `json:"url"`
	Title   string                 `json:"title,omitempty"`
	Content string                 `json:"content,omitempty"`
	Time    map[string]interface{} `json:"time"`
}

type Worktree_Directory struct {
	Directory string `json:"directory"`
	Strategy  string `json:"strategy,omitempty"`
}

type Worktree_Info struct {
	Directory string `json:"directory"`
}

type WorktreeError struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type Agent_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Catalog_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Command_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Config_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Filesystem_changed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Form_cancelled struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Form_created struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Form_replied struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Installation_update_available struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Installation_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Integration_connection_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Integration_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Mcp_resources_changed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Mcp_status_changed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Models_dev_refreshed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Permission_asked struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Permission_replied struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Plugin_added struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Plugin_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Pty_created struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Pty_deleted struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Pty_exited struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Pty_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Reference_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Session_agent_selected struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_compaction_delta struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_compaction_ended struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_compaction_failed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_compaction_started struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_created struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_deleted struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_execution_failed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_execution_interrupted struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_execution_started struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_execution_succeeded struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_forked struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_idle struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_inbox_cancelled struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_inbox_delivered struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_inbox_delivery_changed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_inbox_enqueued struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_instructions_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_model_selected struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_moved struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_reasoning_delta struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_reasoning_ended struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_reasoning_started struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_renamed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_retry_scheduled struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_revert_cleared struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_revert_committed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_revert_staged struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_shell_ended struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_shell_started struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_skill_activated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_status struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_step_ended struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_step_failed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_step_started struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_synthetic struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_text_delta struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_text_ended struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_text_started struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_tool_called struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_tool_failed struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_tool_input_delta struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_tool_input_ended struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_tool_input_started struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_tool_progress struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_tool_success struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_usage_recorded struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Session_usage_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Shell_created struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Shell_deleted struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Shell_exited struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Skill_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Tui_command_execute struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Tui_prompt_append struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Tui_session_select struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Tui_toast_show struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Vcs_branch_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Websearch_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     interface{}            `json:"data"`
}

type Worktree_resolved struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Durable  map[string]interface{} `json:"durable"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type Worktree_updated struct {
	Id       interface{}            `json:"id"`
	Created  float64                `json:"created"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	F_type   string                 `json:"type"`
	Location interface{}            `json:"location,omitempty"`
	Data     map[string]interface{} `json:"data"`
}
