package extractor

import "html/template"

type Out struct {
	H1      string
	H2      []string
	H3      []string
	H4      []string
	H5      []string
	H6      []string
	Content template.HTML
	P       []template.HTML
	Ul      []string
	Ol      []string
	Link    []string
	Image   []string
}
