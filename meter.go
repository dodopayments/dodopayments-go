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
	"github.com/dodopayments/dodopayments-go/shared"
	"github.com/tidwall/gjson"
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
func NewMeterService(opts ...option.RequestOption) (r *MeterService) {
	r = &MeterService{}
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

type Conjunction string

const (
	ConjunctionAnd Conjunction = "and"
	ConjunctionOr  Conjunction = "or"
)

func (r Conjunction) IsKnown() bool {
	switch r {
	case ConjunctionAnd, ConjunctionOr:
		return true
	}
	return false
}

type FilterOperator string

const (
	FilterOperatorEquals              FilterOperator = "equals"
	FilterOperatorNotEquals           FilterOperator = "not_equals"
	FilterOperatorGreaterThan         FilterOperator = "greater_than"
	FilterOperatorGreaterThanOrEquals FilterOperator = "greater_than_or_equals"
	FilterOperatorLessThan            FilterOperator = "less_than"
	FilterOperatorLessThanOrEquals    FilterOperator = "less_than_or_equals"
	FilterOperatorContains            FilterOperator = "contains"
	FilterOperatorDoesNotContain      FilterOperator = "does_not_contain"
)

func (r FilterOperator) IsKnown() bool {
	switch r {
	case FilterOperatorEquals, FilterOperatorNotEquals, FilterOperatorGreaterThan, FilterOperatorGreaterThanOrEquals, FilterOperatorLessThan, FilterOperatorLessThanOrEquals, FilterOperatorContains, FilterOperatorDoesNotContain:
		return true
	}
	return false
}

type Meter struct {
	ID              string           `json:"id" api:"required"`
	Aggregation     MeterAggregation `json:"aggregation" api:"required"`
	BusinessID      string           `json:"business_id" api:"required"`
	CreatedAt       time.Time        `json:"created_at" api:"required" format:"date-time"`
	EventName       string           `json:"event_name" api:"required"`
	MeasurementUnit string           `json:"measurement_unit" api:"required"`
	Name            string           `json:"name" api:"required"`
	UpdatedAt       time.Time        `json:"updated_at" api:"required" format:"date-time"`
	Description     string           `json:"description" api:"nullable"`
	// A filter structure that combines multiple conditions with logical conjunctions
	// (AND/OR).
	//
	// Supports up to 3 levels of nesting to create complex filter expressions. Each
	// filter has a conjunction (and/or) and clauses that can be either direct
	// conditions or nested filters.
	Filter MeterFilter `json:"filter" api:"nullable"`
	JSON   meterJSON   `json:"-"`
}

// meterJSON contains the JSON metadata for the struct [Meter]
type meterJSON struct {
	ID              apijson.Field
	Aggregation     apijson.Field
	BusinessID      apijson.Field
	CreatedAt       apijson.Field
	EventName       apijson.Field
	MeasurementUnit apijson.Field
	Name            apijson.Field
	UpdatedAt       apijson.Field
	Description     apijson.Field
	Filter          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Meter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterJSON) RawJSON() string {
	return r.raw
}

type MeterAggregation struct {
	// Aggregation type for the meter
	Type MeterAggregationType `json:"type" api:"required"`
	// Required when type is not COUNT
	Key  string               `json:"key" api:"nullable"`
	JSON meterAggregationJSON `json:"-"`
}

// meterAggregationJSON contains the JSON metadata for the struct
// [MeterAggregation]
type meterAggregationJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterAggregation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterAggregationJSON) RawJSON() string {
	return r.raw
}

// Aggregation type for the meter
type MeterAggregationType string

const (
	MeterAggregationTypeCount MeterAggregationType = "count"
	MeterAggregationTypeSum   MeterAggregationType = "sum"
	MeterAggregationTypeMax   MeterAggregationType = "max"
	MeterAggregationTypeLast  MeterAggregationType = "last"
)

func (r MeterAggregationType) IsKnown() bool {
	switch r {
	case MeterAggregationTypeCount, MeterAggregationTypeSum, MeterAggregationTypeMax, MeterAggregationTypeLast:
		return true
	}
	return false
}

type MeterAggregationParam struct {
	// Aggregation type for the meter
	Type param.Field[MeterAggregationType] `json:"type" api:"required"`
	// Required when type is not COUNT
	Key param.Field[string] `json:"key"`
}

