package core

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
		{name: "ValidParameters_NoneChildren_returnsForm", fields: fields{
			children: nil,
			params: []Param{
				{Key: "class", Value: "form"},
			},
		}, want: "<Form class=\"form\"></Form>"},
		{name: "ValidParameters_MultiValueInParam_NoneChildren_returnsForm", fields: fields{
			children: nil,
			params: []Param{
				{Key: "class", Value: "form vtb-form new-line"},
			},
		}, want: "<Form class=\"form vtb-form new-line\"></Form>"},
		{name: "NoneParameters_NoneChildren_ReturnEmptyForm", fields: fields{
			children: nil,
			params: nil,
		}, want: "<Form></Form>"},
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
