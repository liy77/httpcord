package httpcord

import "github.com/JustAWaifuHunter/httpcord/endpoints"

type (
	GuildVerificationLevel    int
	GuildMessageNotifications int
	GuildNSFWLevel            int
	PremiumTier               int
)

const (
	PremiumTierNone PremiumTier = 0
	PremiumTier1    PremiumTier = 1
	PremiumTier2    PremiumTier = 2
	PremiumTier3    PremiumTier = 3
)

type Guild struct {
	ID                          Snowflake                 `json:"id"`
	Name                        string                    `json:"name"`
	Icon                        string                    `json:"icon"`
	Splash                      string                    `json:"splash"`
	DiscoverySplash             string                    `json:"discovery_splash"`
	AfkTimeout                  int                       `json:"afk_timeout"`
	MemberCount                 int                       `json:"member_count"`
	VerificationLevel           GuildVerificationLevel    `json:"verification_level"`
	Large                       bool                      `json:"large"`
	DefaultMessageNotifications GuildMessageNotifications `json:"default_message_notifications"`
	Roles                       []*Role                   `json:"roles"`
	Emojis                      []*Emoji                  `json:"emojis"`
	Stickers                    []*Sticker                `json:"stickers"`
	Members                     []*Member                 `json:"members"`
	NSFWLevel                   GuildNSFWLevel            `json:"nsfw_level"`
	Features                    []string                  `json:"features"`
	MfaLevel                    MfaLevel                  `json:"mfa_level"`
	ApplicationID               string                    `json:"application_id"`
	WidgetEnabled               bool                      `json:"widget_enabled"`
	WidgetChannelID             string                    `json:"widget_channel_id"`
	SystemChannelID             string                    `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlag         `json:"system_channel_flags"`
	RulesChannelID              string                    `json:"rules_channel_id"`
	VanityURLCode               string                    `json:"vanity_url_code"`
	Description                 string                    `json:"description"`
	Banner                      string                    `json:"banner"`
	PremiumTier                 PremiumTier               `json:"premium_tier"`
	PreferredLocale             string                    `json:"preferred_locale"`
}

func (g *Guild) IconURL(size string) string {
	if g.Icon == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.GuildIconURL(g.ID.String(), g.Icon), "", size)
}

func (g *Guild) StaticIconURL(size string) string {
	if g.Icon == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.GuildIconURL(g.ID.String(), g.Icon), JpegImageFormat.String(), size)
}

func (g *Guild) DynamicIconURL(format ImageFormat, size string) string {
	if g.Icon == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.GuildIconURL(g.ID.String(), g.Icon), format.String(), size)
}

func (g *Guild) BannerURL(size string) string {
	if g.Banner == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.Banner(g.ID.String(), g.Banner), "", size)
}

func (g *Guild) StaticBannerURL(size string) string {
	if g.Banner == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.Banner(g.ID.String(), g.Banner), JpegImageFormat.String(), size)
}

func (g *Guild) DynamicBannerURL(format ImageFormat, size string) string {
	if g.Banner == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.Banner(g.ID.String(), g.Banner), format.String(), size)
}
