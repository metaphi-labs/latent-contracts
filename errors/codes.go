// Package errors defines the standardized error contract for all Latent platform services.
// This package provides type-safe error codes and structures that ensure consistency
// across Platform API, Media AI, Chat AI, and other services.
package errors

// ErrorCode represents standardized error codes used across all Latent services.
// These codes are the single source of truth for error handling.
type ErrorCode string

// AI Content Violation Errors
// These represent specific content policy violations from AI providers
const (
	// AI_VIOLATION_CHILD_SAFETY indicates content that depicts or relates to children
	AI_VIOLATION_CHILD_SAFETY ErrorCode = "AI_VIOLATION_CHILD_SAFETY"
	
	// AI_VIOLATION_CELEBRITY indicates photorealistic depiction of public figures
	AI_VIOLATION_CELEBRITY ErrorCode = "AI_VIOLATION_CELEBRITY"
	
	// AI_VIOLATION_VIOLENCE indicates violent or harmful content
	AI_VIOLATION_VIOLENCE ErrorCode = "AI_VIOLATION_VIOLENCE"
	
	// AI_VIOLATION_SEXUAL indicates sexual or adult content
	AI_VIOLATION_SEXUAL ErrorCode = "AI_VIOLATION_SEXUAL"
	
	// AI_VIOLATION_HATE_SPEECH indicates hate speech or discriminatory content
	AI_VIOLATION_HATE_SPEECH ErrorCode = "AI_VIOLATION_HATE_SPEECH"
	
	// AI_VIOLATION_PERSONAL_INFO indicates presence of PII data
	AI_VIOLATION_PERSONAL_INFO ErrorCode = "AI_VIOLATION_PERSONAL_INFO"
	
	// AI_VIOLATION_TOXIC indicates toxic or harmful language
	AI_VIOLATION_TOXIC ErrorCode = "AI_VIOLATION_TOXIC"
	
	// AI_VIOLATION_DANGEROUS indicates potentially dangerous content
	AI_VIOLATION_DANGEROUS ErrorCode = "AI_VIOLATION_DANGEROUS"
	
	// AI_VIOLATION_PROHIBITED indicates other prohibited content
	AI_VIOLATION_PROHIBITED ErrorCode = "AI_VIOLATION_PROHIBITED"
	
	// AI_VIOLATION_VULGAR indicates vulgar or inappropriate content
	AI_VIOLATION_VULGAR ErrorCode = "AI_VIOLATION_VULGAR"
	
	// AI_VIOLATION_OTHER indicates unspecified content violations
	AI_VIOLATION_OTHER ErrorCode = "AI_VIOLATION_OTHER"
)

// Media Validation Errors
// These represent media-specific validation failures
const (
	// MEDIA_INVALID_DIMENSIONS indicates image/video dimensions are invalid
	MEDIA_INVALID_DIMENSIONS ErrorCode = "MEDIA_INVALID_DIMENSIONS"
	
	// MEDIA_INVALID_ASPECT_RATIO indicates aspect ratio is not supported
	MEDIA_INVALID_ASPECT_RATIO ErrorCode = "MEDIA_INVALID_ASPECT_RATIO"
	
	// MEDIA_INVALID_DURATION indicates video duration is out of range
	MEDIA_INVALID_DURATION ErrorCode = "MEDIA_INVALID_DURATION"
	
	// MEDIA_INVALID_FRAME_RATE indicates unsupported frame rate
	MEDIA_INVALID_FRAME_RATE ErrorCode = "MEDIA_INVALID_FRAME_RATE"
	
	// MEDIA_UNSUPPORTED_FORMAT indicates file format is not supported
	MEDIA_UNSUPPORTED_FORMAT ErrorCode = "MEDIA_UNSUPPORTED_FORMAT"
	
	// MEDIA_SIZE_TOO_LARGE indicates file size exceeds limits
	MEDIA_SIZE_TOO_LARGE ErrorCode = "MEDIA_SIZE_TOO_LARGE"
	
	// MEDIA_PROCESSING_FAILED indicates media processing error
	MEDIA_PROCESSING_FAILED ErrorCode = "MEDIA_PROCESSING_FAILED"
	
	// MEDIA_CORRUPTED indicates media file is corrupted
	MEDIA_CORRUPTED ErrorCode = "MEDIA_CORRUPTED"
)

// AI Model and Generation Errors
const (
	// AI_MODEL_UNAVAILABLE indicates the requested model is not available
	AI_MODEL_UNAVAILABLE ErrorCode = "AI_MODEL_UNAVAILABLE"
	
	// AI_MODEL_OVERLOADED indicates the model is at capacity
	AI_MODEL_OVERLOADED ErrorCode = "AI_MODEL_OVERLOADED"
	
	// AI_CONTEXT_LENGTH_EXCEEDED indicates input exceeds context window
	AI_CONTEXT_LENGTH_EXCEEDED ErrorCode = "AI_CONTEXT_LENGTH_EXCEEDED"
	
	// AI_GENERATION_FAILED indicates generation process failed
	AI_GENERATION_FAILED ErrorCode = "AI_GENERATION_FAILED"
	
	// AI_OPERATION_FAILED indicates async operation failed
	AI_OPERATION_FAILED ErrorCode = "AI_OPERATION_FAILED"
	
	// AI_INVALID_MODEL indicates model name/version is invalid
	AI_INVALID_MODEL ErrorCode = "AI_INVALID_MODEL"
)