func (r MeterAggregationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Clauses MeterFilterClausesUnion `json:"clauses" api:"required"`
	// Logical conjunction to apply between clauses (and/or)
	Conjunction Conjunction     `json:"conjunction" api:"required"`
	JSON        meterFilterJSON `json:"-"`
}

// meterFilterJSON contains the JSON metadata for the struct [MeterFilter]
type meterFilterJSON struct {
	Clauses     apijson.Field
	Conjunction apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterJSON) RawJSON() string {
	return r.raw
}

// Filter clauses - can be direct conditions or nested filters (up to 3 levels
// deep)
//
// Union satisfied by [MeterFilterClausesDirectFilterConditions] or
// [MeterFilterClausesNestedMeterFilters].
type MeterFilterClausesUnion interface {
	implementsMeterFilterClausesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterFilterClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MeterFilterClausesDirectFilterConditions{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MeterFilterClausesNestedMeterFilters{}),
		},
	)
}

type MeterFilterClausesDirectFilterConditions []MeterFilterClausesDirectFilterCondition

func (r MeterFilterClausesDirectFilterConditions) implementsMeterFilterClausesUnion() {}

// Filter condition with key, operator, and value
type MeterFilterClausesDirectFilterCondition struct {
	// Filter key to apply
	Key      string         `json:"key" api:"required"`
	Operator FilterOperator `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesDirectFilterConditionsValueUnion `json:"value" api:"required"`
	JSON  meterFilterClausesDirectFilterConditionJSON        `json:"-"`
}

// meterFilterClausesDirectFilterConditionJSON contains the JSON metadata for the
// struct [MeterFilterClausesDirectFilterCondition]
type meterFilterClausesDirectFilterConditionJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilterClausesDirectFilterCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterClausesDirectFilterConditionJSON) RawJSON() string {
	return r.raw
}

// Filter value - can be string, number, or boolean
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type MeterFilterClausesDirectFilterConditionsValueUnion interface {
	ImplementsMeterFilterClausesDirectFilterConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterFilterClausesDirectFilterConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type MeterFilterClausesNestedMeterFilters []MeterFilterClausesNestedMeterFilter

func (r MeterFilterClausesNestedMeterFilters) implementsMeterFilterClausesUnion() {}

// Level 1 nested filter - can contain Level 2 filters
type MeterFilterClausesNestedMeterFilter struct {
	// Level 1: Can be conditions or nested filters (2 more levels allowed)
	Clauses     MeterFilterClausesNestedMeterFiltersClausesUnion `json:"clauses" api:"required"`
	Conjunction Conjunction                                      `json:"conjunction" api:"required"`
	JSON        meterFilterClausesNestedMeterFilterJSON          `json:"-"`
}

// meterFilterClausesNestedMeterFilterJSON contains the JSON metadata for the
// struct [MeterFilterClausesNestedMeterFilter]
type meterFilterClausesNestedMeterFilterJSON struct {
	Clauses     apijson.Field
	Conjunction apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilterClausesNestedMeterFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterClausesNestedMeterFilterJSON) RawJSON() string {
	return r.raw
}

// Level 1: Can be conditions or nested filters (2 more levels allowed)
//
// Union satisfied by
// [MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditions] or
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilters].
type MeterFilterClausesNestedMeterFiltersClausesUnion interface {
	implementsMeterFilterClausesNestedMeterFiltersClausesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterFilterClausesNestedMeterFiltersClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditions{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilters{}),
		},
	)
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditions []MeterFilterClausesNestedMeterFiltersClausesLevel1FilterCondition

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditions) implementsMeterFilterClausesNestedMeterFiltersClausesUnion() {
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFiltersClausesLevel1FilterCondition struct {
	// Filter key to apply
	Key      string         `json:"key" api:"required"`
	Operator FilterOperator `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnion `json:"value" api:"required"`
	JSON  meterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionJSON        `json:"-"`
}

// meterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionJSON contains
// the JSON metadata for the struct
// [MeterFilterClausesNestedMeterFiltersClausesLevel1FilterCondition]
type meterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilterClausesNestedMeterFiltersClausesLevel1FilterCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionJSON) RawJSON() string {
	return r.raw
}

