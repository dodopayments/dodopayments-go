// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// ModerationService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModerationService] method instead.
type ModerationService struct {
	Options []option.RequestOption
}

// NewModerationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModerationService(opts ...option.RequestOption) (r *ModerationService) {
	r = &ModerationService{}
	r.Options = opts
	return
}

// Shows how many billable screens you made and how close you are to your next
// charge.
//
// **Billing.** A billable screen is a live-mode screen that returns a verdict.
// Dodo Payments charges $0.30 for each full block of 1000 billable screens and
// debits the fee from your balance. Each full block is charged within one hour.
// Screens that do not fill a block stay unbilled until they do. Errors and
// test-mode screens are free and are not counted.
func (r *ModerationService) GetUsage(ctx context.Context, opts ...option.RequestOption) (res *ModerationGetUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "moderation/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Screens text, an image, or both, and returns a verdict: `allow`, `flag` or
// `deny`. The API is fail-closed: do not generate when you get no verdict.
//
// **Pricing.** Dodo Payments charges $0.30 per 1000 billable screens and debits
// the fee from your balance. A billable screen is a live-mode screen that returns
// a verdict. Errors and test-mode screens are free.
//
// **429.** Honour `Retry-After` and retry. A 429 is a throughput limit, not a
// verdict.
//
// **Test mode** returns mock verdicts and never calls the model. The default
// verdict is `allow`. Put one of these strings in `text` to select another
// outcome: `dodo_mock_flag` (`flag`), `dodo_mock_deny` (`deny`),
// `dodo_mock_overloaded` (429) or `dodo_mock_not_ready` (503).
func (r *ModerationService) Screen(ctx context.Context, body ModerationScreenParams, opts ...option.RequestOption) (res *ModerationScreenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "moderation/screen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A moderation category.
type ModerationCategory string

const (
	ModerationCategoryViolentCrimes                ModerationCategory = "violent_crimes"
	ModerationCategorySexRelatedCrimes             ModerationCategory = "sex_related_crimes"
	ModerationCategoryChildSexualExploitation      ModerationCategory = "child_sexual_exploitation"
	ModerationCategorySuicideAndSelfHarm           ModerationCategory = "suicide_and_self_harm"
	ModerationCategoryIndiscriminateWeapons        ModerationCategory = "indiscriminate_weapons"
	ModerationCategoryIntellectualProperty         ModerationCategory = "intellectual_property"
	ModerationCategoryDefamation                   ModerationCategory = "defamation"
	ModerationCategoryNonViolentCrimes             ModerationCategory = "non_violent_crimes"
	ModerationCategoryHate                         ModerationCategory = "hate"
	ModerationCategoryPrivacy                      ModerationCategory = "privacy"
	ModerationCategorySpecializedAdvice            ModerationCategory = "specialized_advice"
	ModerationCategorySexualContent                ModerationCategory = "sexual_content"
	ModerationCategoryNonConsensualIntimateImagery ModerationCategory = "non_consensual_intimate_imagery"
	ModerationCategoryMinorCodedLanguage           ModerationCategory = "minor_coded_language"
	ModerationCategoryRealPersonLikeness           ModerationCategory = "real_person_likeness"
	ModerationCategoryLivingArtistStyle            ModerationCategory = "living_artist_style"
	ModerationCategoryPromptInjection              ModerationCategory = "prompt_injection"
)

func (r ModerationCategory) IsKnown() bool {
	switch r {
	case ModerationCategoryViolentCrimes, ModerationCategorySexRelatedCrimes, ModerationCategoryChildSexualExploitation, ModerationCategorySuicideAndSelfHarm, ModerationCategoryIndiscriminateWeapons, ModerationCategoryIntellectualProperty, ModerationCategoryDefamation, ModerationCategoryNonViolentCrimes, ModerationCategoryHate, ModerationCategoryPrivacy, ModerationCategorySpecializedAdvice, ModerationCategorySexualContent, ModerationCategoryNonConsensualIntimateImagery, ModerationCategoryMinorCodedLanguage, ModerationCategoryRealPersonLikeness, ModerationCategoryLivingArtistStyle, ModerationCategoryPromptInjection:
		return true
	}
	return false
}

// How each score in `categories` was measured.
type ModerationCategoryProvenance struct {
	// Child sexual exploitation.
	ChildSexualExploitation ModerationProvenance `json:"child_sexual_exploitation" api:"required"`
	// False depiction that is likely to injure the reputation of a real person.
	Defamation ModerationProvenance `json:"defamation" api:"required"`
	// Demeaning people because of a protected characteristic.
	Hate ModerationProvenance `json:"hate" api:"required"`
	// Chemical, biological, radiological, nuclear or explosive weapons.
	IndiscriminateWeapons ModerationProvenance `json:"indiscriminate_weapons" api:"required"`
	// Copyright or trademark infringement.
	IntellectualProperty ModerationProvenance `json:"intellectual_property" api:"required"`
	// Imitation of the signature style of a specific living artist.
	LivingArtistStyle ModerationProvenance `json:"living_artist_style" api:"required"`
	// Age-coded language that suggests the subject is a minor.
	MinorCodedLanguage ModerationProvenance `json:"minor_coded_language" api:"required"`
	// Non-consensual intimate imagery: undressing, nudifying or sexualising a real
	// person.
	NonConsensualIntimateImagery ModerationProvenance `json:"non_consensual_intimate_imagery" api:"required"`
	// Non-violent crimes.
	NonViolentCrimes ModerationProvenance `json:"non_violent_crimes" api:"required"`
	// Sensitive private information about a person.
	Privacy ModerationProvenance `json:"privacy" api:"required"`
	// An attempt to override or manipulate the instructions of the system.
	PromptInjection ModerationProvenance `json:"prompt_injection" api:"required"`
	// The likeness of a real, identifiable, named person.
	RealPersonLikeness ModerationProvenance `json:"real_person_likeness" api:"required"`
	// Sex-related crimes.
	SexRelatedCrimes ModerationProvenance `json:"sex_related_crimes" api:"required"`
	// Sexually explicit or pornographic content.
	SexualContent ModerationProvenance `json:"sexual_content" api:"required"`
	// Unqualified financial, medical, legal or electoral advice.
	SpecializedAdvice ModerationProvenance `json:"specialized_advice" api:"required"`
	// Suicide and self-harm.
	SuicideAndSelfHarm ModerationProvenance `json:"suicide_and_self_harm" api:"required"`
	// Violent crimes.
	ViolentCrimes ModerationProvenance             `json:"violent_crimes" api:"required"`
	JSON          moderationCategoryProvenanceJSON `json:"-"`
}

// moderationCategoryProvenanceJSON contains the JSON metadata for the struct
// [ModerationCategoryProvenance]
type moderationCategoryProvenanceJSON struct {
	ChildSexualExploitation      apijson.Field
	Defamation                   apijson.Field
	Hate                         apijson.Field
	IndiscriminateWeapons        apijson.Field
	IntellectualProperty         apijson.Field
	LivingArtistStyle            apijson.Field
	MinorCodedLanguage           apijson.Field
	NonConsensualIntimateImagery apijson.Field
	NonViolentCrimes             apijson.Field
	Privacy                      apijson.Field
	PromptInjection              apijson.Field
	RealPersonLikeness           apijson.Field
	SexRelatedCrimes             apijson.Field
	SexualContent                apijson.Field
	SpecializedAdvice            apijson.Field
	SuicideAndSelfHarm           apijson.Field
	ViolentCrimes                apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ModerationCategoryProvenance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moderationCategoryProvenanceJSON) RawJSON() string {
	return r.raw
}

// The probability, from 0 to 1, that the screen falls in each category.
type ModerationCategoryScores struct {
	// Child sexual exploitation.
	ChildSexualExploitation float64 `json:"child_sexual_exploitation" api:"required"`
	// False depiction that is likely to injure the reputation of a real person.
	Defamation float64 `json:"defamation" api:"required"`
	// Demeaning people because of a protected characteristic.
	Hate float64 `json:"hate" api:"required"`
	// Chemical, biological, radiological, nuclear or explosive weapons.
	IndiscriminateWeapons float64 `json:"indiscriminate_weapons" api:"required"`
	// Copyright or trademark infringement.
	IntellectualProperty float64 `json:"intellectual_property" api:"required"`
	// Imitation of the signature style of a specific living artist.
	LivingArtistStyle float64 `json:"living_artist_style" api:"required"`
	// Age-coded language that suggests the subject is a minor.
	MinorCodedLanguage float64 `json:"minor_coded_language" api:"required"`
	// Non-consensual intimate imagery: undressing, nudifying or sexualising a real
	// person.
	NonConsensualIntimateImagery float64 `json:"non_consensual_intimate_imagery" api:"required"`
	// Non-violent crimes.
	NonViolentCrimes float64 `json:"non_violent_crimes" api:"required"`
	// Sensitive private information about a person.
	Privacy float64 `json:"privacy" api:"required"`
	// An attempt to override or manipulate the instructions of the system.
	PromptInjection float64 `json:"prompt_injection" api:"required"`
	// The likeness of a real, identifiable, named person.
	RealPersonLikeness float64 `json:"real_person_likeness" api:"required"`
	// Sex-related crimes.
	SexRelatedCrimes float64 `json:"sex_related_crimes" api:"required"`
	// Sexually explicit or pornographic content.
	SexualContent float64 `json:"sexual_content" api:"required"`
	// Unqualified financial, medical, legal or electoral advice.
	SpecializedAdvice float64 `json:"specialized_advice" api:"required"`
	// Suicide and self-harm.
	SuicideAndSelfHarm float64 `json:"suicide_and_self_harm" api:"required"`
	// Violent crimes.
	ViolentCrimes float64                      `json:"violent_crimes" api:"required"`
	JSON          moderationCategoryScoresJSON `json:"-"`
}

// moderationCategoryScoresJSON contains the JSON metadata for the struct
// [ModerationCategoryScores]
type moderationCategoryScoresJSON struct {
	ChildSexualExploitation      apijson.Field
	Defamation                   apijson.Field
	Hate                         apijson.Field
	IndiscriminateWeapons        apijson.Field
	IntellectualProperty         apijson.Field
	LivingArtistStyle            apijson.Field
	MinorCodedLanguage           apijson.Field
	NonConsensualIntimateImagery apijson.Field
	NonViolentCrimes             apijson.Field
	Privacy                      apijson.Field
	PromptInjection              apijson.Field
	RealPersonLikeness           apijson.Field
	SexRelatedCrimes             apijson.Field
	SexualContent                apijson.Field
	SpecializedAdvice            apijson.Field
	SuicideAndSelfHarm           apijson.Field
	ViolentCrimes                apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ModerationCategoryScores) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moderationCategoryScoresJSON) RawJSON() string {
	return r.raw
}

// The verdict. `allow` means the content passed. `deny` means block the content.
// `flag` means apply your own judgement. It is not a soft deny.
type ModerationDecision string

const (
	ModerationDecisionAllow ModerationDecision = "allow"
	ModerationDecisionFlag  ModerationDecision = "flag"
	ModerationDecisionDeny  ModerationDecision = "deny"
)

func (r ModerationDecision) IsKnown() bool {
	switch r {
	case ModerationDecisionAllow, ModerationDecisionFlag, ModerationDecisionDeny:
		return true
	}
	return false
}

// How a score was measured. `targeted` means a check for that one category
// measured it. `broad` means the general check that covers all categories measured
// it.
type ModerationProvenance string

const (
	ModerationProvenanceTargeted ModerationProvenance = "targeted"
	ModerationProvenanceBroad    ModerationProvenance = "broad"
)

func (r ModerationProvenance) IsKnown() bool {
	switch r {
	case ModerationProvenanceTargeted, ModerationProvenanceBroad:
		return true
	}
	return false
}

// Your moderation usage.
type ModerationGetUsageResponse struct {
	// Your billable screens per UTC day for the last 30 days, charged or not. A day
	// with no screens is not in the list.
	Daily []ModerationGetUsageResponseDaily `json:"daily" api:"required"`
	// Billable screens still needed to fill the next block of 1000. A full block is
	// charged within one hour, so this value is 1000 when your unbilled screens fill
	// whole blocks.
	ScreensToNextBlock int64 `json:"screens_to_next_block" api:"required"`
	// Billable screens that Dodo Payments has not charged for yet.
	UnbilledScreens int64                          `json:"unbilled_screens" api:"required"`
	JSON            moderationGetUsageResponseJSON `json:"-"`
}

// moderationGetUsageResponseJSON contains the JSON metadata for the struct
// [ModerationGetUsageResponse]
type moderationGetUsageResponseJSON struct {
	Daily              apijson.Field
	ScreensToNextBlock apijson.Field
	UnbilledScreens    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ModerationGetUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moderationGetUsageResponseJSON) RawJSON() string {
	return r.raw
}

type ModerationGetUsageResponseDaily struct {
	// The UTC day.
	Date time.Time `json:"date" api:"required" format:"date"`
	// Billable screens on that day.
	Screens int64                               `json:"screens" api:"required"`
	JSON    moderationGetUsageResponseDailyJSON `json:"-"`
}

// moderationGetUsageResponseDailyJSON contains the JSON metadata for the struct
// [ModerationGetUsageResponseDaily]
type moderationGetUsageResponseDailyJSON struct {
	Date        apijson.Field
	Screens     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModerationGetUsageResponseDaily) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moderationGetUsageResponseDailyJSON) RawJSON() string {
	return r.raw
}

