package callbacks

import (
	"github.com/metaphi-labs/latent-contracts/errors"
	"github.com/metaphi-labs/latent-contracts/results"
)

// CallbackRequest represents what external services send back to Platform API
// Services (Media AI, Video Processor) send strongly-typed ToolResult directly
type CallbackRequest struct {
	// Operational metadata (routing and context)
	JobID          string `json:"job_id" binding:"required"`
	UserID         string `json:"user_id" binding:"required"`
	ConversationID string `json:"conversation_id" binding:"required"`
	MessageID      string `json:"message_id,omitempty"` // Assistant message that triggered the tool
	Tool           string `json:"tool" binding:"required"` // Tool name from registry

	// Status of the operation
	Status string `json:"status" binding:"required"` // "completed" | "failed" | "processing" | "partial"

	// Result payload - strongly typed from contracts
	// For success: services create ToolResult with appropriate result type (MediaGeneration, VideoProcessing, etc)
	// For failure: services set Error field instead
	Result *results.ToolResult     `json:"result,omitempty"` // Strongly-typed result for successful completion
	Error  *errors.ServiceError    `json:"error,omitempty"`  // Rich error for failures
}

// Status constants for callbacks
const (
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusFailed     = "failed"
	StatusPartial    = "partial"
)