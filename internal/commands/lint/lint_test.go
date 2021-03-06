package lint

import (
	"fmt"
	"github.com/raitonbl/ant/internal"
	"github.com/raitonbl/ant/internal/commands/lint/lint_message"
	"testing"
)

func TestLint_from_json(t *testing.T) {
	doLintTest(t, "index-003.json")
}

func TestLint_from_json_where_schema_has_refers_to(t *testing.T) {
	doLintTest(t, "index-053.json")
}

func TestLint_from_yaml_where_multiple_args(t *testing.T) {
	doLintTest(t, "index-058.yaml")
}

func TestLint_from_yaml_where_multiple_args_with_same_index(t *testing.T) {
	doLintTest(t, "index-059.yaml",Violation{Path: "/commands/1/commands/1/parameters", Message: lint_message.ARGS_INDEX_NOT_UNIQUE})
}

func TestLint_from_json_where_array_schema_has_refers_to(t *testing.T) {
	doLintTest(t, "index-055.json")
}

func TestLint_from_json_where_array_schema_has_refers_to_and_is_refers_to(t *testing.T) {
	doLintTest(t, "index-056.json")
}

func TestLint_from_json_where_refers_to_is_circular_dependency(t *testing.T) {
	doLintTest(t, "index-057.json", Violation{Path: "/schemas/2/refers-to", Message: lint_message.UNRESOLVABLE_FIELD}, Violation{Path: "/schemas/1/refers-to", Message: lint_message.UNRESOLVABLE_FIELD})
}

func TestLint_from_json_where_schema_has_refers_to_and_refers_to_is_unresolvable(t *testing.T) {
	doLintTest(t, "index-054.json", Violation{Path: "/parameters/1/schema/refers-to", Message: lint_message.UNRESOLVABLE_FIELD})
}

func TestLint_from_yaml(t *testing.T) {
	doLintTest(t, "index-003.yaml")
}

func TestLint_where_name_is_missing(t *testing.T) {
	doLintTest(t, "index-005.json", Violation{Path: "/parameters/0/name", Message: lint_message.REQUIRED_FIELD})
}

