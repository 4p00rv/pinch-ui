package head

import (
	"github.com/4p00rv/pinch-ui/pkg/base"
)

type Head struct {
	base.Element
	content         string
	metaCharset     string
	metaDescription string
	metaKeywords    string // Comma separated keywords for seo
	canonicalLink   string // Canonical link to a page with near identical content
	title           string
}

func NewHead() Head {
	h := Head{}
	return h
}

func (h Head) GetContent() string {
	return h.content
}

func (h *Head) SetMetaCharset(a string) {
	h.metaCharset = a
}

func (h *Head) SetMetaDescription(a string) {
	h.metaDescription = a
}

func (h *Head) SetMetaKeywords(a string) {
	h.metaKeywords = a
}

func (h *Head) SetCanonicalLink(a string) {
	h.canonicalLink = a
}

func (h *Head) SetTitle(a string) {
	h.title = a
}
