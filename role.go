package httpcord

import "httpcord/permissions"

type Role struct {
	ID           Snowflake                  `json:"id"`
	Name         string                     `json:"name"`
	Color        int                        `json:"color"`
	Hoist        bool                       `json:"hoist"`
	Icon         string                     `json:"icon,omitempty"`
	UnicodeEmoji string                     `json:"unicode_emoji,omitempty"`
	Position     int                        `json:"position"`
	Permissions  *permissions.PermissionBit `json:"permissions"`
	Managed      bool                       `json:"managed"`
	Mentionable  bool                       `json:"mentionable"`
	Tags         []*RoleTag                 `json:"tags,omitempty"`
}

type RoleTag struct {
	BotID             *Snowflake  `json:"bot_id,omitempty"`
	IntegrationID     *Snowflake  `json:"integration_id,omitempty"`
	PremiumSubscriber interface{} `json:"premium_subscriber,omitempty"`
}
