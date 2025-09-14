# Tools Package - Type-Safe Tool Parameter Contracts

This package provides type-safe parameter definitions and validation for all Latent platform tools.

## Usage

### Chat AI - Validating Gemini Parameters

```go
import (
    "github.com/metaphi-labs/latent-contracts/tools"
)

// When Gemini returns a tool call
func handleGeminiToolCall(toolName string, params map[string]interface{}) error {
    // Validate parameters against contract
    err := tools.ParseAndValidateParams(toolName, params)
    if err != nil {
        // Return user-friendly error
        return fmt.Errorf("Invalid parameters for %s: %w", toolName, err)
    }

    // Send validated params to Platform API
    return platformClient.ExecuteTool(toolName, params)
}
```

### Platform API - Validating Incoming Requests

```go
import (
    "github.com/metaphi-labs/latent-contracts/tools"
)

func (e *ToolExecutor) ExecuteTool(ctx context.Context, req ToolRequest) error {
    // Validate tool exists
    if !tools.IsValidToolName(req.Tool) {
        return fmt.Errorf("unknown tool: %s", req.Tool)
    }

    // Validate parameters
    err := tools.ParseAndValidateParams(req.Tool, req.Params)
    if err != nil {
        return fmt.Errorf("invalid parameters: %w", err)
    }

    // Add signed URLs for video processor tools
    if req.Tool == string(tools.CombineVideos) {
        e.addSignedURLs(req.Params)
    }

    // Forward to appropriate service
    return e.forwardToService(ctx, req)
}
```

### Video Processor - Using Typed Parameters

```go
import (
    "encoding/json"
    "github.com/metaphi-labs/latent-contracts/tools"
)

func handleCombineVideos(body []byte) error {
    var params tools.CombineVideosParams
    if err := json.Unmarshal(body, &params); err != nil {
        return err
    }

    // Validate
    if err := params.Validate(); err != nil {
        return err
    }

    // Use signed URLs if provided, otherwise original URLs
    urls := params.VideoURLs
    if len(params.SignedURLs) > 0 {
        urls = params.SignedURLs
    }

    // Process videos
    return combineVideos(urls, params.Transition, params.Format)
}
```

## Available Tools

### Media Generation
- `generate-image-imagen` - Standard quality image generation
- `generate-image-imagen-fast` - Fast, cost-effective image generation
- `generate-image-imagen-ultra` - Premium ultra-high quality images
- `generate-video-veo3` - Premium video generation with audio
- `generate-video-veo3-fast` - Fast video generation
- `generate-music-lyria` - Music generation

### Video Processing
- `combine-videos` - Combine multiple videos with transitions
- `trim-video` - Trim video by time range
- `image-audio-merge` - Create video from image + audio
- `extract-frame` - Extract frames from videos
- `merge-images` - Merge multiple images

### Content Analysis
- `content-analyzer` - Analyze web content
- `google-search` - Search the internet

## Validation

All parameter structs include validation tags using `github.com/go-playground/validator/v10`:

```go
type GenerateImageImagenParams struct {
    Prompt         string `json:"prompt" validate:"required,min=1,max=2000"`
    AspectRatio    string `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=1:1 16:9 9:16 4:3 3:4"`
    NumberOfImages int    `json:"number_of_images,omitempty" validate:"omitempty,min=1,max=4"`
}
```

## Adding New Tools

1. Add the parameter struct in the appropriate file:
   - `media_generation.go` for AI generation tools
   - `video_processing.go` for video/image processing
   - `content_analysis.go` for analysis tools

2. Add the tool name constant in `registry.go`:
   ```go
   const NewTool ToolName = "new-tool"
   ```

3. Add validation case in `registry.go`:
   ```go
   case NewTool:
       var p NewToolParams
       if err := json.Unmarshal(jsonBytes, &p); err != nil {
           return fmt.Errorf("invalid params for %s: %w", toolName, err)
       }
       return p.Validate()
   ```

4. Implement the `Validate()` method in `validation.go`:
   ```go
   func (p NewToolParams) Validate() error {
       if err := validate.Struct(p); err != nil {
           return fmt.Errorf("validation failed: %w", err)
       }
       return nil
   }
   ```

## Benefits

1. **Type Safety**: Compile-time checking for tool parameters
2. **Single Source of Truth**: One place defines what's valid for each tool
3. **Automatic Validation**: Built-in validation with clear error messages
4. **Documentation**: Struct tags document requirements
5. **Cross-Service Consistency**: All services use the same contracts