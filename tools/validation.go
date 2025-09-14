package tools

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Validate implementations for Media Generation tools

func (p GenerateImageImagenParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GenerateImageImagenFastParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GenerateImageImagenUltraParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GenerateImageFlashParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GeneralImageFlashParams) Validate() error {
	// No required fields for this tool
	return nil
}

func (p GenerateVideoVeo3Params) Validate() error {
	// Custom validation for required_without
	if p.Prompt == "" && p.Image == nil {
		return fmt.Errorf("either prompt or image is required")
	}
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GenerateVideoVeo3FastParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GenerateVideoVeo3FastNoAudioParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GenerateVideoVeo3NoAudioParams) Validate() error {
	// Same as GenerateVideoVeo3Params
	if p.Prompt == "" && p.Image == nil {
		return fmt.Errorf("either prompt or image is required")
	}
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GenerateMusicLyriaParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

// Validate implementations for Video Processing tools

func (p CombineVideosParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p TrimVideoParams) Validate() error {
	// Custom validation for required_without
	if p.EndTime == "" && p.Duration == nil {
		return fmt.Errorf("either end_time or duration is required")
	}
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p ImageAudioMergeParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p ExtractFrameParams) Validate() error {
	// Custom validation for required_without
	if p.Position == "" && p.Timestamp == "" {
		return fmt.Errorf("either position or timestamp is required")
	}
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p MergeImagesParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

// Validate implementations for Content Analysis tools

func (p ContentAnalyzerParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (p GoogleSearchParams) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}