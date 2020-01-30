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

type UserSegmentRule struct {
	Clauses []Clause `json:"clauses,omitempty"`
	Weight int32 `json:"weight,omitempty"`
	BucketBy string `json:"bucketBy,omitempty"`
}
