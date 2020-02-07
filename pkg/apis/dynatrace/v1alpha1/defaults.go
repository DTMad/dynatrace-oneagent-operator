package v1alpha1

func SetDefaults_OneAgentSpec(obj *OneAgentSpec) {
	if obj.WaitReadySeconds == nil {
		obj.WaitReadySeconds = new(uint16)
		*obj.WaitReadySeconds = 300
	}

	if obj.Image == "" {
		obj.Image = "docker.io/dynatrace/oneagent:latest"
	}

	if obj.ServiceAccountName == "" {
		obj.ServiceAccountName = "dynatrace-oneagent"
	}
}
