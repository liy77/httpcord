package httpcord

type Embed struct {
	Title       string          `json:"title,omitempty"`
	// Always "rich" for webhook embeds
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   Time            `json:"timetimestamp"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURl string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURl      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

func NewEmbedBuilder() *Embed {
	return &Embed{}
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

func (e *Embed) SetTimestamp(timestamp Time) *Embed {
	e.Timestamp = timestamp
	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) SetFooter(footer *EmbedFooter) *Embed {
	e.Footer = footer
	return e
}

func (e *Embed) SetImage(image *EmbedImage) *Embed {
	e.Image = image
	return e
}

func (e *Embed) SetThumbnail(thumbnail *EmbedThumbnail) *Embed {
	e.Thumbnail = thumbnail
	return e
}

func (e *Embed) SetAuthor(author *EmbedAuthor) *Embed {
	e.Author = author
	return e
}

func (e *Embed) AddField(field *EmbedField) *Embed {
	e.Fields = append(e.Fields, field)
	return e
}