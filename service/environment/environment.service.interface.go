package environment

type EnvironmentServiceInterface interface {
	GetEnvVars() *EnvironmentVars
}