package utils

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func GetContent(text string, useMd bool, md goldmark.Markdown) *Content {

	content := Content{Text: text}

	if useMd {
		var buf bytes.Buffer
		if err := md.Convert([]byte(text), &buf); err != nil {
			return nil
		}
		html := buf.String()
		content.HTML = &html
	}

	return &content
}
