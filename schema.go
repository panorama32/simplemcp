package simplemcp

import "encoding/json"

// -- supported methods --

const (
	InitializeMethod               = "initialize"
	NotificationsInitializedMethod = "notifications/initialized"
	ToolsListMethod                = "tools/list"
	ToolsCallMethod                = "tools/call"
)

// -- initialize --

type InitializeParams struct {
	ProtocolVersion string         `json:"protocolVersion"`
	ClientInfo      Implementation `json:"clientInfo"`
}

type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      Implementation     `json:"serverInfo"`
}

type Implementation struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ServerCapabilities struct {
	Tools *Tools `json:"tools,omitempty"`
}

type Tools struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

// -- tools/list --

type ListToolsResult struct {
	Tools []*Tool `json:"tools"`
}

type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	InputSchema InputSchema `json:"inputSchema"`
}

type InputSchema struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties,omitempty"`
	Required   []string            `json:"required,omitempty"`
}

const InputSchemaTypeObject = "object"

type Property struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

const (
	PropertyTypeInteger = "integer"
	PropertyTypeNumber  = "number"
	PropertyTypeString  = "string"
)

// -- tools/call --

type CallToolParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
}

type CallToolResult struct {
	Content []any `json:"content"`
	IsError bool  `json:"isError,omitempty"`
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

const ContentTypeText = "text"

func NewCallToolResult() CallToolResult {
	return CallToolResult{
		Content: []any{},
	}
}

func (r *CallToolResult) AddTextContent(text string) {
	r.Content = append(r.Content, TextContent{
		Type: ContentTypeText,
		Text: text,
	})
}
