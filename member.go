package httpcord

import (
	"strconv"

	"httpcord/permissions"
)

type Member struct {
	User                       *User                     `json:"user"`
	Nick                       string                    `json:"nick,omitempty"`
	Avatar                     string                    `json:"avatar,omitempty"`
	Roles                      []*Snowflake              `json:"roles"`
	JoinedAt                   Time                      `json:"joined_at"`
	PremiumSince               Time                      `json:"premium_since"`
	Deaf                       bool                      `json:"deaf"`
	Mute                       bool                      `json:"mute"`
	Pending                    bool                      `json:"pending,omitempty"`
	Permissions                permissions.PermissionBit `json:"permissions,omitempty"`
	CommunicationDisabledUntil Time                      `json:"communication_disabled_until,omitempty"`
}

func ResolveMember(member *APIMember) *Member {
	resolved := &Member{
		Nick:                       member.Nick,
		Avatar:                     member.Avatar,
		JoinedAt:                   member.JoinedAt,
		PremiumSince:               member.PremiumSince,
		Deaf:                       member.Deaf,
		Mute:                       member.Mute,
		Pending:                    member.Pending,
		CommunicationDisabledUntil: member.CommunicationDisabledUntil,
		Roles:                      StringArrayToSnowflakeArray(member.Roles),
	}

	perms, err := strconv.ParseUint(member.Permissions, 10, 64)

	if err != nil {
		panic("Invalid permissions bits")
	}

	resolved.Permissions = permissions.PermissionBit(perms)
	resolved.User = ResolveUser(member.User)

	return resolved
}
