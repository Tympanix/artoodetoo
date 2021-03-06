package meta

import (
	"log"

	"github.com/Tympanix/artoodetoo/unit"
)

// Debug is used to debug program code
type Debug struct {
	Log interface{} `io:"input"`
}

func init() {
	unit.Register(new(Debug))
}

// Describe debugging
func (d *Debug) Describe() string {
	return "Debugger to print statements to the console"
}

// Execute debugging
func (d *Debug) Execute() error {
	log.Println(d.Log)
	return nil
}
