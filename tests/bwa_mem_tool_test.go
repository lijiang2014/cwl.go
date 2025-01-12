package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/lijiang2014/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_bwa_mem_tool(t *testing.T) {
	f := load("bwa-mem-tool.cwl")
	doc := cwl.NewCWL()
	err := doc.Decode(f)
	root := doc.Process.(*cwl.CommandLineTool)
	Expect(t, err).ToBe(nil)
	Expect(t, root.CWLVersion).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	//Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Requirements[0].ClassName()).ToBe("ResourceRequirement")
	Expect(t, int(*root.Requirements[0].(*cwl.ResourceRequirement).CoresMin.Long)).ToBe(2)

	Expect(t, len(root.Inputs)).ToBe(5)
	//Expect(t, root.Inputs).TypeOf("cwl.Inputs")
	sort.Sort(root.Inputs)
	
	//Expect(t, root.Inputs[0].ID).ToBe("args.py")
	//Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	//Expect(t, root.Inputs[0].Binding.Position).ToBe(-1)
	//Expect(t, root.Inputs[1].ID).ToBe("min_std_max_min")
	//Expect(t, root.Inputs[1].Types[0].Type).ToBe("array")
	//Expect(t, root.Inputs[1].Types[0].Items[0].Type).ToBe("int")
	//Expect(t, root.Inputs[1].Binding.Position).ToBe(1)
	//Expect(t, root.Inputs[1].Binding.Separator).ToBe(",")
	//Expect(t, root.Inputs[2].ID).ToBe("minimum_seed_length")
	//Expect(t, root.Inputs[2].Types[0].Type).ToBe("int")
	//Expect(t, root.Inputs[2].Binding.Position).ToBe(1)
	//Expect(t, root.Inputs[2].Binding.Prefix).ToBe("-m")
	//Expect(t, root.Inputs[3].ID).ToBe("reference")
	//Expect(t, root.Inputs[3].Types[0].Type).ToBe("File")
	//Expect(t, root.Inputs[3].Binding.Position).ToBe(2)
	//Expect(t, root.Inputs[4].ID).ToBe("reads")
	//Expect(t, root.Inputs[4].Types[0].Type).ToBe("array")
	//Expect(t, root.Inputs[4].Types[0].Items[0].Type).ToBe("File")
	//Expect(t, root.Inputs[4].Binding.Position).ToBe(3)
	//
	//Expect(t, root.Outputs[0].ID).ToBe("sam")
	//Expect(t, root.Outputs[0].Types[0].Type).ToBe("null")
	//Expect(t, root.Outputs[0].Types[1].Type).ToBe("File")
	//Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"output.sam"})
	//Expect(t, root.Outputs[1].ID).ToBe("args")
	//Expect(t, root.Outputs[1].Types[0].Type).ToBe("array")
	//Expect(t, root.Outputs[1].Types[0].Items[0].Type).ToBe("string")
}
