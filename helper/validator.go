package helper

import (
	// "regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

// NewValidator membuat instance Validator dengan registrasi validator kustom
func NewValidator() *Validator {
	v := validator.New()

	// Daftarkan validasi regex kustom
	// v.RegisterValidation("regex", func(fl validator.FieldLevel) bool {
	// 	pattern := fl.Param() // Ambil pattern regex dari parameter
	// 	value := fl.Field().String()

	// 	matched, err := regexp.MatchString(pattern, value)
	// 	if err != nil {
	// 		return false
	// 	}
	// 	return matched
	// })

	return &Validator{validate: v}
}

// ValidateStruct melakukan validasi terhadap struct yang diberikan
func (v *Validator) ValidateStruct(data interface{}) error {
	return v.validate.Struct(data) // Validasi menggunakan rules yang telah didaftarkan
}

// FormatValidationError mengubah error validasi menjadi pesan yang lebih ramah
func FormatValidationError(err error) string {
	errorMessages := map[string]string{
		"CategoryID_required": "Category ID is required",
		"CategoryID_gt":       "Category ID must be greater than 0",
		"Name_required":       "Name is required",
		"Name_min":            "Name must have at least 3 characters",
		"Name_regex":          "Name must only contain letters and spaces",
		"Images_url":          "Each image must be a valid URL",
		"Stock_gt":            "Stock must be greater than 0",
		"Price_gt":            "Price must be greater than 0",
		"Quantity_gt":         "Quantity must be greater than 0",
		"ProductID_gt":        "ProductID must be greater than 0",
		"Type_oneof":          "Type must be 'in' or 'out'",
	}

	var errMessages []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {

			key := ve.Field() + "_" + ve.Tag()

			if message, found := errorMessages[key]; found {
				errMessages = append(errMessages, message)
			} else {

				errMessages = append(errMessages, ve.Field()+" is invalid: "+ve.Tag())
			}
		}
	}

	return strings.Join(errMessages, ", ")
}
