package request

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	fr_translations "github.com/go-playground/validator/v10/translations/fr"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type T struct {
	Description []string `json:"description"`
	Title       []string `json:"title"`
}

type Request struct {
	body   any
	errors map[string][]string
}

type ValidationErrors map[string][]string

func ValidateStruct(acceptLanguage string) validator.ValidationErrors {
	validate = validator.New()
	if acceptLanguage == "fr" {
		locale := fr.New()
		uni = ut.New(locale, locale)
		trans, _ := uni.GetTranslator("fr")
		fr_translations.RegisterDefaultTranslations(validate, trans)
		// return translateIndividual(trans.go)
		return translateAll(trans)
		// translateOverride(trans.go) // yep you can specify your own in whatever locale you want!
	} else {
		locale := en.New()
		uni = ut.New(locale, locale)
		trans, _ := uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(validate, trans)
		// return translateIndividual(trans.go)
		return translateAll(trans)
		// translateOverride(trans.go) // yep you can specify your own in whatever locale you want!
	}
}

func translateIndividual(trans ut.Translator) validator.ValidationErrors {

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
		return errs
	}
	return nil
}

func translateAll(trans ut.Translator) validator.ValidationErrors {

	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10"`
		Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "Joeybloggs",
		Tagline:  "This tagline is way too long.",
		Tagline2: "1",
	}

	err := validate.Struct(user)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))
		return errs
	}
	return nil
}

func translateOverride(trans ut.Translator) {

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}
