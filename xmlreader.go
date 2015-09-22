package ltxref

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
)

func dummy() {
	fmt.Println("")
}

func ReadXMLFile(filename string) (Ltxref, error) {
	r, err := os.Open(filename)
	if err != nil {
		return Ltxref{}, err
	}
	defer r.Close()
	return ReadXML(r)
}

func ReadXMLData(data []byte) (Ltxref, error) {
	r := bytes.NewReader(data)
	return ReadXML(r)
}

func ReadXML(r io.Reader) (Ltxref, error) {
	lr := Ltxref{}
	dec := xml.NewDecoder(r)

	for {
		t, err := dec.Token()
		if err != nil {
			break
		}
		switch v := t.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case "command":
				cmd := readCommand(v.Attr, dec)
				lr.Commands = append(lr.Commands, cmd)
			case "environment":
				env := readEnvironment(v.Attr, dec)
				lr.Environments = append(lr.Environments, env)
			case "package":
				env := readPackage(v.Attr, dec)
				lr.Packages = append(lr.Packages, env)
			}
		case xml.EndElement:
			switch v.Name.Local {
			case "command":
				return lr, nil
			}
		}
	}
	return lr, nil
}

func readArgument(attributes []xml.Attr, dec *xml.Decoder) Argument {
	argument := Argument{}
	for _, attribute := range attributes {
		switch attribute.Name.Local {
		case "name":
			argument.Name = attribute.Value
		case "optional":
			argument.Optional = attribute.Value == "yes"
		case "type":
			argument.Type = argumenttypemap[attribute.Value]
		}
	}
	return argument
}

func readVariant(attributes []xml.Attr, dec *xml.Decoder) Variant {
	variant := Variant{}
	variant.Description = make(map[string]template.HTML)
	for _, attribute := range attributes {
		if attribute.Name.Local == "name" {
			variant.Name = attribute.Value
		}
	}
	for {
		t, err := dec.Token()
		if err != nil {
			break
		}
		switch v := t.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case "argument":
				variant.Arguments = append(variant.Arguments, readArgument(v.Attr, dec))
			case "description":
				lang, text := readDescription(v.Attr, dec)
				variant.Description[lang] = text
			}
		case xml.EndElement:
			if v.Name.Local == "variant" {
				return variant
			}
		}
	}
	return variant
}

func readPackage(attributes []xml.Attr, dec *xml.Decoder) Package {
	pkg := Package{}
	pkg.ShortDescription = make(map[string]template.HTML)
	pkg.Description = make(map[string]template.HTML)
	for _, attribute := range attributes {
		switch attribute.Name.Local {
		case "name":
			pkg.Name = attribute.Value
		}
	}
	for {
		t, err := dec.Token()
		if err != nil {
			break
		}
		switch v := t.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case "shortdescription":
				lang, text := readDescription(v.Attr, dec)
				pkg.ShortDescription[lang] = text
			case "description":
				lang, text := readDescription(v.Attr, dec)
				pkg.Description[lang] = text
			case "command":
				pkg.Commands = append(pkg.Commands, readCommand(v.Attr, dec))
			}
		case xml.EndElement:
			switch v.Name.Local {
			case "package":
				return pkg
			}
		}

	}
	return pkg
}

func readEnvironment(attributes []xml.Attr, dec *xml.Decoder) Environment {
	env := Environment{}
	env.ShortDescription = make(map[string]template.HTML)
	env.Description = make(map[string]template.HTML)
	for _, attribute := range attributes {
		switch attribute.Name.Local {
		case "name":
			env.Name = attribute.Value
		case "level":
			env.Level = attribute.Value
		case "label":
			env.Label = strings.Split(attribute.Value, ",")
		}
	}
	for {
		t, err := dec.Token()
		if err != nil {
			break
		}
		switch v := t.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case "shortdescription":
				lang, text := readDescription(v.Attr, dec)
				env.ShortDescription[lang] = text
			case "description":
				lang, text := readDescription(v.Attr, dec)
				env.Description[lang] = text
			case "variant":
				variant := readVariant(v.Attr, dec)
				env.Variant = append(env.Variant, variant)
			}
		case xml.EndElement:
			switch v.Name.Local {
			case "environment":
				return env
			}
		}

	}
	return env
}

func readCommand(attributes []xml.Attr, dec *xml.Decoder) Command {
	cmd := Command{}
	cmd.ShortDescription = make(map[string]template.HTML)
	cmd.Description = make(map[string]template.HTML)
	for _, attribute := range attributes {
		switch attribute.Name.Local {
		case "name":
			cmd.Name = attribute.Value
		case "level":
			cmd.Level = attribute.Value
		case "label":
			cmd.Label = strings.Split(attribute.Value, ",")
		}

	}

	for {
		t, err := dec.Token()
		if err != nil {
			break
		}
		switch v := t.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case "shortdescription":
				lang, text := readDescription(v.Attr, dec)
				cmd.ShortDescription[lang] = text
			case "description":
				lang, text := readDescription(v.Attr, dec)
				cmd.Description[lang] = text
			case "variant":
				variant := readVariant(v.Attr, dec)
				cmd.Variant = append(cmd.Variant, variant)
			}
		case xml.EndElement:
			switch v.Name.Local {
			case "command":
				return cmd
			}

		}

	}
	return cmd
}

func readDescription(attributes []xml.Attr, dec *xml.Decoder) (string, template.HTML) {
	var lang string
	for _, attribute := range attributes {
		if attribute.Name.Local == "lang" {
			lang = attribute.Value
		}
	}
	var str string
	for {

		t, err := dec.Token()
		if err != nil {
			break
		}
		switch v := t.(type) {
		case xml.CharData:
			str += string(v)
		case xml.EndElement:
			return lang, template.HTML(str)
		default:
			// huh?
			fmt.Printf("%#v\n", v)
		}
	}
	// never reached!?!?
	return lang, template.HTML(str)
}
