package valueObject

import (
	"testing"

	testHelpers "github.com/speedianet/os/src/devUtils"
)

func TestDataFieldValue(t *testing.T) {
	t.Run("ValidDataFieldValue", func(t *testing.T) {
		validDataFieldValues := []string{
			"This is my username",
			"new_email@mail.net",
			"localhost:8000",
			"https://www.google.com/search",
		}

		for _, dfv := range validDataFieldValues {
			_, err := NewDataFieldValue(dfv)
			if err != nil {
				t.Errorf("Expected no error for %s, got %s", dfv, err.Error())
			}
		}
	})

	t.Run("InvalidDataFieldValue", func(t *testing.T) {
		invalidLength := 70
		invalidDataFieldValues := []string{
			"",
			"a",
			testHelpers.GenerateString(invalidLength),
		}

		for _, dfv := range invalidDataFieldValues {
			_, err := NewDataFieldValue(dfv)
			if err == nil {
				t.Errorf("Expected error for %s, got nil", dfv)
			}
		}
	})
}