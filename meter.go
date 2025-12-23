// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// MeterService contains methods and other services that help with interacting with
// the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeterService] method instead.
type MeterService struct {
	Options []option.RequestOption
}

// NewMeterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeterService(opts ...option.RequestOption) (r MeterService) {
	r = MeterService{}
	r.Options = opts
	return
}

func (r *MeterService) New(ctx context.Context, body MeterNewParams, opts ...option.RequestOption) (res *Meter, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "meters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *MeterService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Meter, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("meters/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *MeterService) List(ctx context.Context, query MeterListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Meter], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "meters"
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

func (r *MeterService) ListAutoPaging(ctx context.Context, query MeterListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Meter] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *MeterService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("meters/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *MeterService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("meters/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type Meter struct {
	ID              string           `json:"id,required"`
	Aggregation     MeterAggregation `json:"aggregation,required"`
	BusinessID      string           `json:"business_id,required"`
	CreatedAt       time.Time        `json:"created_at,required" format:"date-time"`
	EventName       string           `json:"event_name,required"`
	MeasurementUnit string           `json:"measurement_unit,required"`
	Name            string           `json:"name,required"`
	UpdatedAt       time.Time        `json:"updated_at,required" format:"date-time"`
	Description     string           `json:"description,nullable"`
	// A filter structure that combines multiple conditions with logical conjunctions
	// (AND/OR).
	//
	// Supports up to 3 levels of nesting to create complex filter expressions. Each
	// filter has a conjunction (and/or) and clauses that can be either direct
	// conditions or nested filters.
	Filter MeterFilter `json:"filter,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Aggregation     respjson.Field
		BusinessID      respjson.Field
		CreatedAt       respjson.Field
		EventName       respjson.Field
		MeasurementUnit respjson.Field
		Name            respjson.Field
		UpdatedAt       respjson.Field
		Description     respjson.Field
		Filter          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Meter) RawJSON() string { return r.JSON.raw }
func (r *Meter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeterAggregation struct {
	// Aggregation type for the meter
	//
	// Any of "count", "sum", "max", "last".
	Type MeterAggregationType `json:"type,required"`
	// Required when type is not COUNT
	Key string `json:"key,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterAggregation) RawJSON() string { return r.JSON.raw }
func (r *MeterAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MeterAggregation to a MeterAggregationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MeterAggregationParam.Overrides()
func (r MeterAggregation) ToParam() MeterAggregationParam {
	return param.Override[MeterAggregationParam](json.RawMessage(r.RawJSON()))
}

// Aggregation type for the meter
type MeterAggregationType string

const (
	MeterAggregationTypeCount MeterAggregationType = "count"
	MeterAggregationTypeSum   MeterAggregationType = "sum"
	MeterAggregationTypeMax   MeterAggregationType = "max"
	MeterAggregationTypeLast  MeterAggregationType = "last"
)

// The property Type is required.
type MeterAggregationParam struct {
	// Aggregation type for the meter
	//
	// Any of "count", "sum", "max", "last".
	Type MeterAggregationType `json:"type,omitzero,required"`
	// Required when type is not COUNT
	Key param.Opt[string] `json:"key,omitzero"`
	paramObj
}

func (r MeterAggregationParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterAggregationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterAggregationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter structure that combines multiple conditions with logical conjunctions
// (AND/OR).
//
// Supports up to 3 levels of nesting to create complex filter expressions. Each
// filter has a conjunction (and/or) and clauses that can be either direct
// conditions or nested filters.
type MeterFilter struct {
	// Filter clauses - can be direct conditions or nested filters (up to 3 levels
	// deep)
	Clauses MeterFilterClausesUnion `json:"clauses,required"`
	// Logical conjunction to apply between clauses (and/or)
	//
	// Any of "and", "or".
	Conjunction MeterFilterConjunction `json:"conjunction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clauses     respjson.Field
		Conjunction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilter) RawJSON() string { return r.JSON.raw }
func (r *MeterFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MeterFilter to a MeterFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MeterFilterParam.Overrides()
func (r MeterFilter) ToParam() MeterFilterParam {
	return param.Override[MeterFilterParam](json.RawMessage(r.RawJSON()))
}

// MeterFilterClausesUnion contains all possible properties and values from
// [[]MeterFilterClausesDirectFilterCondition],
// [[]MeterFilterClausesNestedMeterFilter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfDirectFilterConditions OfNestedMeterFilters]
type MeterFilterClausesUnion struct {
	// This field will be present if the value is a
	// [[]MeterFilterClausesDirectFilterCondition] instead of an object.
	OfDirectFilterConditions []MeterFilterClausesDirectFilterCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]MeterFilterClausesNestedMeterFilter] instead of an object.
	OfNestedMeterFilters []MeterFilterClausesNestedMeterFilter `json:",inline"`
	JSON                 struct {
		OfDirectFilterConditions respjson.Field
		OfNestedMeterFilters     respjson.Field
		raw                      string
	} `json:"-"`
}

func (u MeterFilterClausesUnion) AsDirectFilterConditions() (v []MeterFilterClausesDirectFilterCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesUnion) AsNestedMeterFilters() (v []MeterFilterClausesNestedMeterFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeterFilterClausesUnion) RawJSON() string { return u.JSON.raw }

func (r *MeterFilterClausesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter condition with key, operator, and value
type MeterFilterClausesDirectFilterCondition struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesDirectFilterConditionValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilterClausesDirectFilterCondition) RawJSON() string { return r.JSON.raw }
func (r *MeterFilterClausesDirectFilterCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeterFilterClausesDirectFilterConditionValueUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type MeterFilterClausesDirectFilterConditionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u MeterFilterClausesDirectFilterConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesDirectFilterConditionValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesDirectFilterConditionValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeterFilterClausesDirectFilterConditionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MeterFilterClausesDirectFilterConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level 1 nested filter - can contain Level 2 filters
type MeterFilterClausesNestedMeterFilter struct {
	// Level 1: Can be conditions or nested filters (2 more levels allowed)
	Clauses MeterFilterClausesNestedMeterFilterClausesUnion `json:"clauses,required"`
	// Any of "and", "or".
	Conjunction string `json:"conjunction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clauses     respjson.Field
		Conjunction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilterClausesNestedMeterFilter) RawJSON() string { return r.JSON.raw }
func (r *MeterFilterClausesNestedMeterFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeterFilterClausesNestedMeterFilterClausesUnion contains all possible properties
// and values from
// [[]MeterFilterClausesNestedMeterFilterClausesLevel1FilterCondition],
// [[]MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfLevel1FilterConditions OfLevel1NestedFilters]
type MeterFilterClausesNestedMeterFilterClausesUnion struct {
	// This field will be present if the value is a
	// [[]MeterFilterClausesNestedMeterFilterClausesLevel1FilterCondition] instead of
	// an object.
	OfLevel1FilterConditions []MeterFilterClausesNestedMeterFilterClausesLevel1FilterCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilter] instead of an
	// object.
	OfLevel1NestedFilters []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilter `json:",inline"`
	JSON                  struct {
		OfLevel1FilterConditions respjson.Field
		OfLevel1NestedFilters    respjson.Field
		raw                      string
	} `json:"-"`
}

func (u MeterFilterClausesNestedMeterFilterClausesUnion) AsLevel1FilterConditions() (v []MeterFilterClausesNestedMeterFilterClausesLevel1FilterCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesUnion) AsLevel1NestedFilters() (v []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeterFilterClausesNestedMeterFilterClausesUnion) RawJSON() string { return u.JSON.raw }

func (r *MeterFilterClausesNestedMeterFilterClausesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFilterClausesLevel1FilterCondition struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilterClausesNestedMeterFilterClausesLevel1FilterCondition) RawJSON() string {
	return r.JSON.raw
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1FilterCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level 2 nested filter
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilter struct {
	// Level 2: Can be conditions or nested filters (1 more level allowed)
	Clauses MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnion `json:"clauses,required"`
	// Any of "and", "or".
	Conjunction string `json:"conjunction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clauses     respjson.Field
		Conjunction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnion
// contains all possible properties and values from
// [[]MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterCondition],
// [[]MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfLevel2FilterConditions OfLevel2NestedFilters]
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnion struct {
	// This field will be present if the value is a
	// [[]MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterCondition]
	// instead of an object.
	OfLevel2FilterConditions []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilter]
	// instead of an object.
	OfLevel2NestedFilters []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilter `json:",inline"`
	JSON                  struct {
		OfLevel2FilterConditions respjson.Field
		OfLevel2NestedFilters    respjson.Field
		raw                      string
	} `json:"-"`
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnion) AsLevel2FilterConditions() (v []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnion) AsLevel2NestedFilters() (v []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterCondition struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterCondition) RawJSON() string {
	return r.JSON.raw
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level 3 nested filter (final nesting level)
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilter struct {
	// Level 3: Filter conditions only (max depth reached)
	Clauses []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClause `json:"clauses,required"`
	// Any of "and", "or".
	Conjunction string `json:"conjunction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clauses     respjson.Field
		Conjunction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClause struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClause) RawJSON() string {
	return r.JSON.raw
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClause) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logical conjunction to apply between clauses (and/or)
type MeterFilterConjunction string

const (
	MeterFilterConjunctionAnd MeterFilterConjunction = "and"
	MeterFilterConjunctionOr  MeterFilterConjunction = "or"
)

// A filter structure that combines multiple conditions with logical conjunctions
// (AND/OR).
//
// Supports up to 3 levels of nesting to create complex filter expressions. Each
// filter has a conjunction (and/or) and clauses that can be either direct
// conditions or nested filters.
//
// The properties Clauses, Conjunction are required.
type MeterFilterParam struct {
	// Filter clauses - can be direct conditions or nested filters (up to 3 levels
	// deep)
	Clauses MeterFilterClausesUnionParam `json:"clauses,omitzero,required"`
	// Logical conjunction to apply between clauses (and/or)
	//
	// Any of "and", "or".
	Conjunction MeterFilterConjunction `json:"conjunction,omitzero,required"`
	paramObj
}

func (r MeterFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeterFilterClausesUnionParam struct {
	OfDirectFilterConditions []MeterFilterClausesDirectFilterConditionParam `json:",omitzero,inline"`
	OfNestedMeterFilters     []MeterFilterClausesNestedMeterFilterParam     `json:",omitzero,inline"`
	paramUnion
}

func (u MeterFilterClausesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDirectFilterConditions, u.OfNestedMeterFilters)
}
func (u *MeterFilterClausesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeterFilterClausesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDirectFilterConditions) {
		return &u.OfDirectFilterConditions
	} else if !param.IsOmitted(u.OfNestedMeterFilters) {
		return &u.OfNestedMeterFilters
	}
	return nil
}

// Filter condition with key, operator, and value
//
// The properties Key, Operator, Value are required.
type MeterFilterClausesDirectFilterConditionParam struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,omitzero,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesDirectFilterConditionValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r MeterFilterClausesDirectFilterConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterClausesDirectFilterConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterClausesDirectFilterConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeterFilterClausesDirectFilterConditionParam](
		"operator", "equals", "not_equals", "greater_than", "greater_than_or_equals", "less_than", "less_than_or_equals", "contains", "does_not_contain",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeterFilterClausesDirectFilterConditionValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u MeterFilterClausesDirectFilterConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *MeterFilterClausesDirectFilterConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeterFilterClausesDirectFilterConditionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Level 1 nested filter - can contain Level 2 filters
//
// The properties Clauses, Conjunction are required.
type MeterFilterClausesNestedMeterFilterParam struct {
	// Level 1: Can be conditions or nested filters (2 more levels allowed)
	Clauses MeterFilterClausesNestedMeterFilterClausesUnionParam `json:"clauses,omitzero,required"`
	// Any of "and", "or".
	Conjunction string `json:"conjunction,omitzero,required"`
	paramObj
}

func (r MeterFilterClausesNestedMeterFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterClausesNestedMeterFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterClausesNestedMeterFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeterFilterClausesNestedMeterFilterParam](
		"conjunction", "and", "or",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeterFilterClausesNestedMeterFilterClausesUnionParam struct {
	OfLevel1FilterConditions []MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionParam `json:",omitzero,inline"`
	OfLevel1NestedFilters    []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterParam    `json:",omitzero,inline"`
	paramUnion
}

func (u MeterFilterClausesNestedMeterFilterClausesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLevel1FilterConditions, u.OfLevel1NestedFilters)
}
func (u *MeterFilterClausesNestedMeterFilterClausesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeterFilterClausesNestedMeterFilterClausesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfLevel1FilterConditions) {
		return &u.OfLevel1FilterConditions
	} else if !param.IsOmitted(u.OfLevel1NestedFilters) {
		return &u.OfLevel1NestedFilters
	}
	return nil
}

// Filter condition with key, operator, and value
//
// The properties Key, Operator, Value are required.
type MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionParam struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,omitzero,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionParam](
		"operator", "equals", "not_equals", "greater_than", "greater_than_or_equals", "less_than", "less_than_or_equals", "contains", "does_not_contain",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeterFilterClausesNestedMeterFilterClausesLevel1FilterConditionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Level 2 nested filter
//
// The properties Clauses, Conjunction are required.
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterParam struct {
	// Level 2: Can be conditions or nested filters (1 more level allowed)
	Clauses MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnionParam `json:"clauses,omitzero,required"`
	// Any of "and", "or".
	Conjunction string `json:"conjunction,omitzero,required"`
	paramObj
}

func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterParam](
		"conjunction", "and", "or",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnionParam struct {
	OfLevel2FilterConditions []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionParam `json:",omitzero,inline"`
	OfLevel2NestedFilters    []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterParam    `json:",omitzero,inline"`
	paramUnion
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLevel2FilterConditions, u.OfLevel2NestedFilters)
}
func (u *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfLevel2FilterConditions) {
		return &u.OfLevel2FilterConditions
	} else if !param.IsOmitted(u.OfLevel2NestedFilters) {
		return &u.OfLevel2NestedFilters
	}
	return nil
}

// Filter condition with key, operator, and value
//
// The properties Key, Operator, Value are required.
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionParam struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,omitzero,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionParam](
		"operator", "equals", "not_equals", "greater_than", "greater_than_or_equals", "less_than", "less_than_or_equals", "contains", "does_not_contain",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2FilterConditionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Level 3 nested filter (final nesting level)
//
// The properties Clauses, Conjunction are required.
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterParam struct {
	// Level 3: Filter conditions only (max depth reached)
	Clauses []MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseParam `json:"clauses,omitzero,required"`
	// Any of "and", "or".
	Conjunction string `json:"conjunction,omitzero,required"`
	paramObj
}

func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterParam](
		"conjunction", "and", "or",
	)
}

