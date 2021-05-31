package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-tutorial/basic_validator/model"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	validateStruct()
	validateVariable()
}

func validateStruct() {

	address := &model.Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &model.User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*model.Address{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("Invalid Error", err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println("Namespace : ", err.Namespace())
			fmt.Println("Field : ",err.Field())
			fmt.Println("StructNamespace : ",err.StructNamespace())
			fmt.Println("StructField : ",err.StructField())
			fmt.Println("Tag : ",err.Tag())
			fmt.Println("ActualTag : ",err.ActualTag())
			fmt.Println("Kind : ",err.Kind())
			fmt.Println("Type : ",err.Type())
			fmt.Println("Value : ",err.Value())
			fmt.Println("Param : ",err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}