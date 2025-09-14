package tools

// SignableFields defines which fields need URL signing for each tool
// This keeps the signing logic centralized in contracts
var SignableFields = map[ToolName][]string{
	// Video processing tools that need signed URLs
	CombineVideos:   {"video_urls"},
	TrimVideo:       {"video_url"},
	ImageAudioMerge: {"image_url", "audio_url"},
	ExtractFrame:    {"video_url"},
	MergeImages:     {"images"},
}

// NeedsSignedURLs checks if a tool requires URL signing
func NeedsSignedURLs(toolName ToolName) bool {
	_, exists := SignableFields[toolName]
	return exists
}

// GetSignableFields returns the fields that need signing for a tool
func GetSignableFields(toolName ToolName) []string {
	return SignableFields[toolName]
}