package contracts_test

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/canonical/ubuntu-pro-for-windows/mocks/contractserver/contractsmockserver"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/contracts"
	"github.com/stretchr/testify/require"
)

func TestProToken(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		// Msft store
		expired      bool
		jwtError     bool
		expDateError bool

		// Contract server
		getServerAccessTokenErr bool
		getProTokenErr          bool

		wantErr bool
	}{
		"Success": {},

		"Error when the subscription has expired":                     {expired: true, wantErr: true},
		"Error when the store's GenerateUserJWT fails":                {jwtError: true, wantErr: true},
		"Error when the store's GetSubscriptionExpirationDate fails":  {expDateError: true, wantErr: true},
		"Error when the contract server's GetServerAccessToken fails": {getServerAccessTokenErr: true, wantErr: true},
		"Error when the contract server's GetProToken fails":          {getProTokenErr: true, wantErr: true},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			store := mockMSStore{
				expirationDate:    time.Now().Add(8760 * time.Hour), // Next year
				expirationDateErr: tc.expDateError,

				jwt:            "JWT_123",
				jwtWantADToken: "AZURE_AD_TOKEN",
				jwtErr:         tc.jwtError,
			}

			if tc.expired {
				store.expirationDate = time.Now().Add(-8760 * time.Hour) // Last year
			}

			settings := contractsmockserver.DefaultSettings()

			settings.Token.OnSuccess.Value = "AZURE_AD_TOKEN"
			settings.Subscription.OnSuccess.Value = "UBUNTU_PRO_TOKEN"

			settings.Token.Disabled = tc.getServerAccessTokenErr
			settings.Subscription.Disabled = tc.getProTokenErr

			server := contractsmockserver.NewServer(settings)
			addr, err := server.Serve(ctx)
			require.NoError(t, err, "Setup: Server should return no error")
			//nolint:errcheck // Nothing we can do about it
			defer server.Stop()

			url, err := url.Parse(fmt.Sprintf("http://%s", addr))
			require.NoError(t, err, "Setup: Server URL should have been parsed with no issues")

			token, err := contracts.ProToken(ctx, contracts.WithProURL(url), contracts.WithMockMicrosoftStore(store))
			if tc.wantErr {
				require.Error(t, err, "ProToken should return an error")
				return
			}
			require.NoError(t, err, "ProToken should return no error")

			require.Equal(t, "UBUNTU_PRO_TOKEN", token, "Unexpected value for the pro token")
		})
	}
}

type mockMSStore struct {
	jwt            string
	jwtWantADToken string
	jwtErr         bool

	expirationDate    time.Time
	expirationDateErr bool
}

func (s mockMSStore) GenerateUserJWT(azureADToken string) (jwt string, err error) {
	if s.jwtErr {
		return "", errors.New("mock error")
	}

	if azureADToken != s.jwtWantADToken {
		return "", fmt.Errorf("azure AD token does not match. Want %q and got %q", s.jwtWantADToken, azureADToken)
	}

	return s.jwt, nil
}

func (s mockMSStore) GetSubscriptionExpirationDate() (tm time.Time, err error) {
	if s.expirationDateErr {
		return time.Time{}, errors.New("mock error")
	}

	return s.expirationDate, nil
}
