package httpcord

import (
	"github.com/JustAWaifuHunter/httpcord/endpoints"
	"strconv"
)

type (
	UserFlags   uint
	PremiumType uint
)

const (
	DiscordEmployee UserFlags = 1 << iota
	PartneredServerOwner
	HypesquadEvents
	BugHunterLevel1
	HouseBravery
	HouseBrilliance
	HouseBalance
	EarlySupporter
	TeamUser
	_
	_
	_
	System
	_
	BugHunterLevel2
	_
	VerifiedBot
	EarlyVerifiedBotDeveloper
)

const (
	PremiumTypeNone PremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
)

type User struct {
	ID            Snowflake   `json:"id"`
	Username      string      `json:"username"`
	Discriminator string      `json:"discriminator"`
	Avatar        string      `json:"avatar,omitempty"`
	Bot           bool        `json:"bot,omitempty"`
	System        bool        `json:"system,omitempty"`
	MfaEnabled    bool        `json:"mfa_enabled,omitempty"`
	Banner        string      `json:"banner,omitempty"`
	AccentColor   int         `json:"accent_color"`
	Locale        Locale      `json:"locale,omitempty"`
	Verified      bool        `json:"verified,omitempty"`
	Email         string      `json:"email,omitempty"`
	Flags         UserFlags   `json:"flags"`
	PremiumType   PremiumType `json:"premium_type,omitempty"`
	PublicFlags   UserFlags   `json:"public_flags,omitempty"`
}

func (u *User) DefaultAvatar() int {
	i, _ := strconv.Atoi(u.Discriminator)
	return i % 5
}

func (u *User) DefaultAvatarURL() string {
	return endpoints.DiscordCDN + endpoints.DefaultUserAvatar(u.DefaultAvatar())
}

func (u *User) AvatarURL(size string) string {
	if u.Avatar == "" {
		return u.DefaultAvatarURL()
	}

	return endpoints.FormatImage(endpoints.AvatarURL(u.ID.String(), u.Avatar), "", size)
}

func (u *User) StaticAvatarURL(size string) string {
	if u.Avatar == "" {
		return u.DefaultAvatarURL()
	}

	return endpoints.FormatImage(endpoints.AvatarURL(u.ID.String(), u.Avatar), JpegImageFormat.String(), size)
}

func (u *User) DynamicAvatarURL(format ImageFormat, size string) string {
	if u.Avatar == "" {
		return u.DefaultAvatarURL()
	}

	return endpoints.FormatImage(endpoints.AvatarURL(u.ID.String(), u.Avatar), format.String(), size)
}

func (u *User) BannerURL(size string) string {
	if u.Banner == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.Banner(u.ID.String(), u.Banner), "", size)
}

func (u *User) StaticBannerURL(size string) string {
	if u.Banner == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.Banner(u.ID.String(), u.Banner), JpegImageFormat.String(), size)
}

func (u *User) DynamicBannerURL(format ImageFormat, size string) string {
	if u.Banner == "" {
		return ""
	}

	return endpoints.FormatImage(endpoints.Banner(u.ID.String(), u.Banner), format.String(), size)
}

func (u *User) Mention() string {
	return "<@" + u.ID.String() + ">"
}

func (u *User) String() string {
	return u.Mention()
}

func ResolveUser(rawUser *APIUser) *User {
	UserAccentColor, _ := strconv.Atoi(rawUser.AccentColor)

	return &User{
		ID:            Snowflake(rawUser.ID),
		Username:      rawUser.Username,
		Discriminator: rawUser.Discriminator,
		Bot:           rawUser.Bot,
		System:        rawUser.System,
		MfaEnabled:    rawUser.MfaEnabled,
		Banner:        rawUser.Banner,
		AccentColor:   UserAccentColor,
		Locale:        rawUser.Locale,
		Verified:      rawUser.Verified,
		Email:         rawUser.Email,
		Flags:         rawUser.Flags,
		PremiumType:   rawUser.PremiumType,
		PublicFlags:   rawUser.PublicFlags,
	}
}
