package ltxref

import (
	"html/template"
)

type Argumenttype int

const (
	_ Argumenttype = iota
	MANDARG
	MANDLIST
	OPTARG
	OPTLIST
	TODIMENORSPREADDIMEN
)

var argumenttypemap map[string]Argumenttype

func init() {
	argumenttypemap = map[string]Argumenttype{
		"mandarg":              MANDARG,
		"mandlist":             MANDLIST,
		"optarg":               OPTARG,
		"optlist":              OPTLIST,
		"todimenorspreaddimen": TODIMENORSPREADDIMEN,
	}
}

// The LaTeX reference knows about commands, environments and packages
type Ltxref struct {
	commands        []Command
	Environments    []Environment
	Documentclasses []Documentclass
	Packages        []Package
}

type Documentclass struct {
	Name             string
	Label            []string
	ShortDescription map[string]template.HTML
	Description      map[string]template.HTML
	Optiongroup      []Optiongroup
}

type Optiongroup struct {
	ShortDescription map[string]template.HTML
	Classoption      []Classoption
}

type Classoption struct {
	Name             string
	Default          bool
	ShortDescription map[string]template.HTML
	Description      map[string]template.HTML
}

type Command struct {
	Name             string
	Level            string
	Label            []string
	ShortDescription map[string]template.HTML
	Description      map[string]template.HTML
	Variant          []Variant
}

type Packageoption struct {
	Name             string
	Default          bool
	ShortDescription map[string]template.HTML
	Description      map[string]template.HTML
}

type Package struct {
	Name             string
	Level            string
	Label            []string
	ShortDescription map[string]template.HTML
	Description      map[string]template.HTML
	Commands         []Command
	Options          []Packageoption
}

type Environment struct {
	Name             string
	Level            string
	Label            []string
	ShortDescription map[string]template.HTML
	Description      map[string]template.HTML
	Variant          []Variant
}

// Some commands can have variants, such as \section or \section*.
// These commands are similar, so they should be documented together.
type Variant struct {
	Name        string
	Arguments   []Argument
	Description map[string]template.HTML
}

// Argument of a command or an environment
type Argument struct {
	Optional bool
	Name     string
	Type     Argumenttype
}
