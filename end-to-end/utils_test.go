package endtoend_test

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/canonical/ubuntu-pro-for-windows/common/wsltestutils"
	"github.com/stretchr/testify/require"
	"github.com/ubuntu/gowsl"
)

func testSetup(t *testing.T) {
	t.Helper()
	ctx := context.Background()

	err := gowsl.Shutdown(ctx)
	require.NoError(t, err, "Setup: could not shut WSL down")

	err = assertCleanRegistry()
	require.NoError(t, err, "Setup: registry is polluted, potentially by a previous test")

	err = assertCleanLocalAppData()
	require.NoError(t, err, "Setup: local app data is polluted, potentially by a previous test")

	err = stopAgent(ctx)
	require.NoError(t, err, "Setup: could not stop the agent")

	t.Cleanup(func() {
		err := errors.Join(
			cleanupRegistry(),
			cleanupLocalAppData(),
			stopAgent(ctx),
		)
		// Cannot assert: the test is finished already
		log.Printf("Cleanup: %v", err)
	})
}

func registerFromGoldenImage(t *testing.T, ctx context.Context) string {
	t.Helper()

	distroName := wsltestutils.RandomDistroName(t)
	_ = wsltestutils.PowershellInstallDistro(t, ctx, distroName, goldenImagePath)
	return distroName
}

func startAgent(t *testing.T, ctx context.Context) {
	out, err := powershellf(ctx, "(Get-AppxPackage CanonicalGroupLimited.UbuntuProForWindows).InstallLocation").CombinedOutput()
	require.NoError(t, err, "could not locate ubuntupro.exe: %v. %s", err, out)

	ctx, cancel := context.WithCancel(ctx)

	ubuntupro := filepath.Join(strings.TrimSpace(string(out)), "gui", "ubuntupro.exe")
	cmd := exec.CommandContext(ctx, ubuntupro)

	t.Cleanup(func() {
		cancel()
		cmd.Wait()
	})

	cmd.Start()

	require.NoError(t, err, "Setup: could not start agent")

	require.Eventually(t, func() bool {
		localAppData := os.Getenv("LocalAppData")
		require.NotEmpty(t, localAppData, "$env:LocalAppData should not be empty")

		_, err := os.Stat(filepath.Join(localAppData, "Ubuntu Pro", "addr"))
		if err == nil {
			return true
		}
		require.ErrorIsf(t, err, fs.ErrNotExist, "could not read addr file")

		return false
	}, 5*time.Second, 100*time.Millisecond, "Agent never started serving")
}

func stopAgent(ctx context.Context) error {
	const process = "ubuntu-pro-agent"

	out, err := powershellf(ctx, "Stop-Process -ProcessName %q", process).CombinedOutput()
	if err == nil {
		return nil
	}

	if strings.Contains(string(out), "NoProcessFoundForGivenName") {
		return nil
	}

	return fmt.Errorf("could not stop process %q: %v. %s", process, err, out)
}
