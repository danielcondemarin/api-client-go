/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.1.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type DestinationBody struct {
	// A human-readable name for your data export destination.
	Name string `json:"name"`
	// The data export destination type. Available choices are kinesis, google-pubsub, mparticle, or segment.
	Kind string `json:"kind"`
	// destination-specific configuration.
	Config *interface{} `json:"config"`
	// Whether the data export destination is on or not.
	On bool `json:"on,omitempty"`
}
