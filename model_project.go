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

type Project struct {
	Links *Links `json:"_links,omitempty"`
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	Key string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
	IncludeInSnippetByDefault bool `json:"includeInSnippetByDefault,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
	// An array of tags for this project.
	Tags []string `json:"tags,omitempty"`
	DefaultClientSideAvailability *ClientSideAvailability `json:"defaultClientSideAvailability,omitempty"`
}
