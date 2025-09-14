package tools

import "fmt"

// ContextMessage for multi-turn conversations with multimodal models
type ContextMessage struct {
	Role  string        `json:"role" validate:"required,oneof=user model"`
	Parts []MessagePart `json:"parts" validate:"required,min=1"`
}

// MessagePart can be text or image reference
type MessagePart struct {
	Text       string `json:"text,omitempty"`
	StorageURL string `json:"storage_url,omitempty"` // GCS URL for image
}

// Validate ensures the message part is well-formed
func (m *MessagePart) Validate() error {
	// At least one field must be populated
	if m.Text == "" && m.StorageURL == "" {
		return fmt.Errorf("message part must have either text or storage_url")
	}
	return nil
}

// Validate ensures the context message is well-formed
func (c *ContextMessage) Validate() error {
	if c.Role != "user" && c.Role != "model" {
		return fmt.Errorf("role must be 'user' or 'model'")
	}

	if len(c.Parts) == 0 {
		return fmt.Errorf("context message must have at least one part")
	}

	for i, part := range c.Parts {
		if err := part.Validate(); err != nil {
			return fmt.Errorf("part[%d]: %w", i, err)
		}
	}

	return nil
}