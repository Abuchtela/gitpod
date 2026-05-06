// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package util_test

import (
	"fmt"
	"testing"

	"github.com/gitpod-io/gitpod/common-go/util"
)

func TestInLeewayBuild_False(t *testing.T) {
	// Ensure the env var is not set to "true" during this test.
	t.Setenv("LEEWAY_BUILD", "")
	if util.InLeewayBuild() {
		t.Error("expected InLeewayBuild() to return false when LEEWAY_BUILD is not \"true\"")
	}
}

func TestInLeewayBuild_True(t *testing.T) {
	t.Setenv("LEEWAY_BUILD", "true")
	if !util.InLeewayBuild() {
		t.Error("expected InLeewayBuild() to return true when LEEWAY_BUILD=true")
	}
}

func TestGetSupervisorAddress_Default(t *testing.T) {
	t.Setenv("SUPERVISOR_ADDR", "")
	want := fmt.Sprintf("127.0.0.1:%d", util.SupervisorPort)
	if got := util.GetSupervisorAddress(); got != want {
		t.Errorf("GetSupervisorAddress() = %q, want %q", got, want)
	}
}

func TestGetSupervisorAddress_CustomValue(t *testing.T) {
	const custom = "10.0.0.1:12345"
	t.Setenv("SUPERVISOR_ADDR", custom)
	if got := util.GetSupervisorAddress(); got != custom {
		t.Errorf("GetSupervisorAddress() = %q, want %q", got, custom)
	}
}