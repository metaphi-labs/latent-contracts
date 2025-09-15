// Package messages defines contracts for inter-service messaging
package messages

import (
	"fmt"
	"time"

	"github.com/metaphi-labs/latent-contracts/results"
)

// ToolResultMessage represents a tool execution result being delivered to an agent
// Purpose: Deliver tool results from Platform API to Chat AI for conversation continuation
// This contract ensures type safety across service boundaries
type ToolResultMessage struct {
	// Message routing
	ConversationID string `json:"conversationId"` // Conversation this belongs to
	MessageID      string `json:"messageId"`      // System message ID in platform
	UserID         string `json:"userId"`         // User who owns the conversation

	// Tool execution reference
	ToolRef  string `json:"toolRef"`         // Reference ID like "step_1", "tool_1"
	ToolName string `json:"toolName"`       // Tool that was executed
	JobID    string `json:"jobId,omitempty"` // Platform's job tracking ID (optional)

	// The actual result - required, full typed result
	Result *results.ToolResult `json:"result"`

	// Metadata
	DeliveredAt time.Time `json:"deliveredAt"`

	// Plan continuation context (optional)
	// Helps Chat AI understand where in a multi-step plan we are
	PlanContext *PlanContext `json:"planContext,omitempty"`
}

// PlanContext provides context about multi-step plan execution
type PlanContext struct {
	PlanID     string `json:"planId"`     // Unique plan identifier
	StepIndex  int    `json:"stepIndex"`  // Current step (0-based)
	TotalSteps int    `json:"totalSteps"` // Total number of steps
	IsLastStep bool   `json:"isLastStep"` // Convenience flag for final step
}

// Validate ensures the message is well-formed
func (m *ToolResultMessage) Validate() error {
	if m.ConversationID == "" {
		return fmt.Errorf("conversationId is required")
	}
	if m.MessageID == "" {
		return fmt.Errorf("messageId is required")
	}
	if m.UserID == "" {
		return fmt.Errorf("userId is required")
	}
	if m.ToolRef == "" {
		return fmt.Errorf("toolRef is required")
	}
	if m.ToolName == "" {
		return fmt.Errorf("toolName is required")
	}
	if m.Result == nil {
		return fmt.Errorf("result is required")
	}

	// Validate the result itself
	if err := m.Result.Validate(); err != nil {
		return fmt.Errorf("invalid result: %w", err)
	}

	// Validate plan context if present
	if m.PlanContext != nil {
		if m.PlanContext.PlanID == "" {
			return fmt.Errorf("planContext.planId is required when planContext is present")
		}
		if m.PlanContext.StepIndex < 0 {
			return fmt.Errorf("planContext.stepIndex must be non-negative")
		}
		if m.PlanContext.StepIndex >= m.PlanContext.TotalSteps {
			return fmt.Errorf("planContext.stepIndex must be less than totalSteps")
		}
	}

	return nil
}