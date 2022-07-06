package httpcord

type WebhookEdit struct {
	Content         string          `json:"content,omitempty"`
	Components      *[]AnyComponent     `json:"components,omitempty"`
	Embeds          *[]*Embed        `json:"embeds,omitempty"`
	Files           []*DiscordFile   `json:"-"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
}
