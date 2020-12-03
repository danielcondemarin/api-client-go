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

type RelayProxyConfig struct {
	// The unique resource id.
	Id string `json:"_id"`
	Creator *Member `json:"_creator"`
	// A human-friendly name for the relay proxy configuration
	Name string `json:"name"`
	Policy []Policy `json:"policy"`
	// Full secret key. Only included if creating or resetting the relay proxy configuration
	FullKey string `json:"fullKey,omitempty"`
	// The last 4 digits of the unique secret key for this relay proxy configuration
	DisplayKey string `json:"displayKey"`
	// A unix epoch time in milliseconds specifying the creation time of this relay proxy configuration
	CreationDate int64 `json:"creationDate"`
	// A unix epoch time in milliseconds specifying the last time this relay proxy configuration was modified
	LastModified int64 `json:"lastModified"`
}
