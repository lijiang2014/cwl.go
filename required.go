package cwl

// RequiredInputs represents "inputs" field in CWL.
type RequiredInputs []RequiredInput

// New constructs new "RequiredInputs" struct.
func (inputs RequiredInputs) New(i interface{}) RequiredInputs {
	dest := RequiredInputs{}
	switch x := i.(type) {
	case []interface{}:
		for _, elm := range x {
			input := RequiredInput{}.New(elm)
			dest = append(dest, input)
		}
	case map[string]interface{}:
		for key, val := range x {
			input := RequiredInput{}.New(val)
			input.ID = key
			dest = append(dest, input)
		}
	}
	return dest
}

// RequiredInput represents an element of "inputs" in "CWL".
type RequiredInput struct {
	ID      string
	Type    *InputType
	Doc     string
	Label   string
	Binding *InputBinding
	Default *InputDefault
}

// New constructs "RequiredInput" struct from interface{}.
func (input RequiredInput) New(i interface{}) RequiredInput {
	dict, ok := i.(map[string]interface{})
	if !ok {
		return RequiredInput{}
	}
	return input.NewFromDict(dict)
}

// NewFromDict constructs "RequiredInput" from dictionary formed map.
func (input RequiredInput) NewFromDict(dict map[string]interface{}) RequiredInput {
	dest := RequiredInput{}
	for key, val := range dict {
		switch key {
		case "id":
			dest.ID = val.(string)
		case "type":
			dest.Type = dest.NewType(val)
		case "label":
			dest.Label = val.(string)
		case "doc":
			dest.Doc = val.(string)
		case "inputBinding":
			dest.Binding = dest.NewBinding(val)
		case "default":
			dest.Default = dest.NewDefault(val)
		}
	}
	return dest
}

// InputType represents "type" field in an element of "inputs".
type InputType struct {
	Type    string
	Items   string
	Binding *InputBinding
}

// NewType constructs new "InputType".
func (input RequiredInput) NewType(i interface{}) *InputType {
	dest := new(InputType)
	switch x := i.(type) {
	case string:
		dest.Type = x
	case map[string]interface{}:
		if val, ok := x["type"]; ok {
			dest.Type = val.(string)
		}
		if val, ok := x["items"]; ok {
			dest.Items = val.(string)
		}
		if val, ok := x["inputBinding"]; ok {
			dest.Binding = RequiredInput{}.NewBinding(val)
		}
	}
	return dest
}

// InputBinding represents "inputBinding" field in an element of "inputs".
type InputBinding struct {
	Position  int
	Prefix    string
	Separator string
}

// NewBinding constructs new "InputBinding".
func (input RequiredInput) NewBinding(i interface{}) *InputBinding {
	dest := new(InputBinding)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "position":
				dest.Position = int(val.(float64))
			case "prefix":
				dest.Prefix = val.(string)
			case "itemSeparator":
				dest.Separator = val.(string)
			}
		}
	}
	return dest
}

// InputDefault represents "default" field in an element of "inputs".
type InputDefault struct {
	Class    string
	Location string
}

// NewDefault constructs new "InputDefault".
func (input RequiredInput) NewDefault(i interface{}) *InputDefault {
	dest := new(InputDefault)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "class":
				dest.Class = val.(string)
			case "location":
				dest.Location = val.(string)
			}
		}
	}
	return dest
}