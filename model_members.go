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

type Members struct {
	Links *Links `json:"_links,omitempty"`
	Items []Member `json:"items,omitempty"`
	TotalCount float32 `json:"totalCount,omitempty"`
}
