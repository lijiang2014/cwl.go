package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_count_lines10_wf(t *testing.T) {
	f := cwl("count-lines10-wf.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")

	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].ID).ToBe("count_output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Outputs[0].Source).ToBe([]string{"step1/count_output"})

	Expect(t, root.Requirements[0].Class).ToBe("SubworkflowFeatureRequirement")

	Expect(t, root.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("file1")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("count_output")

	Expect(t, root.Steps[0].Run.Class).ToBe("Workflow")
	Expect(t, root.Steps[0].Run.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].Run.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Steps[0].Run.Outputs[0].ID).ToBe("count_output")
	Expect(t, root.Steps[0].Run.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Steps[0].Run.Outputs[0].Source).ToBe([]string{"step2/output"})
	// Recursive steps
	Expect(t, root.Steps[0].Run.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].Run.Steps[0].Run.ID).ToBe("wc-tool.cwl")
	Expect(t, root.Steps[0].Run.Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].Run.Steps[0].In[0].Source[0]).ToBe("file1")
	Expect(t, root.Steps[0].Run.Steps[0].Out[0].ID).ToBe("output")
	Expect(t, root.Steps[0].Run.Steps[1].ID).ToBe("step2")
	Expect(t, root.Steps[0].Run.Steps[1].Run.ID).ToBe("parseInt-tool.cwl")
	Expect(t, root.Steps[0].Run.Steps[1].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].Run.Steps[1].In[0].Source[0]).ToBe("step1/output")
	Expect(t, root.Steps[0].Run.Steps[1].Out[0].ID).ToBe("output")
}