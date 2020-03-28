package utils

import "github.com/razzkumar/PR-Automation/logger"

func EnvLoadError(env string, envName string) {

	if env == "" {
		logger.FailOnNoFlag("Unbale to load " + envName)
	}

}