func TestLint_where_in_flags_and_index_is_defined(t *testing.T) {
	doLintTest(t, "index-002.json", Violation{Path: "/parameters/1/index", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_in_arguments_and_shortForm_is_defined(t *testing.T) {
	doLintTest(t, "index-001.json", Violation{Path: "/parameters/0/short-form", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_schema_is_missing(t *testing.T) {
	object := Violation{Path: "/parameters/0/schema", Message: lint_message.REQUIRED_FIELD}
	doLintTest(t, "index-008.json", object)
}

func TestLint_where_index_Is_Negative(t *testing.T) {
	doLintTest(t, "index-009.json", Violation{Path: "/parameters/0/index",
		Message: lint_message.FIELD_INDEX_GT_ZERO},
		Violation{Path: "/commands/0/parameters", Message: lint_message.ARGS_INDEX_NOT_ORDERED},
		Violation{Path: "/commands/1/commands/0/parameters", Message: lint_message.ARGS_INDEX_NOT_ORDERED},
		Violation{Path: "/commands/1/commands/1/parameters", Message: lint_message.ARGS_INDEX_NOT_ORDERED})
}

func TestLint_where_description_is_missing(t *testing.T) {
	doLintTest(t, "index-006.json", Violation{Path: "/parameters/1/description", Message: lint_message.REQUIRED_FIELD})
}

func TestLint_where_index_is_missing_and_argument_Is_Null(t *testing.T) {
	doLintTest(t, "index-004.json", Violation{Path: "/parameters/0/index", Message: lint_message.FIELD_WHEN_IN_ARGUMENTS})
}

func TestLint_where_type_Is_string_and_format_is_int64(t *testing.T) {
	doLintTest(t, "index-012.json", Violation{Path: "/parameters/0/schema/format", Message: lint_message.FIELD_FORMAT_NOT_ALLOWED_IN_TYPE_STRING})
}

func TestLint_where_type_Is_number_and_format_is_date(t *testing.T) {
	doLintTest(t, "index-013.json", Violation{Path: "/parameters/0/schema/format", Message: lint_message.FIELD_FORMAT_IS_ONLY_ALLOWED_IN_TYPE_STRING})
}

func TestLint_where_type_Is_number_and_format_is_datetime(t *testing.T) {
	doLintTest(t, "index-014.json", Violation{Path: "/parameters/0/schema/format", Message: lint_message.FIELD_FORMAT_IS_ONLY_ALLOWED_IN_TYPE_STRING})
}

func TestLint_where_type_enum_and_example_not_part_of_enum(t *testing.T) {
	doLintTest(t, "index-015.json", Violation{Path: "/parameters/1/schema/examples/0", Message: lint_message.FIELD_EXAMPLE_MUST_BE_PART_OF_ENUM})
}

func TestLint_where_type_number_and_max_length_is_two(t *testing.T) {
	doLintTest(t, "index-016.json", Violation{Path: "/parameters/2/schema/max-length", Message: lint_message.FIELD_FORMAT_IS_ONLY_ALLOWED_IN_TYPE_STRING})
}

func TestLint_where_type_number_and_min_length_is_two(t *testing.T) {
	doLintTest(t, "index-017.json", Violation{Path: "/parameters/2/schema/min-length", Message: lint_message.FIELD_FORMAT_IS_ONLY_ALLOWED_IN_TYPE_STRING})
}

func TestLint_where_type_string_and_min_length_is_gt_max_length(t *testing.T) {
	doLintTest(t, "index-018.json", Violation{Path: "/parameters/0/schema/min-length", Message: lint_message.FIELD_MIN_LENGTH_MUST_NOT_BE_GT_MAX_LENGTH})
}

func TestLint_where_type_string_and_min_length_lt_zero(t *testing.T) {
	doLintTest(t, "index-020.json", Violation{Path: "/parameters/0/schema/min-length", Message: lint_message.FIELD_MIN_LENGTH_GT_ZERO})
}

func TestLint_where_type_string_and_max_length_lt_zero(t *testing.T) {
	doLintTest(t, "index-021.json", Violation{Path: "/parameters/0/schema/max-length", Message: lint_message.FIELD_MAX_LENGTH_GT_ZERO})
}

func TestLint_where_type_string_and_multiple_of(t *testing.T) {
	doLintTest(t, "index-022.json", Violation{Path: "/parameters/0/schema/multiple-of", Message: lint_message.FIELD_FORMAT_IS_ONLY_ALLOWED_IN_TYPE_NUMBER})
}

func TestLint_where_type_string_and_maximum(t *testing.T) {
	doLintTest(t, "index-023.json", Violation{Path: "/parameters/0/schema/maximum", Message: lint_message.FIELD_FORMAT_IS_ONLY_ALLOWED_IN_TYPE_NUMBER})
}

func TestLint_where_type_string_and_minimum(t *testing.T) {
	doLintTest(t, "index-024.json", Violation{Path: "/parameters/0/schema/minimum", Message: lint_message.FIELD_FORMAT_IS_ONLY_ALLOWED_IN_TYPE_NUMBER})
}

func TestLint_where_type_number_and_minimum_gt_maximum(t *testing.T) {
	doLintTest(t, "index-025.json", Violation{Path: "/parameters/2/schema/minimum", Message: lint_message.FIELD_MIN_MUST_NOT_BE_GT_MAX})
}

func TestLint_where_type_number_and_max_items(t *testing.T) {
	doLintTest(t, "index-030.json", Violation{Path: "/parameters/0/schema/max-items", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_type_number_and_min_items(t *testing.T) {
	doLintTest(t, "index-031.json", Violation{Path: "/parameters/0/schema/min-items", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_type_array_and_array_schema_is_undefined(t *testing.T) {
	doLintTest(t, "index-033.json", Violation{Path: "/parameters/2/schema/items", Message: lint_message.REQUIRED_FIELD})
}

func TestLint_where_type_array_and_format_not_defined(t *testing.T) {
	doLintTest(t, "index-034.json", Violation{Path: "/parameters/2/schema/format", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_type_array_and_min_items_gt_max_items(t *testing.T) {
	doLintTest(t, "index-035.json", Violation{Path: "/parameters/2/schema/min-items", Message: lint_message.FIELD_MIN_ITEMS_MUST_NOT_BE_GT_MAX_ITEMS})
}

func TestLint_where_type_array_and_min_items_lt_zero(t *testing.T) {
	doLintTest(t, "index-037.json", Violation{Path: "/parameters/2/schema/min-items", Message: lint_message.FIELD_MIN_ITEMS_GT_ZERO})
}

func TestLint_where_type_array_and_max_items_lt_zero(t *testing.T) {
	doLintTest(t, "index-036.json", Violation{Path: "/parameters/2/schema/max-items", Message: lint_message.FIELD_MAX_ITEMS_GT_ZERO})
}

func TestLint_where_type_array_and_array_type_array(t *testing.T) {
	doLintTest(t, "index-039.json", Violation{Path: "/parameters/2/schema/items/type", Message: lint_message.ARRAY_FIELD_TYPE_NOT_ALLOWED})
}

func TestLintCommand_where_command_exit_code_is_missing(t *testing.T) {
	doLintTest(t, "index-046.json", Violation{Path: "/commands/0/exit/0/code", Message: lint_message.REQUIRED_FIELD})
}

func TestLintCommand_where_command_exit_has_id(t *testing.T) {
	doLintTest(t, "index-047.json", Violation{Path: "/commands/0/exit/0", Message: "did not match any of the specified OneOf schemas"})
}

func TestLintCommand_where_command_exit_message_is_missing(t *testing.T) {
	doLintTest(t, "index-048.json", Violation{Path: "/commands/0/exit/0/message", Message: lint_message.REQUIRED_FIELD})
}

func TestLintCommand_where_command_exit_refers_to_is_unresolvable(t *testing.T) {
	doLintTest(t, "index-049.json", Violation{Path: "/commands/0/exit/1/refers-to", Message: lint_message.UNRESOLVABLE_FIELD})
}

func TestLintCommand_where_command_parameter_refers_to_is_unresolvable(t *testing.T) {
	doLintTest(t, "index-050.json", Violation{Path: "/commands/0/parameters/0/refers-to", Message: lint_message.UNRESOLVABLE_FIELD})
}

func TestLintCommand_where_command_parameter_in_is_null_and_index_is_defined(t *testing.T) {
	doLintTest(t, "index-051.json", Violation{Path: "/commands/0/parameters/0/refers-to", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLintCommand_where_command_parameter_in_flags_and_index_is_defined(t *testing.T) {
	doLintTest(t, "index-052.json", Violation{Path: "/commands/0/parameters/0/refers-to", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLintExit_where_id_is_missing(t *testing.T) {
	doLintTest(t, "index-040.json", Violation{Path: "/exit/0", Message: "\"id\" value is required"})
}

func TestLintExit_where_id_is_blank(t *testing.T) {
	doLintTest(t, "index-042.json", Violation{Path: "/exit/0/id", Message: lint_message.REQUIRED_FIELD})
}

func TestLintExit_where_code_is_missing(t *testing.T) {
	doLintTest(t, "index-041.json", Violation{Path: "/exit/0", Message: "\"code\" value is required"})
}

func TestLintExit_where_message_is_missing(t *testing.T) {
	doLintTest(t, "index-043.json", Violation{Path: "/exit/0", Message: "\"message\" value is required"})
}

func TestLintExit_where_message_is_blank(t *testing.T) {
	doLintTest(t, "index-044.json", Violation{Path: "/exit/0/message", Message: lint_message.REQUIRED_FIELD})
}

func TestLintExit_where_refers_to_is_defined(t *testing.T) {
	doLintTest(t, "index-045.json", Violation{Path: "/exit/0", Message: "additional properties are not allowed"})
}

func TestLint_where_type_number_and_pattern_defined(t *testing.T) {
	doLintTest(t, "index-019.json", Violation{Path: "/parameters/2/schema/pattern", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_type_string_and_exclusive_minimum_and_minimum_missing(t *testing.T) {
	doLintTest(t, "index-026.json", Violation{Path: "/parameters/0/schema/exclusive-minimum", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_type_string_and_exclusive_maximum_and_maximum_missing(t *testing.T) {
	doLintTest(t, "index-027.json", Violation{Path: "/parameters/0/schema/exclusive-maximum", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_type_number_and_exclusive_minimum_and_minimum_missing(t *testing.T) {
	doLintTest(t, "index-028.json", Violation{Path: "/parameters/2/schema/minimum", Message: lint_message.REQUIRED_FIELD})
}

func TestLint_where_type_number_and_exclusive_maximum_and_maximum_missing(t *testing.T) {
	doLintTest(t, "index-029.json", Violation{Path: "/parameters/2/schema/maximum", Message: lint_message.REQUIRED_FIELD})
}

func TestLint_where_type_number_and_unique_items(t *testing.T) {
	doLintTest(t, "index-032.json", Violation{Path: "/parameters/0/schema/unique-items", Message: lint_message.FIELD_NOT_ALLOWED})
}

func TestLint_where_type_array_and_array_type_undefined(t *testing.T) {
	doLintTest(t, "index-038.json", Violation{Path: "/parameters/2/schema/items/type", Message: lint_message.REQUIRED_FIELD})
}

func doLintTest(t *testing.T, filename string, seq ...Violation) {
	doLintFrom(t, filename, func(array []Violation) {

		if (array == nil || len(array) == 0) && (seq != nil && len(seq) > 0) {
			t.Fatal(fmt.Sprintf("\nExpected:%s\nActual:[nil]", toText(seq)))
		}

		if len(array) > len(seq) {
			t.Fatal(fmt.Sprintf("\nExpected:%s\nActual:%s", toText(seq), toText(array)))
		}

		if (seq == nil || len(seq) == 0) && len(array) == 0 {
			return
		}

		if seq == nil || len(seq) == 0 {
			t.Fatal(fmt.Sprintf("\nExpected:[]\nActual:%s", toText(array)))
		}

		for index, singleValue := range array {
			if singleValue.Path != seq[index].Path || singleValue.Message != seq[index].Message {
				t.Fatal(fmt.Sprintf("\nExpected:%s\nActual:%s", toText(seq), toText(array)))
			}
		}

	})
}

func doLintFrom(t *testing.T, filename string, afterLint func(array []Violation)) {

	ctx, err := internal.GetContext(fmt.Sprintf("testdata/%s", filename))

	if err != nil {
		t.Fatal(err)
	}

	array, err := Lint(ctx)

	if err != nil {
		t.Fatal(err)
	}

	afterLint(array)
}
