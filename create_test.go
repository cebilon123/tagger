package tagger

import "testing"

func TestTag_Render(t *testing.T) {
	type fields struct {
		Type     TagType
		Value    string
		children []Tag
		params   []Param
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Form_WithInputAndParagraph", fields: fields{
			Type:     Form,
			Value:    "",
			children: []Tag{
				{
					Type:     Paragraph,
					Value:    "Hello text",
					children: nil,
					params: []Param{
						{
							Key:   "class",
							Value: "paragraph-extended",
						},
					},
				},
				{
					Type:     Input,
					Value:    "",
					children: nil,
					params: []Param{
						{
							Key:   "class",
							Value: "input-class",
						},
						{
							Key:   "body",
							Value: "vr.input",
						},
					},
				},
			},
			params: []Param{
				{
					Key:   "class",
					Value: "form form-expanded form-hello",
				},
			},
		}, want: "<form class=\"form form-expanded form-hello\">\n<p class=\"paragraph-extended\">\nHello text</p>\n<input class=\"input-class\" body=\"vr.input\">\n</input>\n</form>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Tag{
				Type:     tt.fields.Type,
				Value:    tt.fields.Value,
				children: tt.fields.children,
				params:   tt.fields.params,
			}
			if got := f.Render(); got != tt.want {
				t.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
