// Package deps exists to force go mod tooling to recognize these as direct
// imports, so that gazelle also sees them as direct imports
package deps

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/longrunning"
	_ "google.golang.org/genproto/googleapis/rpc/status"
)