// Filter condition with key, operator, and value
//
// The properties Key, Operator, Value are required.
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseParam struct {
	// Filter key to apply
	Key string `json:"key,required"`
	// Any of "equals", "not_equals", "greater_than", "greater_than_or_equals",
	// "less_than", "less_than_or_equals", "contains", "does_not_contain".
	Operator string `json:"operator,omitzero,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseParam) MarshalJSON() (data []byte, err error) {
	type shadow MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseParam](
		"operator", "equals", "not_equals", "greater_than", "greater_than_or_equals", "less_than", "less_than_or_equals", "contains", "does_not_contain",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeterFilterClausesNestedMeterFilterClausesLevel1NestedFilterClausesLevel2NestedFilterClauseValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type MeterNewParams struct {
	// Aggregation configuration for the meter
	Aggregation MeterAggregationParam `json:"aggregation,omitzero,required"`
	// Event name to track
	EventName string `json:"event_name,required"`
	// measurement unit
	MeasurementUnit string `json:"measurement_unit,required"`
	// Name of the meter
	Name string `json:"name,required"`
	// Optional description of the meter
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional filter to apply to the meter
	Filter MeterFilterParam `json:"filter,omitzero"`
	paramObj
}

func (r MeterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MeterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeterListParams struct {
	// List archived meters
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeterListParams]'s query parameters as `url.Values`.
func (r MeterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