// Validation Errors
const (
	// VAL_INVALID_REQUEST indicates malformed request structure
	VAL_INVALID_REQUEST ErrorCode = "VAL_INVALID_REQUEST"
	
	// VAL_MISSING_PARAMETER indicates required parameter is missing
	VAL_MISSING_PARAMETER ErrorCode = "VAL_MISSING_PARAMETER"
	
	// VAL_INVALID_PARAMETER indicates parameter value is invalid
	VAL_INVALID_PARAMETER ErrorCode = "VAL_INVALID_PARAMETER"
	
	// VAL_INVALID_FORMAT indicates data format is incorrect
	VAL_INVALID_FORMAT ErrorCode = "VAL_INVALID_FORMAT"
	
	// VAL_OUT_OF_RANGE indicates value is outside acceptable range
	VAL_OUT_OF_RANGE ErrorCode = "VAL_OUT_OF_RANGE"
)

// System and Infrastructure Errors
const (
	// SYS_INTERNAL_ERROR indicates internal server error
	SYS_INTERNAL_ERROR ErrorCode = "SYS_INTERNAL_ERROR"
	
	// SYS_SERVICE_UNAVAILABLE indicates service is temporarily unavailable
	SYS_SERVICE_UNAVAILABLE ErrorCode = "SYS_SERVICE_UNAVAILABLE"
	
	// SYS_TIMEOUT indicates operation timed out
	SYS_TIMEOUT ErrorCode = "SYS_TIMEOUT"
	
	// SYS_NETWORK_ERROR indicates network connectivity issues
	SYS_NETWORK_ERROR ErrorCode = "SYS_NETWORK_ERROR"
	
	// SYS_DATABASE_ERROR indicates database operation failed
	SYS_DATABASE_ERROR ErrorCode = "SYS_DATABASE_ERROR"
	
	// SYS_STORAGE_ERROR indicates storage operation failed
	SYS_STORAGE_ERROR ErrorCode = "SYS_STORAGE_ERROR"
)

// Rate Limiting Errors
const (
	// RATE_LIMIT_EXCEEDED indicates rate limit has been exceeded
	RATE_LIMIT_EXCEEDED ErrorCode = "RATE_LIMIT_EXCEEDED"
	
	// RATE_QUOTA_EXCEEDED indicates quota has been exhausted
	RATE_QUOTA_EXCEEDED ErrorCode = "RATE_QUOTA_EXCEEDED"
)

// Authentication and Authorization Errors
const (
	// AUTH_UNAUTHORIZED indicates missing or invalid authentication
	AUTH_UNAUTHORIZED ErrorCode = "AUTH_UNAUTHORIZED"
	
	// AUTH_FORBIDDEN indicates authenticated but not authorized
	AUTH_FORBIDDEN ErrorCode = "AUTH_FORBIDDEN"
	
	// AUTH_TOKEN_EXPIRED indicates authentication token has expired
	AUTH_TOKEN_EXPIRED ErrorCode = "AUTH_TOKEN_EXPIRED"
	
	// AUTH_INVALID_TOKEN indicates token is malformed or invalid
	AUTH_INVALID_TOKEN ErrorCode = "AUTH_INVALID_TOKEN"
)

// Billing and Credits Errors
const (
	// BILL_INSUFFICIENT_CREDITS indicates not enough credits
	BILL_INSUFFICIENT_CREDITS ErrorCode = "BILL_INSUFFICIENT_CREDITS"
	
	// BILL_PAYMENT_REQUIRED indicates payment is needed
	BILL_PAYMENT_REQUIRED ErrorCode = "BILL_PAYMENT_REQUIRED"
	
	// BILL_SUBSCRIPTION_EXPIRED indicates subscription has expired
	BILL_SUBSCRIPTION_EXPIRED ErrorCode = "BILL_SUBSCRIPTION_EXPIRED"
)

// Tool Execution Errors (for Chat AI and Platform API)
const (
	// TOOL_NOT_FOUND indicates requested tool doesn't exist
	TOOL_NOT_FOUND ErrorCode = "TOOL_NOT_FOUND"
	
	// TOOL_EXECUTION_FAILED indicates tool execution failed
	TOOL_EXECUTION_FAILED ErrorCode = "TOOL_EXECUTION_FAILED"
	
	// TOOL_TIMEOUT indicates tool execution timed out
	TOOL_TIMEOUT ErrorCode = "TOOL_TIMEOUT"
	
	// TOOL_INVALID_PARAMS indicates invalid tool parameters
	TOOL_INVALID_PARAMS ErrorCode = "TOOL_INVALID_PARAMS"
)

// Conversation and Message Errors (for Chat AI)
const (
	// CONV_NOT_FOUND indicates conversation doesn't exist
	CONV_NOT_FOUND ErrorCode = "CONV_NOT_FOUND"
	
	// CONV_MESSAGE_NOT_FOUND indicates message doesn't exist
	CONV_MESSAGE_NOT_FOUND ErrorCode = "CONV_MESSAGE_NOT_FOUND"
	
	// CONV_MESSAGE_TOO_LONG indicates message exceeds length limit
	CONV_MESSAGE_TOO_LONG ErrorCode = "CONV_MESSAGE_TOO_LONG"
)