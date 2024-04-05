package valueObject

import (
	"encoding/json"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMarketplaceItemType(t *testing.T) {
	t.Run("ValidMarketplaceItemType", func(t *testing.T) {
		validMarketplaceItemTypes := []string{
			"app",
			"framework",
			"stack",
		}
		for _, mit := range validMarketplaceItemTypes {
			_, err := NewMarketplaceItemType(mit)
			if err != nil {
				t.Errorf("Expected no error for %s, got %s", mit, err.Error())
			}
		}
	})

	t.Run("InvalidMarketplaceItemType", func(t *testing.T) {
		invalidMarketplaceItemTypes := []string{
			"",
			"service",
			"mobile",
			"ml-model",
			"repository",
		}
		for _, mit := range invalidMarketplaceItemTypes {
			_, err := NewMarketplaceItemType(mit)
			if err == nil {
				t.Errorf("Expected error for %s, got nil", mit)
			}
		}
	})

	t.Run("ValidUnmarshalJSON", func(t *testing.T) {
		var testStruct struct {
			DataToTest MarketplaceItemType
		}

		dataToTest := "framework"
		mapToTest := map[string]string{
			"dataToTest": dataToTest,
		}
		mapBytesToTest, _ := json.Marshal(mapToTest)

		reader := strings.NewReader(string(mapBytesToTest))
		jsonDecoder := json.NewDecoder(reader)
		err := jsonDecoder.Decode(&testStruct)
		if err != nil {
			t.Fatalf("Expected no error on UnmarshalJSON valid test, got %s", err.Error())
		}

		dataToTestFromStructStr := testStruct.DataToTest.String()
		if dataToTestFromStructStr != dataToTest {
			t.Errorf(
				"VO data '%s' after UnmarshalJSON is not the same as the original data '%s'",
				dataToTestFromStructStr,
				dataToTest,
			)
		}
	})

	t.Run("InvalidUnmarshalJSON", func(t *testing.T) {
		var testStruct struct {
			DataToTest MarketplaceItemType
		}

		dataToTest := "service"
		mapToTest := map[string]string{
			"dataToTest": dataToTest,
		}
		mapBytesToTest, _ := json.Marshal(mapToTest)

		reader := strings.NewReader(string(mapBytesToTest))
		jsonDecoder := json.NewDecoder(reader)
		err := jsonDecoder.Decode(&testStruct)
		if err == nil {
			t.Fatal("Expected error on UnmarshalJSON invalid test, got nil")
		}
	})

	t.Run("ValidUnmarshalYAML", func(t *testing.T) {
		var testStruct struct {
			DataToTest MarketplaceItemType `yaml:"dataToTest"`
		}

		dataToTest := "framework"
		mapToTest := map[string]string{
			"dataToTest": dataToTest,
		}
		mapBytesToTest, _ := yaml.Marshal(mapToTest)

		reader := strings.NewReader(string(mapBytesToTest))
		yamlDecoder := yaml.NewDecoder(reader)
		err := yamlDecoder.Decode(&testStruct)
		if err != nil {
			t.Fatalf("Expected no error on UnmarshalYAML valid test, got %s", err.Error())
		}

		dataToTestFromStructStr := testStruct.DataToTest.String()
		if dataToTestFromStructStr != dataToTest {
			t.Errorf(
				"VO data '%s' after UnmarshalYAML is not the same as the original data '%s'",
				dataToTestFromStructStr,
				dataToTest,
			)
		}
	})

	t.Run("InvalidUnmarshalYAML", func(t *testing.T) {
		var testStruct struct {
			DataToTest MarketplaceItemType `yaml:"dataToTest"`
		}

		dataToTest := "service"
		mapToTest := map[string]string{
			"dataToTest": dataToTest,
		}
		mapBytesToTest, _ := yaml.Marshal(mapToTest)

		reader := strings.NewReader(string(mapBytesToTest))
		yamlDecoder := yaml.NewDecoder(reader)
		err := yamlDecoder.Decode(&testStruct)
		if err == nil {
			t.Fatal("Expected error on UnmarshalYAML invalid test, got nil")
		}
	})
}