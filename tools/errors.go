package tools

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/metaphi-labs/latent-contracts/errors"
)

// ValidationErrorToServiceError converts validation errors to user-friendly ServiceErrors
func ValidationErrorToServiceError(err error, toolName string, paramName string) *errors.ServiceError {
	if err == nil {
		return nil
	}

	// Handle validator.ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			field := fieldErr.Field()
			tag := fieldErr.Tag()
			param := fieldErr.Param()

			// Map validator tags to error codes and messages
			switch tag {
			case "required":
				return errors.NewServiceError(
					errors.VAL_MISSING_PARAMETER,
					fmt.Sprintf("Required parameter '%s' is missing for tool '%s'",
						toSnakeCase(field), toolName),
					"tool-validation",
					false,
				)

			case "min":
				if fieldErr.Kind().String() == "string" {
					return errors.NewServiceError(
						errors.VAL_STRING_TOO_SHORT,
						fmt.Sprintf("Parameter '%s' must be at least %s characters long",
							toSnakeCase(field), param),
						"tool-validation",
						false,
					)
				}
				return errors.NewServiceError(
					errors.VAL_OUT_OF_RANGE,
					fmt.Sprintf("Parameter '%s' must be at least %s",
						toSnakeCase(field), param),
					"tool-validation",
					false,
				)

			case "max":
				if fieldErr.Kind().String() == "string" {
					return errors.NewServiceError(
						errors.VAL_STRING_TOO_LONG,
						fmt.Sprintf("Parameter '%s' must be at most %s characters long",
							toSnakeCase(field), param),
						"tool-validation",
						false,
					)
				}
				if fieldErr.Kind().String() == "slice" {
					return errors.NewServiceError(
						errors.VAL_ARRAY_TOO_LONG,
						fmt.Sprintf("Parameter '%s' must have at most %s items",
							toSnakeCase(field), param),
						"tool-validation",
						false,
					)
				}
				return errors.NewServiceError(
					errors.VAL_OUT_OF_RANGE,
					fmt.Sprintf("Parameter '%s' must be at most %s",
						toSnakeCase(field), param),
					"tool-validation",
					false,
				)

			case "oneof":
				return errors.NewServiceError(
					errors.VAL_INVALID_ENUM,
					fmt.Sprintf("Parameter '%s' must be one of: %s",
						toSnakeCase(field), param),
					"tool-validation",
					false,
				)

			case "eq":
				return errors.NewServiceError(
					errors.VAL_INVALID_PARAMETER,
					fmt.Sprintf("Parameter '%s' must be exactly %s",
						toSnakeCase(field), param),
					"tool-validation",
					false,
				)

			case "url":
				return errors.NewServiceError(
					errors.VAL_INVALID_URL,
					fmt.Sprintf("Parameter '%s' must be a valid URL",
						toSnakeCase(field)),
					"tool-validation",
					false,
				)

			case "required_without":
				return errors.NewServiceError(
					errors.VAL_DEPENDENCY_MISSING,
					fmt.Sprintf("Parameter '%s' is required when '%s' is not provided",
						toSnakeCase(field), toSnakeCase(param)),
					"tool-validation",
					false,
				)

			default:
				return errors.NewServiceError(
					errors.VAL_INVALID_PARAMETER,
					fmt.Sprintf("Parameter '%s' is invalid: %s validation failed",
						toSnakeCase(field), tag),
					"tool-validation",
					false,
				)
			}
		}
	}

	// Handle custom validation errors
	if strings.Contains(err.Error(), "either") && strings.Contains(err.Error(), "required") {
		return errors.NewServiceError(
			errors.VAL_MUTUALLY_EXCLUSIVE,
			err.Error(),
			"tool-validation",
			false,
		)
	}

	// Handle JSON unmarshal errors
	if strings.Contains(err.Error(), "cannot unmarshal") {
		return errors.NewServiceError(
			errors.VAL_INVALID_FORMAT,
			fmt.Sprintf("Invalid parameter format for tool '%s': %s", toolName, err.Error()),
			"tool-validation",
			false,
		)
	}

	// Default case
	return errors.NewServiceError(
		errors.VAL_INVALID_REQUEST,
		fmt.Sprintf("Invalid parameters for tool '%s': %s", toolName, err.Error()),
		"tool-validation",
		false,
	)
}

// toSnakeCase converts PascalCase to snake_case
func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, rune(strings.ToLower(string(r))[0]))
	}
	return string(result)
}

// ValidateAndGetError is a convenience function that validates and returns a ServiceError if invalid
func ValidateAndGetError(toolName string, params map[string]interface{}) *errors.ServiceError {
	err := ParseAndValidateParams(toolName, params)
	if err != nil {
		return ValidationErrorToServiceError(err, toolName, "")
	}
	return nil
}