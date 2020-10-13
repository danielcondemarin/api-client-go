/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.7.1
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type UserSettingsBody struct {
	// The variation value to set for the user. Must match the variation type of the flag. 
	Setting bool `json:"setting,omitempty"`
}
