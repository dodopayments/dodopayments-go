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
	Type MeterAggregationType `json:"type,required"`
	// Required when type is not COUNT
	Key  string               `json:"key,nullable"`
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
	Type param.Field[MeterAggregationType] `json:"type,required"`
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
	Clauses MeterFilterClausesUnion `json:"clauses,required"`
	// Logical conjunction to apply between clauses (and/or)
	Conjunction MeterFilterConjunction `json:"conjunction,required"`
	JSON        meterFilterJSON        `json:"-"`
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
	Key      string                                           `json:"key,required"`
	Operator MeterFilterClausesDirectFilterConditionsOperator `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesDirectFilterConditionsValueUnion `json:"value,required"`
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

type MeterFilterClausesDirectFilterConditionsOperator string

const (
	MeterFilterClausesDirectFilterConditionsOperatorEquals              MeterFilterClausesDirectFilterConditionsOperator = "equals"
	MeterFilterClausesDirectFilterConditionsOperatorNotEquals           MeterFilterClausesDirectFilterConditionsOperator = "not_equals"
	MeterFilterClausesDirectFilterConditionsOperatorGreaterThan         MeterFilterClausesDirectFilterConditionsOperator = "greater_than"
	MeterFilterClausesDirectFilterConditionsOperatorGreaterThanOrEquals MeterFilterClausesDirectFilterConditionsOperator = "greater_than_or_equals"
	MeterFilterClausesDirectFilterConditionsOperatorLessThan            MeterFilterClausesDirectFilterConditionsOperator = "less_than"
	MeterFilterClausesDirectFilterConditionsOperatorLessThanOrEquals    MeterFilterClausesDirectFilterConditionsOperator = "less_than_or_equals"
	MeterFilterClausesDirectFilterConditionsOperatorContains            MeterFilterClausesDirectFilterConditionsOperator = "contains"
	MeterFilterClausesDirectFilterConditionsOperatorDoesNotContain      MeterFilterClausesDirectFilterConditionsOperator = "does_not_contain"
)

func (r MeterFilterClausesDirectFilterConditionsOperator) IsKnown() bool {
	switch r {
	case MeterFilterClausesDirectFilterConditionsOperatorEquals, MeterFilterClausesDirectFilterConditionsOperatorNotEquals, MeterFilterClausesDirectFilterConditionsOperatorGreaterThan, MeterFilterClausesDirectFilterConditionsOperatorGreaterThanOrEquals, MeterFilterClausesDirectFilterConditionsOperatorLessThan, MeterFilterClausesDirectFilterConditionsOperatorLessThanOrEquals, MeterFilterClausesDirectFilterConditionsOperatorContains, MeterFilterClausesDirectFilterConditionsOperatorDoesNotContain:
		return true
	}
	return false
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
	Clauses     MeterFilterClausesNestedMeterFiltersClausesUnion `json:"clauses,required"`
	Conjunction MeterFilterClausesNestedMeterFiltersConjunction  `json:"conjunction,required"`
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
	Key      string                                                                    `json:"key,required"`
	Operator MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnion `json:"value,required"`
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

type MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator string

const (
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorEquals              MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorNotEquals           MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "not_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorGreaterThan         MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "greater_than"
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorGreaterThanOrEquals MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "greater_than_or_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorLessThan            MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "less_than"
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorLessThanOrEquals    MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "less_than_or_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorContains            MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "contains"
	MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorDoesNotContain      MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator = "does_not_contain"
)

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator) IsKnown() bool {
	switch r {
	case MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorNotEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorGreaterThan, MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorGreaterThanOrEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorLessThan, MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorLessThanOrEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorContains, MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperatorDoesNotContain:
		return true
	}
	return false
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
	Clauses     MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnion `json:"clauses,required"`
	Conjunction MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunction  `json:"conjunction,required"`
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
	Key      string                                                                                              `json:"key,required"`
	Operator MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnion `json:"value,required"`
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

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator string

const (
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorEquals              MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorNotEquals           MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "not_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorGreaterThan         MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "greater_than"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorGreaterThanOrEquals MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "greater_than_or_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorLessThan            MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "less_than"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorLessThanOrEquals    MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "less_than_or_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorContains            MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "contains"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorDoesNotContain      MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator = "does_not_contain"
)

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator) IsKnown() bool {
	switch r {
	case MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorNotEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorGreaterThan, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorGreaterThanOrEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorLessThan, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorLessThanOrEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorContains, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperatorDoesNotContain:
		return true
	}
	return false
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
	Clauses     []MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClause    `json:"clauses,required"`
	Conjunction MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunction `json:"conjunction,required"`
	JSON        meterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterJSON         `json:"-"`
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
	Key      string                                                                                                  `json:"key,required"`
	Operator MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnion `json:"value,required"`
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

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator string

const (
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorEquals              MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorNotEquals           MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "not_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorGreaterThan         MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "greater_than"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorGreaterThanOrEquals MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "greater_than_or_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorLessThan            MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "less_than"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorLessThanOrEquals    MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "less_than_or_equals"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorContains            MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "contains"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorDoesNotContain      MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator = "does_not_contain"
)

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator) IsKnown() bool {
	switch r {
	case MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorNotEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorGreaterThan, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorGreaterThanOrEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorLessThan, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorLessThanOrEquals, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorContains, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperatorDoesNotContain:
		return true
	}
	return false
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

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunction string

const (
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunctionAnd MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunction = "and"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunctionOr  MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunction = "or"
)

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunction) IsKnown() bool {
	switch r {
	case MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunctionAnd, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunctionOr:
		return true
	}
	return false
}

type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunction string

const (
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunctionAnd MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunction = "and"
	MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunctionOr  MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunction = "or"
)

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunction) IsKnown() bool {
	switch r {
	case MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunctionAnd, MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunctionOr:
		return true
	}
	return false
}

type MeterFilterClausesNestedMeterFiltersConjunction string

const (
	MeterFilterClausesNestedMeterFiltersConjunctionAnd MeterFilterClausesNestedMeterFiltersConjunction = "and"
	MeterFilterClausesNestedMeterFiltersConjunctionOr  MeterFilterClausesNestedMeterFiltersConjunction = "or"
)

func (r MeterFilterClausesNestedMeterFiltersConjunction) IsKnown() bool {
	switch r {
	case MeterFilterClausesNestedMeterFiltersConjunctionAnd, MeterFilterClausesNestedMeterFiltersConjunctionOr:
		return true
	}
	return false
}

// Logical conjunction to apply between clauses (and/or)
type MeterFilterConjunction string

const (
	MeterFilterConjunctionAnd MeterFilterConjunction = "and"
	MeterFilterConjunctionOr  MeterFilterConjunction = "or"
)

func (r MeterFilterConjunction) IsKnown() bool {
	switch r {
	case MeterFilterConjunctionAnd, MeterFilterConjunctionOr:
		return true
	}
	return false
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
	Clauses param.Field[MeterFilterClausesUnionParam] `json:"clauses,required"`
	// Logical conjunction to apply between clauses (and/or)
	Conjunction param.Field[MeterFilterConjunction] `json:"conjunction,required"`
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
	Key      param.Field[string]                                           `json:"key,required"`
	Operator param.Field[MeterFilterClausesDirectFilterConditionsOperator] `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesDirectFilterConditionsValueUnionParam] `json:"value,required"`
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
	Clauses     param.Field[MeterFilterClausesNestedMeterFiltersClausesUnionParam] `json:"clauses,required"`
	Conjunction param.Field[MeterFilterClausesNestedMeterFiltersConjunction]       `json:"conjunction,required"`
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
	Key      param.Field[string]                                                                    `json:"key,required"`
	Operator param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsOperator] `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1FilterConditionsValueUnionParam] `json:"value,required"`
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
	Clauses     param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesUnionParam] `json:"clauses,required"`
	Conjunction param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersConjunction]       `json:"conjunction,required"`
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
	Key      param.Field[string]                                                                                              `json:"key,required"`
	Operator param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsOperator] `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2FilterConditionsValueUnionParam] `json:"value,required"`
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
	Clauses     param.Field[[]MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseParam] `json:"clauses,required"`
	Conjunction param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersConjunction]   `json:"conjunction,required"`
}

func (r MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter condition with key, operator, and value
type MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClauseParam struct {
	// Filter key to apply
	Key      param.Field[string]                                                                                                  `json:"key,required"`
	Operator param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesOperator] `json:"operator,required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[MeterFilterClausesNestedMeterFiltersClausesLevel1NestedFiltersClausesLevel2NestedFiltersClausesValueUnionParam] `json:"value,required"`
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
	Aggregation param.Field[MeterAggregationParam] `json:"aggregation,required"`
	// Event name to track
	EventName param.Field[string] `json:"event_name,required"`
	// measurement unit
	MeasurementUnit param.Field[string] `json:"measurement_unit,required"`
	// Name of the meter
	Name param.Field[string] `json:"name,required"`
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
