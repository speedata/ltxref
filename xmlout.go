package ltxref

import (
	"encoding/xml"
	"errors"
	"fmt"
	"html/template"
	"strings"
)

type description struct {
	XMLName xml.Name
	Lang    string `xml:"lang,attr"`
	Text    string `xml:",innerxml"`
}

func marshalDescription(eltname string, e *xml.Encoder, desc map[string]template.HTML) error {
	var err error
	for lang, text := range desc {
		d := description{}
		d.XMLName = xml.Name{Local: eltname}
		d.Lang = lang
		d.Text = "<![CDATA[" + string(text) + "]]>"
		err = e.Encode(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Command) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error

	cmdstartelt := xml.StartElement{Name: xml.Name{Local: "command"}}
	cmdstartelt.Attr = []xml.Attr{
		xml.Attr{Name: xml.Name{Local: "name"}, Value: c.Name},
		xml.Attr{Name: xml.Name{Local: "label"}, Value: strings.Join(c.Label, ",")},
		xml.Attr{Name: xml.Name{Local: "level"}, Value: c.Level},
	}
	err = e.EncodeToken(cmdstartelt)
	if err != nil {
		return err
	}
	err = marshalDescription("shortdescription", e, c.ShortDescription)
	if err != nil {
		return err
	}

	err = marshalDescription("description", e, c.Description)
	if err != nil {
		return err
	}

	err = e.Encode(c.Variant)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: cmdstartelt.Name})
}

func (node *Environment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	startElt := xml.StartElement{Name: xml.Name{Local: "environment"}}

	startElt.Attr = []xml.Attr{
		xml.Attr{Name: xml.Name{Local: "name"}, Value: node.Name},
		xml.Attr{Name: xml.Name{Local: "level"}, Value: node.Level},
		xml.Attr{Name: xml.Name{Local: "label"}, Value: strings.Join(node.Label, ",")},
	}

	err = e.EncodeToken(startElt)
	if err != nil {
		return err
	}

	err = marshalDescription("shortdescription", e, node.ShortDescription)
	if err != nil {
		return err
	}

	err = marshalDescription("description", e, node.Description)
	if err != nil {
		return err
	}

	err = e.Encode(node.Variant)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: startElt.Name})
	if err != nil {
		return err
	}

	return nil
}

func (node *Package) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	startElt := xml.StartElement{Name: xml.Name{Local: "package"}}

	startElt.Attr = []xml.Attr{
		xml.Attr{Name: xml.Name{Local: "name"}, Value: node.Name},
		xml.Attr{Name: xml.Name{Local: "level"}, Value: node.Level},
		xml.Attr{Name: xml.Name{Local: "label"}, Value: strings.Join(node.Label, ",")},
		xml.Attr{Name: xml.Name{Local: "loadspackages"}, Value: strings.Join(node.LoadsPackages, ",")},
	}

	err = e.EncodeToken(startElt)
	if err != nil {
		return err
	}

	err = marshalDescription("shortdescription", e, node.ShortDescription)
	if err != nil {
		return err
	}

	err = marshalDescription("description", e, node.Description)
	if err != nil {
		return err
	}

	err = e.Encode(node.Options)
	if err != nil {
		return err
	}

	err = e.Encode(node.Commands)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: startElt.Name})
	if err != nil {
		return err
	}

	return nil
}

func (node *Documentclass) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	startElt := xml.StartElement{Name: xml.Name{Local: "documentclass"}}

	startElt.Attr = []xml.Attr{
		xml.Attr{Name: xml.Name{Local: "name"}, Value: node.Name},
		xml.Attr{Name: xml.Name{Local: "level"}, Value: node.Level},
		xml.Attr{Name: xml.Name{Local: "label"}, Value: strings.Join(node.Label, ",")},
	}

	err = e.EncodeToken(startElt)
	if err != nil {
		return err
	}

	err = marshalDescription("shortdescription", e, node.ShortDescription)
	if err != nil {
		return err
	}

	err = marshalDescription("description", e, node.Description)
	if err != nil {
		return err
	}

	err = e.Encode(node.Optiongroup)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: startElt.Name})
	if err != nil {
		return err
	}

	return nil
}

