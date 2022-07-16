package httpcord

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"

	"github.com/valyala/fasthttp"
	"golang.org/x/exp/maps"
	"httpcord/endpoints"
)

const (
	ReasonHeaderKey        = "X-Audit-Log-Reason"
	AuthorizationHeaderKey = "Authorization"
	UserAgentHeaderKey     = "User-Agent"
)

var DefaultUserAgent = fmt.Sprintf("HttpInteractionsBot (http-cord, %s)", VERSION)

// Request Create a request
func Request(URI, method string, body interface{}, clientToken string, headers map[string]string) []byte {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(URI)

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		if !slices.Contains(maps.Keys(headers), UserAgentHeaderKey) {
			req.Header.Set(UserAgentHeaderKey, DefaultUserAgent)
		}
	} else {
		req.Header.Set(UserAgentHeaderKey, DefaultUserAgent)
	}

	if body != nil {
		b, _ := json.Marshal(body)
		req.SetBody(b)
		req.Header.SetContentType("application/json")
	}

	req.Header.SetMethod(method)

	if clientToken != "" {
		req.Header.Set(AuthorizationHeaderKey, "Bot "+clientToken)
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	err := fasthttp.Do(req, res)

	if err != nil {
		panic("Error in request: " + err.Error())
	}

	return res.Body()
}

// ApplicationCommandsBulkOverwrite Overwrite all application commands
// (In case guild id is not provided to edit commands globally)
func ApplicationCommandsBulkOverwrite(applicationID string, commands []*ApplicationCommand, guildID, clientToken string) (createdCommands []*ApplicationCommand, err error) {
	uri := endpoints.FormatAPIURI(endpoints.ApplicationCommandsGlobal(applicationID))

	if guildID != "" {
		uri = endpoints.FormatAPIURI(endpoints.ApplicationCommandsGuild(applicationID, guildID))
	}

	res := Request(uri, fasthttp.MethodPut, commands, clientToken, nil)
	err = json.Unmarshal(res, &createdCommands)
	return
}

// CreateMessage Send a message to a channel
func CreateMessage(clientToken string, channelID string, messageData *Message) {
	Request(endpoints.Messages(channelID), fasthttp.MethodPost, messageData, clientToken, nil)
}

// FetchMessage Fetch a message in channel
func FetchMessage(clientToken string, channelID, messageID string) {
	Request(endpoints.Message(channelID, messageID), fasthttp.MethodGet, nil, clientToken, nil)
}

// DeleteMessage Delete a message in channel
func DeleteMessage(clientToken, channelID, messageID string) {
	Request(endpoints.Message(channelID, messageID), fasthttp.MethodDelete, nil, clientToken, nil)
}

// EditMessage Edit a message in channel
func EditMessage(clientToken, channelID, messageID string, messageData *Message) {
	Request(endpoints.Message(channelID, messageID), fasthttp.MethodPatch, messageData, clientToken, nil)
}

// CreateReaction Create a reaction in the message
func CreateReaction(clientToken, channelID, messageID, reaction string) {
	Request(
		endpoints.UserReaction(channelID, messageID, reaction, "@me"),
		fasthttp.MethodPut,
		nil, clientToken,
		nil,
	)
}

// RemoveReaction Remove a reaction from message
func RemoveReaction(clientToken, channelID, messageID, reaction, userID string) {
	Request(
		endpoints.UserReaction(channelID, messageID, reaction, userID),
		fasthttp.MethodDelete,
		nil, clientToken,
		nil,
	)
}

// RemoveAllReactions Remove all reactions from message
func RemoveAllReactions(clientToken, channelID, messageID string) {
	Request(
		endpoints.Reactions(channelID, messageID),
		fasthttp.MethodDelete,
		nil, clientToken,
		nil,
	)
}

// MessageReactions Get all reactions from message
func MessageReactions(clientToken, channelID, messageID string, options SearchQueryParams) {
	var query string

	if options.Limit != 0 {
		query = fmt.Sprintf("?limit=%d", options.Limit)
	}

	if options.After.String() != "" {
		if query != "" {
			query += fmt.Sprintf("&after=%s", options.After.String())
		} else {
			query = fmt.Sprintf("?after=%s", options.After.String())
		}
	}

	uri := endpoints.Reactions(channelID, messageID)

	if query != "" {
		uri += query
	}

	Request(
		uri,
		fasthttp.MethodGet,
		nil, clientToken,
		nil,
	)
}

func EditOriginalInteractionResponse(applicationID, interactionToken string, data *WebhookEdit) {
	Request(
		endpoints.FormatAPIURI(endpoints.WebhookMessage(applicationID, interactionToken, "@original")),
		fasthttp.MethodPatch,
		data, "", nil,
	)
}

func DeleteOriginalInteractionResponse(applicationID, interactionToken string) {
	Request(
		endpoints.FormatAPIURI(endpoints.WebhookMessage(applicationID, interactionToken, "@original")),
		fasthttp.MethodDelete,
		nil, "", nil,
	)
}

func FollowUpInteractionResponse(applicationID, interactionToken string, data *WebhookEdit) {
	Request(
		endpoints.FormatAPIURI(endpoints.WebhookExecute(applicationID, interactionToken)),
		fasthttp.MethodDelete,
		data, "", nil,
	)
}

type SearchQueryParams struct {
	Limit int
	After Snowflake
}
