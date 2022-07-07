package httpcord

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func Request(URI, method string, body interface{}, clientToken string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.SetRequestURI(URI)

	if body != nil {
		b, _ := json.Marshal(body)
		req.SetBody(b)
		req.Header.SetContentType("application/json")
	}

	req.Header.SetMethod(method)
	req.Header.Set("User-Agent", "HttpInteractionsBot (httpcord, 0.1.0)")

	if clientToken != "" {
		req.Header.Set("Authorization", "Bot "+clientToken)
	}

	err := fasthttp.Do(req, res)

	if err != nil {
		panic("Error in request: " + err.Error())
	}

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	return res
}