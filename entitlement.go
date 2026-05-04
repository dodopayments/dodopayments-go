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

// Detailed view of a single entitlement: identity, integration type,
// integration-specific configuration, and metadata.
type Entitlement struct {
	// Unique identifier of the entitlement.
	ID string `json:"id" api:"required"`
	// Identifier of the business that owns this entitlement.
	BusinessID string `json:"business_id" api:"required"`
	// Timestamp when the entitlement was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Integration-specific configuration. For `digital_files` entitlements this
	// includes presigned download URLs for each attached file.
	IntegrationConfig IntegrationConfigResponse `json:"integration_config" api:"required"`
	// Platform integration this entitlement uses.
	IntegrationType EntitlementIntegrationType `json:"integration_type" api:"required"`
	// Always `true` for entitlements returned by the public API; soft-deleted
	// entitlements are not returned.
	IsActive bool `json:"is_active" api:"required"`
	// Arbitrary key-value metadata supplied at creation or via PATCH.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Display name supplied at creation.
	Name string `json:"name" api:"required"`
	// Timestamp when the entitlement was last modified.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Optional description supplied at creation.
	Description string          `json:"description" api:"nullable"`
	JSON        entitlementJSON `json:"-"`
}

// entitlementJSON contains the JSON metadata for the struct [Entitlement]
type entitlementJSON struct {
	ID                apijson.Field
	BusinessID        apijson.Field
	CreatedAt         apijson.Field
	IntegrationConfig apijson.Field
	IntegrationType   apijson.Field
	IsActive          apijson.Field
	Metadata          apijson.Field
	Name              apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
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

// Integration-specific configuration supplied when creating or updating an
// entitlement. The shape required matches the entitlement's `integration_type`.
type IntegrationConfigParam struct {
	// Optional message displayed when a customer activates the license key (≤ 2500
	// characters).
	ActivationMessage param.Field[string] `json:"activation_message"`
	// Maximum activations allowed per issued license key. Omit for unlimited.
	ActivationsLimit param.Field[int64] `json:"activations_limit"`
	// Telegram chat ID. For groups this is typically a negative integer.
	ChatID         param.Field[string]      `json:"chat_id"`
	DigitalFileIDs param.Field[interface{}] `json:"digital_file_ids"`
	// Validity duration of issued license keys. Provide both `duration_count` and
	// `duration_interval` together for a fixed duration; omit both for non-expiring
	// keys.
	DurationCount param.Field[int64] `json:"duration_count"`
	// Unit of `duration_count`.
	DurationInterval param.Field[TimeInterval] `json:"duration_interval"`
	// Optional external URL shown to the customer alongside the files.
	ExternalURL param.Field[string] `json:"external_url"`
	// Figma file identifier to grant access to.
	FigmaFileID param.Field[string] `json:"figma_file_id"`
	// Framer template identifier to grant access to.
	FramerTemplateID param.Field[string] `json:"framer_template_id"`
	// Discord guild (server) ID.
	GuildID param.Field[string] `json:"guild_id"`
	// Optional human-readable delivery instructions shown to the customer alongside
	// the files.
	Instructions  param.Field[string]      `json:"instructions"`
	LegacyFileIDs param.Field[interface{}] `json:"legacy_file_ids"`
	// Notion template identifier to grant access to.
	NotionTemplateID param.Field[string] `json:"notion_template_id"`
	// Permission to grant on the repository.
	Permission param.Field[IntegrationConfigPermission] `json:"permission"`
	// Optional Discord role to assign within the guild.
	RoleID param.Field[string] `json:"role_id"`
	// Repository or organisation slug to grant access to.
	TargetID param.Field[string] `json:"target_id"`
}

func (r IntegrationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigParam) implementsIntegrationConfigUnionParam() {}

// Integration-specific configuration supplied when creating or updating an
// entitlement. The shape required matches the entitlement's `integration_type`.
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
	// Permission to grant on the repository.
	Permission param.Field[IntegrationConfigGitHubConfigPermission] `json:"permission" api:"required"`
	// Repository or organisation slug to grant access to.
	TargetID param.Field[string] `json:"target_id" api:"required"`
}

func (r IntegrationConfigGitHubConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigGitHubConfigParam) implementsIntegrationConfigUnionParam() {}

// Permission to grant on the repository.
type IntegrationConfigGitHubConfigPermission string

