package validator

import (
	"reflect"
	"strings"

	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/uk"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	uk_translations "github.com/go-playground/validator/v10/translations/uk"
)

var uni *ut.UniversalTranslator

func New() *validator.Validate {
	en := en.New()
	uk := uk.New()
	uni = ut.New(en, uk)

	enTrans, _ := uni.GetTranslator("en")
	ukTrans, _ := uni.GetTranslator("uk")

	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	en_translations.RegisterDefaultTranslations(v, enTrans)
	uk_translations.RegisterDefaultTranslations(v, ukTrans)

	return v
}

func Parse(err error, trans ...ut.Translator) map[string]string {
	var t ut.Translator

	if len(trans) == 0 {
		t, _ = uni.GetTranslator("en")
	} else {
		t = trans[0]
	}

	errs := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		splited := strings.SplitN(err.Namespace(), ".", 2)
		var fullPath string
		switch len(splited) {
		case 1:
			fullPath = splited[0]
		case 2:
			fullPath = splited[1]
		default:
			fullPath = "unknownProperty"
		}
		errs[fullPath] = err.Translate(t)
	}

	return errs
}

type StructValidator struct {
	validator *validator.Validate
}

func NewStructValidator(v *validator.Validate) *StructValidator {
	return &StructValidator{
		validator: v,
	}
}

func (v *StructValidator) Validate(value any) error {
  t, _ := uni.GetTranslator("en")

	err := v.validator.Struct(value)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		firstErr := errs[0]
		msg := firstErr.Field() + ": " + firstErr.Translate(t)
		return apperror.New(apperror.CodeInvalidInput, msg, err)
	}

	return nil
}

