/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.26
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type UsageLinks struct {
	Parent *Link `json:"parent,omitempty"`
	Self *Link `json:"self,omitempty"`
	// The following links that are in the response.
	Subseries []Link `json:"subseries,omitempty"`
}
