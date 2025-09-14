package tools

import (
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
)

// GetJSONSchema returns the JSON schema for a tool's parameters
// This is used by Chat AI to configure Gemini's function declarations
func GetJSONSchema(toolName ToolName) (map[string]interface{}, error) {
	reflector := &jsonschema.Reflector{
		RequiredFromJSONSchemaTags: true,
		AllowAdditionalProperties:  false,
	}

	var schema *jsonschema.Schema

	switch toolName {
	// Media Generation Tools
	case GenerateImageImagen:
		schema = reflector.Reflect(&GenerateImageImagenParams{})
	case GenerateImageImagenFast:
		schema = reflector.Reflect(&GenerateImageImagenFastParams{})
	case GenerateImageImagenUltra:
		schema = reflector.Reflect(&GenerateImageImagenUltraParams{})
	case GenerateImageFlash:
		schema = reflector.Reflect(&GenerateImageFlashParams{})
	case NanoBanana:
		schema = reflector.Reflect(&NanoBananaParams{})
	case GenerateVideoVeo3:
		schema = reflector.Reflect(&GenerateVideoVeo3Params{})
	case GenerateVideoVeo3Fast:
		schema = reflector.Reflect(&GenerateVideoVeo3FastParams{})
	case GenerateVideoVeo3FastNoAudio:
		schema = reflector.Reflect(&GenerateVideoVeo3FastNoAudioParams{})
	case GenerateVideoVeo3NoAudio:
		schema = reflector.Reflect(&GenerateVideoVeo3NoAudioParams{})
	case GenerateMusicLyria:
		schema = reflector.Reflect(&GenerateMusicLyriaParams{})

	// Video Processing Tools
	case CombineVideos:
		schema = reflector.Reflect(&CombineVideosParams{})
	case TrimVideo:
		schema = reflector.Reflect(&TrimVideoParams{})
	case ImageAudioMerge:
		schema = reflector.Reflect(&ImageAudioMergeParams{})
	case ExtractFrame:
		schema = reflector.Reflect(&ExtractFrameParams{})
	case MergeImages:
		schema = reflector.Reflect(&MergeImagesParams{})

	// Content Analysis Tools
	case ContentAnalyzer:
		schema = reflector.Reflect(&ContentAnalyzerParams{})
	case GoogleSearch:
		schema = reflector.Reflect(&GoogleSearchParams{})

	default:
		return nil, fmt.Errorf("unknown tool: %s", toolName)
	}

	// Convert to map for easier manipulation
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal schema: %w", err)
	}

	var schemaMap map[string]interface{}
	if err := json.Unmarshal(schemaBytes, &schemaMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal schema: %w", err)
	}

	// Remove $schema and $id fields that Gemini doesn't need
	delete(schemaMap, "$schema")
	delete(schemaMap, "$id")

	return schemaMap, nil
}

// GetAllSchemas returns JSON schemas for all tools
// Useful for Chat AI to cache at startup
func GetAllSchemas() map[ToolName]map[string]interface{} {
	schemas := make(map[ToolName]map[string]interface{})

	tools := []ToolName{
		GenerateImageImagen, GenerateImageImagenFast, GenerateImageImagenUltra,
		GenerateImageFlash, NanoBanana,
		GenerateVideoVeo3, GenerateVideoVeo3Fast, GenerateVideoVeo3FastNoAudio, GenerateVideoVeo3NoAudio,
		GenerateMusicLyria,
		CombineVideos, TrimVideo, ImageAudioMerge, ExtractFrame, MergeImages,
		ContentAnalyzer, GoogleSearch,
	}

	for _, tool := range tools {
		if schema, err := GetJSONSchema(tool); err == nil {
			schemas[tool] = schema
		}
	}

	return schemas
}