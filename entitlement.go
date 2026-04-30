// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// EntitlementService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitlementService] method instead.
type EntitlementService struct {
	Options []option.RequestOption
	Files   *EntitlementFileService
	Grants  *EntitlementGrantService
}

// NewEntitlementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntitlementService(opts ...option.RequestOption) (r *EntitlementService) {
	r = &EntitlementService{}
	r.Options = opts
	r.Files = NewEntitlementFileService(opts...)
	r.Grants = NewEntitlementGrantService(opts...)
	return
}

// POST /entitlements
func (r *EntitlementService) New(ctx context.Context, body EntitlementNewParams, opts ...option.RequestOption) (res *Entitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// GET /entitlements/{id}
func (r *EntitlementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Entitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// PATCH /entitlements/{id}
func (r *EntitlementService) Update(ctx context.Context, id string, body EntitlementUpdateParams, opts ...option.RequestOption) (res *Entitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// GET /entitlements
func (r *EntitlementService) List(ctx context.Context, query EntitlementListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Entitlement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "entitlements"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// GET /entitlements
func (r *EntitlementService) ListAutoPaging(ctx context.Context, query EntitlementListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Entitlement] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// DELETE /entitlements/{id} (soft-delete)
func (r *EntitlementService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Entitlement struct {
	ID         string    `json:"id" api:"required"`
	BusinessID string    `json:"business_id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
	// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
	// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
	// ID-only via [`IntegrationConfig`]; this enum is response-only.
	IntegrationConfig IntegrationConfigResponse  `json:"integration_config" api:"required"`
	IntegrationType   EntitlementIntegrationType `json:"integration_type" api:"required"`
	IsActive          bool                       `json:"is_active" api:"required"`
	Name              string                     `json:"name" api:"required"`
	UpdatedAt         time.Time                  `json:"updated_at" api:"required" format:"date-time"`
	Description       string                     `json:"description" api:"nullable"`
	Metadata          interface{}                `json:"metadata"`
	JSON              entitlementJSON            `json:"-"`
}

// entitlementJSON contains the JSON metadata for the struct [Entitlement]
type entitlementJSON struct {
	ID                apijson.Field
	BusinessID        apijson.Field
	CreatedAt         apijson.Field
	IntegrationConfig apijson.Field
	IntegrationType   apijson.Field
	IsActive          apijson.Field
	Name              apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
	Metadata          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Entitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementJSON) RawJSON() string {
	return r.raw
}

type EntitlementIntegrationType string

const (
	EntitlementIntegrationTypeDiscord      EntitlementIntegrationType = "discord"
	EntitlementIntegrationTypeTelegram     EntitlementIntegrationType = "telegram"
	EntitlementIntegrationTypeGitHub       EntitlementIntegrationType = "github"
	EntitlementIntegrationTypeFigma        EntitlementIntegrationType = "figma"
	EntitlementIntegrationTypeFramer       EntitlementIntegrationType = "framer"
	EntitlementIntegrationTypeNotion       EntitlementIntegrationType = "notion"
	EntitlementIntegrationTypeDigitalFiles EntitlementIntegrationType = "digital_files"
	EntitlementIntegrationTypeLicenseKey   EntitlementIntegrationType = "license_key"
)

func (r EntitlementIntegrationType) IsKnown() bool {
	switch r {
	case EntitlementIntegrationTypeDiscord, EntitlementIntegrationTypeTelegram, EntitlementIntegrationTypeGitHub, EntitlementIntegrationTypeFigma, EntitlementIntegrationTypeFramer, EntitlementIntegrationTypeNotion, EntitlementIntegrationTypeDigitalFiles, EntitlementIntegrationTypeLicenseKey:
		return true
	}
	return false
}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
type IntegrationConfigParam struct {
	ActivationMessage param.Field[string]       `json:"activation_message"`
	ActivationsLimit  param.Field[int64]        `json:"activations_limit"`
	ChatID            param.Field[string]       `json:"chat_id"`
	DigitalFileIDs    param.Field[interface{}]  `json:"digital_file_ids"`
	DurationCount     param.Field[int64]        `json:"duration_count"`
	DurationInterval  param.Field[TimeInterval] `json:"duration_interval"`
	ExternalURL       param.Field[string]       `json:"external_url"`
	FigmaFileID       param.Field[string]       `json:"figma_file_id"`
	FramerTemplateID  param.Field[string]       `json:"framer_template_id"`
	GuildID           param.Field[string]       `json:"guild_id"`
	Instructions      param.Field[string]       `json:"instructions"`
	LegacyFileIDs     param.Field[interface{}]  `json:"legacy_file_ids"`
	NotionTemplateID  param.Field[string]       `json:"notion_template_id"`
	// One of: pull, push, admin, maintain, triage
	Permission param.Field[string] `json:"permission"`
	RoleID     param.Field[string] `json:"role_id"`
	TargetID   param.Field[string] `json:"target_id"`
}

func (r IntegrationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigParam) implementsIntegrationConfigUnionParam() {}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
//
// Satisfied by [IntegrationConfigGitHubConfigParam],
// [IntegrationConfigDiscordConfigParam], [IntegrationConfigTelegramConfigParam],
// [IntegrationConfigFigmaConfigParam], [IntegrationConfigFramerConfigParam],
// [IntegrationConfigNotionConfigParam],
// [IntegrationConfigDigitalFilesConfigParam],
// [IntegrationConfigLicenseKeyConfigParam], [IntegrationConfigParam].
type IntegrationConfigUnionParam interface {
	implementsIntegrationConfigUnionParam()
}

type IntegrationConfigGitHubConfigParam struct {
	// One of: pull, push, admin, maintain, triage
	Permission param.Field[string] `json:"permission" api:"required"`
	TargetID   param.Field[string] `json:"target_id" api:"required"`
}

func (r IntegrationConfigGitHubConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigGitHubConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigDiscordConfigParam struct {
	GuildID param.Field[string] `json:"guild_id" api:"required"`
	RoleID  param.Field[string] `json:"role_id"`
}

func (r IntegrationConfigDiscordConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigDiscordConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigTelegramConfigParam struct {
	ChatID param.Field[string] `json:"chat_id" api:"required"`
}

func (r IntegrationConfigTelegramConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigTelegramConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigFigmaConfigParam struct {
	FigmaFileID param.Field[string] `json:"figma_file_id" api:"required"`
}

func (r IntegrationConfigFigmaConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigFigmaConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigFramerConfigParam struct {
	FramerTemplateID param.Field[string] `json:"framer_template_id" api:"required"`
}

func (r IntegrationConfigFramerConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigFramerConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigNotionConfigParam struct {
	NotionTemplateID param.Field[string] `json:"notion_template_id" api:"required"`
}

func (r IntegrationConfigNotionConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigNotionConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigDigitalFilesConfigParam struct {
	DigitalFileIDs param.Field[[]string] `json:"digital_file_ids" api:"required"`
	ExternalURL    param.Field[string]   `json:"external_url"`
	Instructions   param.Field[string]   `json:"instructions"`
	// Three-way patchable field (mirrors the credit_entitlements pattern):
	//
	// - omitted → preserve persisted (`None`)
	// - `null` → clear (`Some(None)`)
	// - `[...]` → replace (`Some(Some(...))`)
	//
	// On Create / storage we collapse "clear" and empty-array to `None` so the
	// persisted JSONB never carries a `null` legacy_file_ids key.
	LegacyFileIDs param.Field[[]string] `json:"legacy_file_ids"`
}

func (r IntegrationConfigDigitalFilesConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigDigitalFilesConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigLicenseKeyConfigParam struct {
	ActivationMessage param.Field[string]       `json:"activation_message"`
	ActivationsLimit  param.Field[int64]        `json:"activations_limit"`
	DurationCount     param.Field[int64]        `json:"duration_count"`
	DurationInterval  param.Field[TimeInterval] `json:"duration_interval"`
}

func (r IntegrationConfigLicenseKeyConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigLicenseKeyConfigParam) implementsIntegrationConfigUnionParam() {}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
type IntegrationConfigResponse struct {
	ActivationMessage string `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64  `json:"activations_limit" api:"nullable"`
	ChatID            string `json:"chat_id"`
	// This field can have the runtime type of
	// [IntegrationConfigResponseDigitalFilesConfigDigitalFiles].
	DigitalFiles     interface{}                   `json:"digital_files"`
	DurationCount    int64                         `json:"duration_count" api:"nullable"`
	DurationInterval TimeInterval                  `json:"duration_interval" api:"nullable"`
	FigmaFileID      string                        `json:"figma_file_id"`
	FramerTemplateID string                        `json:"framer_template_id"`
	GuildID          string                        `json:"guild_id"`
	NotionTemplateID string                        `json:"notion_template_id"`
	Permission       string                        `json:"permission"`
	RoleID           string                        `json:"role_id" api:"nullable"`
	TargetID         string                        `json:"target_id"`
	JSON             integrationConfigResponseJSON `json:"-"`
	union            IntegrationConfigResponseUnion
}

// integrationConfigResponseJSON contains the JSON metadata for the struct
// [IntegrationConfigResponse]
type integrationConfigResponseJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	ChatID            apijson.Field
	DigitalFiles      apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	FigmaFileID       apijson.Field
	FramerTemplateID  apijson.Field
	GuildID           apijson.Field
	NotionTemplateID  apijson.Field
	Permission        apijson.Field
	RoleID            apijson.Field
	TargetID          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r integrationConfigResponseJSON) RawJSON() string {
	return r.raw
}

func (r *IntegrationConfigResponse) UnmarshalJSON(data []byte) (err error) {
	*r = IntegrationConfigResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IntegrationConfigResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [IntegrationConfigResponseGitHubConfig],
// [IntegrationConfigResponseDiscordConfig],
// [IntegrationConfigResponseTelegramConfig],
// [IntegrationConfigResponseFigmaConfig], [IntegrationConfigResponseFramerConfig],
// [IntegrationConfigResponseNotionConfig],
// [IntegrationConfigResponseDigitalFilesConfig],
// [IntegrationConfigResponseLicenseKeyConfig].
func (r IntegrationConfigResponse) AsUnion() IntegrationConfigResponseUnion {
	return r.union
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
//
// Union satisfied by [IntegrationConfigResponseGitHubConfig],
// [IntegrationConfigResponseDiscordConfig],
// [IntegrationConfigResponseTelegramConfig],
// [IntegrationConfigResponseFigmaConfig], [IntegrationConfigResponseFramerConfig],
// [IntegrationConfigResponseNotionConfig],
// [IntegrationConfigResponseDigitalFilesConfig] or
// [IntegrationConfigResponseLicenseKeyConfig].
type IntegrationConfigResponseUnion interface {
	implementsIntegrationConfigResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntegrationConfigResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseGitHubConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseDiscordConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseTelegramConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseFigmaConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseFramerConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseNotionConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseDigitalFilesConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationConfigResponseLicenseKeyConfig{}),
		},
	)
}

type IntegrationConfigResponseGitHubConfig struct {
	Permission string                                    `json:"permission" api:"required"`
	TargetID   string                                    `json:"target_id" api:"required"`
	JSON       integrationConfigResponseGitHubConfigJSON `json:"-"`
}

// integrationConfigResponseGitHubConfigJSON contains the JSON metadata for the
// struct [IntegrationConfigResponseGitHubConfig]
type integrationConfigResponseGitHubConfigJSON struct {
	Permission  apijson.Field
	TargetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationConfigResponseGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseGitHubConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseGitHubConfig) implementsIntegrationConfigResponse() {}

type IntegrationConfigResponseDiscordConfig struct {
	GuildID string                                     `json:"guild_id" api:"required"`
	RoleID  string                                     `json:"role_id" api:"nullable"`
	JSON    integrationConfigResponseDiscordConfigJSON `json:"-"`
}

// integrationConfigResponseDiscordConfigJSON contains the JSON metadata for the
// struct [IntegrationConfigResponseDiscordConfig]
type integrationConfigResponseDiscordConfigJSON struct {
	GuildID     apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationConfigResponseDiscordConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseDiscordConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseDiscordConfig) implementsIntegrationConfigResponse() {}

type IntegrationConfigResponseTelegramConfig struct {
	ChatID string                                      `json:"chat_id" api:"required"`
	JSON   integrationConfigResponseTelegramConfigJSON `json:"-"`
}

// integrationConfigResponseTelegramConfigJSON contains the JSON metadata for the
// struct [IntegrationConfigResponseTelegramConfig]
type integrationConfigResponseTelegramConfigJSON struct {
	ChatID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationConfigResponseTelegramConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseTelegramConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseTelegramConfig) implementsIntegrationConfigResponse() {}

type IntegrationConfigResponseFigmaConfig struct {
	FigmaFileID string                                   `json:"figma_file_id" api:"required"`
	JSON        integrationConfigResponseFigmaConfigJSON `json:"-"`
}

// integrationConfigResponseFigmaConfigJSON contains the JSON metadata for the
// struct [IntegrationConfigResponseFigmaConfig]
type integrationConfigResponseFigmaConfigJSON struct {
	FigmaFileID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationConfigResponseFigmaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseFigmaConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseFigmaConfig) implementsIntegrationConfigResponse() {}

type IntegrationConfigResponseFramerConfig struct {
	FramerTemplateID string                                    `json:"framer_template_id" api:"required"`
	JSON             integrationConfigResponseFramerConfigJSON `json:"-"`
}

// integrationConfigResponseFramerConfigJSON contains the JSON metadata for the
// struct [IntegrationConfigResponseFramerConfig]
type integrationConfigResponseFramerConfigJSON struct {
	FramerTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IntegrationConfigResponseFramerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseFramerConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseFramerConfig) implementsIntegrationConfigResponse() {}

type IntegrationConfigResponseNotionConfig struct {
	NotionTemplateID string                                    `json:"notion_template_id" api:"required"`
	JSON             integrationConfigResponseNotionConfigJSON `json:"-"`
}

// integrationConfigResponseNotionConfigJSON contains the JSON metadata for the
// struct [IntegrationConfigResponseNotionConfig]
type integrationConfigResponseNotionConfigJSON struct {
	NotionTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IntegrationConfigResponseNotionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseNotionConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseNotionConfig) implementsIntegrationConfigResponse() {}

type IntegrationConfigResponseDigitalFilesConfig struct {
	// Populated digital-files payload for entitlement read surfaces. Mirrors
	// `DigitalProductDelivery` but is sourced from an entitlement's
	// `integration_config` (not a grant) and tags each file with its origin (`legacy`
	// vs `ee`).
	DigitalFiles IntegrationConfigResponseDigitalFilesConfigDigitalFiles `json:"digital_files" api:"required"`
	JSON         integrationConfigResponseDigitalFilesConfigJSON         `json:"-"`
}

// integrationConfigResponseDigitalFilesConfigJSON contains the JSON metadata for
// the struct [IntegrationConfigResponseDigitalFilesConfig]
type integrationConfigResponseDigitalFilesConfigJSON struct {
	DigitalFiles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntegrationConfigResponseDigitalFilesConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseDigitalFilesConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseDigitalFilesConfig) implementsIntegrationConfigResponse() {}

// Populated digital-files payload for entitlement read surfaces. Mirrors
// `DigitalProductDelivery` but is sourced from an entitlement's
// `integration_config` (not a grant) and tags each file with its origin (`legacy`
// vs `ee`).
type IntegrationConfigResponseDigitalFilesConfigDigitalFiles struct {
	Files        []IntegrationConfigResponseDigitalFilesConfigDigitalFilesFile `json:"files" api:"required"`
	ExternalURL  string                                                        `json:"external_url" api:"nullable"`
	Instructions string                                                        `json:"instructions" api:"nullable"`
	JSON         integrationConfigResponseDigitalFilesConfigDigitalFilesJSON   `json:"-"`
}

// integrationConfigResponseDigitalFilesConfigDigitalFilesJSON contains the JSON
// metadata for the struct
// [IntegrationConfigResponseDigitalFilesConfigDigitalFiles]
type integrationConfigResponseDigitalFilesConfigDigitalFilesJSON struct {
	Files        apijson.Field
	ExternalURL  apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntegrationConfigResponseDigitalFilesConfigDigitalFiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseDigitalFilesConfigDigitalFilesJSON) RawJSON() string {
	return r.raw
}

type IntegrationConfigResponseDigitalFilesConfigDigitalFilesFile struct {
	DownloadURL string `json:"download_url" api:"required"`
	// Seconds until `download_url` expires.
	ExpiresIn int64  `json:"expires_in" api:"required"`
	FileID    string `json:"file_id" api:"required"`
	Filename  string `json:"filename" api:"required"`
	// `"legacy"` for files in `product_files`, `"ee"` for files managed by the
	// Entitlements Engine.
	Source      string                                                          `json:"source" api:"required"`
	ContentType string                                                          `json:"content_type" api:"nullable"`
	FileSize    int64                                                           `json:"file_size" api:"nullable"`
	JSON        integrationConfigResponseDigitalFilesConfigDigitalFilesFileJSON `json:"-"`
}

// integrationConfigResponseDigitalFilesConfigDigitalFilesFileJSON contains the
// JSON metadata for the struct
// [IntegrationConfigResponseDigitalFilesConfigDigitalFilesFile]
type integrationConfigResponseDigitalFilesConfigDigitalFilesFileJSON struct {
	DownloadURL apijson.Field
	ExpiresIn   apijson.Field
	FileID      apijson.Field
	Filename    apijson.Field
	Source      apijson.Field
	ContentType apijson.Field
	FileSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationConfigResponseDigitalFilesConfigDigitalFilesFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseDigitalFilesConfigDigitalFilesFileJSON) RawJSON() string {
	return r.raw
}

type IntegrationConfigResponseLicenseKeyConfig struct {
	ActivationMessage string                                        `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64                                         `json:"activations_limit" api:"nullable"`
	DurationCount     int64                                         `json:"duration_count" api:"nullable"`
	DurationInterval  TimeInterval                                  `json:"duration_interval" api:"nullable"`
	JSON              integrationConfigResponseLicenseKeyConfigJSON `json:"-"`
}

// integrationConfigResponseLicenseKeyConfigJSON contains the JSON metadata for the
// struct [IntegrationConfigResponseLicenseKeyConfig]
type integrationConfigResponseLicenseKeyConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *IntegrationConfigResponseLicenseKeyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigResponseLicenseKeyConfigJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationConfigResponseLicenseKeyConfig) implementsIntegrationConfigResponse() {}

type EntitlementNewParams struct {
	// Platform-specific configuration (validated per integration_type)
	IntegrationConfig param.Field[IntegrationConfigUnionParam] `json:"integration_config" api:"required"`
	// Which platform integration this entitlement uses
	IntegrationType param.Field[EntitlementIntegrationType] `json:"integration_type" api:"required"`
	// Display name for this entitlement
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description
	Description param.Field[string] `json:"description"`
	// Optional user-facing metadata
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r EntitlementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntitlementUpdateParams struct {
	Description param.Field[string] `json:"description"`
	// Platform-specific configuration for an entitlement. Each variant uses unique
	// field names so `#[serde(untagged)]` can disambiguate correctly.
	IntegrationConfig param.Field[IntegrationConfigUnionParam] `json:"integration_config"`
	Metadata          param.Field[map[string]string]           `json:"metadata"`
	Name              param.Field[string]                      `json:"name"`
}

func (r EntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntitlementListParams struct {
	// Filter by integration type
	IntegrationType param.Field[EntitlementListParamsIntegrationType] `query:"integration_type"`
	// Page number (default 0)
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size (default 10, max 100)
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [EntitlementListParams]'s query parameters as `url.Values`.
func (r EntitlementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by integration type
type EntitlementListParamsIntegrationType string

const (
	EntitlementListParamsIntegrationTypeDiscord      EntitlementListParamsIntegrationType = "discord"
	EntitlementListParamsIntegrationTypeTelegram     EntitlementListParamsIntegrationType = "telegram"
	EntitlementListParamsIntegrationTypeGitHub       EntitlementListParamsIntegrationType = "github"
	EntitlementListParamsIntegrationTypeFigma        EntitlementListParamsIntegrationType = "figma"
	EntitlementListParamsIntegrationTypeFramer       EntitlementListParamsIntegrationType = "framer"
	EntitlementListParamsIntegrationTypeNotion       EntitlementListParamsIntegrationType = "notion"
	EntitlementListParamsIntegrationTypeDigitalFiles EntitlementListParamsIntegrationType = "digital_files"
	EntitlementListParamsIntegrationTypeLicenseKey   EntitlementListParamsIntegrationType = "license_key"
)

func (r EntitlementListParamsIntegrationType) IsKnown() bool {
	switch r {
	case EntitlementListParamsIntegrationTypeDiscord, EntitlementListParamsIntegrationTypeTelegram, EntitlementListParamsIntegrationTypeGitHub, EntitlementListParamsIntegrationTypeFigma, EntitlementListParamsIntegrationTypeFramer, EntitlementListParamsIntegrationTypeNotion, EntitlementListParamsIntegrationTypeDigitalFiles, EntitlementListParamsIntegrationTypeLicenseKey:
		return true
	}
	return false
}
