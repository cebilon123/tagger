package tagger

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
	html += " " + params + ">"

	for i := range f.children {
		child := f.children[i]

		//append with rendered child and new line
		html += child.Render() + "\n"
	}

	html += "\n" + "</Form>"

	return html
}

func (f Form) Render() string {
	return f.String()
}
