package tagger

import "testing"

func TestForm_String(t *testing.T) {
	type fields struct {
		children []Creator
		params   []Param
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "ValidParameters_NoneChildren", fields: fields{
			children: nil,
			params: []Param{
				{Key: "class", Value: "form"},
			},
		}, want: "<Form class=\"form\"></Form>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Form{
				children: tt.fields.children,
				params:   tt.fields.params,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
