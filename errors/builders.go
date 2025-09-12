package errors

import (
	"fmt"
	"time"
)

// ContentViolationError creates a standardized content violation error
func ContentViolationError(service string, violations []ViolationDetail) *ServiceError {
	// Determine primary error code based on first violation
	code := AI_VIOLATION_OTHER
	if len(violations) > 0 {
		code = mapViolationTypeToCode(violations[0].Type)
	}
	
	// Build message
	message := "Content violated platform policies"
	if len(violations) > 0 {
		message = violations[0].Description
		if len(violations) > 1 {
			message += fmt.Sprintf(" (+%d more violations)", len(violations)-1)
		}
	}
	
	return NewServiceError(code, message, service, false).
		WithViolations(violations)
}

// ValidationError creates a validation error with field details
func ValidationError(service string, details []ValidationDetail) *ServiceError {
	message := "Validation failed"
	if len(details) > 0 {
		message = fmt.Sprintf("Validation failed for field '%s': %s", 
			details[0].Field, details[0].Reason)
	}
	
	return NewServiceError(VAL_INVALID_REQUEST, message, service, false).
		WithValidationErrors(details)
}

// MediaDimensionError creates an error for invalid media dimensions
func MediaDimensionError(service string, provided, expected interface{}) *ServiceError {
	return NewServiceError(
		MEDIA_INVALID_DIMENSIONS,
		fmt.Sprintf("Invalid media dimensions: %v", provided),
		service,
		false,
	).WithValidationErrors([]ValidationDetail{{
		Field:    "dimensions",
		Provided: provided,
		Expected: expected,
		Reason:   "Dimensions do not meet requirements",
	}})
}

// RateLimitError creates a rate limit error with retry information
func RateLimitError(service string, retryAfter time.Duration) *ServiceError {
	return NewServiceError(
		RATE_LIMIT_EXCEEDED,
		fmt.Sprintf("Rate limit exceeded. Retry after %v", retryAfter),
		service,
		true,
	).WithRetryAfter(retryAfter)
}

// ModelOverloadedError creates an error for overloaded AI models
func ModelOverloadedError(service string, provider string) *ServiceError {
	err := NewServiceError(
		AI_MODEL_OVERLOADED,
		"The AI model is currently at capacity. Please try again later.",
		service,
		true,
	)
	
	if provider != "" {
		err.WithProvider(provider, "")
	}
	
	return err
}

// TimeoutError creates a timeout error
func TimeoutError(service string, operation string) *ServiceError {
	message := "Operation timed out"
	if operation != "" {
		message = fmt.Sprintf("Operation '%s' timed out", operation)
	}
	
	return NewServiceError(SYS_TIMEOUT, message, service, true)
}

// InternalError creates an internal server error
func InternalError(service string, details string) *ServiceError {
	message := "An internal error occurred"
	if details != "" {
		message += ": " + details
	}
	
	return NewServiceError(SYS_INTERNAL_ERROR, message, service, true)
}

// Helper function to map violation types to error codes
func mapViolationTypeToCode(violationType string) ErrorCode {
	mapping := map[string]ErrorCode{
		"CHILD_SAFETY":  AI_VIOLATION_CHILD_SAFETY,
		"CELEBRITY":     AI_VIOLATION_CELEBRITY,
		"VIOLENCE":      AI_VIOLATION_VIOLENCE,
		"SEXUAL":        AI_VIOLATION_SEXUAL,
		"HATE_SPEECH":   AI_VIOLATION_HATE_SPEECH,
		"PERSONAL_INFO": AI_VIOLATION_PERSONAL_INFO,
		"TOXIC":         AI_VIOLATION_TOXIC,
		"DANGEROUS":     AI_VIOLATION_DANGEROUS,
		"PROHIBITED":    AI_VIOLATION_PROHIBITED,
		"VULGAR":        AI_VIOLATION_VULGAR,
	}
	
	if code, exists := mapping[violationType]; exists {
		return code
	}
	return AI_VIOLATION_OTHER
}

// IsRetryableCode returns whether an error code is typically retryable
func IsRetryableCode(code ErrorCode) bool {
	retryableCodes := map[ErrorCode]bool{
		SYS_SERVICE_UNAVAILABLE: true,
		SYS_TIMEOUT:             true,
		SYS_NETWORK_ERROR:       true,
		AI_MODEL_OVERLOADED:     true,
		AI_MODEL_UNAVAILABLE:    true,
		RATE_LIMIT_EXCEEDED:     true,
	}
	
	return retryableCodes[code]
}