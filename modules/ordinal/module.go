package ordinal

import (
	"html/template"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/templates"
)

func init() {
	caddy.RegisterModule(Ordinal{})
}

// Interface guards.
var (
	_ templates.CustomFunctions = (*Ordinal)(nil)
)

// Ordinal provides the conversion of integers to ordinals to Caddy's template directive.
type Ordinal struct{}

// CaddyModule returns info about the module to Caddy.
func (Ordinal) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.templates.functions.ordinal",
		New: func() caddy.Module { return new(Ordinal) },
	}
}

// UnmarshalCaddyfile is a nop.
func (*Ordinal) UnmarshalCaddyfile(_ *caddyfile.Dispenser) error {
	return nil
}

// CustomTemplateFunctions returns a map of custom template functions.
func (o *Ordinal) CustomTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"ordinal": intToOrdinal,
	}
}
