package tools

// === Content Analysis Tools ===

// ContentAnalyzerParams for content-analyzer tool
type ContentAnalyzerParams struct {
	URL          string   `json:"url" validate:"required"`
	FocusAreas   []string `json:"focus_areas,omitempty"`
	AnalysisType string   `json:"analysis_type,omitempty" validate:"omitempty,oneof=summary insights transcript sentiment key-points critique"`
	OutputFormat string   `json:"output_format,omitempty" validate:"omitempty,oneof=detailed concise bullet-points narrative"`
}

// GoogleSearchParams for google-search tool
type GoogleSearchParams struct {
	Query      string `json:"query" validate:"required"`
	Site       string `json:"site,omitempty"`
	Recency    string `json:"recency,omitempty" validate:"omitempty,oneof=hour day week month year all"`
	MaxResults int    `json:"max_results,omitempty" validate:"omitempty,min=1,max=10"`
}