package controller

import (
	"github.com/speak2jc/keeva-operator/pkg/controller/keevakind"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, keevakind.Add)
}
