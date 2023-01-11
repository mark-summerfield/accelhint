// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package accelhints

import (
    "fmt"
    _ "embed"
    )

//go:embed Version.dat
var Version string

func Hello() string {
    return fmt.Sprintf("Hello accelhints v%s", Version)
}
