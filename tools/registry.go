package tools

import (
	"encoding/json"
	"fmt"
)

// ToolName represents valid tool names as constants
type ToolName string

const (
	// Media Generation Tools
	GenerateImageImagen          ToolName = "generate-image-imagen"
	GenerateImageImagenFast      ToolName = "generate-image-imagen-fast"
	GenerateImageImagenUltra     ToolName = "generate-image-imagen-ultra"
	GenerateImageFlash           ToolName = "generate-image-flash"
	NanoBanana                   ToolName = "nano-banana"
	GenerateVideoVeo3            ToolName = "generate-video-veo3"
	GenerateVideoVeo3Fast        ToolName = "generate-video-veo3-fast"
	GenerateVideoVeo3FastNoAudio ToolName = "generate-video-veo3-fast-no-audio"
	GenerateVideoVeo3NoAudio     ToolName = "generate-video-veo3-no-audio"
	GenerateMusicLyria           ToolName = "generate-music-lyria"

	// Video Processing Tools
	CombineVideos    ToolName = "combine-videos"
	TrimVideo        ToolName = "trim-video"
	ImageAudioMerge  ToolName = "image-audio-merge"
	ExtractFrame     ToolName = "extract-frame"
	MergeImages      ToolName = "merge-images"

	// Content Analysis Tools
	ContentAnalyzer ToolName = "content-analyzer"
	GoogleSearch    ToolName = "google-search"
)

// ParseAndValidateParams takes raw params and validates them for a specific tool
func ParseAndValidateParams(toolName string, params map[string]interface{}) error {
	// Convert map to JSON for validation
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal params: %w", err)
	}

	// Parse based on tool name
	switch ToolName(toolName) {
	case GenerateImageImagen:
		var p GenerateImageImagenParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateImageImagenFast:
		var p GenerateImageImagenFastParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateImageImagenUltra:
		var p GenerateImageImagenUltraParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateImageFlash:
		var p GenerateImageFlashParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case NanoBanana:
		var p NanoBananaParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateVideoVeo3:
		var p GenerateVideoVeo3Params
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateVideoVeo3Fast:
		var p GenerateVideoVeo3FastParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateVideoVeo3FastNoAudio:
		var p GenerateVideoVeo3FastNoAudioParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateVideoVeo3NoAudio:
		var p GenerateVideoVeo3NoAudioParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GenerateMusicLyria:
		var p GenerateMusicLyriaParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case CombineVideos:
		var p CombineVideosParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case TrimVideo:
		var p TrimVideoParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case ImageAudioMerge:
		var p ImageAudioMergeParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case ExtractFrame:
		var p ExtractFrameParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case MergeImages:
		var p MergeImagesParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case ContentAnalyzer:
		var p ContentAnalyzerParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	case GoogleSearch:
		var p GoogleSearchParams
		if err := json.Unmarshal(jsonBytes, &p); err != nil {
			return fmt.Errorf("invalid params for %s: %w", toolName, err)
		}
		return p.Validate()

	default:
		return fmt.Errorf("unknown tool: %s", toolName)
	}
}

// IsValidToolName checks if a string is a valid tool name
func IsValidToolName(name string) bool {
	switch ToolName(name) {
	case GenerateImageImagen, GenerateImageImagenFast, GenerateImageImagenUltra,
		GenerateImageFlash, NanoBanana,
		GenerateVideoVeo3, GenerateVideoVeo3Fast, GenerateVideoVeo3FastNoAudio, GenerateVideoVeo3NoAudio,
		GenerateMusicLyria,
		CombineVideos, TrimVideo, ImageAudioMerge, ExtractFrame, MergeImages,
		ContentAnalyzer, GoogleSearch:
		return true
	default:
		return false
	}
}