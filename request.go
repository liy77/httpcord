package httpcord

import (
	"encoding/json"

	"github.com/JustAWaifuHunter/httpcord/endpoints"
	"github.com/valyala/fasthttp"
)

func Request(URI, method string, body interface{}, clientToken string) []byte {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(URI)

	if body != nil {
		b, _ := json.Marshal(body)
		req.SetBody(b)
		req.Header.SetContentType("application/json")
	}

	req.Header.SetMethod(method)
	req.Header.Set("User-Agent", "HttpInteractionsBot (http-cord, 0.1.0)")

	if clientToken != "" {
		req.Header.Set("Authorization", "Bot "+clientToken)
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	err := fasthttp.Do(req, res)

	if err != nil {
		panic("Error in request: " + err.Error())
	}

	return res.Body()
}

func ApplicationCommandsBulkOverwrite(applicationID string, commands []*ApplicationCommand, guildID, clientToken string) (createdCommands []*ApplicationCommand, err error) {
	uri := endpoints.DiscordURL + endpoints.DiscordAPI + endpoints.ApplicationCommandsGlobal(applicationID)

	if guildID != "" {
		uri = endpoints.DiscordURL + endpoints.DiscordAPI + endpoints.ApplicationCommandsGuild(applicationID, guildID)
	}

	res := Request(uri, fasthttp.MethodPut, commands, clientToken)
	err = json.Unmarshal(res, &createdCommands)
	return
}
