package environment

import "os"

var envVars *EnvironmentVars

type EnvironmentService struct {
}

func NewEnvironmentService() *EnvironmentService {
	return &EnvironmentService{}
}

func (e *EnvironmentService) GetEnvVars() *EnvironmentVars {
	if envVars == nil {
		envVars = &EnvironmentVars{
			Test: os.Getenv("test"),
		}
	}

	return envVars
}
