package errors

import "strings"

// determineCategory auto-determines the category from error code
func determineCategory(code ErrorCode) ErrorCategory {
	codeStr := string(code)
	
	if strings.HasPrefix(codeStr, "AI_") {
		return CategoryAI
	}
	if strings.HasPrefix(codeStr, "MEDIA_") {
		return CategoryMedia
	}
	if strings.HasPrefix(codeStr, "VAL_") {
		return CategoryValidation
	}
	if strings.HasPrefix(codeStr, "AUTH_") {
		return CategoryAuth
	}
	if strings.HasPrefix(codeStr, "SYS_") {
		return CategorySystem
	}
	if strings.HasPrefix(codeStr, "BILL_") {
		return CategoryBilling
	}
	if strings.HasPrefix(codeStr, "RATE_") {
		return CategoryRate
	}
	
	return CategorySystem // default
}

// determineSeverity auto-determines severity from error code
func determineSeverity(code ErrorCode) Severity {
	// Critical errors
	criticalCodes := map[ErrorCode]bool{
		SYS_INTERNAL_ERROR:   true,
		SYS_DATABASE_ERROR:   true,
		AUTH_UNAUTHORIZED:    true,
		AUTH_FORBIDDEN:       true,
	}
	
	if criticalCodes[code] {
		return SeverityCritical
	}
	
	// High severity
	highCodes := map[ErrorCode]bool{
		AI_VIOLATION_CHILD_SAFETY:   true,
		AI_VIOLATION_PERSONAL_INFO:  true,
		BILL_INSUFFICIENT_CREDITS:   true,
		SYS_SERVICE_UNAVAILABLE:     true,
	}
	
	if highCodes[code] {
		return SeverityHigh
	}
	
	// Low severity
	lowCodes := map[ErrorCode]bool{
		VAL_INVALID_PARAMETER: true,
		VAL_MISSING_PARAMETER: true,
		RATE_LIMIT_EXCEEDED:   true,
	}
	
	if lowCodes[code] {
		return SeverityLow
	}
	
	// Default to medium
	return SeverityMedium
}

// determineHTTPStatus auto-determines HTTP status from error code
func determineHTTPStatus(code ErrorCode) int {
	// Map error codes to HTTP status codes
	statusMap := map[ErrorCode]int{
		// 400 Bad Request
		VAL_INVALID_REQUEST:          400,
		VAL_INVALID_PARAMETER:        400,
		VAL_MISSING_PARAMETER:        400,
		VAL_INVALID_FORMAT:           400,
		VAL_OUT_OF_RANGE:             400,
		MEDIA_INVALID_DIMENSIONS:     400,
		MEDIA_INVALID_ASPECT_RATIO:   400,
		MEDIA_INVALID_DURATION:       400,
		AI_CONTEXT_LENGTH_EXCEEDED:   400,
		
		// 401 Unauthorized
		AUTH_UNAUTHORIZED:    401,
		AUTH_INVALID_TOKEN:   401,
		AUTH_TOKEN_EXPIRED:   401,
		
		// 402 Payment Required
		BILL_INSUFFICIENT_CREDITS:   402,
		BILL_PAYMENT_REQUIRED:       402,
		BILL_SUBSCRIPTION_EXPIRED:   402,
		
		// 403 Forbidden
		AUTH_FORBIDDEN:               403,
		AI_VIOLATION_CHILD_SAFETY:   403,
		AI_VIOLATION_CELEBRITY:      403,
		AI_VIOLATION_VIOLENCE:       403,
		AI_VIOLATION_SEXUAL:         403,
		AI_VIOLATION_HATE_SPEECH:    403,
		AI_VIOLATION_PERSONAL_INFO:  403,
		AI_VIOLATION_TOXIC:          403,
		AI_VIOLATION_DANGEROUS:      403,
		AI_VIOLATION_PROHIBITED:     403,
		AI_VIOLATION_VULGAR:         403,
		AI_VIOLATION_OTHER:          403,
		
		// 404 Not Found
		TOOL_NOT_FOUND:          404,
		CONV_NOT_FOUND:          404,
		CONV_MESSAGE_NOT_FOUND:  404,
		
		// 413 Payload Too Large
		MEDIA_SIZE_TOO_LARGE:    413,
		CONV_MESSAGE_TOO_LONG:   413,
		
		// 422 Unprocessable Entity
		MEDIA_UNSUPPORTED_FORMAT: 422,
		MEDIA_CORRUPTED:          422,
		AI_INVALID_MODEL:         422,
		
		// 429 Too Many Requests
		RATE_LIMIT_EXCEEDED:     429,
		RATE_QUOTA_EXCEEDED:     429,
		
		// 500 Internal Server Error
		SYS_INTERNAL_ERROR:      500,
		AI_GENERATION_FAILED:    500,
		AI_OPERATION_FAILED:     500,
		MEDIA_PROCESSING_FAILED: 500,
		TOOL_EXECUTION_FAILED:   500,
		SYS_DATABASE_ERROR:      500,
		SYS_STORAGE_ERROR:       500,
		
		// 502 Bad Gateway
		SYS_NETWORK_ERROR:       502,
		
		// 503 Service Unavailable
		SYS_SERVICE_UNAVAILABLE: 503,
		AI_MODEL_UNAVAILABLE:    503,
		AI_MODEL_OVERLOADED:     503,
		
		// 504 Gateway Timeout
		SYS_TIMEOUT:             504,
		TOOL_TIMEOUT:            504,
	}
	
	if status, exists := statusMap[code]; exists {
		return status
	}
	
	// Default based on category
	codeStr := string(code)
	if strings.HasPrefix(codeStr, "VAL_") {
		return 400
	}
	if strings.HasPrefix(codeStr, "AUTH_") {
		return 401
	}
	if strings.HasPrefix(codeStr, "RATE_") {
		return 429
	}
	
	// Default to 500
	return 500
}