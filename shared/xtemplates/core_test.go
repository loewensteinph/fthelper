package xtemplates_test

import (
	"testing"

	"github.com/loewensteinph/fthelper/shared/maps"
	"github.com/loewensteinph/fthelper/shared/xtemplates"
	"github.com/loewensteinph/fthelper/shared/xtests"
)

func TestXtemplate(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("invalid template").
		WithExpected("function \"invalid\" not defined").
		WithActualAndError(xtemplates.Text("{{ invalid \"function\" }}", maps.New())).
		MustContainError()
}
