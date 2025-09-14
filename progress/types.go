// Package progress defines types for tracking long-running operations
package progress

import (
	"time"
)

// Status represents the current state of an operation
type Status string

const (
	StatusQueued      Status = "queued"
	StatusProcessing  Status = "processing"
	StatusGenerating  Status = "generating"
	StatusUploading   Status = "uploading"
	StatusFinalizing  Status = "finalizing"
	StatusCompleted   Status = "completed"
	StatusFailed      Status = "failed"
	StatusCancelled   Status = "cancelled"
)

// Update represents a progress update for a long-running operation
type Update struct {
	// Core fields
	JobID    string    `json:"job_id"`
	Tool     string    `json:"tool"`
	Status   Status    `json:"status"`
	Progress int       `json:"progress"` // 0-100

	// Optional details
	Message       string     `json:"message,omitempty"`
	CurrentStep   string     `json:"current_step,omitempty"`
	TotalSteps    int        `json:"total_steps,omitempty"`
	StepProgress  int        `json:"step_progress,omitempty"` // Progress within current step

	// Timing
	UpdatedAt     time.Time  `json:"updated_at"`
	StartedAt     *time.Time `json:"started_at,omitempty"`
	EstimatedEnd  *time.Time `json:"estimated_end,omitempty"`

	// For partial results
	PartialResult interface{} `json:"partial_result,omitempty"`
}

// BatchUpdate for operations processing multiple items
type BatchUpdate struct {
	JobID          string    `json:"job_id"`
	Tool           string    `json:"tool"`
	TotalItems     int       `json:"total_items"`
	CompletedItems int       `json:"completed_items"`
	FailedItems    int       `json:"failed_items"`
	CurrentItem    string    `json:"current_item,omitempty"`
	Progress       int       `json:"progress"` // Overall progress 0-100
	UpdatedAt      time.Time `json:"updated_at"`

	// Item-level details
	ItemStatuses   []ItemStatus `json:"item_statuses,omitempty"`
}

// ItemStatus for tracking individual items in a batch
type ItemStatus struct {
	ID       string    `json:"id"`
	Status   Status    `json:"status"`
	Progress int       `json:"progress"`
	Error    *string   `json:"error,omitempty"`
	Result   interface{} `json:"result,omitempty"`
}

// Constructor functions

// NewUpdate creates a standard progress update
func NewUpdate(jobID, tool string, status Status, progress int, message string) *Update {
	return &Update{
		JobID:     jobID,
		Tool:      tool,
		Status:    status,
		Progress:  progress,
		Message:   message,
		UpdatedAt: time.Now(),
	}
}

// NewProcessingUpdate creates an update for processing status
func NewProcessingUpdate(jobID, tool string, progress int, currentStep string) *Update {
	return &Update{
		JobID:       jobID,
		Tool:        tool,
		Status:      StatusProcessing,
		Progress:    progress,
		CurrentStep: currentStep,
		UpdatedAt:   time.Now(),
	}
}

// NewBatchUpdate creates a batch progress update
func NewBatchUpdate(jobID, tool string, totalItems, completedItems, failedItems int) *BatchUpdate {
	progress := 0
	if totalItems > 0 {
		progress = (completedItems * 100) / totalItems
	}

	return &BatchUpdate{
		JobID:          jobID,
		Tool:           tool,
		TotalItems:     totalItems,
		CompletedItems: completedItems,
		FailedItems:    failedItems,
		Progress:       progress,
		UpdatedAt:      time.Now(),
	}
}

// Helper methods

// IsTerminal returns true if this is a final status
func (s Status) IsTerminal() bool {
	return s == StatusCompleted || s == StatusFailed || s == StatusCancelled
}

// IsActive returns true if the operation is still running
func (s Status) IsActive() bool {
	return !s.IsTerminal()
}

// CalculateETA estimates completion time based on progress and elapsed time
func (u *Update) CalculateETA() *time.Time {
	if u.StartedAt == nil || u.Progress <= 0 || u.Progress >= 100 {
		return nil
	}

	elapsed := time.Since(*u.StartedAt)
	totalEstimated := time.Duration(float64(elapsed) * (100.0 / float64(u.Progress)))
	remaining := totalEstimated - elapsed
	eta := time.Now().Add(remaining)

	return &eta
}