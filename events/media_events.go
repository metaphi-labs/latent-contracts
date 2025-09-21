package events

import (
	"fmt"
	"time"
)

// EventType represents the type of media event
type EventType string

const (
	// Media service event types
	EventToolStarted   EventType = "media.tool.started"
	EventToolProgress  EventType = "media.tool.progress"
	EventToolCompleted EventType = "media.tool.completed"
	EventToolRejected  EventType = "media.tool.rejected"
	EventToolFailed    EventType = "media.tool.failed"
)

// MediaEvent represents events from media-ai service matching platform structure exactly
type MediaEvent struct {
	ID        string      `json:"id"`
	Type      EventType   `json:"eventType"` // Note: "eventType" to match platform
	UserID    string      `json:"userId"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// MediaEventData contains the event-specific data
type MediaEventData struct {
	JobID          string `json:"jobId"`
	ConversationID string `json:"conversationId"`
	ToolName       string `json:"toolName"`
	Service        string `json:"service"`

	// Progress events
	Progress int    `json:"progress,omitempty"`  // 0-100
	Message  string `json:"message,omitempty"`   // Status message

	// Error/rejection events
	Reason  string `json:"reason,omitempty"`    // Why it failed/rejected
	Details string `json:"details,omitempty"`  // Additional context
}

// Helper functions for creating events

// NewMediaStarted creates a tool started event
func NewMediaStarted(jobID, toolName, userID, conversationID, service string) *MediaEvent {
	return &MediaEvent{
		ID:        fmt.Sprintf("%s-started-%d", jobID, time.Now().UnixNano()),
		Type:      EventToolStarted,
		UserID:    userID,
		Timestamp: time.Now().Unix(),
		Data: MediaEventData{
			JobID:          jobID,
			ConversationID: conversationID,
			ToolName:       toolName,
			Service:        service,
		},
	}
}

// NewMediaProgress creates a progress event
func NewMediaProgress(jobID, toolName, userID, conversationID, service string, progress int, message string) *MediaEvent {
	return &MediaEvent{
		ID:        fmt.Sprintf("%s-progress-%d", jobID, time.Now().UnixNano()),
		Type:      EventToolProgress,
		UserID:    userID,
		Timestamp: time.Now().Unix(),
		Data: MediaEventData{
			JobID:          jobID,
			ConversationID: conversationID,
			ToolName:       toolName,
			Service:        service,
			Progress:       progress,
			Message:        message,
		},
	}
}

// NewMediaCompleted creates a completion event
func NewMediaCompleted(jobID, toolName, userID, conversationID, service string) *MediaEvent {
	return &MediaEvent{
		ID:        fmt.Sprintf("%s-completed-%d", jobID, time.Now().UnixNano()),
		Type:      EventToolCompleted,
		UserID:    userID,
		Timestamp: time.Now().Unix(),
		Data: MediaEventData{
			JobID:          jobID,
			ConversationID: conversationID,
			ToolName:       toolName,
			Service:        service,
		},
	}
}

// NewMediaRejected creates a rejection event
func NewMediaRejected(jobID, toolName, userID, conversationID, service, reason, details string) *MediaEvent {
	return &MediaEvent{
		ID:        fmt.Sprintf("%s-rejected-%d", jobID, time.Now().UnixNano()),
		Type:      EventToolRejected,
		UserID:    userID,
		Timestamp: time.Now().Unix(),
		Data: MediaEventData{
			JobID:          jobID,
			ConversationID: conversationID,
			ToolName:       toolName,
			Service:        service,
			Reason:         reason,
			Details:        details,
		},
	}
}

// NewMediaFailed creates a failure event
func NewMediaFailed(jobID, toolName, userID, conversationID, service, reason, details string) *MediaEvent {
	return &MediaEvent{
		ID:        fmt.Sprintf("%s-failed-%d", jobID, time.Now().UnixNano()),
		Type:      EventToolFailed,
		UserID:    userID,
		Timestamp: time.Now().Unix(),
		Data: MediaEventData{
			JobID:          jobID,
			ConversationID: conversationID,
			ToolName:       toolName,
			Service:        service,
			Reason:         reason,
			Details:        details,
		},
	}
}