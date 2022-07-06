# HttpCord Go
Simply go project for discord http interactions

Inspired by [Discord Interactions](https://github.com/discord/discord-interactions-js), [Discord.JS](https://github.com/discordjs/discord.js) And [DiscordGo](https://github.com/bwmarrin/discordgo)

## Example
```go
package main

import "httpcord"

func main() {
	connection := httpcord.NewConnection(httpcord.ConnectionOptions{
		HttpConnection: httpcord.FastHttpConnection,
		PublicKey: "Your Discord Application Public Key Here",
	})

	connection.AddInteractionHandler(func(ctx httpcord.ConnectionContext) {
		ctx.ReplyInteraction(&httpcord.InteractionCallbackData{
			Content: "Hello World",
		})
	})

	connection.Connect(":8080")
}
```