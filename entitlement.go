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
func (r *EntitlementService) New(ctx context.Context, body EntitlementNewParams, opts ...option.RequestOption) (res *EntitlementNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// GET /entitlements/{id}
func (r *EntitlementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EntitlementGetResponse, err error) {
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
func (r *EntitlementService) Update(ctx context.Context, id string, body EntitlementUpdateParams, opts ...option.RequestOption) (res *EntitlementUpdateResponse, err error) {
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
func (r *EntitlementService) List(ctx context.Context, query EntitlementListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[EntitlementListResponse], err error) {
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
func (r *EntitlementService) ListAutoPaging(ctx context.Context, query EntitlementListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[EntitlementListResponse] {
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

type EntitlementNewResponse struct {
	ID         string    `json:"id" api:"required"`
	BusinessID string    `json:"business_id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
	// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
	// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
	// ID-only via [`IntegrationConfig`]; this enum is response-only.
	IntegrationConfig EntitlementNewResponseIntegrationConfig `json:"integration_config" api:"required"`
	IntegrationType   EntitlementNewResponseIntegrationType   `json:"integration_type" api:"required"`
	IsActive          bool                                    `json:"is_active" api:"required"`
	Name              string                                  `json:"name" api:"required"`
	UpdatedAt         time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                  `json:"description" api:"nullable"`
	Metadata          interface{}                             `json:"metadata"`
	JSON              entitlementNewResponseJSON              `json:"-"`
}

// entitlementNewResponseJSON contains the JSON metadata for the struct
// [EntitlementNewResponse]
type entitlementNewResponseJSON struct {
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

func (r *EntitlementNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseJSON) RawJSON() string {
	return r.raw
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
type EntitlementNewResponseIntegrationConfig struct {
	ActivationMessage string `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64  `json:"activations_limit" api:"nullable"`
	ChatID            string `json:"chat_id"`
	// This field can have the runtime type of
	// [EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFiles].
	DigitalFiles     interface{}                                 `json:"digital_files"`
	DurationCount    int64                                       `json:"duration_count" api:"nullable"`
	DurationInterval TimeInterval                                `json:"duration_interval" api:"nullable"`
	FigmaFileID      string                                      `json:"figma_file_id"`
	FramerTemplateID string                                      `json:"framer_template_id"`
	GuildID          string                                      `json:"guild_id"`
	NotionTemplateID string                                      `json:"notion_template_id"`
	Permission       string                                      `json:"permission"`
	RoleID           string                                      `json:"role_id" api:"nullable"`
	TargetID         string                                      `json:"target_id"`
	JSON             entitlementNewResponseIntegrationConfigJSON `json:"-"`
	union            EntitlementNewResponseIntegrationConfigUnion
}

// entitlementNewResponseIntegrationConfigJSON contains the JSON metadata for the
// struct [EntitlementNewResponseIntegrationConfig]
type entitlementNewResponseIntegrationConfigJSON struct {
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

func (r entitlementNewResponseIntegrationConfigJSON) RawJSON() string {
	return r.raw
}

func (r *EntitlementNewResponseIntegrationConfig) UnmarshalJSON(data []byte) (err error) {
	*r = EntitlementNewResponseIntegrationConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EntitlementNewResponseIntegrationConfigUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EntitlementNewResponseIntegrationConfigGitHubConfig],
// [EntitlementNewResponseIntegrationConfigDiscordConfig],
// [EntitlementNewResponseIntegrationConfigTelegramConfig],
// [EntitlementNewResponseIntegrationConfigFigmaConfig],
// [EntitlementNewResponseIntegrationConfigFramerConfig],
// [EntitlementNewResponseIntegrationConfigNotionConfig],
// [EntitlementNewResponseIntegrationConfigDigitalFilesConfig],
// [EntitlementNewResponseIntegrationConfigLicenseKeyConfig].
func (r EntitlementNewResponseIntegrationConfig) AsUnion() EntitlementNewResponseIntegrationConfigUnion {
	return r.union
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
//
// Union satisfied by [EntitlementNewResponseIntegrationConfigGitHubConfig],
// [EntitlementNewResponseIntegrationConfigDiscordConfig],
// [EntitlementNewResponseIntegrationConfigTelegramConfig],
// [EntitlementNewResponseIntegrationConfigFigmaConfig],
// [EntitlementNewResponseIntegrationConfigFramerConfig],
// [EntitlementNewResponseIntegrationConfigNotionConfig],
// [EntitlementNewResponseIntegrationConfigDigitalFilesConfig] or
// [EntitlementNewResponseIntegrationConfigLicenseKeyConfig].
type EntitlementNewResponseIntegrationConfigUnion interface {
	implementsEntitlementNewResponseIntegrationConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EntitlementNewResponseIntegrationConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigGitHubConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigDiscordConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigTelegramConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigFigmaConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigFramerConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigNotionConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigDigitalFilesConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementNewResponseIntegrationConfigLicenseKeyConfig{}),
		},
	)
}

type EntitlementNewResponseIntegrationConfigGitHubConfig struct {
	Permission string                                                  `json:"permission" api:"required"`
	TargetID   string                                                  `json:"target_id" api:"required"`
	JSON       entitlementNewResponseIntegrationConfigGitHubConfigJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigGitHubConfigJSON contains the JSON
// metadata for the struct [EntitlementNewResponseIntegrationConfigGitHubConfig]
type entitlementNewResponseIntegrationConfigGitHubConfigJSON struct {
	Permission  apijson.Field
	TargetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigGitHubConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigGitHubConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

type EntitlementNewResponseIntegrationConfigDiscordConfig struct {
	GuildID string                                                   `json:"guild_id" api:"required"`
	RoleID  string                                                   `json:"role_id" api:"nullable"`
	JSON    entitlementNewResponseIntegrationConfigDiscordConfigJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigDiscordConfigJSON contains the JSON
// metadata for the struct [EntitlementNewResponseIntegrationConfigDiscordConfig]
type entitlementNewResponseIntegrationConfigDiscordConfigJSON struct {
	GuildID     apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigDiscordConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigDiscordConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigDiscordConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

type EntitlementNewResponseIntegrationConfigTelegramConfig struct {
	ChatID string                                                    `json:"chat_id" api:"required"`
	JSON   entitlementNewResponseIntegrationConfigTelegramConfigJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigTelegramConfigJSON contains the JSON
// metadata for the struct [EntitlementNewResponseIntegrationConfigTelegramConfig]
type entitlementNewResponseIntegrationConfigTelegramConfigJSON struct {
	ChatID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigTelegramConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigTelegramConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigTelegramConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

type EntitlementNewResponseIntegrationConfigFigmaConfig struct {
	FigmaFileID string                                                 `json:"figma_file_id" api:"required"`
	JSON        entitlementNewResponseIntegrationConfigFigmaConfigJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigFigmaConfigJSON contains the JSON
// metadata for the struct [EntitlementNewResponseIntegrationConfigFigmaConfig]
type entitlementNewResponseIntegrationConfigFigmaConfigJSON struct {
	FigmaFileID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigFigmaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigFigmaConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigFigmaConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

type EntitlementNewResponseIntegrationConfigFramerConfig struct {
	FramerTemplateID string                                                  `json:"framer_template_id" api:"required"`
	JSON             entitlementNewResponseIntegrationConfigFramerConfigJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigFramerConfigJSON contains the JSON
// metadata for the struct [EntitlementNewResponseIntegrationConfigFramerConfig]
type entitlementNewResponseIntegrationConfigFramerConfigJSON struct {
	FramerTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigFramerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigFramerConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigFramerConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

type EntitlementNewResponseIntegrationConfigNotionConfig struct {
	NotionTemplateID string                                                  `json:"notion_template_id" api:"required"`
	JSON             entitlementNewResponseIntegrationConfigNotionConfigJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigNotionConfigJSON contains the JSON
// metadata for the struct [EntitlementNewResponseIntegrationConfigNotionConfig]
type entitlementNewResponseIntegrationConfigNotionConfigJSON struct {
	NotionTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigNotionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigNotionConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigNotionConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

type EntitlementNewResponseIntegrationConfigDigitalFilesConfig struct {
	// Populated digital-files payload for entitlement read surfaces. Mirrors
	// `DigitalProductDelivery` but is sourced from an entitlement's
	// `integration_config` (not a grant) and tags each file with its origin (`legacy`
	// vs `ee`).
	DigitalFiles EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFiles `json:"digital_files" api:"required"`
	JSON         entitlementNewResponseIntegrationConfigDigitalFilesConfigJSON         `json:"-"`
}

// entitlementNewResponseIntegrationConfigDigitalFilesConfigJSON contains the JSON
// metadata for the struct
// [EntitlementNewResponseIntegrationConfigDigitalFilesConfig]
type entitlementNewResponseIntegrationConfigDigitalFilesConfigJSON struct {
	DigitalFiles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigDigitalFilesConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigDigitalFilesConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigDigitalFilesConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

// Populated digital-files payload for entitlement read surfaces. Mirrors
// `DigitalProductDelivery` but is sourced from an entitlement's
// `integration_config` (not a grant) and tags each file with its origin (`legacy`
// vs `ee`).
type EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFiles struct {
	Files        []EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile `json:"files" api:"required"`
	ExternalURL  string                                                                      `json:"external_url" api:"nullable"`
	Instructions string                                                                      `json:"instructions" api:"nullable"`
	JSON         entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON   `json:"-"`
}

// entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON
// contains the JSON metadata for the struct
// [EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFiles]
type entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON struct {
	Files        apijson.Field
	ExternalURL  apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON) RawJSON() string {
	return r.raw
}

type EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile struct {
	DownloadURL string `json:"download_url" api:"required"`
	// Seconds until `download_url` expires.
	ExpiresIn int64  `json:"expires_in" api:"required"`
	FileID    string `json:"file_id" api:"required"`
	Filename  string `json:"filename" api:"required"`
	// `"legacy"` for files in `product_files`, `"ee"` for files managed by the
	// Entitlements Engine.
	Source      string                                                                        `json:"source" api:"required"`
	ContentType string                                                                        `json:"content_type" api:"nullable"`
	FileSize    int64                                                                         `json:"file_size" api:"nullable"`
	JSON        entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON
// contains the JSON metadata for the struct
// [EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile]
type entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON struct {
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

func (r *EntitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON) RawJSON() string {
	return r.raw
}

type EntitlementNewResponseIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage string                                                      `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64                                                       `json:"activations_limit" api:"nullable"`
	DurationCount     int64                                                       `json:"duration_count" api:"nullable"`
	DurationInterval  TimeInterval                                                `json:"duration_interval" api:"nullable"`
	JSON              entitlementNewResponseIntegrationConfigLicenseKeyConfigJSON `json:"-"`
}

// entitlementNewResponseIntegrationConfigLicenseKeyConfigJSON contains the JSON
// metadata for the struct
// [EntitlementNewResponseIntegrationConfigLicenseKeyConfig]
type entitlementNewResponseIntegrationConfigLicenseKeyConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EntitlementNewResponseIntegrationConfigLicenseKeyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementNewResponseIntegrationConfigLicenseKeyConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementNewResponseIntegrationConfigLicenseKeyConfig) implementsEntitlementNewResponseIntegrationConfig() {
}

type EntitlementNewResponseIntegrationType string

const (
	EntitlementNewResponseIntegrationTypeDiscord      EntitlementNewResponseIntegrationType = "discord"
	EntitlementNewResponseIntegrationTypeTelegram     EntitlementNewResponseIntegrationType = "telegram"
	EntitlementNewResponseIntegrationTypeGitHub       EntitlementNewResponseIntegrationType = "github"
	EntitlementNewResponseIntegrationTypeFigma        EntitlementNewResponseIntegrationType = "figma"
	EntitlementNewResponseIntegrationTypeFramer       EntitlementNewResponseIntegrationType = "framer"
	EntitlementNewResponseIntegrationTypeNotion       EntitlementNewResponseIntegrationType = "notion"
	EntitlementNewResponseIntegrationTypeDigitalFiles EntitlementNewResponseIntegrationType = "digital_files"
	EntitlementNewResponseIntegrationTypeLicenseKey   EntitlementNewResponseIntegrationType = "license_key"
)

func (r EntitlementNewResponseIntegrationType) IsKnown() bool {
	switch r {
	case EntitlementNewResponseIntegrationTypeDiscord, EntitlementNewResponseIntegrationTypeTelegram, EntitlementNewResponseIntegrationTypeGitHub, EntitlementNewResponseIntegrationTypeFigma, EntitlementNewResponseIntegrationTypeFramer, EntitlementNewResponseIntegrationTypeNotion, EntitlementNewResponseIntegrationTypeDigitalFiles, EntitlementNewResponseIntegrationTypeLicenseKey:
		return true
	}
	return false
}

type EntitlementGetResponse struct {
	ID         string    `json:"id" api:"required"`
	BusinessID string    `json:"business_id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
	// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
	// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
	// ID-only via [`IntegrationConfig`]; this enum is response-only.
	IntegrationConfig EntitlementGetResponseIntegrationConfig `json:"integration_config" api:"required"`
	IntegrationType   EntitlementGetResponseIntegrationType   `json:"integration_type" api:"required"`
	IsActive          bool                                    `json:"is_active" api:"required"`
	Name              string                                  `json:"name" api:"required"`
	UpdatedAt         time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                  `json:"description" api:"nullable"`
	Metadata          interface{}                             `json:"metadata"`
	JSON              entitlementGetResponseJSON              `json:"-"`
}

// entitlementGetResponseJSON contains the JSON metadata for the struct
// [EntitlementGetResponse]
type entitlementGetResponseJSON struct {
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

func (r *EntitlementGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseJSON) RawJSON() string {
	return r.raw
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
type EntitlementGetResponseIntegrationConfig struct {
	ActivationMessage string `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64  `json:"activations_limit" api:"nullable"`
	ChatID            string `json:"chat_id"`
	// This field can have the runtime type of
	// [EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFiles].
	DigitalFiles     interface{}                                 `json:"digital_files"`
	DurationCount    int64                                       `json:"duration_count" api:"nullable"`
	DurationInterval TimeInterval                                `json:"duration_interval" api:"nullable"`
	FigmaFileID      string                                      `json:"figma_file_id"`
	FramerTemplateID string                                      `json:"framer_template_id"`
	GuildID          string                                      `json:"guild_id"`
	NotionTemplateID string                                      `json:"notion_template_id"`
	Permission       string                                      `json:"permission"`
	RoleID           string                                      `json:"role_id" api:"nullable"`
	TargetID         string                                      `json:"target_id"`
	JSON             entitlementGetResponseIntegrationConfigJSON `json:"-"`
	union            EntitlementGetResponseIntegrationConfigUnion
}

// entitlementGetResponseIntegrationConfigJSON contains the JSON metadata for the
// struct [EntitlementGetResponseIntegrationConfig]
type entitlementGetResponseIntegrationConfigJSON struct {
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

func (r entitlementGetResponseIntegrationConfigJSON) RawJSON() string {
	return r.raw
}

func (r *EntitlementGetResponseIntegrationConfig) UnmarshalJSON(data []byte) (err error) {
	*r = EntitlementGetResponseIntegrationConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EntitlementGetResponseIntegrationConfigUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EntitlementGetResponseIntegrationConfigGitHubConfig],
// [EntitlementGetResponseIntegrationConfigDiscordConfig],
// [EntitlementGetResponseIntegrationConfigTelegramConfig],
// [EntitlementGetResponseIntegrationConfigFigmaConfig],
// [EntitlementGetResponseIntegrationConfigFramerConfig],
// [EntitlementGetResponseIntegrationConfigNotionConfig],
// [EntitlementGetResponseIntegrationConfigDigitalFilesConfig],
// [EntitlementGetResponseIntegrationConfigLicenseKeyConfig].
func (r EntitlementGetResponseIntegrationConfig) AsUnion() EntitlementGetResponseIntegrationConfigUnion {
	return r.union
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
//
// Union satisfied by [EntitlementGetResponseIntegrationConfigGitHubConfig],
// [EntitlementGetResponseIntegrationConfigDiscordConfig],
// [EntitlementGetResponseIntegrationConfigTelegramConfig],
// [EntitlementGetResponseIntegrationConfigFigmaConfig],
// [EntitlementGetResponseIntegrationConfigFramerConfig],
// [EntitlementGetResponseIntegrationConfigNotionConfig],
// [EntitlementGetResponseIntegrationConfigDigitalFilesConfig] or
// [EntitlementGetResponseIntegrationConfigLicenseKeyConfig].
type EntitlementGetResponseIntegrationConfigUnion interface {
	implementsEntitlementGetResponseIntegrationConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EntitlementGetResponseIntegrationConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigGitHubConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigDiscordConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigTelegramConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigFigmaConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigFramerConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigNotionConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigDigitalFilesConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementGetResponseIntegrationConfigLicenseKeyConfig{}),
		},
	)
}

type EntitlementGetResponseIntegrationConfigGitHubConfig struct {
	Permission string                                                  `json:"permission" api:"required"`
	TargetID   string                                                  `json:"target_id" api:"required"`
	JSON       entitlementGetResponseIntegrationConfigGitHubConfigJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigGitHubConfigJSON contains the JSON
// metadata for the struct [EntitlementGetResponseIntegrationConfigGitHubConfig]
type entitlementGetResponseIntegrationConfigGitHubConfigJSON struct {
	Permission  apijson.Field
	TargetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigGitHubConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigGitHubConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

type EntitlementGetResponseIntegrationConfigDiscordConfig struct {
	GuildID string                                                   `json:"guild_id" api:"required"`
	RoleID  string                                                   `json:"role_id" api:"nullable"`
	JSON    entitlementGetResponseIntegrationConfigDiscordConfigJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigDiscordConfigJSON contains the JSON
// metadata for the struct [EntitlementGetResponseIntegrationConfigDiscordConfig]
type entitlementGetResponseIntegrationConfigDiscordConfigJSON struct {
	GuildID     apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigDiscordConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigDiscordConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigDiscordConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

type EntitlementGetResponseIntegrationConfigTelegramConfig struct {
	ChatID string                                                    `json:"chat_id" api:"required"`
	JSON   entitlementGetResponseIntegrationConfigTelegramConfigJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigTelegramConfigJSON contains the JSON
// metadata for the struct [EntitlementGetResponseIntegrationConfigTelegramConfig]
type entitlementGetResponseIntegrationConfigTelegramConfigJSON struct {
	ChatID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigTelegramConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigTelegramConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigTelegramConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

type EntitlementGetResponseIntegrationConfigFigmaConfig struct {
	FigmaFileID string                                                 `json:"figma_file_id" api:"required"`
	JSON        entitlementGetResponseIntegrationConfigFigmaConfigJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigFigmaConfigJSON contains the JSON
// metadata for the struct [EntitlementGetResponseIntegrationConfigFigmaConfig]
type entitlementGetResponseIntegrationConfigFigmaConfigJSON struct {
	FigmaFileID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigFigmaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigFigmaConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigFigmaConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

type EntitlementGetResponseIntegrationConfigFramerConfig struct {
	FramerTemplateID string                                                  `json:"framer_template_id" api:"required"`
	JSON             entitlementGetResponseIntegrationConfigFramerConfigJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigFramerConfigJSON contains the JSON
// metadata for the struct [EntitlementGetResponseIntegrationConfigFramerConfig]
type entitlementGetResponseIntegrationConfigFramerConfigJSON struct {
	FramerTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigFramerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigFramerConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigFramerConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

type EntitlementGetResponseIntegrationConfigNotionConfig struct {
	NotionTemplateID string                                                  `json:"notion_template_id" api:"required"`
	JSON             entitlementGetResponseIntegrationConfigNotionConfigJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigNotionConfigJSON contains the JSON
// metadata for the struct [EntitlementGetResponseIntegrationConfigNotionConfig]
type entitlementGetResponseIntegrationConfigNotionConfigJSON struct {
	NotionTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigNotionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigNotionConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigNotionConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

type EntitlementGetResponseIntegrationConfigDigitalFilesConfig struct {
	// Populated digital-files payload for entitlement read surfaces. Mirrors
	// `DigitalProductDelivery` but is sourced from an entitlement's
	// `integration_config` (not a grant) and tags each file with its origin (`legacy`
	// vs `ee`).
	DigitalFiles EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFiles `json:"digital_files" api:"required"`
	JSON         entitlementGetResponseIntegrationConfigDigitalFilesConfigJSON         `json:"-"`
}

// entitlementGetResponseIntegrationConfigDigitalFilesConfigJSON contains the JSON
// metadata for the struct
// [EntitlementGetResponseIntegrationConfigDigitalFilesConfig]
type entitlementGetResponseIntegrationConfigDigitalFilesConfigJSON struct {
	DigitalFiles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigDigitalFilesConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigDigitalFilesConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigDigitalFilesConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

// Populated digital-files payload for entitlement read surfaces. Mirrors
// `DigitalProductDelivery` but is sourced from an entitlement's
// `integration_config` (not a grant) and tags each file with its origin (`legacy`
// vs `ee`).
type EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFiles struct {
	Files        []EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile `json:"files" api:"required"`
	ExternalURL  string                                                                      `json:"external_url" api:"nullable"`
	Instructions string                                                                      `json:"instructions" api:"nullable"`
	JSON         entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON   `json:"-"`
}

// entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON
// contains the JSON metadata for the struct
// [EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFiles]
type entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON struct {
	Files        apijson.Field
	ExternalURL  apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON) RawJSON() string {
	return r.raw
}

type EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile struct {
	DownloadURL string `json:"download_url" api:"required"`
	// Seconds until `download_url` expires.
	ExpiresIn int64  `json:"expires_in" api:"required"`
	FileID    string `json:"file_id" api:"required"`
	Filename  string `json:"filename" api:"required"`
	// `"legacy"` for files in `product_files`, `"ee"` for files managed by the
	// Entitlements Engine.
	Source      string                                                                        `json:"source" api:"required"`
	ContentType string                                                                        `json:"content_type" api:"nullable"`
	FileSize    int64                                                                         `json:"file_size" api:"nullable"`
	JSON        entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON
// contains the JSON metadata for the struct
// [EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile]
type entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON struct {
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

func (r *EntitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON) RawJSON() string {
	return r.raw
}

type EntitlementGetResponseIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage string                                                      `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64                                                       `json:"activations_limit" api:"nullable"`
	DurationCount     int64                                                       `json:"duration_count" api:"nullable"`
	DurationInterval  TimeInterval                                                `json:"duration_interval" api:"nullable"`
	JSON              entitlementGetResponseIntegrationConfigLicenseKeyConfigJSON `json:"-"`
}

// entitlementGetResponseIntegrationConfigLicenseKeyConfigJSON contains the JSON
// metadata for the struct
// [EntitlementGetResponseIntegrationConfigLicenseKeyConfig]
type entitlementGetResponseIntegrationConfigLicenseKeyConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EntitlementGetResponseIntegrationConfigLicenseKeyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseIntegrationConfigLicenseKeyConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementGetResponseIntegrationConfigLicenseKeyConfig) implementsEntitlementGetResponseIntegrationConfig() {
}

type EntitlementGetResponseIntegrationType string

const (
	EntitlementGetResponseIntegrationTypeDiscord      EntitlementGetResponseIntegrationType = "discord"
	EntitlementGetResponseIntegrationTypeTelegram     EntitlementGetResponseIntegrationType = "telegram"
	EntitlementGetResponseIntegrationTypeGitHub       EntitlementGetResponseIntegrationType = "github"
	EntitlementGetResponseIntegrationTypeFigma        EntitlementGetResponseIntegrationType = "figma"
	EntitlementGetResponseIntegrationTypeFramer       EntitlementGetResponseIntegrationType = "framer"
	EntitlementGetResponseIntegrationTypeNotion       EntitlementGetResponseIntegrationType = "notion"
	EntitlementGetResponseIntegrationTypeDigitalFiles EntitlementGetResponseIntegrationType = "digital_files"
	EntitlementGetResponseIntegrationTypeLicenseKey   EntitlementGetResponseIntegrationType = "license_key"
)

func (r EntitlementGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case EntitlementGetResponseIntegrationTypeDiscord, EntitlementGetResponseIntegrationTypeTelegram, EntitlementGetResponseIntegrationTypeGitHub, EntitlementGetResponseIntegrationTypeFigma, EntitlementGetResponseIntegrationTypeFramer, EntitlementGetResponseIntegrationTypeNotion, EntitlementGetResponseIntegrationTypeDigitalFiles, EntitlementGetResponseIntegrationTypeLicenseKey:
		return true
	}
	return false
}

type EntitlementUpdateResponse struct {
	ID         string    `json:"id" api:"required"`
	BusinessID string    `json:"business_id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
	// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
	// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
	// ID-only via [`IntegrationConfig`]; this enum is response-only.
	IntegrationConfig EntitlementUpdateResponseIntegrationConfig `json:"integration_config" api:"required"`
	IntegrationType   EntitlementUpdateResponseIntegrationType   `json:"integration_type" api:"required"`
	IsActive          bool                                       `json:"is_active" api:"required"`
	Name              string                                     `json:"name" api:"required"`
	UpdatedAt         time.Time                                  `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                     `json:"description" api:"nullable"`
	Metadata          interface{}                                `json:"metadata"`
	JSON              entitlementUpdateResponseJSON              `json:"-"`
}

// entitlementUpdateResponseJSON contains the JSON metadata for the struct
// [EntitlementUpdateResponse]
type entitlementUpdateResponseJSON struct {
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

func (r *EntitlementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
type EntitlementUpdateResponseIntegrationConfig struct {
	ActivationMessage string `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64  `json:"activations_limit" api:"nullable"`
	ChatID            string `json:"chat_id"`
	// This field can have the runtime type of
	// [EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFiles].
	DigitalFiles     interface{}                                    `json:"digital_files"`
	DurationCount    int64                                          `json:"duration_count" api:"nullable"`
	DurationInterval TimeInterval                                   `json:"duration_interval" api:"nullable"`
	FigmaFileID      string                                         `json:"figma_file_id"`
	FramerTemplateID string                                         `json:"framer_template_id"`
	GuildID          string                                         `json:"guild_id"`
	NotionTemplateID string                                         `json:"notion_template_id"`
	Permission       string                                         `json:"permission"`
	RoleID           string                                         `json:"role_id" api:"nullable"`
	TargetID         string                                         `json:"target_id"`
	JSON             entitlementUpdateResponseIntegrationConfigJSON `json:"-"`
	union            EntitlementUpdateResponseIntegrationConfigUnion
}

// entitlementUpdateResponseIntegrationConfigJSON contains the JSON metadata for
// the struct [EntitlementUpdateResponseIntegrationConfig]
type entitlementUpdateResponseIntegrationConfigJSON struct {
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

func (r entitlementUpdateResponseIntegrationConfigJSON) RawJSON() string {
	return r.raw
}

func (r *EntitlementUpdateResponseIntegrationConfig) UnmarshalJSON(data []byte) (err error) {
	*r = EntitlementUpdateResponseIntegrationConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EntitlementUpdateResponseIntegrationConfigUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EntitlementUpdateResponseIntegrationConfigGitHubConfig],
// [EntitlementUpdateResponseIntegrationConfigDiscordConfig],
// [EntitlementUpdateResponseIntegrationConfigTelegramConfig],
// [EntitlementUpdateResponseIntegrationConfigFigmaConfig],
// [EntitlementUpdateResponseIntegrationConfigFramerConfig],
// [EntitlementUpdateResponseIntegrationConfigNotionConfig],
// [EntitlementUpdateResponseIntegrationConfigDigitalFilesConfig],
// [EntitlementUpdateResponseIntegrationConfigLicenseKeyConfig].
func (r EntitlementUpdateResponseIntegrationConfig) AsUnion() EntitlementUpdateResponseIntegrationConfigUnion {
	return r.union
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
//
// Union satisfied by [EntitlementUpdateResponseIntegrationConfigGitHubConfig],
// [EntitlementUpdateResponseIntegrationConfigDiscordConfig],
// [EntitlementUpdateResponseIntegrationConfigTelegramConfig],
// [EntitlementUpdateResponseIntegrationConfigFigmaConfig],
// [EntitlementUpdateResponseIntegrationConfigFramerConfig],
// [EntitlementUpdateResponseIntegrationConfigNotionConfig],
// [EntitlementUpdateResponseIntegrationConfigDigitalFilesConfig] or
// [EntitlementUpdateResponseIntegrationConfigLicenseKeyConfig].
type EntitlementUpdateResponseIntegrationConfigUnion interface {
	implementsEntitlementUpdateResponseIntegrationConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EntitlementUpdateResponseIntegrationConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigGitHubConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigDiscordConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigTelegramConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigFigmaConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigFramerConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigNotionConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigDigitalFilesConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementUpdateResponseIntegrationConfigLicenseKeyConfig{}),
		},
	)
}

type EntitlementUpdateResponseIntegrationConfigGitHubConfig struct {
	Permission string                                                     `json:"permission" api:"required"`
	TargetID   string                                                     `json:"target_id" api:"required"`
	JSON       entitlementUpdateResponseIntegrationConfigGitHubConfigJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigGitHubConfigJSON contains the JSON
// metadata for the struct [EntitlementUpdateResponseIntegrationConfigGitHubConfig]
type entitlementUpdateResponseIntegrationConfigGitHubConfigJSON struct {
	Permission  apijson.Field
	TargetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigGitHubConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigGitHubConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

type EntitlementUpdateResponseIntegrationConfigDiscordConfig struct {
	GuildID string                                                      `json:"guild_id" api:"required"`
	RoleID  string                                                      `json:"role_id" api:"nullable"`
	JSON    entitlementUpdateResponseIntegrationConfigDiscordConfigJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigDiscordConfigJSON contains the JSON
// metadata for the struct
// [EntitlementUpdateResponseIntegrationConfigDiscordConfig]
type entitlementUpdateResponseIntegrationConfigDiscordConfigJSON struct {
	GuildID     apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigDiscordConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigDiscordConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigDiscordConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

type EntitlementUpdateResponseIntegrationConfigTelegramConfig struct {
	ChatID string                                                       `json:"chat_id" api:"required"`
	JSON   entitlementUpdateResponseIntegrationConfigTelegramConfigJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigTelegramConfigJSON contains the JSON
// metadata for the struct
// [EntitlementUpdateResponseIntegrationConfigTelegramConfig]
type entitlementUpdateResponseIntegrationConfigTelegramConfigJSON struct {
	ChatID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigTelegramConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigTelegramConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigTelegramConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

type EntitlementUpdateResponseIntegrationConfigFigmaConfig struct {
	FigmaFileID string                                                    `json:"figma_file_id" api:"required"`
	JSON        entitlementUpdateResponseIntegrationConfigFigmaConfigJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigFigmaConfigJSON contains the JSON
// metadata for the struct [EntitlementUpdateResponseIntegrationConfigFigmaConfig]
type entitlementUpdateResponseIntegrationConfigFigmaConfigJSON struct {
	FigmaFileID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigFigmaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigFigmaConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigFigmaConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

type EntitlementUpdateResponseIntegrationConfigFramerConfig struct {
	FramerTemplateID string                                                     `json:"framer_template_id" api:"required"`
	JSON             entitlementUpdateResponseIntegrationConfigFramerConfigJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigFramerConfigJSON contains the JSON
// metadata for the struct [EntitlementUpdateResponseIntegrationConfigFramerConfig]
type entitlementUpdateResponseIntegrationConfigFramerConfigJSON struct {
	FramerTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigFramerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigFramerConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigFramerConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

type EntitlementUpdateResponseIntegrationConfigNotionConfig struct {
	NotionTemplateID string                                                     `json:"notion_template_id" api:"required"`
	JSON             entitlementUpdateResponseIntegrationConfigNotionConfigJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigNotionConfigJSON contains the JSON
// metadata for the struct [EntitlementUpdateResponseIntegrationConfigNotionConfig]
type entitlementUpdateResponseIntegrationConfigNotionConfigJSON struct {
	NotionTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigNotionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigNotionConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigNotionConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

type EntitlementUpdateResponseIntegrationConfigDigitalFilesConfig struct {
	// Populated digital-files payload for entitlement read surfaces. Mirrors
	// `DigitalProductDelivery` but is sourced from an entitlement's
	// `integration_config` (not a grant) and tags each file with its origin (`legacy`
	// vs `ee`).
	DigitalFiles EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFiles `json:"digital_files" api:"required"`
	JSON         entitlementUpdateResponseIntegrationConfigDigitalFilesConfigJSON         `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigDigitalFilesConfigJSON contains the
// JSON metadata for the struct
// [EntitlementUpdateResponseIntegrationConfigDigitalFilesConfig]
type entitlementUpdateResponseIntegrationConfigDigitalFilesConfigJSON struct {
	DigitalFiles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigDigitalFilesConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigDigitalFilesConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigDigitalFilesConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

// Populated digital-files payload for entitlement read surfaces. Mirrors
// `DigitalProductDelivery` but is sourced from an entitlement's
// `integration_config` (not a grant) and tags each file with its origin (`legacy`
// vs `ee`).
type EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFiles struct {
	Files        []EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile `json:"files" api:"required"`
	ExternalURL  string                                                                         `json:"external_url" api:"nullable"`
	Instructions string                                                                         `json:"instructions" api:"nullable"`
	JSON         entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON   `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON
// contains the JSON metadata for the struct
// [EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFiles]
type entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON struct {
	Files        apijson.Field
	ExternalURL  apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON) RawJSON() string {
	return r.raw
}

type EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile struct {
	DownloadURL string `json:"download_url" api:"required"`
	// Seconds until `download_url` expires.
	ExpiresIn int64  `json:"expires_in" api:"required"`
	FileID    string `json:"file_id" api:"required"`
	Filename  string `json:"filename" api:"required"`
	// `"legacy"` for files in `product_files`, `"ee"` for files managed by the
	// Entitlements Engine.
	Source      string                                                                           `json:"source" api:"required"`
	ContentType string                                                                           `json:"content_type" api:"nullable"`
	FileSize    int64                                                                            `json:"file_size" api:"nullable"`
	JSON        entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON
// contains the JSON metadata for the struct
// [EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile]
type entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON struct {
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

func (r *EntitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON) RawJSON() string {
	return r.raw
}

type EntitlementUpdateResponseIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage string                                                         `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64                                                          `json:"activations_limit" api:"nullable"`
	DurationCount     int64                                                          `json:"duration_count" api:"nullable"`
	DurationInterval  TimeInterval                                                   `json:"duration_interval" api:"nullable"`
	JSON              entitlementUpdateResponseIntegrationConfigLicenseKeyConfigJSON `json:"-"`
}

// entitlementUpdateResponseIntegrationConfigLicenseKeyConfigJSON contains the JSON
// metadata for the struct
// [EntitlementUpdateResponseIntegrationConfigLicenseKeyConfig]
type entitlementUpdateResponseIntegrationConfigLicenseKeyConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EntitlementUpdateResponseIntegrationConfigLicenseKeyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementUpdateResponseIntegrationConfigLicenseKeyConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementUpdateResponseIntegrationConfigLicenseKeyConfig) implementsEntitlementUpdateResponseIntegrationConfig() {
}

type EntitlementUpdateResponseIntegrationType string

const (
	EntitlementUpdateResponseIntegrationTypeDiscord      EntitlementUpdateResponseIntegrationType = "discord"
	EntitlementUpdateResponseIntegrationTypeTelegram     EntitlementUpdateResponseIntegrationType = "telegram"
	EntitlementUpdateResponseIntegrationTypeGitHub       EntitlementUpdateResponseIntegrationType = "github"
	EntitlementUpdateResponseIntegrationTypeFigma        EntitlementUpdateResponseIntegrationType = "figma"
	EntitlementUpdateResponseIntegrationTypeFramer       EntitlementUpdateResponseIntegrationType = "framer"
	EntitlementUpdateResponseIntegrationTypeNotion       EntitlementUpdateResponseIntegrationType = "notion"
	EntitlementUpdateResponseIntegrationTypeDigitalFiles EntitlementUpdateResponseIntegrationType = "digital_files"
	EntitlementUpdateResponseIntegrationTypeLicenseKey   EntitlementUpdateResponseIntegrationType = "license_key"
)

func (r EntitlementUpdateResponseIntegrationType) IsKnown() bool {
	switch r {
	case EntitlementUpdateResponseIntegrationTypeDiscord, EntitlementUpdateResponseIntegrationTypeTelegram, EntitlementUpdateResponseIntegrationTypeGitHub, EntitlementUpdateResponseIntegrationTypeFigma, EntitlementUpdateResponseIntegrationTypeFramer, EntitlementUpdateResponseIntegrationTypeNotion, EntitlementUpdateResponseIntegrationTypeDigitalFiles, EntitlementUpdateResponseIntegrationTypeLicenseKey:
		return true
	}
	return false
}

type EntitlementListResponse struct {
	ID         string    `json:"id" api:"required"`
	BusinessID string    `json:"business_id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
	// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
	// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
	// ID-only via [`IntegrationConfig`]; this enum is response-only.
	IntegrationConfig EntitlementListResponseIntegrationConfig `json:"integration_config" api:"required"`
	IntegrationType   EntitlementListResponseIntegrationType   `json:"integration_type" api:"required"`
	IsActive          bool                                     `json:"is_active" api:"required"`
	Name              string                                   `json:"name" api:"required"`
	UpdatedAt         time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                   `json:"description" api:"nullable"`
	Metadata          interface{}                              `json:"metadata"`
	JSON              entitlementListResponseJSON              `json:"-"`
}

// entitlementListResponseJSON contains the JSON metadata for the struct
// [EntitlementListResponse]
type entitlementListResponseJSON struct {
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

func (r *EntitlementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseJSON) RawJSON() string {
	return r.raw
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
type EntitlementListResponseIntegrationConfig struct {
	ActivationMessage string `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64  `json:"activations_limit" api:"nullable"`
	ChatID            string `json:"chat_id"`
	// This field can have the runtime type of
	// [EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFiles].
	DigitalFiles     interface{}                                  `json:"digital_files"`
	DurationCount    int64                                        `json:"duration_count" api:"nullable"`
	DurationInterval TimeInterval                                 `json:"duration_interval" api:"nullable"`
	FigmaFileID      string                                       `json:"figma_file_id"`
	FramerTemplateID string                                       `json:"framer_template_id"`
	GuildID          string                                       `json:"guild_id"`
	NotionTemplateID string                                       `json:"notion_template_id"`
	Permission       string                                       `json:"permission"`
	RoleID           string                                       `json:"role_id" api:"nullable"`
	TargetID         string                                       `json:"target_id"`
	JSON             entitlementListResponseIntegrationConfigJSON `json:"-"`
	union            EntitlementListResponseIntegrationConfigUnion
}

// entitlementListResponseIntegrationConfigJSON contains the JSON metadata for the
// struct [EntitlementListResponseIntegrationConfig]
type entitlementListResponseIntegrationConfigJSON struct {
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

func (r entitlementListResponseIntegrationConfigJSON) RawJSON() string {
	return r.raw
}

func (r *EntitlementListResponseIntegrationConfig) UnmarshalJSON(data []byte) (err error) {
	*r = EntitlementListResponseIntegrationConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EntitlementListResponseIntegrationConfigUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EntitlementListResponseIntegrationConfigGitHubConfig],
// [EntitlementListResponseIntegrationConfigDiscordConfig],
// [EntitlementListResponseIntegrationConfigTelegramConfig],
// [EntitlementListResponseIntegrationConfigFigmaConfig],
// [EntitlementListResponseIntegrationConfigFramerConfig],
// [EntitlementListResponseIntegrationConfigNotionConfig],
// [EntitlementListResponseIntegrationConfigDigitalFilesConfig],
// [EntitlementListResponseIntegrationConfigLicenseKeyConfig].
func (r EntitlementListResponseIntegrationConfig) AsUnion() EntitlementListResponseIntegrationConfigUnion {
	return r.union
}

// Public-facing variant of [`IntegrationConfig`]. Mirrors every variant shape on
// the wire EXCEPT `DigitalFiles`, which is replaced with a hydrated
// `digital_files` object (resolved download URLs etc.). The persisted JSONB stays
// ID-only via [`IntegrationConfig`]; this enum is response-only.
//
// Union satisfied by [EntitlementListResponseIntegrationConfigGitHubConfig],
// [EntitlementListResponseIntegrationConfigDiscordConfig],
// [EntitlementListResponseIntegrationConfigTelegramConfig],
// [EntitlementListResponseIntegrationConfigFigmaConfig],
// [EntitlementListResponseIntegrationConfigFramerConfig],
// [EntitlementListResponseIntegrationConfigNotionConfig],
// [EntitlementListResponseIntegrationConfigDigitalFilesConfig] or
// [EntitlementListResponseIntegrationConfigLicenseKeyConfig].
type EntitlementListResponseIntegrationConfigUnion interface {
	implementsEntitlementListResponseIntegrationConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EntitlementListResponseIntegrationConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigGitHubConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigDiscordConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigTelegramConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigFigmaConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigFramerConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigNotionConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigDigitalFilesConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseIntegrationConfigLicenseKeyConfig{}),
		},
	)
}

type EntitlementListResponseIntegrationConfigGitHubConfig struct {
	Permission string                                                   `json:"permission" api:"required"`
	TargetID   string                                                   `json:"target_id" api:"required"`
	JSON       entitlementListResponseIntegrationConfigGitHubConfigJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigGitHubConfigJSON contains the JSON
// metadata for the struct [EntitlementListResponseIntegrationConfigGitHubConfig]
type entitlementListResponseIntegrationConfigGitHubConfigJSON struct {
	Permission  apijson.Field
	TargetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigGitHubConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigGitHubConfig) implementsEntitlementListResponseIntegrationConfig() {
}

type EntitlementListResponseIntegrationConfigDiscordConfig struct {
	GuildID string                                                    `json:"guild_id" api:"required"`
	RoleID  string                                                    `json:"role_id" api:"nullable"`
	JSON    entitlementListResponseIntegrationConfigDiscordConfigJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigDiscordConfigJSON contains the JSON
// metadata for the struct [EntitlementListResponseIntegrationConfigDiscordConfig]
type entitlementListResponseIntegrationConfigDiscordConfigJSON struct {
	GuildID     apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigDiscordConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigDiscordConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigDiscordConfig) implementsEntitlementListResponseIntegrationConfig() {
}

type EntitlementListResponseIntegrationConfigTelegramConfig struct {
	ChatID string                                                     `json:"chat_id" api:"required"`
	JSON   entitlementListResponseIntegrationConfigTelegramConfigJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigTelegramConfigJSON contains the JSON
// metadata for the struct [EntitlementListResponseIntegrationConfigTelegramConfig]
type entitlementListResponseIntegrationConfigTelegramConfigJSON struct {
	ChatID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigTelegramConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigTelegramConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigTelegramConfig) implementsEntitlementListResponseIntegrationConfig() {
}

type EntitlementListResponseIntegrationConfigFigmaConfig struct {
	FigmaFileID string                                                  `json:"figma_file_id" api:"required"`
	JSON        entitlementListResponseIntegrationConfigFigmaConfigJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigFigmaConfigJSON contains the JSON
// metadata for the struct [EntitlementListResponseIntegrationConfigFigmaConfig]
type entitlementListResponseIntegrationConfigFigmaConfigJSON struct {
	FigmaFileID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigFigmaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigFigmaConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigFigmaConfig) implementsEntitlementListResponseIntegrationConfig() {
}

type EntitlementListResponseIntegrationConfigFramerConfig struct {
	FramerTemplateID string                                                   `json:"framer_template_id" api:"required"`
	JSON             entitlementListResponseIntegrationConfigFramerConfigJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigFramerConfigJSON contains the JSON
// metadata for the struct [EntitlementListResponseIntegrationConfigFramerConfig]
type entitlementListResponseIntegrationConfigFramerConfigJSON struct {
	FramerTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigFramerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigFramerConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigFramerConfig) implementsEntitlementListResponseIntegrationConfig() {
}

type EntitlementListResponseIntegrationConfigNotionConfig struct {
	NotionTemplateID string                                                   `json:"notion_template_id" api:"required"`
	JSON             entitlementListResponseIntegrationConfigNotionConfigJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigNotionConfigJSON contains the JSON
// metadata for the struct [EntitlementListResponseIntegrationConfigNotionConfig]
type entitlementListResponseIntegrationConfigNotionConfigJSON struct {
	NotionTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigNotionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigNotionConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigNotionConfig) implementsEntitlementListResponseIntegrationConfig() {
}

type EntitlementListResponseIntegrationConfigDigitalFilesConfig struct {
	// Populated digital-files payload for entitlement read surfaces. Mirrors
	// `DigitalProductDelivery` but is sourced from an entitlement's
	// `integration_config` (not a grant) and tags each file with its origin (`legacy`
	// vs `ee`).
	DigitalFiles EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFiles `json:"digital_files" api:"required"`
	JSON         entitlementListResponseIntegrationConfigDigitalFilesConfigJSON         `json:"-"`
}

// entitlementListResponseIntegrationConfigDigitalFilesConfigJSON contains the JSON
// metadata for the struct
// [EntitlementListResponseIntegrationConfigDigitalFilesConfig]
type entitlementListResponseIntegrationConfigDigitalFilesConfigJSON struct {
	DigitalFiles apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigDigitalFilesConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigDigitalFilesConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigDigitalFilesConfig) implementsEntitlementListResponseIntegrationConfig() {
}

// Populated digital-files payload for entitlement read surfaces. Mirrors
// `DigitalProductDelivery` but is sourced from an entitlement's
// `integration_config` (not a grant) and tags each file with its origin (`legacy`
// vs `ee`).
type EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFiles struct {
	Files        []EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile `json:"files" api:"required"`
	ExternalURL  string                                                                       `json:"external_url" api:"nullable"`
	Instructions string                                                                       `json:"instructions" api:"nullable"`
	JSON         entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON   `json:"-"`
}

// entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON
// contains the JSON metadata for the struct
// [EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFiles]
type entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON struct {
	Files        apijson.Field
	ExternalURL  apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesJSON) RawJSON() string {
	return r.raw
}

type EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile struct {
	DownloadURL string `json:"download_url" api:"required"`
	// Seconds until `download_url` expires.
	ExpiresIn int64  `json:"expires_in" api:"required"`
	FileID    string `json:"file_id" api:"required"`
	Filename  string `json:"filename" api:"required"`
	// `"legacy"` for files in `product_files`, `"ee"` for files managed by the
	// Entitlements Engine.
	Source      string                                                                         `json:"source" api:"required"`
	ContentType string                                                                         `json:"content_type" api:"nullable"`
	FileSize    int64                                                                          `json:"file_size" api:"nullable"`
	JSON        entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON
// contains the JSON metadata for the struct
// [EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile]
type entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON struct {
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

func (r *EntitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigDigitalFilesConfigDigitalFilesFileJSON) RawJSON() string {
	return r.raw
}

type EntitlementListResponseIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage string                                                       `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64                                                        `json:"activations_limit" api:"nullable"`
	DurationCount     int64                                                        `json:"duration_count" api:"nullable"`
	DurationInterval  TimeInterval                                                 `json:"duration_interval" api:"nullable"`
	JSON              entitlementListResponseIntegrationConfigLicenseKeyConfigJSON `json:"-"`
}

// entitlementListResponseIntegrationConfigLicenseKeyConfigJSON contains the JSON
// metadata for the struct
// [EntitlementListResponseIntegrationConfigLicenseKeyConfig]
type entitlementListResponseIntegrationConfigLicenseKeyConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EntitlementListResponseIntegrationConfigLicenseKeyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseIntegrationConfigLicenseKeyConfigJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseIntegrationConfigLicenseKeyConfig) implementsEntitlementListResponseIntegrationConfig() {
}

type EntitlementListResponseIntegrationType string

const (
	EntitlementListResponseIntegrationTypeDiscord      EntitlementListResponseIntegrationType = "discord"
	EntitlementListResponseIntegrationTypeTelegram     EntitlementListResponseIntegrationType = "telegram"
	EntitlementListResponseIntegrationTypeGitHub       EntitlementListResponseIntegrationType = "github"
	EntitlementListResponseIntegrationTypeFigma        EntitlementListResponseIntegrationType = "figma"
	EntitlementListResponseIntegrationTypeFramer       EntitlementListResponseIntegrationType = "framer"
	EntitlementListResponseIntegrationTypeNotion       EntitlementListResponseIntegrationType = "notion"
	EntitlementListResponseIntegrationTypeDigitalFiles EntitlementListResponseIntegrationType = "digital_files"
	EntitlementListResponseIntegrationTypeLicenseKey   EntitlementListResponseIntegrationType = "license_key"
)

func (r EntitlementListResponseIntegrationType) IsKnown() bool {
	switch r {
	case EntitlementListResponseIntegrationTypeDiscord, EntitlementListResponseIntegrationTypeTelegram, EntitlementListResponseIntegrationTypeGitHub, EntitlementListResponseIntegrationTypeFigma, EntitlementListResponseIntegrationTypeFramer, EntitlementListResponseIntegrationTypeNotion, EntitlementListResponseIntegrationTypeDigitalFiles, EntitlementListResponseIntegrationTypeLicenseKey:
		return true
	}
	return false
}

type EntitlementNewParams struct {
	// Platform-specific configuration (validated per integration_type)
	IntegrationConfig param.Field[EntitlementNewParamsIntegrationConfigUnion] `json:"integration_config" api:"required"`
	// Which platform integration this entitlement uses
	IntegrationType param.Field[EntitlementNewParamsIntegrationType] `json:"integration_type" api:"required"`
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

// Platform-specific configuration (validated per integration_type)
type EntitlementNewParamsIntegrationConfig struct {
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

func (r EntitlementNewParamsIntegrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

// Platform-specific configuration (validated per integration_type)
//
// Satisfied by [EntitlementNewParamsIntegrationConfigGitHubConfig],
// [EntitlementNewParamsIntegrationConfigDiscordConfig],
// [EntitlementNewParamsIntegrationConfigTelegramConfig],
// [EntitlementNewParamsIntegrationConfigFigmaConfig],
// [EntitlementNewParamsIntegrationConfigFramerConfig],
// [EntitlementNewParamsIntegrationConfigNotionConfig],
// [EntitlementNewParamsIntegrationConfigDigitalFilesConfig],
// [EntitlementNewParamsIntegrationConfigLicenseKeyConfig],
// [EntitlementNewParamsIntegrationConfig].
type EntitlementNewParamsIntegrationConfigUnion interface {
	implementsEntitlementNewParamsIntegrationConfigUnion()
}

type EntitlementNewParamsIntegrationConfigGitHubConfig struct {
	// One of: pull, push, admin, maintain, triage
	Permission param.Field[string] `json:"permission" api:"required"`
	TargetID   param.Field[string] `json:"target_id" api:"required"`
}

func (r EntitlementNewParamsIntegrationConfigGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigGitHubConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

type EntitlementNewParamsIntegrationConfigDiscordConfig struct {
	GuildID param.Field[string] `json:"guild_id" api:"required"`
	RoleID  param.Field[string] `json:"role_id"`
}

func (r EntitlementNewParamsIntegrationConfigDiscordConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigDiscordConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

type EntitlementNewParamsIntegrationConfigTelegramConfig struct {
	ChatID param.Field[string] `json:"chat_id" api:"required"`
}

func (r EntitlementNewParamsIntegrationConfigTelegramConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigTelegramConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

type EntitlementNewParamsIntegrationConfigFigmaConfig struct {
	FigmaFileID param.Field[string] `json:"figma_file_id" api:"required"`
}

func (r EntitlementNewParamsIntegrationConfigFigmaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigFigmaConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

type EntitlementNewParamsIntegrationConfigFramerConfig struct {
	FramerTemplateID param.Field[string] `json:"framer_template_id" api:"required"`
}

func (r EntitlementNewParamsIntegrationConfigFramerConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigFramerConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

type EntitlementNewParamsIntegrationConfigNotionConfig struct {
	NotionTemplateID param.Field[string] `json:"notion_template_id" api:"required"`
}

func (r EntitlementNewParamsIntegrationConfigNotionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigNotionConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

type EntitlementNewParamsIntegrationConfigDigitalFilesConfig struct {
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

func (r EntitlementNewParamsIntegrationConfigDigitalFilesConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigDigitalFilesConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

type EntitlementNewParamsIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage param.Field[string]       `json:"activation_message"`
	ActivationsLimit  param.Field[int64]        `json:"activations_limit"`
	DurationCount     param.Field[int64]        `json:"duration_count"`
	DurationInterval  param.Field[TimeInterval] `json:"duration_interval"`
}

func (r EntitlementNewParamsIntegrationConfigLicenseKeyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementNewParamsIntegrationConfigLicenseKeyConfig) implementsEntitlementNewParamsIntegrationConfigUnion() {
}

// Which platform integration this entitlement uses
type EntitlementNewParamsIntegrationType string

const (
	EntitlementNewParamsIntegrationTypeDiscord      EntitlementNewParamsIntegrationType = "discord"
	EntitlementNewParamsIntegrationTypeTelegram     EntitlementNewParamsIntegrationType = "telegram"
	EntitlementNewParamsIntegrationTypeGitHub       EntitlementNewParamsIntegrationType = "github"
	EntitlementNewParamsIntegrationTypeFigma        EntitlementNewParamsIntegrationType = "figma"
	EntitlementNewParamsIntegrationTypeFramer       EntitlementNewParamsIntegrationType = "framer"
	EntitlementNewParamsIntegrationTypeNotion       EntitlementNewParamsIntegrationType = "notion"
	EntitlementNewParamsIntegrationTypeDigitalFiles EntitlementNewParamsIntegrationType = "digital_files"
	EntitlementNewParamsIntegrationTypeLicenseKey   EntitlementNewParamsIntegrationType = "license_key"
)

func (r EntitlementNewParamsIntegrationType) IsKnown() bool {
	switch r {
	case EntitlementNewParamsIntegrationTypeDiscord, EntitlementNewParamsIntegrationTypeTelegram, EntitlementNewParamsIntegrationTypeGitHub, EntitlementNewParamsIntegrationTypeFigma, EntitlementNewParamsIntegrationTypeFramer, EntitlementNewParamsIntegrationTypeNotion, EntitlementNewParamsIntegrationTypeDigitalFiles, EntitlementNewParamsIntegrationTypeLicenseKey:
		return true
	}
	return false
}

type EntitlementUpdateParams struct {
	Description param.Field[string] `json:"description"`
	// Platform-specific configuration for an entitlement. Each variant uses unique
	// field names so `#[serde(untagged)]` can disambiguate correctly.
	IntegrationConfig param.Field[EntitlementUpdateParamsIntegrationConfigUnion] `json:"integration_config"`
	Metadata          param.Field[map[string]string]                             `json:"metadata"`
	Name              param.Field[string]                                        `json:"name"`
}

func (r EntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
type EntitlementUpdateParamsIntegrationConfig struct {
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

func (r EntitlementUpdateParamsIntegrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
//
// Satisfied by [EntitlementUpdateParamsIntegrationConfigGitHubConfig],
// [EntitlementUpdateParamsIntegrationConfigDiscordConfig],
// [EntitlementUpdateParamsIntegrationConfigTelegramConfig],
// [EntitlementUpdateParamsIntegrationConfigFigmaConfig],
// [EntitlementUpdateParamsIntegrationConfigFramerConfig],
// [EntitlementUpdateParamsIntegrationConfigNotionConfig],
// [EntitlementUpdateParamsIntegrationConfigDigitalFilesConfig],
// [EntitlementUpdateParamsIntegrationConfigLicenseKeyConfig],
// [EntitlementUpdateParamsIntegrationConfig].
type EntitlementUpdateParamsIntegrationConfigUnion interface {
	implementsEntitlementUpdateParamsIntegrationConfigUnion()
}

type EntitlementUpdateParamsIntegrationConfigGitHubConfig struct {
	// One of: pull, push, admin, maintain, triage
	Permission param.Field[string] `json:"permission" api:"required"`
	TargetID   param.Field[string] `json:"target_id" api:"required"`
}

func (r EntitlementUpdateParamsIntegrationConfigGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigGitHubConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

type EntitlementUpdateParamsIntegrationConfigDiscordConfig struct {
	GuildID param.Field[string] `json:"guild_id" api:"required"`
	RoleID  param.Field[string] `json:"role_id"`
}

func (r EntitlementUpdateParamsIntegrationConfigDiscordConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigDiscordConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

type EntitlementUpdateParamsIntegrationConfigTelegramConfig struct {
	ChatID param.Field[string] `json:"chat_id" api:"required"`
}

func (r EntitlementUpdateParamsIntegrationConfigTelegramConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigTelegramConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

type EntitlementUpdateParamsIntegrationConfigFigmaConfig struct {
	FigmaFileID param.Field[string] `json:"figma_file_id" api:"required"`
}

func (r EntitlementUpdateParamsIntegrationConfigFigmaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigFigmaConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

type EntitlementUpdateParamsIntegrationConfigFramerConfig struct {
	FramerTemplateID param.Field[string] `json:"framer_template_id" api:"required"`
}

func (r EntitlementUpdateParamsIntegrationConfigFramerConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigFramerConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

type EntitlementUpdateParamsIntegrationConfigNotionConfig struct {
	NotionTemplateID param.Field[string] `json:"notion_template_id" api:"required"`
}

func (r EntitlementUpdateParamsIntegrationConfigNotionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigNotionConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

type EntitlementUpdateParamsIntegrationConfigDigitalFilesConfig struct {
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

func (r EntitlementUpdateParamsIntegrationConfigDigitalFilesConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigDigitalFilesConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
}

type EntitlementUpdateParamsIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage param.Field[string]       `json:"activation_message"`
	ActivationsLimit  param.Field[int64]        `json:"activations_limit"`
	DurationCount     param.Field[int64]        `json:"duration_count"`
	DurationInterval  param.Field[TimeInterval] `json:"duration_interval"`
}

func (r EntitlementUpdateParamsIntegrationConfigLicenseKeyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EntitlementUpdateParamsIntegrationConfigLicenseKeyConfig) implementsEntitlementUpdateParamsIntegrationConfigUnion() {
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
