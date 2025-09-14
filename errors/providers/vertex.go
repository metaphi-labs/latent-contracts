package providers

import (
	"strings"
	"github.com/metaphi-labs/latent-contracts/errors"
)

// Vertex AI specific error codes
const (
	// Content violation codes
	VertexChildSafety     = "58061214"
	VertexCelebrity       = "58061215"
	VertexViolence        = "58061216"
	VertexSexualContent   = "58061217"
	VertexHateSpeech      = "58061218"
	VertexDangerousContent = "58061219"

	// Resource and quota errors
	VertexQuotaExceeded   = "RESOURCE_EXHAUSTED"
	VertexRateLimited     = "RATE_LIMIT_EXCEEDED"
	VertexModelOverloaded = "MODEL_OVERLOADED"

	// Request errors
	VertexInvalidRequest  = "INVALID_ARGUMENT"
	VertexRequestTooLarge = "REQUEST_TOO_LARGE"

	// System errors
	VertexInternalError   = "INTERNAL"
	VertexUnavailable     = "UNAVAILABLE"
	VertexTimeout         = "DEADLINE_EXCEEDED"
)

// MapVertexError maps Vertex AI error codes to standard error codes
func MapVertexError(code string, message string) errors.ErrorCode {
	// Check specific error codes
	switch code {
	case VertexChildSafety:
		return errors.AI_VIOLATION_CHILD_SAFETY
	case VertexCelebrity:
		return errors.AI_VIOLATION_CELEBRITY
	case VertexViolence:
		return errors.AI_VIOLATION_VIOLENCE
	case VertexSexualContent:
		return errors.AI_VIOLATION_SEXUAL
	case VertexHateSpeech:
		return errors.AI_VIOLATION_HATE_SPEECH
	case VertexDangerousContent:
		return errors.AI_VIOLATION_DANGEROUS
	case VertexQuotaExceeded, VertexRateLimited:
		return errors.RATE_QUOTA_EXCEEDED
	case VertexModelOverloaded:
		return errors.AI_MODEL_OVERLOADED
	case VertexInvalidRequest:
		return errors.VAL_INVALID_REQUEST
	case VertexRequestTooLarge:
		return errors.AI_CONTEXT_LENGTH_EXCEEDED
	case VertexInternalError:
		return errors.SYS_INTERNAL_ERROR
	case VertexUnavailable:
		return errors.SYS_SERVICE_UNAVAILABLE
	case VertexTimeout:
		return errors.SYS_TIMEOUT
	}

	// Check message patterns for additional context
	if strings.Contains(strings.ToLower(message), "child") {
		return errors.AI_VIOLATION_CHILD_SAFETY
	}
	if strings.Contains(strings.ToLower(message), "celebrity") || strings.Contains(strings.ToLower(message), "public figure") {
		return errors.AI_VIOLATION_CELEBRITY
	}
	if strings.Contains(strings.ToLower(message), "violent") || strings.Contains(strings.ToLower(message), "violence") {
		return errors.AI_VIOLATION_VIOLENCE
	}
	if strings.Contains(strings.ToLower(message), "sexual") || strings.Contains(strings.ToLower(message), "adult") {
		return errors.AI_VIOLATION_SEXUAL
	}
	if strings.Contains(strings.ToLower(message), "hate") || strings.Contains(strings.ToLower(message), "discriminat") {
		return errors.AI_VIOLATION_HATE_SPEECH
	}
	if strings.Contains(strings.ToLower(message), "quota") || strings.Contains(strings.ToLower(message), "limit") {
		return errors.RATE_QUOTA_EXCEEDED
	}
	if strings.Contains(strings.ToLower(message), "overload") || strings.Contains(strings.ToLower(message), "capacity") {
		return errors.AI_MODEL_OVERLOADED
	}

	// Default to generation failed
	return errors.AI_GENERATION_FAILED
}

// IsRetryableVertexError determines if a Vertex error is retryable
func IsRetryableVertexError(code string) bool {
	switch code {
	case VertexQuotaExceeded, VertexRateLimited, VertexModelOverloaded:
		return true
	case VertexUnavailable, VertexTimeout:
		return true
	case VertexInternalError:
		return true
	default:
		return false
	}
}