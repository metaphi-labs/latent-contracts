package providers

import (
	"strings"
	"github.com/metaphi-labs/latent-contracts/errors"
)

// OpenAI specific error codes and patterns
const (
	// Rate limiting
	OpenAIRateLimitError = "rate_limit_exceeded"
	OpenAIQuotaExceeded  = "insufficient_quota"

	// Model errors
	OpenAIModelNotFound  = "model_not_found"
	OpenAIModelOverloaded = "engine_overloaded"

	// Request errors
	OpenAIInvalidRequest = "invalid_request_error"
	OpenAIContextLength  = "context_length_exceeded"
	OpenAIMaxTokens      = "max_tokens_exceeded"

	// Content policy
	OpenAIContentPolicy  = "content_policy_violation"
	OpenAIContentFilter  = "content_filter"

	// System errors
	OpenAIServerError    = "server_error"
	OpenAIServiceUnavailable = "service_unavailable"
	OpenAITimeout        = "timeout"
)

// MapOpenAIError maps OpenAI error codes to standard error codes
func MapOpenAIError(errorType string, errorCode string, message string) errors.ErrorCode {
	// Check error codes first
	switch errorCode {
	case OpenAIRateLimitError:
		return errors.RATE_LIMIT_EXCEEDED
	case OpenAIQuotaExceeded:
		return errors.RATE_QUOTA_EXCEEDED
	case OpenAIModelNotFound:
		return errors.AI_INVALID_MODEL
	case OpenAIModelOverloaded:
		return errors.AI_MODEL_OVERLOADED
	case OpenAIInvalidRequest:
		return errors.VAL_INVALID_REQUEST
	case OpenAIContextLength, OpenAIMaxTokens:
		return errors.AI_CONTEXT_LENGTH_EXCEEDED
	case OpenAIContentPolicy, OpenAIContentFilter:
		return mapOpenAIContentViolation(message)
	case OpenAIServerError:
		return errors.SYS_INTERNAL_ERROR
	case OpenAIServiceUnavailable:
		return errors.SYS_SERVICE_UNAVAILABLE
	case OpenAITimeout:
		return errors.SYS_TIMEOUT
	}

	// Check error type
	switch errorType {
	case "invalid_request_error":
		return errors.VAL_INVALID_REQUEST
	case "authentication_error":
		return errors.AUTH_UNAUTHORIZED
	case "permission_error":
		return errors.AUTH_FORBIDDEN
	case "not_found_error":
		return errors.AI_INVALID_MODEL
	case "rate_limit_error":
		return errors.RATE_LIMIT_EXCEEDED
	case "api_connection_error":
		return errors.SYS_NETWORK_ERROR
	case "timeout_error":
		return errors.SYS_TIMEOUT
	}

	// Check message patterns
	if strings.Contains(strings.ToLower(message), "rate limit") {
		return errors.RATE_LIMIT_EXCEEDED
	}
	if strings.Contains(strings.ToLower(message), "quota") {
		return errors.RATE_QUOTA_EXCEEDED
	}
	if strings.Contains(strings.ToLower(message), "context") || strings.Contains(strings.ToLower(message), "token") {
		return errors.AI_CONTEXT_LENGTH_EXCEEDED
	}
	if strings.Contains(strings.ToLower(message), "content") || strings.Contains(strings.ToLower(message), "policy") {
		return mapOpenAIContentViolation(message)
	}

	// Default
	return errors.AI_GENERATION_FAILED
}

// mapOpenAIContentViolation maps content policy violations to specific error codes
func mapOpenAIContentViolation(message string) errors.ErrorCode {
	msg := strings.ToLower(message)

	if strings.Contains(msg, "child") || strings.Contains(msg, "minor") {
		return errors.AI_VIOLATION_CHILD_SAFETY
	}
	if strings.Contains(msg, "violence") || strings.Contains(msg, "violent") {
		return errors.AI_VIOLATION_VIOLENCE
	}
	if strings.Contains(msg, "sexual") || strings.Contains(msg, "adult") {
		return errors.AI_VIOLATION_SEXUAL
	}
	if strings.Contains(msg, "hate") || strings.Contains(msg, "discriminat") {
		return errors.AI_VIOLATION_HATE_SPEECH
	}
	if strings.Contains(msg, "self-harm") || strings.Contains(msg, "dangerous") {
		return errors.AI_VIOLATION_DANGEROUS
	}
	if strings.Contains(msg, "personal") || strings.Contains(msg, "pii") {
		return errors.AI_VIOLATION_PERSONAL_INFO
	}

	// Default content violation
	return errors.AI_VIOLATION_OTHER
}

// IsRetryableOpenAIError determines if an OpenAI error is retryable
func IsRetryableOpenAIError(errorType string, errorCode string) bool {
	// Retryable error codes
	switch errorCode {
	case OpenAIRateLimitError:
		return true
	case OpenAIModelOverloaded:
		return true
	case OpenAIServerError:
		return true
	case OpenAIServiceUnavailable:
		return true
	case OpenAITimeout:
		return true
	}

	// Retryable error types
	switch errorType {
	case "rate_limit_error":
		return true
	case "api_connection_error":
		return true
	case "timeout_error":
		return true
	case "server_error":
		return true
	}

	return false
}