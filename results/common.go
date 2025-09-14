// Package results defines the canonical result types for all tool executions.
// Services MUST create these types directly - no adapters or extraction needed.
package results

import (
	"fmt"
	"time"
)

// ToolResult is the standard response format for ALL tool executions
// Services create this directly with the appropriate result type
type ToolResult struct {
	// Core fields - always required
	Success bool   `json:"success"`
	Tool    string `json:"tool"`
	JobID   string `json:"job_id"`

	// Result payload - exactly one of these based on tool category
	MediaGeneration  *MediaGenerationResult  `json:"media_generation,omitempty"`
	VideoProcessing  *VideoProcessingResult  `json:"video_processing,omitempty"`
	ContentAnalysis  *ContentAnalysisResult  `json:"content_analysis,omitempty"`

	// Error information if Success is false
	Error *ErrorInfo `json:"error,omitempty"`

	// Execution metadata
	Metadata ExecutionMetadata `json:"metadata"`
}

// MediaAsset represents a single media file (used across all result types)
type MediaAsset struct {
	// Identifiers
	ID    string `json:"id"`    // Unique asset ID
	Index int    `json:"index"` // For multi-generation (0-based)

	// URLs - services MUST provide these explicitly
	StorageURL string `json:"storage_url"` // gs://bucket/path for persistence
	PublicURL  string `json:"public_url"`  // https:// URL for access

	// Media properties
	MimeType string `json:"mime_type"`
	FileSize int64  `json:"file_size_bytes"`

	// Type-specific properties
	Width     *int     `json:"width,omitempty"`              // For images/video
	Height    *int     `json:"height,omitempty"`             // For images/video
	Duration  *float64 `json:"duration_seconds,omitempty"`  // For video/audio
	FrameRate *float64 `json:"frame_rate,omitempty"`        // For video
	BitRate   *int     `json:"bit_rate,omitempty"`          // For video/audio

	// Optional
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`
}

// ErrorInfo when Success is false
type ErrorInfo struct {
	Code       string      `json:"code"`    // Standardized error code from errors package
	Message    string      `json:"message"` // Human-readable message
	Retryable  bool        `json:"retryable"`
	RetryAfter *time.Time  `json:"retry_after,omitempty"`
	Details    interface{} `json:"details,omitempty"`
}

// ExecutionMetadata about the tool execution
type ExecutionMetadata struct {
	// Timing
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	DurationMs int64     `json:"duration_ms"`

	// Resources
	CreditsUsed int  `json:"credits_used"`
	TokensUsed  *int `json:"tokens_used,omitempty"`

	// Provider info
	Provider      string `json:"provider"`                // "vertex-ai", "openai", etc
	ProviderJobID string `json:"provider_job_id,omitempty"`
	Region        string `json:"region,omitempty"`

	// Request tracking
	RequestID      string `json:"request_id"`
	ConversationID string `json:"conversation_id"`
	UserID         string `json:"user_id"`
}

// Validation methods

// Validate ensures the result is well-formed
func (r *ToolResult) Validate() error {
	if r.Tool == "" {
		return fmt.Errorf("tool name is required")
	}
	if r.JobID == "" {
		return fmt.Errorf("job ID is required")
	}

	if r.Success {
		// Must have exactly one result type
		count := 0
		if r.MediaGeneration != nil {
			count++
		}
		if r.VideoProcessing != nil {
			count++
		}
		if r.ContentAnalysis != nil {
			count++
		}
		if count != 1 {
			return fmt.Errorf("successful result must have exactly one result type, got %d", count)
		}
	} else {
		// Must have error info
		if r.Error == nil {
			return fmt.Errorf("failed result must have error information")
		}
		if r.Error.Code == "" {
			return fmt.Errorf("error must have a code")
		}
	}

	return nil
}

// ValidateMediaAsset ensures a media asset is well-formed
func ValidateMediaAsset(asset MediaAsset, index int) error {
	if asset.ID == "" {
		return fmt.Errorf("asset[%d]: ID is required", index)
	}
	if asset.StorageURL == "" {
		return fmt.Errorf("asset[%d]: storage URL is required", index)
	}
	if asset.PublicURL == "" {
		return fmt.Errorf("asset[%d]: public URL is required", index)
	}
	if asset.MimeType == "" {
		return fmt.Errorf("asset[%d]: MIME type is required", index)
	}
	return nil
}