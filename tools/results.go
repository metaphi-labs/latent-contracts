package tools

import (
	"time"
)

// ToolResult represents the result of a tool execution
type ToolResult struct {
	// Success indicates if the tool executed successfully
	Success bool `json:"success"`

	// Tool that was executed
	Tool string `json:"tool"`

	// Result contains the actual output from the tool
	Result interface{} `json:"result,omitempty"`

	// Error message if the tool failed
	Error string `json:"error,omitempty"`

	// ErrorCode for programmatic error handling
	ErrorCode string `json:"error_code,omitempty"`

	// Retryable indicates if the operation can be retried
	Retryable bool `json:"retryable,omitempty"`

	// ExecutionTime in milliseconds
	ExecutionTime int64 `json:"execution_time_ms,omitempty"`

	// Metadata about the execution
	Metadata *ResultMetadata `json:"metadata,omitempty"`
}

// ResultMetadata contains additional information about tool execution
type ResultMetadata struct {
	// JobID for async operations
	JobID string `json:"job_id,omitempty"`

	// MediaAssets created by the tool
	MediaAssets []MediaAsset `json:"media_assets,omitempty"`

	// CreditsUsed by this execution
	CreditsUsed int `json:"credits_used,omitempty"`

	// Provider that executed the tool (e.g., "vertex-ai", "openai")
	Provider string `json:"provider,omitempty"`

	// Model used (for AI tools)
	Model string `json:"model,omitempty"`

	// RequestID for tracing
	RequestID string `json:"request_id,omitempty"`

	// Additional provider-specific data
	ProviderData map[string]interface{} `json:"provider_data,omitempty"`
}

// MediaAsset represents a media file created by a tool
type MediaAsset struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"` // "image", "video", "audio"
	PublicURL   string    `json:"public_url"`
	StorageURL  string    `json:"storage_url"`
	SignedURL   string    `json:"signed_url,omitempty"`
	SignedExpiry time.Time `json:"signed_expiry,omitempty"`
	MimeType    string    `json:"mime_type"`
	Width       int       `json:"width,omitempty"`
	Height      int       `json:"height,omitempty"`
	Duration    float64   `json:"duration,omitempty"` // For video/audio in seconds
	FileSize    int64     `json:"file_size,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// NewSuccessResult creates a successful tool result
func NewSuccessResult(tool string, result interface{}) *ToolResult {
	return &ToolResult{
		Success: true,
		Tool:    tool,
		Result:  result,
	}
}

// NewErrorResult creates an error tool result
func NewErrorResult(tool string, errorMsg string, errorCode string) *ToolResult {
	return &ToolResult{
		Success:   false,
		Tool:      tool,
		Error:     errorMsg,
		ErrorCode: errorCode,
		Retryable: isRetryableError(errorCode),
	}
}

// isRetryableError determines if an error code represents a retryable error
func isRetryableError(errorCode string) bool {
	retryableCodes := []string{
		"TIMEOUT",
		"NETWORK_ERROR",
		"SERVICE_UNAVAILABLE",
		"RATE_LIMIT",
		"PROVIDER_OVERLOADED",
	}

	for _, code := range retryableCodes {
		if code == errorCode {
			return true
		}
	}
	return false
}

// HasMediaAssets checks if the result contains media assets
func (r *ToolResult) HasMediaAssets() bool {
	return r.Metadata != nil && len(r.Metadata.MediaAssets) > 0
}

// GetMediaAssets returns media assets from the result
func (r *ToolResult) GetMediaAssets() []MediaAsset {
	if r.Metadata == nil {
		return nil
	}
	return r.Metadata.MediaAssets
}

// IsRetryable checks if the tool execution can be retried
func (r *ToolResult) IsRetryable() bool {
	return !r.Success && r.Retryable
}