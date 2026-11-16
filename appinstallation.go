// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// AppInstallationService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppInstallationService] method instead.
type AppInstallationService struct {
	Options []option.RequestOption
}

// NewAppInstallationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppInstallationService(opts ...option.RequestOption) (r *AppInstallationService) {
	r = &AppInstallationService{}
	r.Options = opts
	return
}

// The settings of the App Installation of the app token, with the secret values.
// The token grant gives the installation. The token needs `app_settings:read`.
func (r *AppInstallationService) GetSettings(ctx context.Context, opts ...option.RequestOption) (res *AppInstallationGetSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "app-installation/settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The settings of the App Installation of the app token, with the secret values.
type AppInstallationGetSettingsResponse struct {
	// `true` when a required field has no value, or a stored value fails a live rule.
	ActionNeeded   bool   `json:"action_needed" api:"required"`
	AppID          string `json:"app_id" api:"required"`
	InstallationID string `json:"installation_id" api:"required"`
	// The keys that need an action.
	InvalidKeys []string `json:"invalid_keys" api:"required"`
	// The values of the secret fields of the live manifest, by key.
	Secrets interface{} `json:"secrets" api:"required"`
	// The values of the fields of the live manifest that are not secrets.
	Values interface{} `json:"values" api:"required"`
	// The version of the live settings manifest. It is `null` when the app has no live
	// manifest. Then `values` and `secrets` are empty.
	LiveManifestVersion int64 `json:"live_manifest_version" api:"nullable"`
	// The manifest version of the last save. It is `null` when the installation has no
	// settings.
	ManifestVersion int64 `json:"manifest_version" api:"nullable"`
	// The time of the last save.
	UpdatedAt time.Time                              `json:"updated_at" api:"nullable" format:"date-time"`
	JSON      appInstallationGetSettingsResponseJSON `json:"-"`
}

// appInstallationGetSettingsResponseJSON contains the JSON metadata for the struct
// [AppInstallationGetSettingsResponse]
type appInstallationGetSettingsResponseJSON struct {
	ActionNeeded        apijson.Field
	AppID               apijson.Field
	InstallationID      apijson.Field
	InvalidKeys         apijson.Field
	Secrets             apijson.Field
	Values              apijson.Field
	LiveManifestVersion apijson.Field
	ManifestVersion     apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AppInstallationGetSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appInstallationGetSettingsResponseJSON) RawJSON() string {
	return r.raw
}
