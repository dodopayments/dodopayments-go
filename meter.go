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
	return res, err
}

func (r *MeterService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Meter, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meters/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
		return err
	}
	path := fmt.Sprintf("meters/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *MeterService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("meters/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
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

// Filter clauses — either a flat list of `MeterFilterCondition`s or a list of
// nested `MeterFilter`s. Up to 3 levels of nesting are accepted; the limit is
// enforced at runtime.
//
// Union satisfied by [FilterTypeMeterFilterConditionList] or
// [FilterTypeNestedMeterFilterList].
type FilterTypeUnion interface {
	implementsFilterTypeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FilterTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FilterTypeMeterFilterConditionList{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FilterTypeNestedMeterFilterList{}),
		},
	)
}

type FilterTypeMeterFilterConditionList []FilterTypeMeterFilterConditionListItem

func (r FilterTypeMeterFilterConditionList) implementsFilterTypeUnion() {}

type FilterTypeMeterFilterConditionListItem struct {
	// Filter key to apply
	Key string `json:"key" api:"required"`
	// Filter operator
	Operator FilterOperator `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value FilterTypeMeterFilterConditionListValueUnion `json:"value" api:"required"`
	JSON  filterTypeMeterFilterConditionListItemJSON   `json:"-"`
}

// filterTypeMeterFilterConditionListItemJSON contains the JSON metadata for the
// struct [FilterTypeMeterFilterConditionListItem]
type filterTypeMeterFilterConditionListItemJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterTypeMeterFilterConditionListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterTypeMeterFilterConditionListItemJSON) RawJSON() string {
	return r.raw
}

// Filter value - can be string, number, or boolean
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type FilterTypeMeterFilterConditionListValueUnion interface {
	ImplementsFilterTypeMeterFilterConditionListValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FilterTypeMeterFilterConditionListValueUnion)(nil)).Elem(),
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

type FilterTypeNestedMeterFilterList []MeterFilter

func (r FilterTypeNestedMeterFilterList) implementsFilterTypeUnion() {}

// Filter clauses — either a flat list of `MeterFilterCondition`s or a list of
// nested `MeterFilter`s. Up to 3 levels of nesting are accepted; the limit is
// enforced at runtime.
//
// Satisfied by [FilterTypeMeterFilterConditionListParam],
// [FilterTypeNestedMeterFilterListParam].
type FilterTypeUnionParam interface {
	implementsFilterTypeUnionParam()
}

type FilterTypeMeterFilterConditionListParam []FilterTypeMeterFilterConditionListItemParam

func (r FilterTypeMeterFilterConditionListParam) implementsFilterTypeUnionParam() {}

type FilterTypeMeterFilterConditionListItemParam struct {
	// Filter key to apply
	Key param.Field[string] `json:"key" api:"required"`
	// Filter operator
	Operator param.Field[FilterOperator] `json:"operator" api:"required"`
	// Filter value - can be string, number, or boolean
	Value param.Field[FilterTypeMeterFilterConditionListValueUnionParam] `json:"value" api:"required"`
}

func (r FilterTypeMeterFilterConditionListItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter value - can be string, number, or boolean
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type FilterTypeMeterFilterConditionListValueUnionParam interface {
	ImplementsFilterTypeMeterFilterConditionListValueUnionParam()
}

type FilterTypeNestedMeterFilterListParam []MeterFilterParam

func (r FilterTypeNestedMeterFilterListParam) implementsFilterTypeUnionParam() {}

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
	Clauses FilterTypeUnion `json:"clauses" api:"required"`
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

// A filter structure that combines multiple conditions with logical conjunctions
// (AND/OR).
//
// Supports up to 3 levels of nesting to create complex filter expressions. Each
// filter has a conjunction (and/or) and clauses that can be either direct
// conditions or nested filters.
type MeterFilterParam struct {
	// Filter clauses - can be direct conditions or nested filters (up to 3 levels
	// deep)
	Clauses param.Field[FilterTypeUnionParam] `json:"clauses" api:"required"`
	// Logical conjunction to apply between clauses (and/or)
	Conjunction param.Field[Conjunction] `json:"conjunction" api:"required"`
}

func (r MeterFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
