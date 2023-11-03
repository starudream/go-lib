package strutil

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestToSnakeCase(t *testing.T) {
	testutil.Equal(t, "test_to_snake_case", ToSnakeCase("TestToSnakeCase"))
}

func TestToCamelCase(t *testing.T) {
	testutil.Equal(t, "testToCamelCase", ToCamelCase("test_to_camel_case"))
}

func TestToPascalCase(t *testing.T) {
	testutil.Equal(t, "TestToPascalCase", ToPascalCase("test_to_pascal_case"))
}