const (
	IntegrationConfigGitHubConfigPermissionPull     IntegrationConfigGitHubConfigPermission = "pull"
	IntegrationConfigGitHubConfigPermissionPush     IntegrationConfigGitHubConfigPermission = "push"
	IntegrationConfigGitHubConfigPermissionAdmin    IntegrationConfigGitHubConfigPermission = "admin"
	IntegrationConfigGitHubConfigPermissionMaintain IntegrationConfigGitHubConfigPermission = "maintain"
	IntegrationConfigGitHubConfigPermissionTriage   IntegrationConfigGitHubConfigPermission = "triage"
)

func (r IntegrationConfigGitHubConfigPermission) IsKnown() bool {
	switch r {
	case IntegrationConfigGitHubConfigPermissionPull, IntegrationConfigGitHubConfigPermissionPush, IntegrationConfigGitHubConfigPermissionAdmin, IntegrationConfigGitHubConfigPermissionMaintain, IntegrationConfigGitHubConfigPermissionTriage:
		return true
	}
	return false
}

type IntegrationConfigDiscordConfigParam struct {
	// Discord guild (server) ID.
	GuildID param.Field[string] `json:"guild_id" api:"required"`
	// Optional Discord role to assign within the guild.
	RoleID param.Field[string] `json:"role_id"`
}

func (r IntegrationConfigDiscordConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigDiscordConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigTelegramConfigParam struct {
	// Telegram chat ID. For groups this is typically a negative integer.
	ChatID param.Field[string] `json:"chat_id" api:"required"`
}

func (r IntegrationConfigTelegramConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigTelegramConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigFigmaConfigParam struct {
	// Figma file identifier to grant access to.
	FigmaFileID param.Field[string] `json:"figma_file_id" api:"required"`
}

func (r IntegrationConfigFigmaConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigFigmaConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigFramerConfigParam struct {
	// Framer template identifier to grant access to.
	FramerTemplateID param.Field[string] `json:"framer_template_id" api:"required"`
}

func (r IntegrationConfigFramerConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigFramerConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigNotionConfigParam struct {
	// Notion template identifier to grant access to.
	NotionTemplateID param.Field[string] `json:"notion_template_id" api:"required"`
}

func (r IntegrationConfigNotionConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigNotionConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigDigitalFilesConfigParam struct {
	// Files attached to this entitlement. Add files via
	// `POST /entitlements/{id}/files` and remove them via
	// `DELETE /entitlements/{id}/files/{file_id}`.
	DigitalFileIDs param.Field[[]string] `json:"digital_file_ids" api:"required"`
	// Optional external URL shown to the customer alongside the files.
	ExternalURL param.Field[string] `json:"external_url"`
	// Optional human-readable delivery instructions shown to the customer alongside
	// the files.
	Instructions param.Field[string] `json:"instructions"`
	// Three-way patchable list of legacy file identifiers:
	//
	// - omitted → preserve the current value
	// - `null` → clear
	// - `[...]` → replace
	//
	// On create, an omitted field, an explicit `null`, or an empty array all result in
	// no legacy files attached.
	LegacyFileIDs param.Field[[]string] `json:"legacy_file_ids"`
}

func (r IntegrationConfigDigitalFilesConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigDigitalFilesConfigParam) implementsIntegrationConfigUnionParam() {}

type IntegrationConfigLicenseKeyConfigParam struct {
	// Optional message displayed when a customer activates the license key (≤ 2500
	// characters).
	ActivationMessage param.Field[string] `json:"activation_message"`
	// Maximum activations allowed per issued license key. Omit for unlimited.
	ActivationsLimit param.Field[int64] `json:"activations_limit"`
	// Validity duration of issued license keys. Provide both `duration_count` and
	// `duration_interval` together for a fixed duration; omit both for non-expiring
	// keys.
	DurationCount param.Field[int64] `json:"duration_count"`
	// Unit of `duration_count`.
	DurationInterval param.Field[TimeInterval] `json:"duration_interval"`
}

func (r IntegrationConfigLicenseKeyConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntegrationConfigLicenseKeyConfigParam) implementsIntegrationConfigUnionParam() {}

// Permission to grant on the repository.
type IntegrationConfigPermission string

const (
	IntegrationConfigPermissionPull     IntegrationConfigPermission = "pull"
	IntegrationConfigPermissionPush     IntegrationConfigPermission = "push"
	IntegrationConfigPermissionAdmin    IntegrationConfigPermission = "admin"
	IntegrationConfigPermissionMaintain IntegrationConfigPermission = "maintain"
	IntegrationConfigPermissionTriage   IntegrationConfigPermission = "triage"
)

func (r IntegrationConfigPermission) IsKnown() bool {
	switch r {
	case IntegrationConfigPermissionPull, IntegrationConfigPermissionPush, IntegrationConfigPermissionAdmin, IntegrationConfigPermissionMaintain, IntegrationConfigPermissionTriage:
		return true
	}
	return false
}

// Integration-specific configuration on an entitlement read response.
//
// For `digital_files` entitlements the response includes presigned download URLs
// for each attached file; other integrations match the shape supplied at creation.
type IntegrationConfigResponse struct {
	// Optional message displayed when a customer activates the license key (≤ 2500
	// characters).
	ActivationMessage string `json:"activation_message" api:"nullable"`
	// Maximum activations allowed per issued license key. Omit for unlimited.
	ActivationsLimit int64 `json:"activations_limit" api:"nullable"`
	// Telegram chat ID. For groups this is typically a negative integer.
	ChatID string `json:"chat_id"`
	// This field can have the runtime type of
	// [IntegrationConfigResponseDigitalFilesConfigDigitalFiles].
	DigitalFiles interface{} `json:"digital_files"`
	// Validity duration of issued license keys. Provide both `duration_count` and
	// `duration_interval` together for a fixed duration; omit both for non-expiring
	// keys.
	DurationCount int64 `json:"duration_count" api:"nullable"`
	// Unit of `duration_count`.
	DurationInterval TimeInterval `json:"duration_interval" api:"nullable"`
	// Figma file identifier to grant access to.
	FigmaFileID string `json:"figma_file_id"`
	// Framer template identifier to grant access to.
	FramerTemplateID string `json:"framer_template_id"`
	// Discord guild (server) ID.
	GuildID string `json:"guild_id"`
	// Notion template identifier to grant access to.
	NotionTemplateID string `json:"notion_template_id"`
	// Permission to grant on the repository.
	Permission IntegrationConfigResponsePermission `json:"permission"`
	// Optional Discord role to assign within the guild.
	RoleID string `json:"role_id" api:"nullable"`
	// Repository or organisation slug to grant access to.
	TargetID string                        `json:"target_id"`
	JSON     integrationConfigResponseJSON `json:"-"`
	union    IntegrationConfigResponseUnion
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

// Integration-specific configuration on an entitlement read response.
//
// For `digital_files` entitlements the response includes presigned download URLs
// for each attached file; other integrations match the shape supplied at creation.
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
	// Permission to grant on the repository.
	Permission IntegrationConfigResponseGitHubConfigPermission `json:"permission" api:"required"`
	// Repository or organisation slug to grant access to.
	TargetID string                                    `json:"target_id" api:"required"`
	JSON     integrationConfigResponseGitHubConfigJSON `json:"-"`
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

// Permission to grant on the repository.
type IntegrationConfigResponseGitHubConfigPermission string

const (
	IntegrationConfigResponseGitHubConfigPermissionPull     IntegrationConfigResponseGitHubConfigPermission = "pull"
	IntegrationConfigResponseGitHubConfigPermissionPush     IntegrationConfigResponseGitHubConfigPermission = "push"
	IntegrationConfigResponseGitHubConfigPermissionAdmin    IntegrationConfigResponseGitHubConfigPermission = "admin"
	IntegrationConfigResponseGitHubConfigPermissionMaintain IntegrationConfigResponseGitHubConfigPermission = "maintain"
	IntegrationConfigResponseGitHubConfigPermissionTriage   IntegrationConfigResponseGitHubConfigPermission = "triage"
)

func (r IntegrationConfigResponseGitHubConfigPermission) IsKnown() bool {
	switch r {
	case IntegrationConfigResponseGitHubConfigPermissionPull, IntegrationConfigResponseGitHubConfigPermissionPush, IntegrationConfigResponseGitHubConfigPermissionAdmin, IntegrationConfigResponseGitHubConfigPermissionMaintain, IntegrationConfigResponseGitHubConfigPermissionTriage:
		return true
	}
	return false
}

type IntegrationConfigResponseDiscordConfig struct {
	// Discord guild (server) ID.
	GuildID string `json:"guild_id" api:"required"`
	// Optional Discord role to assign within the guild.
	RoleID string                                     `json:"role_id" api:"nullable"`
	JSON   integrationConfigResponseDiscordConfigJSON `json:"-"`
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
	// Telegram chat ID. For groups this is typically a negative integer.
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
	// Figma file identifier to grant access to.
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
	// Framer template identifier to grant access to.
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
	// Notion template identifier to grant access to.
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
	// Populated digital-files payload with each file's metadata and a short-lived
	// presigned download URL.
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

// Populated digital-files payload with each file's metadata and a short-lived
// presigned download URL.
type IntegrationConfigResponseDigitalFilesConfigDigitalFiles struct {
	// One entry per attached file.
	Files []IntegrationConfigResponseDigitalFilesConfigDigitalFilesFile `json:"files" api:"required"`
	// Optional external URL, passed through from the entitlement configuration.
	ExternalURL string `json:"external_url" api:"nullable"`
	// Optional human-readable delivery instructions, passed through from the
	// entitlement configuration.
	Instructions string                                                      `json:"instructions" api:"nullable"`
	JSON         integrationConfigResponseDigitalFilesConfigDigitalFilesJSON `json:"-"`
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

// One file in a resolved digital-files payload.
type IntegrationConfigResponseDigitalFilesConfigDigitalFilesFile struct {
	// Short-lived presigned URL for downloading the file.
	DownloadURL string `json:"download_url" api:"required"`
	// Seconds until `download_url` expires.
	ExpiresIn int64 `json:"expires_in" api:"required"`
	// Identifier of the attached file.
	FileID string `json:"file_id" api:"required"`
	// Original filename of the attached file.
	Filename string `json:"filename" api:"required"`
	// Optional content-type declared at upload.
	ContentType string `json:"content_type" api:"nullable"`
	// Optional size of the file in bytes.
	FileSize int64                                                           `json:"file_size" api:"nullable"`
	JSON     integrationConfigResponseDigitalFilesConfigDigitalFilesFileJSON `json:"-"`
}

// integrationConfigResponseDigitalFilesConfigDigitalFilesFileJSON contains the
// JSON metadata for the struct
// [IntegrationConfigResponseDigitalFilesConfigDigitalFilesFile]
type integrationConfigResponseDigitalFilesConfigDigitalFilesFileJSON struct {
	DownloadURL apijson.Field
	ExpiresIn   apijson.Field
	FileID      apijson.Field
	Filename    apijson.Field
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
	// Optional message displayed when a customer activates the license key (≤ 2500
	// characters).
	ActivationMessage string `json:"activation_message" api:"nullable"`
	// Maximum activations allowed per issued license key. Omit for unlimited.
	ActivationsLimit int64 `json:"activations_limit" api:"nullable"`
	// Validity duration of issued license keys. Provide both `duration_count` and
	// `duration_interval` together for a fixed duration; omit both for non-expiring
	// keys.
	DurationCount int64 `json:"duration_count" api:"nullable"`
	// Unit of `duration_count`.
	DurationInterval TimeInterval                                  `json:"duration_interval" api:"nullable"`
	JSON             integrationConfigResponseLicenseKeyConfigJSON `json:"-"`
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

// Permission to grant on the repository.
type IntegrationConfigResponsePermission string

const (
	IntegrationConfigResponsePermissionPull     IntegrationConfigResponsePermission = "pull"
	IntegrationConfigResponsePermissionPush     IntegrationConfigResponsePermission = "push"
	IntegrationConfigResponsePermissionAdmin    IntegrationConfigResponsePermission = "admin"
	IntegrationConfigResponsePermissionMaintain IntegrationConfigResponsePermission = "maintain"
	IntegrationConfigResponsePermissionTriage   IntegrationConfigResponsePermission = "triage"
)

func (r IntegrationConfigResponsePermission) IsKnown() bool {
	switch r {
	case IntegrationConfigResponsePermissionPull, IntegrationConfigResponsePermissionPush, IntegrationConfigResponsePermissionAdmin, IntegrationConfigResponsePermissionMaintain, IntegrationConfigResponsePermissionTriage:
		return true
	}
	return false
}

type EntitlementNewParams struct {
	// Platform-specific configuration (validated per integration_type)
	IntegrationConfig param.Field[IntegrationConfigUnionParam] `json:"integration_config" api:"required"`
	// Which platform integration this entitlement uses
	IntegrationType param.Field[EntitlementIntegrationType] `json:"integration_type" api:"required"`
	// Display name for this entitlement
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description
	Description param.Field[string] `json:"description"`
	// Additional metadata for the entitlement
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r EntitlementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntitlementUpdateParams struct {
	Description param.Field[string] `json:"description"`
	// Integration-specific configuration supplied when creating or updating an
	// entitlement. The shape required matches the entitlement's `integration_type`.
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
