package ui_test

import (
	"context"
	"testing"

	"github.com/hashicorp/boundary/testing/internal/e2e"
	"github.com/hashicorp/boundary/testing/internal/e2e/boundary"
	"github.com/hashicorp/go-secure-stdlib/base62"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestUiCreateTarget(t *testing.T) {
	e2e.MaybeSkipTest(t)
	c, err := loadConfig()
	require.NoError(t, err)

	pw, err := playwright.Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	require.NoError(t, err)
	cc, err := browser.NewContext()
	require.NoError(t, err)
	page, err := cc.NewPage()
	require.NoError(t, err)
	_, err = page.Goto(c.Address)
	require.NoError(t, err)

	orgName, err := base62.Random(16)
	require.NoError(t, err)

	require.NoError(t, page.Type("input[name=identification]", c.AdminLoginName))
	require.NoError(t, page.Type("input[name=password]", c.AdminLoginPassword))
	require.NoError(t, page.Click("button[type=submit]"))
	locator, err := page.Locator("text=" + c.AdminLoginName)
	require.NoError(t, err)
	locatorVisible, err := locator.IsEnabled()
	require.NoError(t, err)
	require.True(t, locatorVisible)

	require.NoError(t, page.Click("//a//span[text()='New']"))
	require.NoError(t, page.Type("input[name=name]", orgName))
	require.NoError(t, page.Type("textarea[name=description]", "This is an automated test"))
	require.NoError(t, page.Click("button[type=submit]"))

	ctx := context.Background()
	boundary.AuthenticateAdminCli(t, ctx)
	output := e2e.RunCommand(ctx, "boundary", e2e.WithArgs("scopes", "list", "-format", "json"))
	require.NoError(t, output.Err, string(output.Stderr))
}
