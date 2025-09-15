package results

import (
	"fmt"
	"time"

	"github.com/metaphi-labs/latent-contracts/errors"
	"github.com/metaphi-labs/latent-contracts/types"
)

// VideoProcessingResult for video manipulation tools
// Used by: combine-videos, trim-video, extract-frame, image-audio-merge, etc.
type VideoProcessingResult struct {
	// Output assets - use the same types as MediaGenerationResult for consistency
	Images []types.OutputImage `json:"images,omitempty"` // For extract-frame producing images
	Videos []types.OutputVideo `json:"videos,omitempty"` // For trim, combine producing videos
	Audio  []types.OutputAudio `json:"audio,omitempty"`  // For audio extraction/processing

	// Input references (what was processed)
	InputAssets []InputReference `json:"input_assets"`

	// Processing operations performed
	Operations []ProcessingOperation `json:"operations"`

	// Processing stats
	ProcessingTime float64 `json:"processing_time_seconds"`
	InputSize      int64   `json:"input_size_bytes"`
	OutputSize     int64   `json:"output_size_bytes"`
}

// InputReference describes an input to the processing
type InputReference struct {
	Type       string `json:"type"`        // "video", "image", "audio"
	SourceURL  string `json:"source_url"`  // Original input URL
	Duration   *float64 `json:"duration,omitempty"`
	Resolution *string  `json:"resolution,omitempty"`
}

// ProcessingOperation describes what was done
type ProcessingOperation struct {
	Type       string                 `json:"type"` // "trim", "combine", "extract", "merge"
	Parameters map[string]interface{} `json:"parameters"`
}

// Constructor functions for Video Processor service to use

// NewTrimVideoResult creates a result for trim-video tool
func NewTrimVideoResult(
	outputVideo types.OutputVideo,
	inputURL string,
	startTime string,
	endTime string,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success:        true,
		Tool:           "trim-video",
		VideoProcessing: &VideoProcessingResult{
			Videos: []types.OutputVideo{outputVideo},
			InputAssets: []InputReference{
				{
					Type:      "video",
					SourceURL: inputURL,
				},
			},
			Operations: []ProcessingOperation{
				{
					Type: "trim",
					Parameters: map[string]interface{}{
						"start_time": startTime,
						"end_time":   endTime,
					},
				},
			},
			ProcessingTime: processingTime,
		},
		Metadata: meta,
	}
}

// NewCombineVideosResult creates a result for combine-videos tool
func NewCombineVideosResult(
	outputVideo types.OutputVideo,
	inputURLs []string,
	transition string,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	// Build input references
	inputs := make([]InputReference, len(inputURLs))
	for i, url := range inputURLs {
		inputs[i] = InputReference{
			Type:      "video",
			SourceURL: url,
		}
	}
	return &ToolResult{
		Success:        true,
		Tool:           "combine-videos",
		VideoProcessing: &VideoProcessingResult{
			Videos:      []types.OutputVideo{outputVideo},
			InputAssets: inputs,
			Operations: []ProcessingOperation{
				{
					Type: "combine",
					Parameters: map[string]interface{}{
						"transition": transition,
						"count":      len(inputURLs),
					},
				},
			},
			ProcessingTime: processingTime,
		},
		Metadata: meta,
	}
}

// NewExtractFrameResult creates a result for extract-frame tool (single frame)
func NewExtractFrameResult(
	outputImage types.OutputImage,
	inputVideoURL string,
	position string,
	timestamp string,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	params := make(map[string]interface{})
	if position != "" {
		params["position"] = position
	}
	if timestamp != "" {
		params["timestamp"] = timestamp
	}
	return &ToolResult{
		Success:        true,
		Tool:           "extract-frame",
		VideoProcessing: &VideoProcessingResult{
			Images: []types.OutputImage{outputImage},
			InputAssets: []InputReference{
				{
					Type:      "video",
					SourceURL: inputVideoURL,
				},
			},
			Operations: []ProcessingOperation{
				{
					Type:       "extract",
					Parameters: params,
				},
			},
			ProcessingTime: processingTime,
		},
		Metadata: meta,
	}
}

