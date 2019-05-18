package traq

import (
	"io"
	"strings"
)

type plainTextProvider struct {
	text string
}

func (p *plainTextProvider) ContentType() string {
	return "text/plain; charset=utf-8"
}

func (p *plainTextProvider) Body() (io.Reader, error) {
	return strings.NewReader(p.text), nil
}
