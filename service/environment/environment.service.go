package environment

import "os"

var envVars *EnvironmentVars

type EnvironmentService struct {

}

func (e *EnvironmentService) GetEnvVars() *EnvironmentVars {
	if envVars == nil {
		envVars = &EnvironmentVars{
			zendeskUsername:              os.Getenv("zendeskUsername"),
			zendeskPassword:              os.Getenv("zendeskPassword"),
			sendSafelyApiKey:             os.Getenv("sendSafelyApiKey"),
			sendSafelyApiSecret:          os.Getenv("sendSafelyApiSecret"),
			sendSafelyAdminEmail:         os.Getenv("sendSafelyAdminEmail"),
			ccloudUsername:               os.Getenv("ccloudUsername"),
			ccloudPassword:               os.Getenv("ccloudPassword"),
			googleApplicationCredentials: os.Getenv("googleApplicationCredentials"),
			jwtSigningKey:                os.Getenv("jwtSigningKey"),
		}
	}

	return envVars
}

