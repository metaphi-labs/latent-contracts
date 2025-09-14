package tools

import "fmt"

// === Image Generation Tools (Imagen) ===

// GenerateImageImagenParams for generate-image-imagen tool
type GenerateImageImagenParams struct {
	Prompt                    string  `json:"prompt" validate:"required,min=1,max=2000"`
	Seed                      *int64  `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	Style                     string  `json:"style,omitempty"`
	Language                  string  `json:"language,omitempty" validate:"omitempty,len=2"`
	ImageSize                 string  `json:"image_size,omitempty" validate:"omitempty,oneof=1K 2K"`
	AspectRatio               string  `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=1:1 16:9 9:16 4:3 3:4"`
	AddWatermark              bool    `json:"add_watermark,omitempty"`
	EnhancePrompt             bool    `json:"enhance_prompt,omitempty"`
	GuidanceScale             float64 `json:"guidance_scale,omitempty" validate:"omitempty,min=1.0,max=20.0"`
	NegativePrompt            string  `json:"negative_prompt,omitempty" validate:"omitempty,max=500"`
	NumberOfImages            int     `json:"number_of_images,omitempty" validate:"omitempty,min=1,max=4"`
	OutputMimeType            string  `json:"output_mime_type,omitempty" validate:"omitempty,oneof=image/png image/jpeg"`
	PersonGeneration          string  `json:"person_generation,omitempty" validate:"omitempty,oneof=DONT_ALLOW ALLOW_ADULT ALLOW_ALL"`
	IncludeRaiReason          bool    `json:"include_rai_reason,omitempty"`
	SafetyFilterLevel         string  `json:"safety_filter_level,omitempty" validate:"omitempty,oneof=BLOCK_MEDIUM_AND_ABOVE BLOCK_MOST BLOCK_SOME BLOCK_FEW"`
	IncludeSafetyAttributes   bool    `json:"include_safety_attributes,omitempty"`
	OutputCompressionQuality  int     `json:"output_compression_quality,omitempty" validate:"omitempty,min=1,max=100"`
}

// GenerateImageImagenFastParams for generate-image-imagen-fast tool
type GenerateImageImagenFastParams struct {
	Prompt                    string  `json:"prompt" validate:"required,min=1,max=2000"`
	Seed                      *int64  `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	Style                     string  `json:"style,omitempty"`
	Language                  string  `json:"language,omitempty" validate:"omitempty,len=2"`
	ImageSize                 string  `json:"image_size,omitempty" validate:"omitempty,eq=1K"`
	AspectRatio               string  `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=1:1 16:9 9:16 4:3 3:4"`
	AddWatermark              bool    `json:"add_watermark,omitempty"`
	EnhancePrompt             bool    `json:"enhance_prompt,omitempty"`
	GuidanceScale             float64 `json:"guidance_scale,omitempty" validate:"omitempty,min=1.0,max=20.0"`
	NegativePrompt            string  `json:"negative_prompt,omitempty" validate:"omitempty,max=500"`
	NumberOfImages            int     `json:"number_of_images,omitempty" validate:"omitempty,min=1,max=4"`
	OutputMimeType            string  `json:"output_mime_type,omitempty" validate:"omitempty,oneof=image/png image/jpeg"`
	PersonGeneration          string  `json:"person_generation,omitempty" validate:"omitempty,oneof=DONT_ALLOW ALLOW_ADULT ALLOW_ALL"`
	IncludeRaiReason          bool    `json:"include_rai_reason,omitempty"`
	SafetyFilterLevel         string  `json:"safety_filter_level,omitempty" validate:"omitempty,oneof=BLOCK_MEDIUM_AND_ABOVE BLOCK_MOST BLOCK_SOME BLOCK_FEW"`
	IncludeSafetyAttributes   bool    `json:"include_safety_attributes,omitempty"`
	OutputCompressionQuality  int     `json:"output_compression_quality,omitempty" validate:"omitempty,min=1,max=100"`
}