func (node *Optiongroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	startElt := xml.StartElement{Name: xml.Name{Local: "optiongroup"}}

	startElt.Attr = []xml.Attr{
		xml.Attr{Name: xml.Name{Local: "name"}, Value: node.Name},
	}

	err = e.EncodeToken(startElt)
	if err != nil {
		return err
	}

	err = marshalDescription("shortdescription", e, node.ShortDescription)
	if err != nil {
		return err
	}

	err = e.Encode(node.Classoption)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: startElt.Name})
	if err != nil {
		return err
	}

	return nil
}

func (node *Classoption) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	startElt := xml.StartElement{Name: xml.Name{Local: "classoption"}}

	var dflt string
	if node.Default {
		dflt = "yes"
	} else {
		dflt = "no"
	}
	startElt.Attr = []xml.Attr{
		xml.Attr{Name: xml.Name{Local: "name"}, Value: node.Name},
		xml.Attr{Name: xml.Name{Local: "default"}, Value: dflt},
	}

	err = e.EncodeToken(startElt)
	if err != nil {
		return err
	}

	err = marshalDescription("shortdescription", e, node.ShortDescription)
	if err != nil {
		return err
	}
	err = marshalDescription("description", e, node.Description)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: startElt.Name})
}

func (node *Packageoption) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	startElt := xml.StartElement{Name: xml.Name{Local: "packageoption"}}
	startElt.Attr = []xml.Attr{
		xml.Attr{Name: xml.Name{Local: "name"}, Value: node.Name},
	}

	err = e.EncodeToken(startElt)
	if err != nil {
		return err
	}
	err = marshalDescription("description", e, node.Description)
	if err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: startElt.Name})
}

func (a *Argument) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	startElt := xml.StartElement{Name: xml.Name{Local: "argument"}}
	var opt string
	if a.Optional {
		opt = "yes"
	} else {
		opt = "no"
	}
	startElt.Attr = append(startElt.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: a.Name})
	startElt.Attr = append(startElt.Attr, xml.Attr{Name: xml.Name{Local: "optional"}, Value: opt})
	startElt.Attr = append(startElt.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: argumentTypeReveseMap[a.Type]})

	err = e.EncodeToken(startElt)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: startElt.Name})
	if err != nil {
		return err
	}

	return nil
}

func (v *Variant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	variantStartElt := xml.StartElement{Name: xml.Name{Local: "variant"}}
	variantStartElt.Attr = append(variantStartElt.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: v.Name})
	err := e.EncodeToken(variantStartElt)
	if err != nil {
		return err
	}
	e.Encode(v.Arguments)
	marshalDescription("description", e, v.Description)

	err = e.EncodeToken(xml.EndElement{Name: variantStartElt.Name})
	if err != nil {
		return err
	}

	return nil
}

func (l *Ltxref) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	eltname := xml.Name{Local: "ltxref", Space: "urn:speedata.de:2015/latexref"}
	startelt := xml.StartElement{Name: eltname}
	startelt.Attr = append(startelt.Attr, xml.Attr{Name: xml.Name{Local: "version"}, Value: l.Version})

	e.Indent("", "  ")
	err := e.EncodeToken(startelt)
	if err != nil {
		return err
	}
	err = e.Encode(l.commands)
	if err != nil {
		return err
	}
	err = e.Encode(l.environments)
	if err != nil {
		return err
	}
	err = e.Encode(l.documentclasses)
	if err != nil {
		return err
	}

	err = e.Encode(l.packages)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: eltname})
	if err != nil {
		return err
	}

	if false {

		return errors.New("dummy error")
	}
	return nil
}

func (l *Ltxref) ToXML() {
	a, err := xml.Marshal(l)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(a))
}
