package results

import (
	"fmt"
	"time"
)

// VideoProcessingResult for video manipulation tools
// Used by: combine-videos, trim-video, extract-frame, image-audio-merge, etc.
type VideoProcessingResult struct {
	// Output assets
	OutputAssets []MediaAsset `json:"output_assets"`

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
	jobID string,
	outputAsset MediaAsset,
	inputURL string,
	startTime string,
	endTime string,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success: true,
		Tool:    "trim-video",
		JobID:   jobID,
		VideoProcessing: &VideoProcessingResult{
			OutputAssets: []MediaAsset{outputAsset},
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
			OutputSize:     outputAsset.FileSize,
		},
		Metadata: meta,
	}
}

// NewCombineVideosResult creates a result for combine-videos tool
func NewCombineVideosResult(
	jobID string,
	outputAsset MediaAsset,
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
		Success: true,
		Tool:    "combine-videos",
		JobID:   jobID,
		VideoProcessing: &VideoProcessingResult{
			OutputAssets: []MediaAsset{outputAsset},
			InputAssets:  inputs,
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
			OutputSize:     outputAsset.FileSize,
		},
		Metadata: meta,
	}
}

// NewExtractFrameResult creates a result for extract-frame tool (single frame)
func NewExtractFrameResult(
	jobID string,
	outputAsset MediaAsset,
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
		Success: true,
		Tool:    "extract-frame",
		JobID:   jobID,
		VideoProcessing: &VideoProcessingResult{
			OutputAssets: []MediaAsset{outputAsset},
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
			OutputSize:     outputAsset.FileSize,
		},
		Metadata: meta,
	}
}

// NewExtractFramesResult creates a result for extract-frame tool (batch extraction)
func NewExtractFramesResult(
	jobID string,
	outputAssets []MediaAsset,
	inputVideoURL string,
	positions []string,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	// Calculate total output size
	var totalOutputSize int64
	for _, asset := range outputAssets {
		totalOutputSize += asset.FileSize
	}

	return &ToolResult{
		Success: true,
		Tool:    "extract-frame",
		JobID:   jobID,
		VideoProcessing: &VideoProcessingResult{
			OutputAssets: outputAssets,
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
			OutputSize:     totalOutputSize,
		},
		Metadata: meta,
	}
}

// NewImageAudioMergeResult creates a result for image-audio-merge tool
func NewImageAudioMergeResult(
	jobID string,
	outputAsset MediaAsset,
	imageURL string,
	audioURL string,
	audioDuration float64,
	processingTime float64,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success: true,
		Tool:    "image-audio-merge",
		JobID:   jobID,
		VideoProcessing: &VideoProcessingResult{
			OutputAssets: []MediaAsset{outputAsset},
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
			OutputSize:     outputAsset.FileSize,
		},
		Metadata: meta,
	}
}

// NewMergeImagesResult creates a result for merge-images tool
func NewMergeImagesResult(
	jobID string,
	outputAsset MediaAsset,
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
		Success: true,
		Tool:    "merge-images",
		JobID:   jobID,
		VideoProcessing: &VideoProcessingResult{
			OutputAssets: []MediaAsset{outputAsset},
			InputAssets:  inputs,
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
			OutputSize:     outputAsset.FileSize,
		},
		Metadata: meta,
	}
}

// NewVideoProcessingError creates an error result for video processing
func NewVideoProcessingError(
	tool string,
	jobID string,
	code string,
	message string,
	retryable bool,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success: false,
		Tool:    tool,
		JobID:   jobID,
		Error: &ErrorInfo{
			Code:      code,
			Message:   message,
			Retryable: retryable,
		},
		Metadata: meta,
	}
}

// Validate ensures video processing result is well-formed
func (v *VideoProcessingResult) Validate() error {
	if len(v.OutputAssets) == 0 {
		return fmt.Errorf("video processing must have at least one output asset")
	}

	if len(v.InputAssets) == 0 {
		return fmt.Errorf("video processing must have at least one input asset")
	}

	if len(v.Operations) == 0 {
		return fmt.Errorf("video processing must specify at least one operation")
	}

	// Validate output assets
	for i, asset := range v.OutputAssets {
		if err := ValidateMediaAsset(asset, i); err != nil {
			return fmt.Errorf("output %s", err)
		}
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
		CreditsUsed:    creditsUsed,
		Provider:       "video-processor",
		RequestID:      requestID,
		ConversationID: conversationID,
		UserID:         userID,
	}
}