package tools

// === Video Processing Tools ===

// CombineVideosParams for combine-videos tool
type CombineVideosParams struct {
	VideoURLs     []string `json:"video_urls" validate:"required,min=2,max=10"`
	SignedURLs    []string `json:"signed_urls,omitempty"` // Platform API adds these
	VideoIDs      []string `json:"video_ids,omitempty"`   // For tracking source videos
	Transition    string   `json:"transition,omitempty" validate:"omitempty,oneof=none fade dissolve"`
	FadeDuration  float64  `json:"fade_duration,omitempty" validate:"omitempty,min=0,max=5"`
	AudioStrategy string   `json:"audio_strategy,omitempty" validate:"omitempty,oneof=crossfade concat none cut_continue"`
	Format        string   `json:"format,omitempty" validate:"omitempty,oneof=mp4 webm"`
	VideoCodec    string   `json:"video_codec,omitempty"` // libx264, libx265, etc.
	AudioCodec    string   `json:"audio_codec,omitempty"` // aac, mp3, etc.
}

// TrimVideoParams for trim-video tool
type TrimVideoParams struct {
	VideoURL  string   `json:"video_url" validate:"required"`
	SignedURL string   `json:"signed_url,omitempty"` // Platform API adds this
	StartTime string   `json:"start_time,omitempty" validate:"omitempty"`
	EndTime   string   `json:"end_time,omitempty" validate:"required_without=Duration"`
	Duration  *float64 `json:"duration,omitempty" validate:"required_without=EndTime,omitempty,min=0.1,max=600"`
	FastMode  *bool    `json:"fast_mode,omitempty"`
	Format    string   `json:"format,omitempty" validate:"omitempty,oneof=mp4 webm"`
}

// ImageAudioMergeParams for image-audio-merge tool
type ImageAudioMergeParams struct {
	ImageURL     string `json:"image_url" validate:"required"`
	AudioURL     string `json:"audio_url" validate:"required"`
	SignedImageURL string `json:"signed_image_url,omitempty"` // Platform API adds this
	SignedAudioURL string `json:"signed_audio_url,omitempty"` // Platform API adds this
	Format       string `json:"format,omitempty" validate:"omitempty,oneof=mp4 webm"`
	Resolution   string `json:"resolution,omitempty" validate:"omitempty,oneof=1920x1080 1280x720 854x480 640x360"`
	VideoBitrate string `json:"video_bitrate,omitempty" validate:"omitempty"`
	AudioBitrate string `json:"audio_bitrate,omitempty" validate:"omitempty"`
}

// ExtractFrameParams for extract-frame tool
type ExtractFrameParams struct {
	VideoURL  string   `json:"video_url" validate:"required"`
	SignedURL string   `json:"signed_url,omitempty"` // Platform API adds this
	Position  string   `json:"position,omitempty" validate:"omitempty,oneof=first last middle,required_without_all=Timestamp Positions"`
	Positions []string `json:"positions,omitempty" validate:"required_without_all=Position Timestamp"` // For batch extraction
	Timestamp string   `json:"timestamp,omitempty" validate:"required_without_all=Position Positions"`
	Format    string   `json:"format,omitempty" validate:"omitempty,oneof=jpg png"`
	Quality   int      `json:"quality,omitempty" validate:"omitempty,min=1,max=100"`
	Width     int      `json:"width,omitempty" validate:"omitempty,min=1"`
	Height    int      `json:"height,omitempty" validate:"omitempty,min=1"`
}

// === Image Processing Tools ===

// MergeImagesParams for merge-images tool
type MergeImagesParams struct {
	Images      []string `json:"images,omitempty" validate:"required,min=2"`
	SignedURLs  []string `json:"signed_urls,omitempty"` // Platform API adds these
	Layout      string   `json:"layout,omitempty" validate:"omitempty,oneof=horizontal vertical grid"`
	Spacing     int      `json:"spacing,omitempty" validate:"omitempty,min=0"`
}