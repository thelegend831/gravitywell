package actions

import (
	"fmt"
	"github.com/AlexsJones/gravitywell/configuration"
	"github.com/AlexsJones/gravitywell/kinds"
	"github.com/AlexsJones/gravitywell/vcs"
	"github.com/google/logger"
	"strings"
)

func ApplicationProcessor(commandFlag configuration.CommandFlag,
	opt configuration.Options, clusterName string, application kinds.Application) {

	executeDeployment(application, opt, clusterName, commandFlag)
}

func executeDeployment(deployment kinds.Application, opt configuration.Options,
	clusterName string, commandFlag configuration.CommandFlag) {
	logger.Info(fmt.Sprintf("Loading deployment %s\n", deployment.Name))

	remoteVCSRepoName, err := vcs.FetchRepo(deployment.Git, opt)
	if err != nil {
		logger.Error(err.Error())

		return
	}

	for _, a := range deployment.Action {
		switch strings.ToLower(a.Execute.Kind) {
		case "shell":
			ExecuteShellAction(a, opt, remoteVCSRepoName)

		case "kubernetes":
			ExecuteKubernetesAction(a, clusterName, deployment, commandFlag, opt, remoteVCSRepoName)
		}

	}
}
