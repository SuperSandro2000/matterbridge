// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AdminConsent undocumented
type AdminConsent struct {
	// Object is the base model of AdminConsent
	Object
	// ShareAPNSData The admin consent state of sharing user and device data to Apple.
	ShareAPNSData *AdminConsentState `json:"shareAPNSData,omitempty"`
	// ShareUserExperienceAnalyticsData Gets or sets the admin consent for sharing User experience analytics data.
	ShareUserExperienceAnalyticsData *AdminConsentState `json:"shareUserExperienceAnalyticsData,omitempty"`
}
