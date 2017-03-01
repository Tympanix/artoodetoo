package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/hub"
	"github.com/Tympanix/automato/task"
)

func main() {
	// Create a new event and give it a name for reference
	event := task.NewComponent(example.PersonEvent{})
	event.SetName("person")

	// Create a new converter, set its name, and give it an ingredient
	converter := task.NewComponent(example.StringConverter{})
	converter.SetName("strcon").AddIngredient(task.Ingredient{
		Type:     task.IngredientStatic,
		Argument: "String",
		Value:    "Person %s would like to say hello",
	}).AddIngredient(task.Ingredient{
		Type:     task.IngredientVar,
		Argument: "Arguments",
	})

	// Create a new action, set its name, and give it an ingredient
	action := task.NewComponent(example.EmailAction{})
	action.SetName("email").AddIngredient(task.Ingredient{
		Type:  task.IngredientVar,
		Value: "strcon",
	})

	task := task.Task{
		Event:   event,
		Actions: []task.Component{},
	}

	task.Run()

	fmt.Println(hub.Events)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(hub.Components)

	fmt.Println("Task completed!")
}
