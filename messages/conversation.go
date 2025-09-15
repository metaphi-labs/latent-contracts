package messages

import (
	"fmt"
	"time"
)

// ConversationMessage represents a message being sent for processing
// Purpose: Standardize all Platform API → Chat AI communication
// This replaces the mix of typed/untyped formats with a single contract
type ConversationMessage struct {
	// Message identification
	ConversationID string `json:"conversationId"`
	MessageID      string `json:"messageId"`
	UserID         string `json:"userId"`

	// Current message
	Role    string `json:"role"`    // "user", "assistant", "system"
	Content string `json:"content,omitempty"`

	// For system messages with tool results
	ToolResultMessage *ToolResultMessage `json:"toolResultMessage,omitempty"`

	// Conversation context
	History       *ConversationHistory `json:"history,omitempty"`
	MediaRegistry []MediaAsset         `json:"mediaRegistry,omitempty"`

	// Settings
	EnableThinking bool `json:"enableThinking,omitempty"`
}

// ConversationHistory represents the conversation context
type ConversationHistory struct {
	Messages []HistoryMessage `json:"messages"`
}

// HistoryMessage represents a single message in conversation history
type HistoryMessage struct {
	Role        string     `json:"role"`
	Content     string     `json:"content"`
	ToolCalls   []ToolCall `json:"toolCalls,omitempty"`
	RawThoughts *string    `json:"rawThoughts,omitempty"`
	Timestamp   time.Time  `json:"timestamp"`
}

// MediaAsset represents a media file in the conversation
type MediaAsset struct {
	RefID      string    `json:"refId"`      // e.g., "image_1", "video_2"
	RefType    string    `json:"refType"`    // "image", "video", "audio"
	RefIndex   int       `json:"refIndex"`   // 1-based index for ordering
	StorageURL string    `json:"storageUrl"` // gs:// URL for AI models
	PublicURL  string    `json:"publicUrl"`  // https:// URL for display
	MimeType   string    `json:"mimeType"`   // e.g., "image/png"
	ToolName   string    `json:"toolName"`   // Tool that created this
	CreatedAt  time.Time `json:"createdAt"`
}

// Validate ensures the conversation message is well-formed
func (cm *ConversationMessage) Validate() error {
	if cm.ConversationID == "" {
		return fmt.Errorf("conversationId is required")
	}
	if cm.MessageID == "" {
		return fmt.Errorf("messageId is required")
	}
	if cm.UserID == "" {
		return fmt.Errorf("userId is required")
	}

	// Validate role
	switch cm.Role {
	case "user", "assistant", "system":
		// Valid role
	default:
		return fmt.Errorf("invalid role: %s", cm.Role)
	}

	// For user and assistant messages, content is required
	if cm.Role != "system" && cm.Content == "" {
		return fmt.Errorf("content is required for %s messages", cm.Role)
	}

	// For system messages, either content or tool result is required
	if cm.Role == "system" && cm.Content == "" && cm.ToolResultMessage == nil {
		return fmt.Errorf("system messages must have content or toolResultMessage")
	}

	// Validate tool result if present
	if cm.ToolResultMessage != nil {
		if err := cm.ToolResultMessage.Validate(); err != nil {
			return fmt.Errorf("invalid toolResultMessage: %w", err)
		}
	}

	// Validate history if present
	if cm.History != nil {
		if err := cm.History.Validate(); err != nil {
			return fmt.Errorf("invalid history: %w", err)
		}
	}

	// Validate media assets
	for i, asset := range cm.MediaRegistry {
		if err := asset.Validate(); err != nil {
			return fmt.Errorf("invalid mediaRegistry[%d]: %w", i, err)
		}
	}

	return nil
}

// Validate ensures the history is well-formed
func (ch *ConversationHistory) Validate() error {
	for i, msg := range ch.Messages {
		if err := msg.Validate(); err != nil {
			return fmt.Errorf("invalid message[%d]: %w", i, err)
		}
	}
	return nil
}

// Validate ensures the history message is well-formed
func (hm *HistoryMessage) Validate() error {
	switch hm.Role {
	case "user", "assistant", "system":
		// Valid role
	default:
		return fmt.Errorf("invalid role: %s", hm.Role)
	}

	if hm.Content == "" && len(hm.ToolCalls) == 0 {
		return fmt.Errorf("message must have content or tool calls")
	}

	// Validate tool calls if present
	for i, tc := range hm.ToolCalls {
		if err := tc.Validate(); err != nil {
			return fmt.Errorf("invalid toolCall[%d]: %w", i, err)
		}
	}

	return nil
}

// Validate ensures the media asset is well-formed
func (ma *MediaAsset) Validate() error {
	if ma.RefID == "" {
		return fmt.Errorf("refId is required")
	}
	if ma.RefType == "" {
		return fmt.Errorf("refType is required")
	}
	if ma.StorageURL == "" {
		return fmt.Errorf("storageUrl is required")
	}
	if ma.PublicURL == "" {
		return fmt.Errorf("publicUrl is required")
	}
	if ma.MimeType == "" {
		return fmt.Errorf("mimeType is required")
	}
	return nil
}