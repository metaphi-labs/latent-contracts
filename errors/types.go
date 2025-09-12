package errors

import (
	"encoding/json"
	"time"
)

// ServiceError is the standard error structure used across all Latent services.
// All services (Platform API, Media AI, Chat AI, etc.) must use this structure
// for error responses to ensure consistency.
type ServiceError struct {
	// Code is the standardized error code from codes.go
	Code ErrorCode `json:"code"`
	
	// Message is a human-readable error message
	Message string `json:"message"`
	
	// Service identifies which service generated the error
	Service string `json:"service"`
	
	// Retryable indicates if the client should retry the request
	Retryable bool `json:"retryable"`
	
	// OccurredAt is when the error occurred
	OccurredAt time.Time `json:"occurred_at"`
	
	// RequestID for tracing (if available)
	RequestID string `json:"request_id,omitempty"`
	
	// JobID for async operations (if applicable)
	JobID string `json:"job_id,omitempty"`
	
	// Metadata contains structured error details
	Metadata *ErrorMetadata `json:"metadata,omitempty"`
}

// ErrorMetadata contains additional structured information about the error
type ErrorMetadata struct {
	// For validation errors
	ValidationDetails []ValidationDetail `json:"validation_details,omitempty"`
	
	// For AI content violations
	ViolationDetails []ViolationDetail `json:"violation_details,omitempty"`
	
	// For rate limiting
	RetryAfter *time.Duration `json:"retry_after,omitempty"`
	
	// For quota errors
	QuotaLimit int `json:"quota_limit,omitempty"`
	QuotaUsed  int `json:"quota_used,omitempty"`
	
	// Provider-specific information (e.g., Vertex AI, OpenAI)
	Provider     string                 `json:"provider,omitempty"`
	ProviderCode string                 `json:"provider_code,omitempty"`
	ProviderData map[string]interface{} `json:"provider_data,omitempty"`
	
	// Additional context
	Details map[string]interface{} `json:"details,omitempty"`
}

// ValidationDetail describes a specific validation failure
type ValidationDetail struct {
	// Field that failed validation
	Field string `json:"field"`
	
	// Value that was provided
	Provided interface{} `json:"provided,omitempty"`
	
	// Expected value or format
	Expected interface{} `json:"expected,omitempty"`
	
	// Human-readable reason for failure
	Reason string `json:"reason"`
}

// ViolationDetail describes a content policy violation
type ViolationDetail struct {
	// Type of violation (maps to specific error codes)
	Type string `json:"type"`
	
	// Human-readable description
	Description string `json:"description"`
	
	// Severity level
	Severity Severity `json:"severity"`
	
	// Confidence score (0.0 to 1.0) if available
	Confidence *float64 `json:"confidence,omitempty"`
	
	// Provider-specific violation code
	ProviderCode string `json:"provider_code,omitempty"`
}

// Severity represents the severity level of an error or violation
type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

// Error implements the error interface
func (e *ServiceError) Error() string {
	return e.Message
}

// IsRetryable returns whether the error is retryable
func (e *ServiceError) IsRetryable() bool {
	return e.Retryable
}

// GetCode returns the error code
func (e *ServiceError) GetCode() ErrorCode {
	return e.Code
}

// HasViolations returns true if the error contains content violations
func (e *ServiceError) HasViolations() bool {
	return e.Metadata != nil && len(e.Metadata.ViolationDetails) > 0
}

// HasValidationErrors returns true if the error contains validation errors
func (e *ServiceError) HasValidationErrors() bool {
	return e.Metadata != nil && len(e.Metadata.ValidationDetails) > 0
}

// ToJSON converts the error to JSON
func (e *ServiceError) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// NewServiceError creates a new ServiceError with required fields
func NewServiceError(code ErrorCode, message string, service string, retryable bool) *ServiceError {
	return &ServiceError{
		Code:       code,
		Message:    message,
		Service:    service,
		Retryable:  retryable,
		OccurredAt: time.Now(),
	}
}

// Builder pattern methods for fluent API

// WithRequestID adds a request ID to the error
func (e *ServiceError) WithRequestID(requestID string) *ServiceError {
	e.RequestID = requestID
	return e
}

// WithJobID adds a job ID to the error
func (e *ServiceError) WithJobID(jobID string) *ServiceError {
	e.JobID = jobID
	return e
}

// WithMetadata adds metadata to the error
func (e *ServiceError) WithMetadata(metadata *ErrorMetadata) *ServiceError {
	e.Metadata = metadata
	return e
}

// WithValidationErrors adds validation details to the error
func (e *ServiceError) WithValidationErrors(details []ValidationDetail) *ServiceError {
	if e.Metadata == nil {
		e.Metadata = &ErrorMetadata{}
	}
	e.Metadata.ValidationDetails = details
	return e
}

// WithViolations adds violation details to the error
func (e *ServiceError) WithViolations(violations []ViolationDetail) *ServiceError {
	if e.Metadata == nil {
		e.Metadata = &ErrorMetadata{}
	}
	e.Metadata.ViolationDetails = violations
	return e
}

// WithRetryAfter adds retry-after duration for rate limiting
func (e *ServiceError) WithRetryAfter(duration time.Duration) *ServiceError {
	if e.Metadata == nil {
		e.Metadata = &ErrorMetadata{}
	}
	e.Metadata.RetryAfter = &duration
	return e
}

// WithProvider adds provider information to the error
func (e *ServiceError) WithProvider(provider string, code string) *ServiceError {
	if e.Metadata == nil {
		e.Metadata = &ErrorMetadata{}
	}
	e.Metadata.Provider = provider
	e.Metadata.ProviderCode = code
	return e
}