package resolution

import "encoding/json"

type TokenMetadataAttribute struct {
	DisplayType string     `json:"display_type"`
	TraitType   string     `json:"trait_type"`
	Value       json.Token `json:"value"`
}

type TokenMetadata struct {
	Name            string                   `json:"name"`
	Description     string                   `json:"description"`
	Image           string                   `json:"image"`
	ExternalUrl     string                   `json:"external_url"`
	ExternalLink    string                   `json:"external_link"`
	ImageData       string                   `json:"image_data"`
	BackgroundColor string                   `json:"background_color"`
	AnimationUrl    string                   `json:"animation_url"`
	YoutubeUrl      string                   `json:"youtube_url"`
	Attributes      []TokenMetadataAttribute `json:"attributes"`
}