// NewExtractFramesResult creates a result for extract-frame tool (batch extraction)
func NewExtractFramesResult(
	outputImages []types.OutputImage,
	inputVideoURL string,
	positions []string,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success:        true,
		Tool:           "extract-frame",
		VideoProcessing: &VideoProcessingResult{
			Images:      outputImages,
			InputAssets: []InputReference{
				{
					Type:      "video",
					SourceURL: inputVideoURL,
				},
			},
			Operations: []ProcessingOperation{
				{
					Type: "extract",
					Parameters: map[string]interface{}{
						"positions": positions,
						"count":     len(positions),
					},
				},
			},
			ProcessingTime: processingTime,
		},
		Metadata: meta,
	}
}

// NewImageAudioMergeResult creates a result for image-audio-merge tool
func NewImageAudioMergeResult(
	outputVideo types.OutputVideo,
	imageURL string,
	audioURL string,
	audioDuration float64,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success:        true,
		Tool:           "image-audio-merge",
		VideoProcessing: &VideoProcessingResult{
			Videos: []types.OutputVideo{outputVideo},
			InputAssets: []InputReference{
				{
					Type:      "image",
					SourceURL: imageURL,
				},
				{
					Type:      "audio",
					SourceURL: audioURL,
					Duration:  &audioDuration,
				},
			},
			Operations: []ProcessingOperation{
				{
					Type: "merge",
					Parameters: map[string]interface{}{
						"output_duration": audioDuration,
					},
				},
			},
			ProcessingTime: processingTime,
		},
		Metadata: meta,
	}
}

// NewMergeImagesResult creates a result for merge-images tool
func NewMergeImagesResult(
	outputImage types.OutputImage,
	inputImageURLs []string,
	layout string,
	spacing int,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	// Build input references
	inputs := make([]InputReference, len(inputImageURLs))
	for i, url := range inputImageURLs {
		inputs[i] = InputReference{
			Type:      "image",
			SourceURL: url,
		}
	}
	return &ToolResult{
		Success:        true,
		Tool:           "merge-images",
		VideoProcessing: &VideoProcessingResult{
			Images:      []types.OutputImage{outputImage},
			InputAssets: inputs,
			Operations: []ProcessingOperation{
				{
					Type: "merge",
					Parameters: map[string]interface{}{
						"layout":  layout,
						"spacing": spacing,
						"count":   len(inputImageURLs),
					},
				},
			},
			ProcessingTime: processingTime,
		},
		Metadata: meta,
	}
}

// NewVideoProcessingError creates an error result for video processing
func NewVideoProcessingError(
	tool string,
	code errors.ErrorCode,
	message string,
	service string,
	retryable bool,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success:        false,
		Tool:           tool,
		Error: errors.NewServiceError(code, message, service, retryable),
		Metadata: meta,
	}
}

// Validate ensures video processing result is well-formed
func (v *VideoProcessingResult) Validate() error {
	// Check that at least one output type has content
	totalOutputs := len(v.Images) + len(v.Videos) + len(v.Audio)
	if totalOutputs == 0 {
		return fmt.Errorf("video processing must have at least one output asset")
	}

	if len(v.InputAssets) == 0 {
		return fmt.Errorf("video processing must have at least one input asset")
	}

	if len(v.Operations) == 0 {
		return fmt.Errorf("video processing must specify at least one operation")
	}

	// Validate input references
	for i, input := range v.InputAssets {
		if input.SourceURL == "" {
			return fmt.Errorf("input[%d]: source URL is required", i)
		}
		if input.Type == "" {
			return fmt.Errorf("input[%d]: type is required", i)
		}
	}

	return nil
}

// Helper to create execution metadata for Video Processor
func NewVideoProcessorMetadata(
	startTime time.Time,
	creditsUsed int,
	requestID string,
	conversationID string,
	userID string,
) ExecutionMetadata {
	endTime := time.Now()
	return ExecutionMetadata{
		StartTime:      startTime,
		EndTime:        endTime,
		DurationMs:     endTime.Sub(startTime).Milliseconds(),
		CreatedAt:      time.Now(), // Add CreatedAt field
		CreditsUsed:    creditsUsed,
		Provider:       "video-processor",
		Model:          "ffmpeg", // Add Model field for consistency
		RequestID:      requestID,
	}
}