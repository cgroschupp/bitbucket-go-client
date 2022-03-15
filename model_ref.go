/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bitbucket

// A ref object, representing a branch or tag in a repository.
type Ref struct {
	Type_ string `json:"type"`
	Links *RefLinks `json:"links,omitempty"`
	// The name of the ref.
	Name string `json:"name,omitempty"`
	Target *Commit `json:"target,omitempty"`
}
