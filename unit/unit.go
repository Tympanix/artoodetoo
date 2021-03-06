package unit

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/Tympanix/artoodetoo/state"
	"github.com/Tympanix/artoodetoo/subject"
	"github.com/Tympanix/artoodetoo/types"
)

const (
	id     = "id"
	output = "output"
	input  = "input"
)

// NewUnit creates a new unit from events, actions and converters
func NewUnit(a Action) *Unit {
	return &Unit{
		Subject: *subject.New(a, new(ActionResolver)),
		Desc:    a.Describe(),
		action:  a,
	}
}

// ActionResolver is a subject resolver which can look up actions
type ActionResolver struct{}

// ResolveSubject resolves an action
func (a *ActionResolver) ResolveSubject(t string) (action interface{}, err error) {
	action, ok := GetActionByID(t)
	if !ok {
		err = fmt.Errorf("Could not resolve action with type %v", t)
	}
	return
}

// Unit wraps the elements of the application and extends it's functionality.
type Unit struct {
	subject.Subject
	Desc   string `json:"description"`
	action Action
}

// Validate makes sure that the unit is set up correctly for execution
func (c *Unit) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("Unit was not given a name")
	}
	return c.Subject.Validate()
}

// Execute executes the unit by evaluating input and assigning output
func (c *Unit) Execute() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New(fmt.Sprint(x))
			}
		}
	}()
	return c.action.Execute()
}

// Action returns the underlying action represented by the unit
func (c *Unit) Action() Action {
	return c.action
}

// RunAsync runs the unit asynchronously
func (c *Unit) RunAsync(waitgroup *sync.WaitGroup, ts types.TupleSpace, errchan chan<- error) {
	go func() {
		defer waitgroup.Done()
		if err := c.AssignInput(ts); err != nil {
			if _, ok := err.(*state.Closed); !ok {
				errchan <- err
			}
			return
		}
		if err := c.Execute(); err != nil {
			errchan <- err
			return
		}
		if err := c.StoreOutput(ts); err != nil {
			errchan <- err
			return
		}
	}()
}

// UnmarshalJSON is used to transform json data into a units
func (c *Unit) UnmarshalJSON(data []byte) error {
	type jsonUnit Unit
	u := jsonUnit(*c)
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}

	*c = Unit(u)

	err := c.RebuildSubject(new(ActionResolver))
	if err != nil {
		return err
	}

	newAction, ok := c.GetSubject().(Action)
	if !ok {
		return fmt.Errorf("Internal error while parsing unit")
	}
	c.action = newAction

	return nil
}
