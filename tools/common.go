package tools

// SignedURLs represents URLs with expiration for video processor tools
type SignedURLs struct {
	URLs      []string `json:"urls" validate:"required,min=1"`
	ExpiresAt string   `json:"expires_at,omitempty"`
}