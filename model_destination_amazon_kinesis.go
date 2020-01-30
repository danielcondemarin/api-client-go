/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.23
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type DestinationAmazonKinesis struct {
	Region string `json:"region,omitempty"`
	RoleArn string `json:"roleArn,omitempty"`
	StreamName string `json:"streamName,omitempty"`
}
