package results

import (
	"fmt"
	"time"

	"github.com/metaphi-labs/latent-contracts/types"
)

// MediaGenerationResult for image/video/audio generation tools
// Used by: generate-image-imagen, generate-video-veo3, generate-music-lyria, nano-banana, etc.
type MediaGenerationResult struct {
	// Generated assets - use specific output types
	Images []types.OutputImage `json:"images,omitempty"`
	Videos []types.OutputVideo `json:"videos,omitempty"`
	Audio  []types.OutputAudio `json:"audio,omitempty"`

	// Generation parameters that were actually used
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
	Seed   *int   `json:"seed,omitempty"`

	// For image generation
	AspectRatio     *string `json:"aspect_ratio,omitempty"`      // "1:1", "16:9", etc.
	NegativePrompt  *string `json:"negative_prompt,omitempty"`
	SafetyLevel     *string `json:"safety_level,omitempty"`
	PersonGeneration *bool   `json:"person_generation,omitempty"`

	// For video generation
	StartImage      *string `json:"start_image_url,omitempty"`   // For image-to-video
	EndImage        *string `json:"end_image_url,omitempty"`     // For controlled endings
	CameraMovement  *string `json:"camera_movement,omitempty"`   // "zoom_in", "pan_left", etc.
	AudioGenerated  *bool   `json:"audio_generated,omitempty"`

	// For audio generation
	Genre           *string `json:"genre,omitempty"`
	Instruments     *string `json:"instruments,omitempty"`
	Mood            *string `json:"mood,omitempty"`

	// Multi-generation info
	TotalRequested  int    `json:"total_requested"`
	TotalGenerated  int    `json:"total_generated"`

}

// Constructor functions for Media AI service to use

// NewImageGenerationResult creates a result for Imagen tools
func NewImageGenerationResult(
	tool string,
	jobID string,
	userID string,
	conversationID string,
	messageID string,
	images []types.OutputImage,
	prompt string,
	model string,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success:        true,
		Tool:           tool,
		JobID:          jobID,
		UserID:         userID,
		ConversationID: conversationID,
		MessageID:      messageID,
		MediaGeneration: &MediaGenerationResult{
			Images:         images,
			Prompt:         prompt,
			Model:          model,
			TotalRequested: len(images),
			TotalGenerated: len(images),
		},
		Metadata: meta,
	}
}

// NewVideoGenerationResult creates a result for Veo tools
func NewVideoGenerationResult(
	tool string,
	jobID string,
	userID string,
	conversationID string,
	messageID string,
	video types.OutputVideo,
	prompt string,
	model string,
	audioGenerated bool,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success:        true,
		Tool:           tool,
		JobID:          jobID,
		UserID:         userID,
		ConversationID: conversationID,
		MessageID:      messageID,
		MediaGeneration: &MediaGenerationResult{
			Videos:         []types.OutputVideo{video},
			Prompt:         prompt,
			Model:          model,
			AudioGenerated: &audioGenerated,
			TotalRequested: 1,
			TotalGenerated: 1,
		},
		Metadata: meta,
	}
}

// NewAudioGenerationResult creates a result for Lyria tools
func NewAudioGenerationResult(
	tool string,
	jobID string,
	userID string,
	conversationID string,
	messageID string,
	audio []types.OutputAudio,
	prompt string,
	model string,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success:        true,
		Tool:           tool,
		JobID:          jobID,
		UserID:         userID,
		ConversationID: conversationID,
		MessageID:      messageID,
		MediaGeneration: &MediaGenerationResult{
			Audio:          audio,
			Prompt:         prompt,
			Model:          model,
			TotalRequested: len(audio),
			TotalGenerated: len(audio),
		},
		Metadata: meta,
	}
}

// NewMediaGenerationError creates an error result for media generation
func NewMediaGenerationError(
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

// Validate ensures media generation result is well-formed
func (m *MediaGenerationResult) Validate() error {
	// Must have at least one type of media
	if len(m.Images) == 0 && len(m.Videos) == 0 && len(m.Audio) == 0 {
		return fmt.Errorf("media generation must have at least one asset")
	}

	if m.Prompt == "" {
		return fmt.Errorf("prompt is required for media generation")
	}

	if m.Model == "" {
		return fmt.Errorf("model is required for media generation")
	}

	// Validate each image
	for i, img := range m.Images {
		if img.StorageURL == "" {
			return fmt.Errorf("image %d missing storage URL", i)
		}
		if img.PublicURL == "" {
			return fmt.Errorf("image %d missing public URL", i)
		}
	}

	// Validate each video
	for i, vid := range m.Videos {
		if vid.StorageURL == "" {
			return fmt.Errorf("video %d missing storage URL", i)
		}
		if vid.PublicURL == "" {
			return fmt.Errorf("video %d missing public URL", i)
		}
	}

	// Validate each audio
	for i, aud := range m.Audio {
		if aud.StorageURL == "" {
			return fmt.Errorf("audio %d missing storage URL", i)
		}
		if aud.PublicURL == "" {
			return fmt.Errorf("audio %d missing public URL", i)
		}
	}

	return nil
}

// Helper to create execution metadata for Media AI
func NewMediaAIMetadata(
	startTime time.Time,
	creditsUsed int,
	provider string,
	providerJobID string,
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
		Provider:       provider,
		ProviderJobID:  providerJobID,
		RequestID:      requestID,
	}
}