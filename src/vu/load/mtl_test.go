// Copyright © 2013-2014 Galvanized Logic Inc.
// Use is governed by a FreeBSD license found in the LICENSE file.

package load

import (
	"fmt"
	"testing"
)

func TestLoadMtl(t *testing.T) {
	load := newLoader().setDir(mod, "../eg/models")
	m, err := load.mtl("cube")
	if m == nil || err != nil {
		t.Fatalf("Should be able to load a valid material file %s", err)
	}
	got, want := fmt.Sprintf("%2.1f %2.1f %2.1f", m.KdR, m.KdG, m.KdB), "0.6 0.6 0.6"
	if got != want {
		t.Errorf(format, got, want)
	}
}