// The verdict of one screen.
type ModerationScreenResponse struct {
	// The probability, from 0 to 1, that the screen falls in each category.
	Categories ModerationCategoryScores `json:"categories" api:"required"`
	// True when real-person likeness and sexual content together crossed their
	// combined threshold, the pattern of a sexual deepfake.
	CompoundTriggered bool `json:"compound_triggered" api:"required"`
	// The verdict. `allow` means the content passed. `deny` means block the content.
	// `flag` means apply your own judgement. It is not a soft deny.
	Decision ModerationDecision `json:"decision" api:"required"`
	// The time the screen took, in milliseconds.
	LatencyMs int64 `json:"latency_ms" api:"required"`
	// True when the text was also screened in a normalized form, with obfuscation such
	// as invisible or look-alike characters removed.
	NormalizedApplied bool `json:"normalized_applied" api:"required"`
	// Human-readable reasons for the decision. The wording can change, so do not parse
	// it.
	Notes []string `json:"notes" api:"required"`
	// The number of yes/no questions the model answered for this screen.
	Passes int64 `json:"passes" api:"required"`
	// How each score in `categories` was measured.
	Provenance ModerationCategoryProvenance `json:"provenance" api:"required"`
	// The `request_id` you sent, or null.
	RequestID string `json:"request_id" api:"required,nullable"`
	// The categories whose score crossed the threshold of the category. It can be
	// empty on a `flag` from the general check. `notes` then gives the reason.
	Triggered []ModerationCategory         `json:"triggered" api:"required"`
	JSON      moderationScreenResponseJSON `json:"-"`
}

// moderationScreenResponseJSON contains the JSON metadata for the struct
// [ModerationScreenResponse]
type moderationScreenResponseJSON struct {
	Categories        apijson.Field
	CompoundTriggered apijson.Field
	Decision          apijson.Field
	LatencyMs         apijson.Field
	NormalizedApplied apijson.Field
	Notes             apijson.Field
	Passes            apijson.Field
	Provenance        apijson.Field
	RequestID         apijson.Field
	Triggered         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ModerationScreenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moderationScreenResponseJSON) RawJSON() string {
	return r.raw
}

type ModerationScreenParams struct {
	// The image to screen, as base64, with or without a `data:image/...;base64,`
	// prefix. The formats are JPEG, PNG, WebP, GIF and BMP. The limit is 6991530
	// base64 characters, and the decoded image must be at most 5 MiB.
	Image param.Field[string] `json:"image"`
	// Your identifier for this screen, up to 128 characters, with no control
	// characters. The response returns it in `request_id`.
	RequestID param.Field[string] `json:"request_id"`
	// The text to screen, up to 8000 characters.
	Text param.Field[string] `json:"text"`
}

func (r ModerationScreenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