// Filter value - can be string, number, or boolean
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnion interface {
	ImplementsMeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilters []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilter

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilters) implementsMeterFilterClausesNestedMeterFiltersClausesUnion() {
}

// Level 2 nested filter
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilter struct {
	// Level 2: Can be conditions or nested filters (1 more level allowed)
	Clauses     MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnion `json:"clauses" api:"required"`
	Conjunction Conjunction                                                                `json:"conjunction" api:"required"`
	JSON        meterFilterClausesNestedMeterFiltersClausesLevel1NestedFilterJSON          `json:"-"`
}

// meterFilterClausesNestedMeterFiltersClausesLevel1NestedFilterJSON contains the
// JSON metadata for the struct
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilter]
type meterFilterClausesNestedMeterFiltersClausesLevel1NestedFilterJSON struct {
	Clauses     apijson.Field
	Conjunction apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterClausesNestedMeterFiltersClausesLevel1NestedFilterJSON) RawJSON() string {
	return r.raw
}

// Level 2: Can be conditions or nested filters (1 more level allowed)
//
// Union satisfied by
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditions]
// or
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilters].
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnion interface {
	implementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditions{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilters{}),
		},
	)
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditions []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterCondition

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditions) implementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnion() {
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterCondition struct {
	// Filter key to apply
	Key      string         `json:"key" api:"required"`
	Operator FilterOperator `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnion `json:"value" api:"required"`
	JSON  meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionJSON        `json:"-"`
}

// meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionJSON
// contains the JSON metadata for the struct
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterCondition]
type meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionJSON) RawJSON() string {
	return r.raw
}

// Filter value - can be string, number, or boolean
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnion interface {
	ImplementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilters []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilter

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilters) implementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnion() {
}

// Level 3 nested filter (final nesting level)
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilter struct {
	// Level 3: Filter conditions only (max depth reached)
	Clauses     []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClause `json:"clauses" api:"required"`
	Conjunction Conjunction                                                                                      `json:"conjunction" api:"required"`
	JSON        meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterJSON      `json:"-"`
}

// meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterJSON
// contains the JSON metadata for the struct
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilter]
type meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterJSON struct {
	Clauses     apijson.Field
	Conjunction apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterJSON) RawJSON() string {
	return r.raw
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClause struct {
	// Filter key to apply
	Key      string         `json:"key" api:"required"`
	Operator FilterOperator `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnion `json:"value" api:"required"`
	JSON  meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseJSON        `json:"-"`
}

// meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseJSON
// contains the JSON metadata for the struct
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClause]
type meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClause) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseJSON) RawJSON() string {
	return r.raw
}

