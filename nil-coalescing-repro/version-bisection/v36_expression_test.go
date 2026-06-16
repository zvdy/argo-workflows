package template

import (
	"bytes"
	"testing"
)

// Drop into util/template/ of a v3.6.x checkout (module .../v3) and run:
//   GOFLAGS=-mod=mod go test ./util/template/ -run Test_v36_NilCoalescing_StrictPath -v
// Proves v3.6's strict expression path (allowUnresolved=false, as withItems uses)
// resolves a guarded missing key to the fallback — no "is missing" error.
func Test_v36_NilCoalescing_StrictPath(t *testing.T) {
	cases := []struct {
		name, expr string
		env        map[string]interface{}
		want       string
	}{
		{"present", `item.optionalKey ?? 'fallback'`,
			map[string]interface{}{"item": map[string]interface{}{"name": "a", "optionalKey": "value"}}, "value"},
		{"missing-coalesce", `item.optionalKey ?? 'fallback'`,
			map[string]interface{}{"item": map[string]interface{}{"name": "b"}}, "fallback"},
		{"missing-optchain", `item?.optionalKey ?? 'fallback'`,
			map[string]interface{}{"item": map[string]interface{}{"name": "b"}}, "fallback"},
		{"missing-bracket", `item['optionalKey'] ?? 'fallback'`,
			map[string]interface{}{"item": map[string]interface{}{"name": "b"}}, "fallback"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var buf bytes.Buffer
			n, err := expressionReplace(&buf, c.expr, c.env, false) // false == strict
			if err != nil {
				t.Fatalf("v3.6 strict eval FAILED: %v", err)
			}
			if buf.String() != c.want {
				t.Fatalf("got %q (%d bytes), want %q", buf.String(), n, c.want)
			}
		})
	}
}
