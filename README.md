[![Github tag](https://badgen.net/github/tag/JustAWaifuHunter/httpcord)](https://github.com/JustAWaifuHunter/httpcord/tags) [![CodeFactor](https://www.codefactor.io/repository/github/justawaifuhunter/httpcord/badge)](https://www.codefactor.io/repository/github/justawaifuhunter/httpcord) [![DeepSource](https://deepsource.io/gh/JustAWaifuHunter/httpcord.svg/?label=active+issues&show_trend=true&token=v43r3QF8bIhSH6ARZQg5Zsz7)](https://deepsource.io/gh/JustAWaifuHunter/httpcord/?ref=repository-badge) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) 
# HttpCord Go
Simply go project for discord http interactions

Inspired by [Discord Interactions](https://github.com/discord/discord-interactions-js), [Discord.JS](https://github.com/discordjs/discord.js) And [DiscordGo](https://github.com/bwmarrin/discordgo)

## Getting Started
### Installation
```bash
go get github.com/denkylabs/discord-api-types-go@v0.0.6
```

### Example
```go
package main

import "github.com/JustAWaifuHunter/httpcord"

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