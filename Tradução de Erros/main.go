package main

import (
	"fmt"

	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	brTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
)

type person struct {
	Name                string `validate:"required,min=4,max=15"`
	Email               string `validate:"required,email"`
	Age                 int    `validate:"required,numeric,min=18"`
	DriverLicenseNumber string `validate:"omitempty,len=12,numeric"`
}

func main() {
	validate := validator.New()

	portuguese := pt.New()
	uni := ut.New(portuguese, portuguese)
	trans, _ := uni.GetTranslator("pt-br")
	_ = brTranslations.RegisterDefaultTranslations(validate, trans)

	// Validação de um campo apenas
	name := "re"
	err := validate.Var(name, "required,min=3")
	errs := translateError(err, trans)
	fmt.Println(errs)

	p := person{
		Name:                "Joe",
		Email:               "dummyemail",
		Age:                 0,
		DriverLicenseNumber: "",
	}
	err = validate.Struct(p)
	errs = translateError(err, trans)
	fmt.Println(errs)
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
