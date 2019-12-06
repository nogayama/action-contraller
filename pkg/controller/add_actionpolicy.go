package controller

import (
	"github.com/nogayama/action-contraller/pkg/controller/actionpolicy"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, actionpolicy.Add)
}
