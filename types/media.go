package types

import "time"

// Input media structures - for tool parameters

// InputImage represents an image provided as input to a tool
type InputImage struct {
	StorageURL string `json:"storage_url"`           // GCS URL (gs://)
	MimeType   string `json:"mime_type"`             // MIME type
	FileSize   int64  `json:"file_size_bytes"`       // Required for input validation
	Width      *int   `json:"width,omitempty"`       // Optional dimensions
	Height     *int   `json:"height,omitempty"`
}

// InputVideo represents a video provided as input to a tool
type InputVideo struct {
	StorageURL string   `json:"storage_url"`         // GCS URL (gs://)
	MimeType   string   `json:"mime_type"`           // MIME type
	FileSize   int64    `json:"file_size_bytes"`     // Required for input validation
	Duration   *float64 `json:"duration_seconds,omitempty"`
	Width      *int     `json:"width,omitempty"`
	Height     *int     `json:"height,omitempty"`
}

// InputAudio represents audio provided as input to a tool
type InputAudio struct {
	StorageURL string   `json:"storage_url"`         // GCS URL (gs://)
	MimeType   string   `json:"mime_type"`           // MIME type
	FileSize   int64    `json:"file_size_bytes"`     // Required for input validation
	Duration   *float64 `json:"duration_seconds,omitempty"`
	SampleRate *int     `json:"sample_rate,omitempty"`
	Channels   *int     `json:"channels,omitempty"`
}

// Output media structures - for tool results

// OutputImage represents a generated image in tool results
type OutputImage struct {
	ID         string     `json:"id"`                   // Unique identifier
	Index      int        `json:"index"`                // Order in batch
	StorageURL string     `json:"storage_url"`          // GCS URL (gs://)
	PublicURL  string     `json:"public_url"`           // HTTPS URL
	MimeType   string     `json:"mime_type"`
	Width      *int       `json:"width,omitempty"`
	Height     *int       `json:"height,omitempty"`

	// Platform API adds these for client access
	SignedURL    *string    `json:"signed_url,omitempty"`
	SignedExpiry *time.Time `json:"signed_expiry,omitempty"`
}

// OutputVideo represents a generated video in tool results
type OutputVideo struct {
	ID         string     `json:"id"`                   // Unique identifier
	Index      int        `json:"index"`                // Order in batch
	StorageURL string     `json:"storage_url"`          // GCS URL (gs://)
	PublicURL  string     `json:"public_url"`           // HTTPS URL
	MimeType   string     `json:"mime_type"`
	Duration   *float64   `json:"duration_seconds,omitempty"`
	Width      *int       `json:"width,omitempty"`
	Height     *int       `json:"height,omitempty"`

	// Platform API adds these for client access
	SignedURL    *string    `json:"signed_url,omitempty"`
	SignedExpiry *time.Time `json:"signed_expiry,omitempty"`
}

// OutputAudio represents generated audio in tool results
type OutputAudio struct {
	ID         string     `json:"id"`                   // Unique identifier
	Index      int        `json:"index"`                // Order in batch
	StorageURL string     `json:"storage_url"`          // GCS URL (gs://)
	PublicURL  string     `json:"public_url"`           // HTTPS URL
	MimeType   string     `json:"mime_type"`
	Duration   *float64   `json:"duration_seconds,omitempty"`
	SampleRate *int       `json:"sample_rate,omitempty"`
	Channels   *int       `json:"channels,omitempty"`

	// Platform API adds these for client access
	SignedURL    *string    `json:"signed_url,omitempty"`
	SignedExpiry *time.Time `json:"signed_expiry,omitempty"`
}