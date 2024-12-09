package web

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

// <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
// <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
// <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
// <link rel="manifest" href="/site.webmanifest">

func LandingRoute() elem.Node {
	return elem.Div(
		attrs.Props{attrs.Lang: "en"},
		elem.Div(
			nil,
			elem.H1(attrs.Props{attrs.Class: "text-4xl"}, elem.Text("Walross")),
			elem.P(
				nil,
				elem.Text("I find a lot of joy in building software."),
			),
			elem.Ul(
				attrs.Props{attrs.Class: "py-8 flex gap-6"},
				elem.Li(
					nil,
					Link("/", "/"),
				),
				elem.Li(
					nil,
					Link("https://git.walross.co", "projects"),
				),
				elem.Li(
					nil,
					Link("https://github.com/rsbear", "github"),
				),
			),
		),
	)
}

func Link(href string, text string) elem.Node {
	return elem.A(
		attrs.Props{attrs.Href: href},
		elem.Span(
			attrs.Props{attrs.Class: "text-lively"},
			elem.Text(text),
		),
	)
}