// GenerateImageImagenUltraParams for generate-image-imagen-ultra tool
type GenerateImageImagenUltraParams struct {
	Prompt                    string  `json:"prompt" validate:"required,min=1,max=2000"`
	Seed                      *int64  `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	Style                     string  `json:"style,omitempty"`
	Language                  string  `json:"language,omitempty" validate:"omitempty,len=2"`
	ImageSize                 string  `json:"image_size,omitempty" validate:"omitempty,oneof=1K 2K"`
	AspectRatio               string  `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=1:1 16:9 9:16 4:3 3:4"`
	AddWatermark              bool    `json:"add_watermark,omitempty"`
	EnhancePrompt             bool    `json:"enhance_prompt,omitempty"`
	GuidanceScale             float64 `json:"guidance_scale,omitempty" validate:"omitempty,min=1.0,max=20.0"`
	NegativePrompt            string  `json:"negative_prompt,omitempty" validate:"omitempty,max=500"`
	NumberOfImages            int     `json:"number_of_images,omitempty" validate:"omitempty,eq=1"`
	OutputMimeType            string  `json:"output_mime_type,omitempty" validate:"omitempty,oneof=image/png image/jpeg"`
	PersonGeneration          string  `json:"person_generation,omitempty" validate:"omitempty,oneof=DONT_ALLOW ALLOW_ADULT ALLOW_ALL"`
	IncludeRaiReason          bool    `json:"include_rai_reason,omitempty"`
	SafetyFilterLevel         string  `json:"safety_filter_level,omitempty" validate:"omitempty,oneof=BLOCK_MEDIUM_AND_ABOVE BLOCK_MOST BLOCK_SOME BLOCK_FEW"`
	IncludeSafetyAttributes   bool    `json:"include_safety_attributes,omitempty"`
	OutputCompressionQuality  int     `json:"output_compression_quality,omitempty" validate:"omitempty,min=1,max=100"`
}

// === Multimodal Image Generation ===

// NanoBananaParams for nano-banana tool (Gemini 2.5 Flash Image Preview)
// Supports native image generation, editing, style transfer, and multi-image composition
type NanoBananaParams struct {
	// Core generation prompt (required for image generation)
	Prompt string `json:"prompt" validate:"required,min=1,max=8000"`

	// Input images for editing, style transfer, or composition (up to 3 recommended)
	StorageURLs []string      `json:"storage_urls,omitempty" validate:"omitempty,max=3"` // GCS URLs for input images
	Context []ContextMessage `json:"context,omitempty" validate:"omitempty,max=10"` // Conversation history for multi-turn

	// Generation parameters
	Temperature     *float64 `json:"temperature,omitempty" validate:"omitempty,min=0,max=2"`
	TopP            *float64 `json:"top_p,omitempty" validate:"omitempty,min=0,max=1"`
	NumberOfImages  int      `json:"number_of_images,omitempty" validate:"omitempty,min=1,max=4"`

	// Safety settings
	SafetyFilterLevel string `json:"safety_filter_level,omitempty" validate:"omitempty,oneof=BLOCK_NONE BLOCK_ONLY_HIGH BLOCK_MEDIUM_AND_ABOVE"`

	// Additional control
	Seed     *int64 `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	Language string `json:"language,omitempty" validate:"omitempty,oneof=en es-MX ja-JP zh-CN hi-IN"` // Best results with these languages
}

// Validate ensures NanoBananaParams is well-formed
func (n *NanoBananaParams) Validate() error {
	// Prompt is required
	if n.Prompt == "" {
		return fmt.Errorf("prompt is required")
	}

	// Validate context messages if provided
	for i, msg := range n.Context {
		if err := msg.Validate(); err != nil {
			return fmt.Errorf("context[%d]: %w", i, err)
		}
	}

	return nil
}

// === Video Generation Tools (Veo3) ===

// GenerateVideoVeo3Params for generate-video-veo3 tool
type GenerateVideoVeo3Params struct {
	Prompt              string          `json:"prompt,omitempty" validate:"required_without=Image,omitempty,min=10,max=2000"`
	Image               *MediaReference `json:"image,omitempty" validate:"required_without=Prompt"`
	LastFrame           *MediaReference `json:"last_frame,omitempty"`
	Seed                *int64          `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	FPS                 int             `json:"fps,omitempty" validate:"omitempty,oneof=24 30 60"`
	Duration            int             `json:"duration,omitempty" validate:"omitempty,min=5,max=8"`
	Resolution          string          `json:"resolution,omitempty" validate:"omitempty,oneof=720p 1080p"`
	AspectRatio         string          `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=16:9 9:16"`
	SampleCount         int             `json:"sample_count,omitempty" validate:"omitempty,min=1,max=4"`
	EnhancePrompt       bool            `json:"enhance_prompt,omitempty"`
	GenerateAudio       bool            `json:"generate_audio,omitempty"`
	NegativePrompt      string          `json:"negative_prompt,omitempty" validate:"omitempty,max=500"`
	PersonGeneration    string          `json:"person_generation,omitempty" validate:"omitempty,oneof=allow_adult dont_allow"`
	CompressionQuality  string          `json:"compression_quality,omitempty" validate:"omitempty,oneof=low medium high"`
}

