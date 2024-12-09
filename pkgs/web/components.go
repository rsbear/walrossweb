package web

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

// <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
// <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
// <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
// <link rel="manifest" href="/site.webmanifest">

func BaseTemplate(content elem.Node) elem.Node {
	return elem.Html(
		attrs.Props{attrs.Lang: "en"},
		elem.Head(
			nil,
			elem.Meta(attrs.Props{attrs.Charset: "UTF-8"}),
			elem.Meta(attrs.Props{
				attrs.Name:    "viewport",
				attrs.Content: "width=device-width, initial-scale=1.0",
			}),
			elem.Link(attrs.Props{
				attrs.Href: "/static/styles.css",
				attrs.Rel:  "stylesheet",
				attrs.Type: "text/css",
			}),
			elem.Link(attrs.Props{
				attrs.Rel:   "apple-touch-icon",
				attrs.Sizes: "180x180",
				attrs.Href:  "/static/apple-touch-icon.png",
			}),
			elem.Link(attrs.Props{
				attrs.Rel:   "icon",
				attrs.Type:  "image/png",
				attrs.Sizes: "32x32",
				attrs.Href:  "/static/favicon-32x32.png",
			}),
			elem.Link(attrs.Props{
				attrs.Rel:   "icon",
				attrs.Type:  "image/png",
				attrs.Sizes: "16x16",
				attrs.Href:  "/static/favicon-16x16.png",
			}),
			elem.Link(attrs.Props{
				attrs.Rel:  "manifest",
				attrs.Href: "/static/site.webmanifest",
			}),
		),
		elem.Body(
			attrs.Props{attrs.Class: "font-mono min-h-screen rose-pine bg-bgcolor text-subtext"},
			elem.Div(attrs.Props{attrs.Class: "p-10"},
				content,
			),
		),
	)
}
