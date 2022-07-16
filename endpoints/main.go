package endpoints

import (
	"fmt"
	"strings"
)

const (
	DiscordCDN = "https://cdn.discordapp.com"
	DiscordAPI = "/api/v10"
	DiscordURL = "https://discord.com"
)

func Messages(channelID string) string {
	return fmt.Sprintf("/channels/%s/messages", channelID)
}

func Message(channelID, messageID string) string {
	return fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID)
}

func Reactions(channelID, messageID string) string {
	return fmt.Sprintf("/channels/%s/messages/%s/reactions", channelID, messageID)
}

func UserReaction(channelID, messageID, reaction, userID string) string {
	return fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/%s", channelID, messageID, reaction, userID)
}

func DefaultUserAvatar(DefaultAvatar int) string {
	return fmt.Sprintf("/embed/avatars/%d", DefaultAvatar)
}

func Banner(ID, hash string) string {
	return fmt.Sprintf("/banners/%s/%s", ID, hash)
}

func AvatarURL(ID, hash string) string {
	return fmt.Sprintf("/avatars/%s/%s", ID, hash)
}

func GuildIconURL(ID, GuildIcon string) string {
	return fmt.Sprintf("/icons/%s/%s", ID, GuildIcon)
}

func WebhookMessage(ID, token, messageID string) string {
	return fmt.Sprintf("/webhooks/%s/%s/messages/%s", ID, token, messageID)
}

func WebhookExecute(ID, token string) string {
	return fmt.Sprintf("/webhooks/%s/%s", ID, token)
}

func ApplicationCommandsGlobal(applicationID string) string {
	return fmt.Sprintf("/applications/%s/commands", applicationID)
}

func ApplicationCommandsGuild(applicationID, GuildID string) string {
	return fmt.Sprintf("/applications/%s/guilds/%s/commands", applicationID, GuildID)
}

func FormatImage(URL, format, size string) string {
	if format == "" {
		if strings.Contains(URL, "/a_") {
			format = "gif"
		} else {
			format = "jpg"
		}
	}

	return DiscordCDN + URL + "." + format + "?size=" + size
}

func FormatAPIURI(URI string) string {
	return fmt.Sprintf("%s%s%s", DiscordURL, DiscordAPI, URI)
}
