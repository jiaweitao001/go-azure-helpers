package commonids

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = UserAssignedIdentityId{}

// UserAssignedIdentityId is a struct representing the Resource ID for a User Assigned Identity
type UserAssignedIdentityId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
}

// NewUserAssignedIdentityID returns a new UserAssignedIdentityId struct
func NewUserAssignedIdentityID(subscriptionId string, resourceGroupName string, resourceName string) UserAssignedIdentityId {
	return UserAssignedIdentityId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourceName:      resourceName,
	}
}

// ParseUserAssignedIdentityID parses 'input' into a UserAssignedIdentityId
func ParseUserAssignedIdentityID(input string) (*UserAssignedIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAssignedIdentityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAssignedIdentityId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseUserAssignedIdentityIDInsensitively parses 'input' case-insensitively into a UserAssignedIdentityId
// note: this method should only be used for API response data and not user input
func ParseUserAssignedIdentityIDInsensitively(input string) (*UserAssignedIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAssignedIdentityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAssignedIdentityId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateUserAssignedIdentityID checks that 'input' can be parsed as a User Assigned Identity ID
func ValidateUserAssignedIdentityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAssignedIdentityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Assigned Identity ID
func (id UserAssignedIdentityId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ManagedIdentity/userAssignedIdentities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this User Assigned Identity ID
func (id UserAssignedIdentityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagedIdentity", "Microsoft.ManagedIdentity", "Microsoft.ManagedIdentity"),
		resourceids.StaticSegment("staticUserAssignedIdentities", "userAssignedIdentities", "userAssignedIdentities"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
	}
}

// String returns a human-readable description of this User Assigned Identities ID
func (id UserAssignedIdentityId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
	}
	return fmt.Sprintf("User Assigned Identity (%s)", strings.Join(components, "\n"))
}
