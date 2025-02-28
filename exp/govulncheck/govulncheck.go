// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package govulncheck has experimental govulncheck API.
package govulncheck

import (
	"context"

	"golang.org/x/vuln/internal/govulncheck"
)

// Config is the configuration for Main.
type Config = govulncheck.Config

// Run is the primary function for the govulncheck command line tool.
func Run(ctx context.Context, cfg Config) error {
	return govulncheck.Run(ctx, cfg)
}
