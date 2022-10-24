package helper

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

var v = validator.New()
var english = en.New()
var uni = ut.New(english)
var trans, _ = uni.GetTranslator("en")

func ParseAndValidateBody[T any](c *fiber.Ctx) (*T, error) {
	var reqBody T
	if err := c.BodyParser(&reqBody); err != nil {
		return nil, err
	}

	if err := Translate(v, v.Struct(reqBody)); err != nil {
		for _, e := range err {
			return nil, e
		}
	}
	return &reqBody, nil
}

func ParseAndValidateQuery[T any](c *fiber.Ctx) (*T, error) {
	var reqQuery T
	if err := c.QueryParser(&reqQuery); err != nil {
		return nil, err
	}

	if err := Translate(v, v.Struct(reqQuery)); err != nil {
		for _, e := range err {
			return nil, e
		}
	}
	return &reqQuery, nil
}

func Translate(v *validator.Validate, err error) (errs []error) {
	_ = enTranslations.RegisterDefaultTranslations(v, trans)
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
