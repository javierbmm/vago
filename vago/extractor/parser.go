package extractor

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
)

func ParseMarkdown(md []byte) []byte {

	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	doc = extractAst(doc)

	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	result := markdown.Render(doc, renderer)

	return result
}

func extractAst(doc ast.Node) ast.Node {
	var out Out
	out.P = append(out.P, "")
	var tracker ContentTracker

	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if entering {
			extractTitles(node, &out)
			if extractContent(node, &out, &tracker) {
				return ast.SkipChildren
			}
		}

		return ast.GoToNext
	})

	return doc
}

func extractContent(node ast.Node, out *Out, tracker *ContentTracker) bool {
	if shouldSkip(node) {
		return false
	}

	if _, ok := node.(*ast.Heading); ok {
		if tracker.IsInside {
			tracker.IsInside = false
			// If the content of P is empty, means there was a title previously, hence no need to increment the index
			// for subsequent titles/headings.
			if out.P[tracker.Index] != "" {
				tracker.Index++
				out.P = append(out.P, "")
			}
		} else {
			tracker.IsInside = true
		}

		return false
	} else {
		tracker.IsInside = true
		// create HTML renderer with extensions
		htmlFlags := html.CommonFlags | html.HrefTargetBlank
		opts := html.RendererOptions{Flags: htmlFlags}
		renderer := html.NewRenderer(opts)
		content := markdown.Render(node, renderer)

		// Save rendered content. This is to also include any container child (URL, bold, italics, etc) already parsed.
		out.P[tracker.Index] += template.HTML(content)

		return true
	}
}

func extractTitles(node ast.Node, out *Out) {
	// Note: There is the limitation of not parsing links within a title. These will be parsed separately.
	if h, ok := node.(*ast.Heading); ok {
		// Only accept one Heading 1
		if h.Level == 1 && out.H1 == "" {
			out.H1 = string(h.Children[0].AsLeaf().Literal)
		}

		if h.Level == 2 {
			h2 := string(h.Children[0].AsLeaf().Literal)
			out.H2 = append(out.H2, h2)
		}

		if h.Level == 3 {
			h3 := string(h.Children[0].AsLeaf().Literal)
			out.H3 = append(out.H3, h3)
		}

		if h.Level == 4 {
			h4 := string(h.Children[0].AsLeaf().Literal)
			out.H4 = append(out.H4, h4)
		}

		if h.Level == 5 {
			h5 := string(h.Children[0].AsLeaf().Literal)
			out.H5 = append(out.H5, h5)
		}

		if h.Level == 6 {
			h6 := string(h.Children[0].AsLeaf().Literal)
			out.H6 = append(out.H6, h6)
		}
	}
}

func shouldSkip(node ast.Node) bool {
	// ignore the two following as these don't contain parsable information.
	if _, ok := node.(*ast.Document); ok {
		return true
	}

	if _, ok := node.(*ast.Text); ok {
		return true
	}

	return false
}
