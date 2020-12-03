/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.9.2
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type CustomRoleBody struct {
	// Name of the custom role.
	Name string `json:"name"`
	// Description of the custom role.
	Description string `json:"description,omitempty"`
	// The 20-hexdigit id or the key for a custom role.
	Key string `json:"key"`
	Policy []Policy `json:"policy"`
}
