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

type FlagListItem struct {
	Name string `json:"name,omitempty"`
	Key string `json:"key,omitempty"`
	Links *Links `json:"_links,omitempty"`
	Site *Site `json:"_site,omitempty"`
}
