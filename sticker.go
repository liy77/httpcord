package httpcord

type (
	StickerType       int
	StickerFormatType int
	StickerTags       []string
)

const (
	StickerFormatTypePNG StickerFormatType = iota + 1
	StickerFormatTypeAPNG
	StickerFormatTypeLOTTIE
)

type StickerItem struct {
	ID         Snowflake         `json:"id"`
	Name       string            `json:"name"`
	FormatType StickerFormatType `json:"format_type"`
}

type Sticker struct {
	StickerItem
	PackID      Snowflake   `json:"pack_id"`
	Description string      `json:"description"`
	Tags        string      `json:"tags"`
	Type        StickerType `json:"type"`
	Available   bool        `json:"available,omitempty"`
	GuildID     Snowflake   `json:"guild_id,omitempty"`
	User        *User       `json:"user,omitempty"`
	SortValue   *int        `json:"sort_value,omitempty"`
}
