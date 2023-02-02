package static_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/boundary/api/authmethods"
	"github.com/hashicorp/boundary/testing/internal/e2e"
	"github.com/hashicorp/boundary/testing/internal/e2e/boundary"
	"github.com/stretchr/testify/require"
)

// TestCliAuthMethodPassword uses the boundary cli to... !!
func TestCliAuthMethodPassword(t *testing.T) {
	e2e.MaybeSkipTest(t)

	ctx := context.Background()
	boundary.AuthenticateAdminCli(t, ctx)
	newOrgId := boundary.CreateNewOrgCli(t, ctx)
	t.Cleanup(func() {
		ctx := context.Background()
		boundary.AuthenticateAdminCli(t, ctx)
		output := e2e.RunCommand(ctx, "boundary", e2e.WithArgs("scopes", "delete", "-id", newOrgId))
		require.NoError(t, output.Err, string(output.Stderr))
	})

	output := e2e.RunCommand(ctx, "boundary",
		e2e.WithArgs(
			"auth-methods", "create", "password",
			"-scope-id", newOrgId,
			"-name", "e2e Auth Method",
			"-format", "json",
		),
	)
	require.NoError(t, output.Err, output.Stderr)
	var newAuthMethodResult authmethods.AuthMethodCreateResult
	err := json.Unmarshal(output.Stdout, &newAuthMethodResult)
	require.NoError(t, err)
	newAuthMethodId := newAuthMethodResult.Item.Id
	t.Logf("Created Auth Method Id: %s", newAuthMethodId)

	newAccountId, _, _ := boundary.CreateNewAccountCli(t, ctx, newAuthMethodId)
	newUserId := boundary.CreateNewUserCli(t, ctx, newOrgId)
	t.Cleanup(func() {
		boundary.AuthenticateAdminCli(t, context.Background())
		output := e2e.RunCommand(ctx, "boundary",
			e2e.WithArgs("users", "delete", "-id", newUserId),
		)
		require.NoError(t, output.Err, string(output.Stderr))
	})

	boundary.SetAccountToUserCli(t, ctx, newUserId, newAccountId)

	// login as user?
}
