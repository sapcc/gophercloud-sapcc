// copied from upstream to bypass internal package
package tools

import (
	"crypto/rand"
	"encoding/json"
	"testing"
)

// copied from https://github.com/gophercloud/gophercloud/blob/v1.7.0/github.com/sapcc/gophercloud-sapcc/internal/acceptance/tools/tools.go#L48-L60
// RandomString generates a string of given length, but random content.
// All content will be within the ASCII graphic character set.
// (Implementation from Even Shaw's contribution on
// http://stackoverflow.com/questions/12771930/what-is-the-fastest-way-to-generate-a-long-random-string-in-go).
func RandomString(prefix string, n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return prefix + string(bytes)
}

// copied from https://github.com/gophercloud/gophercloud/blob/v1.7.0/github.com/sapcc/gophercloud-sapcc/internal/acceptance/tools/tools.go#L77-L81
// PrintResource returns a resource as a readable structure
func PrintResource(t *testing.T, resource interface{}) {
	b, _ := json.MarshalIndent(resource, "", "  ")
	t.Logf(string(b))
}