// Filter value - can be string, number, or boolean
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnion interface {
	ImplementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// A filter structure that combines multiple conditions with logical conjunctions
// (AND/OR).
//
// Supports up to 3 levels of nesting to create complex filter expressions. Each
// filter has a conjunction (and/or) and clauses that can be either direct
// conditions or nested filters.
type MeterFilterParam struct {
	// Filter clauses - can be direct conditions or nested filters (up to 3 levels
	// deep)
	Clauses param.Field[MeterFilterClausesUnionParam] `json:"clauses" api:"required"`
	// Logical conjunction to apply between clauses (and/or)
	Conjunction param.Field[Conjunction] `json:"conjunction" api:"required"`
}

func (r MeterFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter clauses - can be direct conditions or nested filters (up to 3 levels
// deep)
//
// Satisfied by [MeterFilterClausesDirectFilterConditionsParam],
// [MeterFilterClausesNestedMeterFiltersParam].
type MeterFilterClausesUnionParam interface {
	implementsMeterFilterClausesUnionParam()
}

type MeterFilterClausesDirectFilterConditionsParam []MeterFilterClausesDirectFilterConditionParam

func (r MeterFilterClausesDirectFilterConditionsParam) implementsMeterFilterClausesUnionParam() {}

// Filter condition with key, operator, and value
type MeterFilterClausesDirectFilterConditionParam struct {
	// Filter key to apply
	Key      param.Field[string]         `json:"key" api:"required"`
	Operator param.Field[FilterOperator] `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesDirectFilterConditionsValueUnionParam] `json:"value" api:"required"`
}

func (r MeterFilterClausesDirectFilterConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter value - can be string, number, or boolean
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type MeterFilterClausesDirectFilterConditionsValueUnionParam interface {
	ImplementsMeterFilterClausesDirectFilterConditionsValueUnionParam()
}

type MeterFilterClausesNestedMeterFiltersParam []MeterFilterClausesNestedMeterFilterParam

func (r MeterFilterClausesNestedMeterFiltersParam) implementsMeterFilterClausesUnionParam() {}

// Level 1 nested filter - can contain Level 2 filters
type MeterFilterClausesNestedMeterFilterParam struct {
	// Level 1: Can be conditions or nested filters (2 more levels allowed)
	Clauses     param.Field[MeterFilterClausesNestedMeterFiltersClausesUnionParam] `json:"clauses" api:"required"`
	Conjunction param.Field[Conjunction]                                           `json:"conjunction" api:"required"`
}

func (r MeterFilterClausesNestedMeterFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Level 1: Can be conditions or nested filters (2 more levels allowed)
//
// Satisfied by
// [MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsParam],
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersParam].
type MeterFilterClausesNestedMeterFiltersClausesUnionParam interface {
	implementsMeterFilterClausesNestedMeterFiltersClausesUnionParam()
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsParam []MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionParam

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsParam) implementsMeterFilterClausesNestedMeterFiltersClausesUnionParam() {
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionParam struct {
	// Filter key to apply
	Key      param.Field[string]         `json:"key" api:"required"`
	Operator param.Field[FilterOperator] `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnionParam] `json:"value" api:"required"`
}

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter value - can be string, number, or boolean
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnionParam interface {
	ImplementsMeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnionParam()
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersParam []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilterParam

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersParam) implementsMeterFilterClausesNestedMeterFiltersClausesUnionParam() {
}

// Level 2 nested filter
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilterParam struct {
	// Level 2: Can be conditions or nested filters (1 more level allowed)
	Clauses     param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnionParam] `json:"clauses" api:"required"`
	Conjunction param.Field[Conjunction]                                                                     `json:"conjunction" api:"required"`
}

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Level 2: Can be conditions or nested filters (1 more level allowed)
//
// Satisfied by
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsParam],
// [MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersParam].
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnionParam interface {
	implementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnionParam()
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsParam []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionParam

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsParam) implementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnionParam() {
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionParam struct {
	// Filter key to apply
	Key      param.Field[string]         `json:"key" api:"required"`
	Operator param.Field[FilterOperator] `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnionParam] `json:"value" api:"required"`
}

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter value - can be string, number, or boolean
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnionParam interface {
	ImplementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnionParam()
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersParam []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterParam

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersParam) implementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnionParam() {
}

// Level 3 nested filter (final nesting level)
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterParam struct {
	// Level 3: Filter conditions only (max depth reached)
	Clauses     param.Field[[]MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseParam] `json:"clauses" api:"required"`
	Conjunction param.Field[Conjunction]                                                                                           `json:"conjunction" api:"required"`
}

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseParam struct {
	// Filter key to apply
	Key      param.Field[string]         `json:"key" api:"required"`
	Operator param.Field[FilterOperator] `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnionParam] `json:"value" api:"required"`
}

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter value - can be string, number, or boolean
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnionParam interface {
	ImplementsMeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnionParam()
}

type MeterNewParams struct {
	// Aggregation configuration for the meter
	Aggregation param.Field[MeterAggregationParam] `json:"aggregation" api:"required"`
	// Event name to track
	EventName param.Field[string] `json:"event_name" api:"required"`
	// measurement unit
	MeasurementUnit param.Field[string] `json:"measurement_unit" api:"required"`
	// Name of the meter
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description of the meter
	Description param.Field[string] `json:"description"`
	// Optional filter to apply to the meter
	Filter param.Field[MeterFilterParam] `json:"filter"`
}

func (r MeterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeterListParams struct {
	// List archived meters
	Archived param.Field[bool] `query:"archived"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [MeterListParams]'s query parameters as `url.Values`.
func (r MeterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
