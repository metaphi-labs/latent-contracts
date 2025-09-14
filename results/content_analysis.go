package results

import (
	"fmt"
	"time"
)

// ContentAnalysisResult for content analysis and search tools
// Used by: content-analyzer, google-search, etc.
type ContentAnalysisResult struct {
	// Analysis type
	AnalysisType string `json:"analysis_type"` // "webpage", "search", "document", "media"

	// Core content
	Summary string `json:"summary"` // Always present summary
	Content string `json:"content"` // Full extracted/analyzed content

	// For search results
	SearchResults []SearchResult `json:"search_results,omitempty"`
	TotalResults  int            `json:"total_results,omitempty"`

	// For webpage/document analysis
	Sources []ContentSource `json:"sources,omitempty"`

	// Extracted entities and metadata
	Entities  []ExtractedEntity  `json:"entities,omitempty"`
	Topics    []string           `json:"topics,omitempty"`
	Sentiment *SentimentAnalysis `json:"sentiment,omitempty"`

	// Structured data (tool-specific)
	StructuredData map[string]interface{} `json:"structured_data,omitempty"`
}

// SearchResult for google-search tool
type SearchResult struct {
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Snippet     string  `json:"snippet"`
	Position    int     `json:"position"`
	Relevance   float64 `json:"relevance,omitempty"`
	PublishDate *string `json:"publish_date,omitempty"`
	Author      *string `json:"author,omitempty"`
}

// ContentSource for analyzed content
type ContentSource struct {
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ContentType string    `json:"content_type"` // "article", "video", "pdf", etc.
	ExtractedAt time.Time `json:"extracted_at"`
	WordCount   int       `json:"word_count,omitempty"`
}

// ExtractedEntity from content
type ExtractedEntity struct {
	Type       string  `json:"type"`       // "person", "organization", "location", "date", etc.
	Value      string  `json:"value"`
	Context    string  `json:"context,omitempty"`    // Surrounding text
	Confidence float64 `json:"confidence,omitempty"` // 0-1
	Count      int     `json:"count,omitempty"`      // Occurrences
}

// SentimentAnalysis results
type SentimentAnalysis struct {
	Overall    string  `json:"overall"`    // "positive", "negative", "neutral", "mixed"
	Score      float64 `json:"score"`      // -1 to 1
	Confidence float64 `json:"confidence"` // 0-1
}

// Constructor functions for content analysis services

// NewContentAnalyzerResult creates a result for content-analyzer tool
func NewContentAnalyzerResult(
	jobID string,
	urls []string,
	summary string,
	content string,
	sources []ContentSource,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success: true,
		Tool:    "content-analyzer",
		JobID:   jobID,
		ContentAnalysis: &ContentAnalysisResult{
			AnalysisType: "webpage",
			Summary:      summary,
			Content:      content,
			Sources:      sources,
		},
		Metadata: meta,
	}
}

// NewGoogleSearchResult creates a result for google-search tool
func NewGoogleSearchResult(
	jobID string,
	query string,
	results []SearchResult,
	summary string,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success: true,
		Tool:    "google-search",
		JobID:   jobID,
		ContentAnalysis: &ContentAnalysisResult{
			AnalysisType:  "search",
			Summary:       summary,
			Content:       fmt.Sprintf("Search results for: %s", query),
			SearchResults: results,
			TotalResults:  len(results),
		},
		Metadata: meta,
	}
}

// NewContentAnalysisError creates an error result for content analysis
func NewContentAnalysisError(
	tool string,
	jobID string,
	code string,
	message string,
	retryable bool,
	meta ExecutionMetadata,
) *ToolResult {
	return &ToolResult{
		Success: false,
		Tool:    tool,
		JobID:   jobID,
		Error: &ErrorInfo{
			Code:      code,
			Message:   message,
			Retryable: retryable,
		},
		Metadata: meta,
	}
}

// Validate ensures content analysis result is well-formed
func (c *ContentAnalysisResult) Validate() error {
	if c.AnalysisType == "" {
		return fmt.Errorf("analysis type is required")
	}

	if c.Summary == "" {
		return fmt.Errorf("summary is required for content analysis")
	}

	// Type-specific validation
	switch c.AnalysisType {
	case "search":
		if len(c.SearchResults) == 0 {
			return fmt.Errorf("search results cannot be empty for search analysis")
		}
		for i, result := range c.SearchResults {
			if result.Title == "" {
				return fmt.Errorf("search result[%d]: title is required", i)
			}
			if result.URL == "" {
				return fmt.Errorf("search result[%d]: URL is required", i)
			}
		}

	case "webpage", "document":
		if len(c.Sources) == 0 {
			return fmt.Errorf("sources cannot be empty for %s analysis", c.AnalysisType)
		}
		for i, source := range c.Sources {
			if source.URL == "" {
				return fmt.Errorf("source[%d]: URL is required", i)
			}
			if source.Content == "" {
				return fmt.Errorf("source[%d]: content is required", i)
			}
		}
	}

	return nil
}

// Helper to create execution metadata for content analysis services
func NewContentAnalysisMetadata(
	startTime time.Time,
	creditsUsed int,
	provider string,
	requestID string,
	conversationID string,
	userID string,
) ExecutionMetadata {
	endTime := time.Now()
	return ExecutionMetadata{
		StartTime:      startTime,
		EndTime:        endTime,
		DurationMs:     endTime.Sub(startTime).Milliseconds(),
		CreditsUsed:    creditsUsed,
		Provider:       provider,
		RequestID:      requestID,
		ConversationID: conversationID,
		UserID:         userID,
	}
}