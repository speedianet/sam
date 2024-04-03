package valueObject

import (
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

type MktplaceItemInstallStep string

func NewMktplaceItemInstallStep(value string) (MktplaceItemInstallStep, error) {
	miis := MktplaceItemInstallStep(value)
	if !miis.isValid() {
		return "", errors.New("InvalidMktItemInstallStep")
	}

	return miis, nil
}

func NewMktplaceItemInstallStepPanic(value string) MktplaceItemInstallStep {
	miis, err := NewMktplaceItemInstallStep(value)
	if err != nil {
		panic(err)
	}

	return miis
}

func (miis MktplaceItemInstallStep) isValid() bool {
	isTooShort := len(string(miis)) < 1
	isTooLong := len(string(miis)) > 512
	return !isTooShort && !isTooLong
}

func (miis MktplaceItemInstallStep) String() string {
	return string(miis)
}

func (miisPtr *MktplaceItemInstallStep) UnmarshalJSON(value []byte) error {
	valueStr := string(value)
	unquotedValue := strings.Trim(valueStr, "\"")
	valueWithoutBackslash := strings.ReplaceAll(unquotedValue, "\\", "")

	miis, err := NewMktplaceItemInstallStep(valueWithoutBackslash)
	if err != nil {
		return err
	}

	*miisPtr = miis
	return nil
}

func (miisPtr *MktplaceItemInstallStep) UnmarshalYAML(value *yaml.Node) error {
	var valueStr string
	err := value.Decode(&valueStr)
	if err != nil {
		return err
	}

	miis, err := NewMktplaceItemInstallStep(valueStr)
	if err != nil {
		return err
	}

	*miisPtr = miis
	return nil
}