package tools

// MediaReference represents image/video inputs for tools
type MediaReference struct {
	URL      string `json:"url,omitempty"`
	Base64   string `json:"base64,omitempty"`
	MimeType string `json:"mime_type,omitempty" validate:"omitempty,oneof=image/jpeg image/png"`
}

// SignedURLs represents URLs with expiration for video processor tools
type SignedURLs struct {
	URLs      []string `json:"urls" validate:"required,min=1"`
	ExpiresAt string   `json:"expires_at,omitempty"`
}