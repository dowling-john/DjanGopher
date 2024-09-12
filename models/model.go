package models

import "time"

type (
	// Model
	// The Model struct should be used as a composition element in the all defined models, this will give the ability
	// to save a changed model to the database.
	//
	// Example:
	// type ExampleModel struct {
	//     model.Model
	//	   ExampleField  		string		`model:"example_field"`  <- use the "model" struct tag that matches the column from the database
	// }
	Model struct {
		ID        uint      `model:"id"`
		CreatedAt time.Time `model:"createdAt"`
		UpdatedAt time.Time `model:"updatedAt"`
	}
)
