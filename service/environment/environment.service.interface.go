package environment

//go:generate mockery --name=EnvironmentServiceInterface
type EnvironmentServiceInterface interface {
	GetEnvVars() *EnvironmentVars
}
