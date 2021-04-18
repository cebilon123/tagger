package core

import "strings"

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

//Form represents html form.
type Form struct {
	children []Creator
	params   []Param
}

func (f Form) String() string {
	html := "<Form"

	stringifyParams := make([]string, 0)
	for _, param := range f.params {
		stringifyParams = append(stringifyParams, param.String())
	}

	params := strings.Join(stringifyParams, ", ")

	//If there is no params append html without white space
	if len(params) > 0 {
		html += " " + params + ">"
	} else {
		html += ">"
	}

	for i := range f.children {
		child := f.children[i]

		//append with rendered child and new line
		html += child.Render() + "\n"
	}

	if f.children != nil || len(f.children) > 0 {
		html += "\n"
	}
	html += "</Form>"

	return html
}

func (f Form) Render() string {
	return f.String()
}
