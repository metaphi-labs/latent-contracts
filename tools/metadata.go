package tools

// ToolType represents the category of tool
type ToolType string

const (
	ToolTypeMedia           ToolType = "media"
	ToolTypeNative          ToolType = "native"
	ToolTypeImageProcessing ToolType = "image_processing"
)

// ServiceType represents which service handles the tool
type ServiceType string

const (
	ServiceTypeMediaAI        ServiceType = "media-ai"
	ServiceTypeVideoProcessor ServiceType = "video-processor"
	ServiceTypeNative         ServiceType = "native"
)

// OutputType represents what the tool produces
type OutputType string

const (
	OutputTypeImage OutputType = "image"
	OutputTypeVideo OutputType = "video"
	OutputTypeAudio OutputType = "audio"
	OutputTypeJSON  OutputType = "json"
	OutputTypeText  OutputType = "text"
)

// ToolMeta contains all metadata for a tool
type ToolMeta struct {
	Name         ToolName
	Credits      int
	Type         ToolType
	ServiceType  ServiceType
	Description  string
	Examples     []string
	OutputType   OutputType
	EndpointPath string
}

// Metadata contains all tool metadata definitions
var Metadata = map[ToolName]ToolMeta{
	// === Media Generation Tools ===

	GenerateImageImagen: {
		Name:         GenerateImageImagen,
		Credits:      2,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Standard quality image generation using Google Imagen 4.0. Balanced quality and speed. Supports creative controls, multiple images, and various aspect ratios. Best for general use cases.",
		Examples: []string{
			"Generate a professional product photo",
			"Create a marketing banner with text",
			"Design a logo concept",
			"Illustrate a scene from a story",
			"Generate multiple variations of an icon",
		},
		OutputType:   OutputTypeImage,
		EndpointPath: "/api/generate-image-imagen/generate/async",
	},

	GenerateImageImagenFast: {
		Name:         GenerateImageImagenFast,
		Credits:      1,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Fast, cost-effective image generation using Google Imagen 4.0 Fast. Lower quality but quicker generation and reduced cost. Best for drafts, iterations, and testing.",
		Examples: []string{
			"Quick draft of a product mockup",
			"Generate test images for layout",
			"Fast iteration on concept art",
			"Budget-friendly batch image generation",
		},
		OutputType:   OutputTypeImage,
		EndpointPath: "/api/generate-image-imagen-fast/generate/async",
	},

	GenerateImageImagenUltra: {
		Name:         GenerateImageImagenUltra,
		Credits:      3,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Premium ultra-high quality image generation using Google Imagen 4.0 Ultra. Highest quality, photorealistic results. Limited to 1 image per request. Best for hero images, professional photography, and final production.",
		Examples: []string{
			"Ultra-realistic product photography for marketing",
			"Photorealistic portrait for professional use",
			"High-end architectural visualization",
			"Premium quality hero image for website",
			"Museum-quality artistic rendering",
		},
		OutputType:   OutputTypeImage,
		EndpointPath: "/api/generate-image-imagen-ultra/generate/async",
	},

	GenerateImageFlash: {
		Name:         GenerateImageFlash,
		Credits:      10,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Generate images using Flash technology. Fast and efficient image generation.",
		Examples: []string{
			"Generate a quick concept image",
			"Create a draft illustration",
			"Make a simple graphic",
		},
		OutputType:   OutputTypeImage,
		EndpointPath: "/api/generate-image-flash/generate/async",
	},

	NanoBanana: {
		Name:         NanoBanana,
		Credits:      3,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Versatile image generation with Gemini 2.5 Flash Image Preview. Supports text-to-image, image editing, style transfer, and multi-image composition. Best for creative editing and conversational image generation.",
		Examples: []string{
			"Create a photorealistic portrait of an elderly Japanese ceramicist",
			"Edit this image to add a sunset background",
			"Combine these three images into a creative composition",
			"Apply the style of this painting to my photo",
			"Transform this sketch into a detailed illustration",
		},
		OutputType:   OutputTypeImage,
		EndpointPath: "/api/nano-banana/generate/async",
	},

	GenerateVideoVeo3: {
		Name:         GenerateVideoVeo3,
		Credits:      300,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Premium video generation using Google Veo3. Creates high-quality 8-second videos at up to 1080p resolution with optional audio generation. Supports both text-to-video and image-to-video with optional ending frame control. Best for final production quality.",
		Examples: []string{
			"Create a cinematic 8-second video in 1080p",
			"Generate high-quality product showcase video",
			"Transform this image into a dynamic 8-second video",
			"Make a professional marketing video with specific camera movements",
			"Create artistic video from image with zoom effect",
		},
		OutputType:   OutputTypeVideo,
		EndpointPath: "/api/generate-video-veo3/generate/async",
	},

	GenerateVideoVeo3Fast: {
		Name:         GenerateVideoVeo3Fast,
		Credits:      160,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Fast, cost-effective video generation using Google Veo3 Fast. Generates 8-second videos at 720p resolution with optional audio generation. Best for quick iterations and testing.",
		Examples: []string{
			"Create an 8-second video of a sunset",
			"Generate a preview of product demo",
			"Make a social media clip of nature scene",
			"Quick test video of abstract patterns",
		},
		OutputType:   OutputTypeVideo,
		EndpointPath: "/api/generate-video-veo3-fast/generate/async",
	},

	GenerateVideoVeo3FastNoAudio: {
		Name:         GenerateVideoVeo3FastNoAudio,
		Credits:      100,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Fast, cost-effective video generation using Google Veo3 Fast. Generates 8-second videos at 720p resolution without audio generation.",
		Examples: []string{
			"Create an 8-second video of a sunset",
			"Generate a preview of product demo",
			"Make a social media clip of nature scene",
			"Quick test video of abstract patterns",
		},
		OutputType:   OutputTypeVideo,
		EndpointPath: "/api/generate-video-veo3-fast-no-audio/generate/async",
	},

	GenerateVideoVeo3NoAudio: {
		Name:         GenerateVideoVeo3NoAudio,
		Credits:      200,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Premium video generation using Google Veo3. Creates high-quality 8-second videos at up to 1080p resolution without audio generation. Supports both text-to-video and image-to-video with optional ending frame control. Best for final production quality.",
		Examples: []string{
			"Create a cinematic 8-second video in 1080p",
			"Generate high-quality product showcase video",
			"Transform this image into a dynamic 8-second video",
			"Make a professional marketing video with specific camera movements",
			"Create artistic video from image with zoom effect",
		},
		OutputType:   OutputTypeVideo,
		EndpointPath: "/api/generate-video-veo3-no-audio/generate/async",
	},

	GenerateMusicLyria: {
		Name:         GenerateMusicLyria,
		Credits:      3,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Generate high-quality music using Google Lyria. Creates 30-second instrumental tracks in various genres and styles. Supports custom seeds for reproducible generation.",
		Examples: []string{
			"Smooth jazz with mellow brass and piano",
			"Epic orchestral battle music with heavy drums",
			"Lo-fi hip hop beat for studying",
		},
		OutputType:   OutputTypeAudio,
		EndpointPath: "/api/generate-music-lyria/generate/async",
	},

	// === Video Processing Tools ===

	CombineVideos: {
		Name:         CombineVideos,
		Credits:      20,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeVideoProcessor,
		Description:  "Combine multiple videos into a single video with optional transitions. Supports fade and dissolve transitions between clips.",
		Examples: []string{
			"Combine two videos with fade transition",
			"Merge multiple clips into one video",
			"Create a compilation video with transitions",
		},
		OutputType:   OutputTypeVideo,
		EndpointPath: "/api/video/combine/async",
	},

	TrimVideo: {
		Name:         TrimVideo,
		Credits:      10,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeVideoProcessor,
		Description:  "Trim a video by specifying start and end times or duration. Supports fast mode for quick processing.",
		Examples: []string{
			"Extract 30-second clip starting at 10 seconds",
			"Trim video from 0:10 to 0:40",
			"Cut the first 5 seconds from a video",
		},
		OutputType:   OutputTypeVideo,
		EndpointPath: "/api/video/trim/async",
	},

	ImageAudioMerge: {
		Name:         ImageAudioMerge,
		Credits:      15,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeVideoProcessor,
		Description:  "Create a video by combining a static image with an audio track. Perfect for music visualizations or podcast videos.",
		Examples: []string{
			"Create music video with album cover",
			"Make a podcast video with logo",
			"Generate audio visualization with background image",
		},
		OutputType:   OutputTypeVideo,
		EndpointPath: "/api/video/image-audio-merge/async",
	},

	ExtractFrame: {
		Name:         ExtractFrame,
		Credits:      5,
		Type:         ToolTypeMedia,
		ServiceType:  ServiceTypeVideoProcessor,
		Description:  "Extract frames from videos at specific positions or timestamps. Supports first, last, middle positions or exact timestamps. Works with generated videos, uploads, or gallery assets.",
		Examples: []string{
			"Extract the last frame from this video",
			"Get frame at 5 seconds from the video",
			"Extract the middle frame from the generated video",
			"Get the first frame as a thumbnail",
			"Extract frame at 00:00:03.5 from my upload",
		},
		OutputType:   OutputTypeImage,
		EndpointPath: "/api/video/extract-frame/async",
	},

	MergeImages: {
		Name:         MergeImages,
		Credits:      0,
		Type:         ToolTypeImageProcessing,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Merge multiple images into a single composite",
		Examples:     []string{},
		OutputType:   OutputTypeImage,
		EndpointPath: "/api/merge-images/generate/async",
	},

	// === Content Analysis Tools ===

	ContentAnalyzer: {
		Name:         ContentAnalyzer,
		Credits:      1,
		Type:         ToolTypeNative,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Analyzes a given web page(s) for information as needed for the core user task or request",
		Examples: []string{
			"analyze this YouTube video",
			"summarize this article",
			"extract key points from this webpage",
		},
		OutputType:   OutputTypeText,
		EndpointPath: "/api/content-analyzer/generate/async",
	},

	GoogleSearch: {
		Name:         GoogleSearch,
		Credits:      0,
		Type:         ToolTypeNative,
		ServiceType:  ServiceTypeMediaAI,
		Description:  "Proactively or upon instruction search the internet to build context on the user's request or domain",
		Examples: []string{
			"search for AI developments",
			"find information about renewable energy",
			"research competitors",
		},
		OutputType:   OutputTypeJSON,
		EndpointPath: "/api/google-search/generate/async",
	},
}

// GetToolMetadata returns metadata for a specific tool
func GetToolMetadata(toolName ToolName) (ToolMeta, bool) {
	meta, exists := Metadata[toolName]
	return meta, exists
}

// GetToolCredits returns the credit cost for a tool
func GetToolCredits(toolName ToolName) int {
	if meta, exists := Metadata[toolName]; exists {
		return meta.Credits
	}
	return 0
}

// GetToolsByType returns all tools of a specific type
func GetToolsByType(toolType ToolType) []ToolMeta {
	var tools []ToolMeta
	for _, meta := range Metadata {
		if meta.Type == toolType {
			tools = append(tools, meta)
		}
	}
	return tools
}

// GetToolsByService returns all tools handled by a specific service
func GetToolsByService(serviceType ServiceType) []ToolMeta {
	var tools []ToolMeta
	for _, meta := range Metadata {
		if meta.ServiceType == serviceType {
			tools = append(tools, meta)
		}
	}
	return tools
}

// GetAllTools returns metadata for all tools
func GetAllTools() map[ToolName]ToolMeta {
	return Metadata
}