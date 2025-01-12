package cwl

func (p *ProcessBase) RequiresInlineJavascript() *InlineJavascriptRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "InlineJavascriptRequirement" {
			return r.(*InlineJavascriptRequirement)
		}
	}
	return nil
}

func (p *ProcessBase) RequiresSchemaDef() *SchemaDefRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "SchemaDefRequirement" {
			return r.(*SchemaDefRequirement)
		}
	}
	return nil
}

func (p *ProcessBase) RequiresLoadListing() *LoadListingRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "LoadListingRequirement" {
			return r.(*LoadListingRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresDocker() *DockerRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "DockerRequirement" {
			return r.(*DockerRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) HitsDocker() *DockerRequirement {
	for _, r := range p.Hints {
		if r.ClassName() == "DockerRequirement" {
			return r.(*DockerRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresSoftware() *SoftwareRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "SoftwareRequirement" {
			return r.(*SoftwareRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresInitialWorkDir() *InitialWorkDirRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "InitialWorkDirRequirement" {
			return r.(*InitialWorkDirRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresEnvVar() *EnvVarRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "EnvVarRequirement" {
			return r.(*EnvVarRequirement)
		}
	}
	for _, r := range p.Hints {
		if r.ClassName() == "EnvVarRequirement" {
			return r.(*EnvVarRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresShellCommand() bool {
	for _, r := range p.Requirements {
		if r.ClassName() == "ShellCommandRequirement" {
			return true
		}
	}
	for _, r := range p.Hints {
		if r.ClassName() == "ShellCommandRequirement" {
			return true
		}
	}
	return false
}

func (p *CommandLineTool) RequiresResource() *ResourceRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "ResourceRequirement" {
			return r.(*ResourceRequirement)
		}
	}
	for _, r := range p.Hints {
		if r.ClassName() == "ResourceRequirement" {
			return r.(*ResourceRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresWorkReuse() *WorkReuse {
	for _, r := range p.Requirements {
		if r.ClassName() == "WorkReuse" {
			return r.(*WorkReuse)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresNetworkAccess() *NetworkAccess {
	for _, r := range p.Requirements {
		if r.ClassName() == "NetworkAccess" {
			return r.(*NetworkAccess)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresInplaceUpdateRequirement() *InplaceUpdateRequirement {
	for _, r := range p.Requirements {
		if r.ClassName() == "InplaceUpdateRequirement" {
			return r.(*InplaceUpdateRequirement)
		}
	}
	return nil
}

func (p *CommandLineTool) RequiresToolTimeLimit() *ToolTimeLimit {
	for _, r := range p.Requirements {
		if r.ClassName() == "ToolTimeLimit" {
			return r.(*ToolTimeLimit)
		}
	}
	return nil
}

func (p *Workflow) RequiresSubworkflowFeature() bool {
	for _, r := range p.Requirements {
		if r.ClassName() == "SubworkflowFeatureRequirement" {
			return true
		}
	}
	return false
}

func (p *Workflow) RequiresScatterFeature() bool {
	for _, r := range p.Requirements {
		if r.ClassName() == "ScatterFeatureRequirement" {
			return true
		}
	}
	return false
}

func (p *Workflow) RequiresMultipleInputFeature() bool {
	for _, r := range p.Requirements {
		if r.ClassName() == "MultipleInputFeatureRequirement" {
			return true
		}
	}
	return false
}

func (p *Workflow) RequiresStepInputExpression() bool {
	for _, r := range p.Requirements {
		if r.ClassName() == "StepInputExpressionRequirement" {
			return true
		}
	}
	return false
}

func (p *ProcessBase) InheritRequirement(parentRequirements, parentHints Requirements) {
	// 目前我们只根据ClassName确定是否继承
	// 若已有该Class，则不继承
	myRequirements := map[string]bool{}
	if p.Requirements == nil {
		p.Requirements = Requirements{}
	}
	for _, r := range p.Requirements {
		myRequirements[r.ClassName()] = true
	}
	for _, r := range parentRequirements {
		if _, ok := myRequirements[r.ClassName()]; !ok {
			p.Requirements = append(p.Requirements, r)
		}
	}
	myHints := map[string]bool{}
	if p.Hints == nil {
		p.Hints = Requirements{}
	}
	for _, r := range p.Hints {
		myHints[r.ClassName()] = true
	}
	for _, r := range parentHints {
		if _, ok := myHints[r.ClassName()]; !ok {
			p.Hints = append(p.Hints, r)
		}
	}
}