// GenerateVideoVeo3FastParams for generate-video-veo3-fast tool
type GenerateVideoVeo3FastParams struct {
	Prompt              string  `json:"prompt" validate:"required,min=10,max=2000"`
	Seed                *int64  `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	FPS                 int     `json:"fps,omitempty" validate:"omitempty,oneof=24 30 60"`
	Duration            int     `json:"duration,omitempty" validate:"omitempty,min=5,max=8"`
	AspectRatio         string  `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=16:9 9:16"`
	SampleCount         int     `json:"sample_count,omitempty" validate:"omitempty,min=1,max=4"`
	EnhancePrompt       bool    `json:"enhance_prompt,omitempty"`
	GenerateAudio       bool    `json:"generate_audio,omitempty"`
	NegativePrompt      string  `json:"negative_prompt,omitempty" validate:"omitempty,max=500"`
	PersonGeneration    string  `json:"person_generation,omitempty" validate:"omitempty,oneof=allow_adult dont_allow"`
	CompressionQuality  string  `json:"compression_quality,omitempty" validate:"omitempty,oneof=low medium high"`
}

// GenerateVideoVeo3FastNoAudioParams for generate-video-veo3-fast-no-audio tool
type GenerateVideoVeo3FastNoAudioParams struct {
	Prompt              string  `json:"prompt" validate:"required,min=10,max=2000"`
	Seed                *int64  `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	FPS                 int     `json:"fps,omitempty" validate:"omitempty,oneof=24 30 60"`
	Duration            int     `json:"duration,omitempty" validate:"omitempty,min=5,max=8"`
	AspectRatio         string  `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=16:9 9:16"`
	SampleCount         int     `json:"sample_count,omitempty" validate:"omitempty,min=1,max=4"`
	EnhancePrompt       bool    `json:"enhance_prompt,omitempty"`
	NegativePrompt      string  `json:"negative_prompt,omitempty" validate:"omitempty,max=500"`
	PersonGeneration    string  `json:"person_generation,omitempty" validate:"omitempty,oneof=allow_adult dont_allow"`
	CompressionQuality  string  `json:"compression_quality,omitempty" validate:"omitempty,oneof=low medium high"`
}

// GenerateVideoVeo3NoAudioParams for generate-video-veo3-no-audio tool
type GenerateVideoVeo3NoAudioParams struct {
	Prompt              string          `json:"prompt,omitempty" validate:"required_without=Image,omitempty,min=10,max=2000"`
	Image               *MediaReference `json:"image,omitempty" validate:"required_without=Prompt"`
	LastFrame           *MediaReference `json:"last_frame,omitempty"`
	Seed                *int64          `json:"seed,omitempty" validate:"omitempty,min=0,max=4294967295"`
	FPS                 int             `json:"fps,omitempty" validate:"omitempty,oneof=24 30 60"`
	Duration            int             `json:"duration,omitempty" validate:"omitempty,min=5,max=8"`
	Resolution          string          `json:"resolution,omitempty" validate:"omitempty,oneof=720p 1080p"`
	AspectRatio         string          `json:"aspect_ratio,omitempty" validate:"omitempty,oneof=16:9 9:16"`
	SampleCount         int             `json:"sample_count,omitempty" validate:"omitempty,min=1,max=4"`
	EnhancePrompt       bool            `json:"enhance_prompt,omitempty"`
	NegativePrompt      string          `json:"negative_prompt,omitempty" validate:"omitempty,max=500"`
	PersonGeneration    string          `json:"person_generation,omitempty" validate:"omitempty,oneof=allow_adult dont_allow"`
	CompressionQuality  string          `json:"compression_quality,omitempty" validate:"omitempty,oneof=low medium high"`
}

// === Audio Generation Tools (Lyria) ===

// GenerateMusicLyriaParams for generate-music-lyria tool
type GenerateMusicLyriaParams struct {
	Prompt         string `json:"prompt" validate:"required"`
	Seed           *int   `json:"seed,omitempty"`
	SampleCount    int    `json:"sample_count,omitempty" validate:"omitempty,min=1,max=4"`
	NegativePrompt string `json:"negative_prompt,omitempty"`
}