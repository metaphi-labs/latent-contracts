# Latent Contracts

Shared type definitions and contracts for all Latent platform services.

## Purpose

This package provides a single source of truth for:
- Error codes and structures
- API contracts
- Shared types

All Latent services (Platform API, Media AI, Chat AI, etc.) import this package to ensure type safety and consistency.

## Installation

```bash
go get github.com/metaphi-labs/latent-contracts
```

## Usage

### In Media AI

```go
import (
    "github.com/metaphi-labs/latent-contracts/errors"
)

// When Vertex AI returns a child safety violation
func handleVertexAIError(supportCode string) *errors.ServiceError {
    if supportCode == "58061214" {
        return errors.ContentViolationError(
            "media-ai",
            []errors.ViolationDetail{{
                Type:         "CHILD_SAFETY",
                Description:  "Content depicts children",
                Severity:     errors.SeverityHigh,
                ProviderCode: supportCode,
            }},
        )
    }
}
```

### In Platform API

```go
import (
    "github.com/metaphi-labs/latent-contracts/errors"
)

// Receive and handle errors from Media AI
func handleCallback(err *errors.ServiceError) {
    switch err.Code {
    case errors.AI_VIOLATION_CHILD_SAFETY:
        // Handle child safety violation
        logViolation(err.Metadata.ViolationDetails)
        notifyUser("Content blocked due to child safety policies")
        
    case errors.MEDIA_INVALID_DIMENSIONS:
        // Handle dimension error
        expected := err.Metadata.ValidationDetails[0].Expected
        notifyUser(fmt.Sprintf("Invalid dimensions. Expected: %v", expected))
        
    case errors.AI_MODEL_OVERLOADED:
        if err.Retryable {
            scheduleRetry(err.Metadata.RetryAfter)
        }
    }
}
```

### In Chat AI

```go
import (
    "github.com/metaphi-labs/latent-contracts/errors"
)

// Use same error types
func handleToolError() *errors.ServiceError {
    return errors.NewServiceError(
        errors.TOOL_EXECUTION_FAILED,
        "Failed to execute code analysis tool",
        "chat-ai",
        false,
    )
}
```

## Error Structure

All services return errors in this format:

```json
{
    "code": "AI_VIOLATION_CHILD_SAFETY",
    "message": "Content depicts children",
    "service": "media-ai",
    "retryable": false,
    "occurred_at": "2024-01-15T10:30:00Z",
    "job_id": "job-123",
    "metadata": {
        "violation_details": [{
            "type": "CHILD_SAFETY",
            "description": "Content depicts children",
            "severity": "high",
            "provider_code": "58061214"
        }],
        "provider": "vertex_ai"
    }
}
```

## Benefits

1. **Type Safety** - Compile-time checking of error codes
2. **No Mappings** - Services use same types directly
3. **Consistency** - Same error structure everywhere
4. **Documentation** - Self-documenting error codes
5. **Evolution** - Versioned package for gradual updates

## Versioning

This package follows semantic versioning:
- **Patch** (1.0.1): Bug fixes, documentation
- **Minor** (1.1.0): New error codes (backward compatible)
- **Major** (2.0.0): Breaking changes (rare)

## Adding New Error Codes

1. Add the error code to `errors/codes.go`
2. Update services to handle new code
3. Create new version tag
4. Services update at their own pace

## Development

```bash
# Run tests
go test ./...

# Tag new version
git tag v1.0.0
git push --tags
```

## Services Using This Package

- Platform API (`platform-api`)
- Media AI (`media-ai`)
- Chat AI (`chat-ai`) - planned
- Video Processor (`video-processor`) - planned
- OAuth Service (`oauth-service`) - planned

## License

Internal use only - Latent Platform