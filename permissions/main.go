package permissions

type PermissionBit uint64

func (p PermissionBit) Has(bits PermissionBit, checkAdmin bool) bool {
	if checkAdmin {
		return (p&bits) == bits || (p&Administrator) == Administrator
	}

	return (p & bits) == bits
}

const (
	CreateInstantInvite PermissionBit = 1 << iota
	KickMembers
	BanMembers
	Administrator
	ManageChannels
	ManageGuild
	AddReactions
	ViewAuditLog
	PrioritySpeaker
	Stream
	ViewChannel
	SendMessages
	SendTTSMessages
	ManageMessages
	EmbedLinks
	AttachFiles
	ReadMessageHistory
	MentionEveryone
	UseExternalEmojis
	ViewGuildInsights
	Connect
	Speak
	MuteMembers
	DeafenMembers
	MoveMembers
	UseVAD
	ChangeNickname
	ManageNicknames
	ManageRoles
	ManageWebhooks
	ManageEmojisAndStickers
	UseApplicationCommands
	RequestToSpeak
	ManageEvents
	ManageThreads
	CreatePublicThreads
	CreatePrivateThreads
	UseExternalStickers
	SendMessagesInThreads
	StartEmbeddedActivities
	ModerateMembers
)
