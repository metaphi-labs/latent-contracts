package events

import (
	"time"
)

// EventType represents the type of tool event
type EventType string

const (
	// Tool lifecycle events
	EventToolRequested EventType = "tool.requested"
	EventToolStarted   EventType = "tool.started"
	EventToolProgress  EventType = "tool.progress"
	EventToolCompleted EventType = "tool.completed"
	EventToolRejected  EventType = "tool.rejected"
	EventToolFailed    EventType = "tool.failed"
)

// ToolEvent represents a tool execution event emitted by microservices
type ToolEvent struct {
	// Core fields - always required
	Type           EventType `json:"type"`
	JobID          string    `json:"jobId"`
	Tool           string    `json:"tool"`
	UserID         string    `json:"userId"`
	ConversationID string    `json:"conversationId"`
	Timestamp      time.Time `json:"timestamp"`
	Service        string    `json:"service"` // Which service emitted this

	// For rejections and failures
	Reason    string `json:"reason,omitempty"`
	Details   string `json:"details,omitempty"`
	Retryable bool   `json:"retryable,omitempty"`

	// For progress updates
	Progress    int    `json:"progress,omitempty"`    // 0-100
	Message     string `json:"message,omitempty"`     // User-friendly status
	EstimatedMs int64  `json:"estimatedMs,omitempty"` // Estimated time remaining

	// For completed events
	ResultID   string `json:"resultId,omitempty"`   // ID of the result in tool_results table
	CreditUsed int    `json:"creditUsed,omitempty"` // Actual credits consumed
}

// NewToolRejected creates a rejection event
func NewToolRejected(jobID, tool, userID, conversationID, service, reason, details string) *ToolEvent {
	return &ToolEvent{
		Type:           EventToolRejected,
		JobID:          jobID,
		Tool:           tool,
		UserID:         userID,
		ConversationID: conversationID,
		Service:        service,
		Timestamp:      time.Now(),
		Reason:         reason,
		Details:        details,
		Retryable:      false,
	}
}

// NewToolProgress creates a progress event
func NewToolProgress(jobID, tool, userID, conversationID, service string, progress int, message string) *ToolEvent {
	return &ToolEvent{
		Type:           EventToolProgress,
		JobID:          jobID,
		Tool:           tool,
		UserID:         userID,
		ConversationID: conversationID,
		Service:        service,
		Timestamp:      time.Now(),
		Progress:       progress,
		Message:        message,
	}
}

// NewToolCompleted creates a completion event
func NewToolCompleted(jobID, tool, userID, conversationID, service, resultID string, creditUsed int) *ToolEvent {
	return &ToolEvent{
		Type:           EventToolCompleted,
		JobID:          jobID,
		Tool:           tool,
		UserID:         userID,
		ConversationID: conversationID,
		Service:        service,
		Timestamp:      time.Now(),
		ResultID:       resultID,
		CreditUsed:     creditUsed,
	}
}