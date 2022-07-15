package httpcord

type APIInteraction struct {
	ID             string          `json:"id"`
	ApplicationID  string          `json:"application_id"`
	Type           InteractionType `json:"type"`
	Data           interface{}     `json:"data"`
	GuildID        string          `json:"guild_id,omitempty"`
	ChannelID      string          `json:"channel_id,omitempty"`
	Member         *APIMember      `json:"member,omitempty"`
	User           *APIUser        `json:"user,omitempty"`
	Token          string          `json:"token"`
	Version        int             `json:"version"`
	Message        *APIMessage     `json:"message,omitempty"`
	AppPermissions string          `json:"app_permissions,omitempty"`
	Locale         string          `json:"locale,omitempty"`
}

type APIMember struct {
	User                       *APIUser `json:"user,omitempty"`
	Nick                       string   `json:"nick,omitempty"`
	Avatar                     string   `json:"avatar,omitempty"`
	Roles                      []string `json:"roles"`
	JoinedAt                   Time     `json:"joined_at"`
	PremiumSince               Time     `json:"premium_since"`
	Deaf                       bool     `json:"deaf"`
	Mute                       bool     `json:"mute"`
	Pending                    bool     `json:"pending,omitempty"`
	Permissions                string   `json:"permissions,omitempty"`
	CommunicationDisabledUntil Time     `json:"communication_disabled_until"`
}

type APIUser struct {
	ID            string      `json:"id"`
	Username      string      `json:"username"`
	Discriminator string      `json:"discriminator"`
	Avatar        string      `json:"avatar,omitempty"`
	Bot           bool        `json:"bot,omitempty"`
	System        bool        `json:"system"`
	MfaEnabled    bool        `json:"mfa_enabled"`
	Banner        string      `json:"banner,omitempty"`
	AccentColor   string      `json:"accent_color"`
	Locale        Locale      `json:"locale,omitempty"`
	Verified      bool        `json:"verified,omitempty"`
	Email         string      `json:"email,omitempty"`
	Flags         UserFlags   `json:"flags,omitempty"`
	PremiumType   PremiumType `json:"premium_type"`
	PublicFlags   UserFlags   `json:"public_flags,omitempty"`
}

type APIMessage struct {
}
