package tagger

import "strings"

type TagType int

func (t TagType) String() string {
	return [...]string{"form", "input", "div", "button", "img", "video", "p", ""}[t]
}

const (
	Form TagType = iota
	Input
	Div
	Button
	Image
	Video
	Paragraph
	Plain
)

//Creator represents object which can be rendered into html string
//e.g. <div>tagger</div>.
type Creator interface {
	Render() string
}

type Param struct {
	Key   string
	Value string
}

func (p Param) String() string {
	return p.Key + "=" + "\"" + p.Value + "\""
}

type Tag struct {
	Type     TagType
	Value    string
	Children []Tag
	Params   []Param
}

func (f Tag) String() string {
	//If tag is of type plain, just return its value
	if f.Type == Plain {
		return f.Value
	}

	html := "<" + f.Type.String()

	stringifyParams := make([]string, 0)
	for _, param := range f.params {
		stringifyParams = append(stringifyParams, param.String())
	}

	params := strings.Join(stringifyParams, " ")

	//If there is no params append html without white space
	if len(params) > 0 {
		html += " " + params + ">"
	} else {
		html += ">"
	}

	html += "\n"

	for i := range f.children {
		child := f.children[i]

		//append with rendered child and new line
		html += child.Render() + "\n"
	}

	if len(f.Value) > 0 {
		html += f.Value
	}

	html += "</" + f.Type.String() + ">"

	return html
}

func (f Tag) Render() string {
	return f.String()
}
