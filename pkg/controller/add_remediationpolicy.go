package controller

import (
	"github.com/nogayama/action-contraller/pkg/controller/remediationpolicy"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, remediationpolicy.Add)
}
