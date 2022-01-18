package commonids

import "fmt"

func ValidateManagementGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagementGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}
