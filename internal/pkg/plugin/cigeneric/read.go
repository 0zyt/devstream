package cigeneric

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/plugin/installer"
	"github.com/devstream-io/devstream/internal/pkg/plugin/installer/ci/cifile"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/log"
)

func Read(options configmanager.RawOptions) (statemanager.ResourceStatus, error) {
	operator := &installer.Operator{
		PreExecuteOperations: installer.PreExecuteOperations{
			cifile.Validate,
		},
		GetStatusOperation: cifile.GetCIFileStatus,
	}

	status, err := operator.Execute(options)
	if err != nil {
		return nil, err
	}

	log.Debugf("Return map: %v", status)
	return status, nil

}
