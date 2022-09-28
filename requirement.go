package cwl

// Requirement represent an element of "requirements".
type Requirement struct {
	Class string
	InlineJavascriptRequirement
	SchemaDefRequirement
	DockerRequirement
	SoftwareRequirement
	InitialWorkDirRequirement
	EnvVarRequirement
	ShellCommandRequirement
	ResourceRequirement
	Import string
}

// New constructs "Requirement" struct from interface.
func (_ Requirement) New(i interface{}) Requirement {
	dest := Requirement{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "class":
				dest.Class = v.(string)
			case "dockerPull":
				dest.DockerPull = v.(string)
			case "dockerOutputDirectory":
				dest.DockerOutputDirectory = v.(string)
			case "types":
				dest.Types = Type{}.NewList(v)
			case "expressionLib":
				dest.ExpressionLib = JavascriptExpression{}.NewList(v)
			case "envDef":
				dest.EnvDef = EnvDef{}.NewList(v)
			case "listing":
				dest.Listing = Entry{}.NewList(v)
			case "$import":
				dest.Import = v.(string)
			case "coresMin":
				dest.CoresMin = int(v.(float64))
			}
		}
	}
	return dest
}

// Requirements represents "requirements" field in CWL.
type Requirements []Requirement

// New constructs "Requirements" struct from interface.
func (_ Requirements) New(i interface{}) Requirements {
	dest := Requirements{}
	switch x := i.(type) {
	case []interface{}:
		for _, r := range x {
			dest = append(dest, Requirement{}.New(r))
		}
	case map[string]interface{}:
		for key, v := range x {
			r := Requirement{}.New(v)
			r.Class = key
			dest = append(dest, r)
		}
	}
	return dest
}


// RequiresShellCommand
func (r Requirements) RequiresShellCommand() bool {
	for _, x := range r {
		if x.Class == "ShellCommandRequirement" {
			return true
		}
	}
	return false
}

// RequiresShellCommand
func (r Requirements) RequiresResource() *Requirement {
	for i, x := range r {
		if x.Class == "ResourceRequirement" {
			return &r[i]
		}
	}
	return nil
}

func (r Requirements) RequiresInlineJavascript() *Requirement {
	for _, x := range r {
		if x.Class == "InlineJavascriptRequirement" {
			return &x
		}
	}
	return nil
}

func (r Requirements) RequiresSchemaDef() *Requirement {
	for _, x := range r {
		if x.Class == "SchemaDefRequirement" {
			return &x
		}
	}
	return nil
}

func (r Requirements) RequiresInitialWorkDir() *Requirement {
	for _, x := range r {
		if x.Class == "InitialWorkDirRequirement" {
			return &x
		}
	}
	return nil
}

func (r Requirements) RequiresEnvVar() *Requirement {
	for _, x := range r {
		if x.Class == "EnvVarRequirement" {
			return &x
		}
	}
	return nil
}

// InlineJavascriptRequirement is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#InlineJavascriptRequirement
type InlineJavascriptRequirement struct {
	ExpressionLib []JavascriptExpression
}

// JavascriptExpression represents an element of "expressionLib" of InlineJavascriptRequirement.
type JavascriptExpression struct {
	Kind  string // could be "" or "$include"
	Value string
}

func (_ JavascriptExpression) NewList(i interface{}) []JavascriptExpression {
	dest := []JavascriptExpression{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, JavascriptExpression{}.New(v))
		}
	}
	return dest
}

func (_ JavascriptExpression) New(i interface{}) JavascriptExpression {
	dest := JavascriptExpression{}
	switch x := i.(type) {
	case string:
		dest.Kind = "$execute"
		dest.Value = x
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "$include":
				dest.Kind = key
				dest.Value = v.(string)
			}
		}
	}
	return dest
}

// SchemaDefRequirement is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#SchemaDefRequirement
type SchemaDefRequirement struct {
	Types []Type
}

// DockerRequirement is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#DockerRequirement
type DockerRequirement struct {
	DockerPull            string
	DockerLoad            string
	DockerFile            string
	DockerImport          string
	DockerImageID         string
	DockerOutputDirectory string
}

// SoftwareRequirement is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#SoftwareRequirement
type SoftwareRequirement struct {
	Packages []SoftwarePackage
}

// SoftwarePackage represents an element of SoftwarePackage.Packages
type SoftwarePackage struct {
	Package  string
	Versions []string
	Specs    []string
}

// InitialWorkDirRequirement is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#InitialWorkDirRequirement
type InitialWorkDirRequirement struct {
	Listing []Entry
}

// EnvVarRequirement  is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#EnvVarRequirement
type EnvVarRequirement struct {
	EnvDef []EnvDef
}

// ShellCommandRequirement is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#ShellCommandRequirement
type ShellCommandRequirement struct {
}

// ResourceRequirement is supposed to be embeded to Requirement.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#ResourceRequirement
type ResourceRequirement struct {
	CoresMin int
	CoresMax int
}
