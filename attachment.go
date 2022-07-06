package httpcord

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

var QuoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

type DiscordFile struct {
	*bytes.Buffer
	Filename    string
	Description string
	ContentType string
	Spoiler     bool
}

func (f *DiscordFile) MakeAttach(ID Snowflake, m *multipart.Writer) (*Attachment, error) {
	if f.Spoiler && !strings.HasPrefix(f.Filename, "SPOILER_") {
		f.Filename = "SPOILER_" + f.Filename
	}

	attach := &Attachment{
		ID:          ID,
		Filename:    f.Filename,
		Description: f.Description,
	}

	contentType := "application/octet-stream"

	if f.ContentType != "" {
		contentType = f.ContentType
	}

	headers := make(textproto.MIMEHeader)
	headers.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"%s\"; filename=\"%s\"",
		QuoteEscaper.Replace(fmt.Sprintf("files[%d]", ID.Uint64())),
		QuoteEscaper.Replace(f.Filename),
	))
	headers.Set("Content-Type", contentType)

	w, err := m.CreatePart(headers)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(w, f)

	if err != nil {
		return nil, err
	}

	return attach, nil
}
