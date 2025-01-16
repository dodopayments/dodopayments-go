// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
)

// RefundService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefundService] method instead.
type RefundService struct {
	Options []option.RequestOption
}

// NewRefundService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRefundService(opts ...option.RequestOption) (r *RefundService) {
	r = &RefundService{}
	r.Options = opts
	return
}

func (r *RefundService) New(ctx context.Context, body RefundNewParams, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	path := "refunds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *RefundService) Get(ctx context.Context, refundID string, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	if refundID == "" {
		err = errors.New("missing required refund_id parameter")
		return
	}
	path := fmt.Sprintf("refunds/%s", refundID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *RefundService) List(ctx context.Context, query RefundListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Refund], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "refunds"
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

func (r *RefundService) ListAutoPaging(ctx context.Context, query RefundListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Refund] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type Refund struct {
	// The unique identifier of the business issuing the refund.
	BusinessID string `json:"business_id,required"`
	// The timestamp of when the refund was created in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The unique identifier of the payment associated with the refund.
	PaymentID string `json:"payment_id,required"`
	// The unique identifier of the refund.
	RefundID string       `json:"refund_id,required"`
	Status   RefundStatus `json:"status,required"`
	// The refunded amount.
	Amount   int64          `json:"amount,nullable"`
	Currency RefundCurrency `json:"currency,nullable"`
	// The reason provided for the refund, if any. Optional.
	Reason string     `json:"reason,nullable"`
	JSON   refundJSON `json:"-"`
}

// refundJSON contains the JSON metadata for the struct [Refund]
type refundJSON struct {
	BusinessID  apijson.Field
	CreatedAt   apijson.Field
	PaymentID   apijson.Field
	RefundID    apijson.Field
	Status      apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Refund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refundJSON) RawJSON() string {
	return r.raw
}

type RefundStatus string

const (
	RefundStatusSucceeded RefundStatus = "succeeded"
	RefundStatusFailed    RefundStatus = "failed"
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusReview    RefundStatus = "review"
)

func (r RefundStatus) IsKnown() bool {
	switch r {
	case RefundStatusSucceeded, RefundStatusFailed, RefundStatusPending, RefundStatusReview:
		return true
	}
	return false
}

type RefundCurrency string

const (
	RefundCurrencyAed RefundCurrency = "AED"
	RefundCurrencyAll RefundCurrency = "ALL"
	RefundCurrencyAmd RefundCurrency = "AMD"
	RefundCurrencyAng RefundCurrency = "ANG"
	RefundCurrencyAoa RefundCurrency = "AOA"
	RefundCurrencyArs RefundCurrency = "ARS"
	RefundCurrencyAud RefundCurrency = "AUD"
	RefundCurrencyAwg RefundCurrency = "AWG"
	RefundCurrencyAzn RefundCurrency = "AZN"
	RefundCurrencyBam RefundCurrency = "BAM"
	RefundCurrencyBbd RefundCurrency = "BBD"
	RefundCurrencyBdt RefundCurrency = "BDT"
	RefundCurrencyBgn RefundCurrency = "BGN"
	RefundCurrencyBhd RefundCurrency = "BHD"
	RefundCurrencyBif RefundCurrency = "BIF"
	RefundCurrencyBmd RefundCurrency = "BMD"
	RefundCurrencyBnd RefundCurrency = "BND"
	RefundCurrencyBob RefundCurrency = "BOB"
	RefundCurrencyBrl RefundCurrency = "BRL"
	RefundCurrencyBsd RefundCurrency = "BSD"
	RefundCurrencyBwp RefundCurrency = "BWP"
	RefundCurrencyByn RefundCurrency = "BYN"
	RefundCurrencyBzd RefundCurrency = "BZD"
	RefundCurrencyCad RefundCurrency = "CAD"
	RefundCurrencyChf RefundCurrency = "CHF"
	RefundCurrencyClp RefundCurrency = "CLP"
	RefundCurrencyCny RefundCurrency = "CNY"
	RefundCurrencyCop RefundCurrency = "COP"
	RefundCurrencyCrc RefundCurrency = "CRC"
	RefundCurrencyCup RefundCurrency = "CUP"
	RefundCurrencyCve RefundCurrency = "CVE"
	RefundCurrencyCzk RefundCurrency = "CZK"
	RefundCurrencyDjf RefundCurrency = "DJF"
	RefundCurrencyDkk RefundCurrency = "DKK"
	RefundCurrencyDop RefundCurrency = "DOP"
	RefundCurrencyDzd RefundCurrency = "DZD"
	RefundCurrencyEgp RefundCurrency = "EGP"
	RefundCurrencyEtb RefundCurrency = "ETB"
	RefundCurrencyEur RefundCurrency = "EUR"
	RefundCurrencyFjd RefundCurrency = "FJD"
	RefundCurrencyFkp RefundCurrency = "FKP"
	RefundCurrencyGbp RefundCurrency = "GBP"
	RefundCurrencyGel RefundCurrency = "GEL"
	RefundCurrencyGhs RefundCurrency = "GHS"
	RefundCurrencyGip RefundCurrency = "GIP"
	RefundCurrencyGmd RefundCurrency = "GMD"
	RefundCurrencyGnf RefundCurrency = "GNF"
	RefundCurrencyGtq RefundCurrency = "GTQ"
	RefundCurrencyGyd RefundCurrency = "GYD"
	RefundCurrencyHkd RefundCurrency = "HKD"
	RefundCurrencyHnl RefundCurrency = "HNL"
	RefundCurrencyHrk RefundCurrency = "HRK"
	RefundCurrencyHtg RefundCurrency = "HTG"
	RefundCurrencyHuf RefundCurrency = "HUF"
	RefundCurrencyIdr RefundCurrency = "IDR"
	RefundCurrencyIls RefundCurrency = "ILS"
	RefundCurrencyInr RefundCurrency = "INR"
	RefundCurrencyIqd RefundCurrency = "IQD"
	RefundCurrencyJmd RefundCurrency = "JMD"
	RefundCurrencyJod RefundCurrency = "JOD"
	RefundCurrencyJpy RefundCurrency = "JPY"
	RefundCurrencyKes RefundCurrency = "KES"
	RefundCurrencyKgs RefundCurrency = "KGS"
	RefundCurrencyKhr RefundCurrency = "KHR"
	RefundCurrencyKmf RefundCurrency = "KMF"
	RefundCurrencyKrw RefundCurrency = "KRW"
	RefundCurrencyKwd RefundCurrency = "KWD"
	RefundCurrencyKyd RefundCurrency = "KYD"
	RefundCurrencyKzt RefundCurrency = "KZT"
	RefundCurrencyLak RefundCurrency = "LAK"
	RefundCurrencyLbp RefundCurrency = "LBP"
	RefundCurrencyLkr RefundCurrency = "LKR"
	RefundCurrencyLrd RefundCurrency = "LRD"
	RefundCurrencyLsl RefundCurrency = "LSL"
	RefundCurrencyLyd RefundCurrency = "LYD"
	RefundCurrencyMad RefundCurrency = "MAD"
	RefundCurrencyMdl RefundCurrency = "MDL"
	RefundCurrencyMga RefundCurrency = "MGA"
	RefundCurrencyMkd RefundCurrency = "MKD"
	RefundCurrencyMmk RefundCurrency = "MMK"
	RefundCurrencyMnt RefundCurrency = "MNT"
	RefundCurrencyMop RefundCurrency = "MOP"
	RefundCurrencyMru RefundCurrency = "MRU"
	RefundCurrencyMur RefundCurrency = "MUR"
	RefundCurrencyMvr RefundCurrency = "MVR"
	RefundCurrencyMwk RefundCurrency = "MWK"
	RefundCurrencyMxn RefundCurrency = "MXN"
	RefundCurrencyMyr RefundCurrency = "MYR"
	RefundCurrencyMzn RefundCurrency = "MZN"
	RefundCurrencyNad RefundCurrency = "NAD"
	RefundCurrencyNgn RefundCurrency = "NGN"
	RefundCurrencyNio RefundCurrency = "NIO"
	RefundCurrencyNok RefundCurrency = "NOK"
	RefundCurrencyNpr RefundCurrency = "NPR"
	RefundCurrencyNzd RefundCurrency = "NZD"
	RefundCurrencyOmr RefundCurrency = "OMR"
	RefundCurrencyPab RefundCurrency = "PAB"
	RefundCurrencyPen RefundCurrency = "PEN"
	RefundCurrencyPgk RefundCurrency = "PGK"
	RefundCurrencyPhp RefundCurrency = "PHP"
	RefundCurrencyPkr RefundCurrency = "PKR"
	RefundCurrencyPln RefundCurrency = "PLN"
	RefundCurrencyPyg RefundCurrency = "PYG"
	RefundCurrencyQar RefundCurrency = "QAR"
	RefundCurrencyRon RefundCurrency = "RON"
	RefundCurrencyRsd RefundCurrency = "RSD"
	RefundCurrencyRub RefundCurrency = "RUB"
	RefundCurrencyRwf RefundCurrency = "RWF"
	RefundCurrencySar RefundCurrency = "SAR"
	RefundCurrencySbd RefundCurrency = "SBD"
	RefundCurrencyScr RefundCurrency = "SCR"
	RefundCurrencySek RefundCurrency = "SEK"
	RefundCurrencySgd RefundCurrency = "SGD"
	RefundCurrencyShp RefundCurrency = "SHP"
	RefundCurrencySle RefundCurrency = "SLE"
	RefundCurrencySll RefundCurrency = "SLL"
	RefundCurrencySos RefundCurrency = "SOS"
	RefundCurrencySrd RefundCurrency = "SRD"
	RefundCurrencySsp RefundCurrency = "SSP"
	RefundCurrencyStn RefundCurrency = "STN"
	RefundCurrencySvc RefundCurrency = "SVC"
	RefundCurrencySzl RefundCurrency = "SZL"
	RefundCurrencyThb RefundCurrency = "THB"
	RefundCurrencyTnd RefundCurrency = "TND"
	RefundCurrencyTop RefundCurrency = "TOP"
	RefundCurrencyTry RefundCurrency = "TRY"
	RefundCurrencyTtd RefundCurrency = "TTD"
	RefundCurrencyTwd RefundCurrency = "TWD"
	RefundCurrencyTzs RefundCurrency = "TZS"
	RefundCurrencyUah RefundCurrency = "UAH"
	RefundCurrencyUgx RefundCurrency = "UGX"
	RefundCurrencyUsd RefundCurrency = "USD"
	RefundCurrencyUyu RefundCurrency = "UYU"
	RefundCurrencyUzs RefundCurrency = "UZS"
	RefundCurrencyVes RefundCurrency = "VES"
	RefundCurrencyVnd RefundCurrency = "VND"
	RefundCurrencyVuv RefundCurrency = "VUV"
	RefundCurrencyWst RefundCurrency = "WST"
	RefundCurrencyXaf RefundCurrency = "XAF"
	RefundCurrencyXcd RefundCurrency = "XCD"
	RefundCurrencyXof RefundCurrency = "XOF"
	RefundCurrencyXpf RefundCurrency = "XPF"
	RefundCurrencyYer RefundCurrency = "YER"
	RefundCurrencyZar RefundCurrency = "ZAR"
	RefundCurrencyZmw RefundCurrency = "ZMW"
)

func (r RefundCurrency) IsKnown() bool {
	switch r {
	case RefundCurrencyAed, RefundCurrencyAll, RefundCurrencyAmd, RefundCurrencyAng, RefundCurrencyAoa, RefundCurrencyArs, RefundCurrencyAud, RefundCurrencyAwg, RefundCurrencyAzn, RefundCurrencyBam, RefundCurrencyBbd, RefundCurrencyBdt, RefundCurrencyBgn, RefundCurrencyBhd, RefundCurrencyBif, RefundCurrencyBmd, RefundCurrencyBnd, RefundCurrencyBob, RefundCurrencyBrl, RefundCurrencyBsd, RefundCurrencyBwp, RefundCurrencyByn, RefundCurrencyBzd, RefundCurrencyCad, RefundCurrencyChf, RefundCurrencyClp, RefundCurrencyCny, RefundCurrencyCop, RefundCurrencyCrc, RefundCurrencyCup, RefundCurrencyCve, RefundCurrencyCzk, RefundCurrencyDjf, RefundCurrencyDkk, RefundCurrencyDop, RefundCurrencyDzd, RefundCurrencyEgp, RefundCurrencyEtb, RefundCurrencyEur, RefundCurrencyFjd, RefundCurrencyFkp, RefundCurrencyGbp, RefundCurrencyGel, RefundCurrencyGhs, RefundCurrencyGip, RefundCurrencyGmd, RefundCurrencyGnf, RefundCurrencyGtq, RefundCurrencyGyd, RefundCurrencyHkd, RefundCurrencyHnl, RefundCurrencyHrk, RefundCurrencyHtg, RefundCurrencyHuf, RefundCurrencyIdr, RefundCurrencyIls, RefundCurrencyInr, RefundCurrencyIqd, RefundCurrencyJmd, RefundCurrencyJod, RefundCurrencyJpy, RefundCurrencyKes, RefundCurrencyKgs, RefundCurrencyKhr, RefundCurrencyKmf, RefundCurrencyKrw, RefundCurrencyKwd, RefundCurrencyKyd, RefundCurrencyKzt, RefundCurrencyLak, RefundCurrencyLbp, RefundCurrencyLkr, RefundCurrencyLrd, RefundCurrencyLsl, RefundCurrencyLyd, RefundCurrencyMad, RefundCurrencyMdl, RefundCurrencyMga, RefundCurrencyMkd, RefundCurrencyMmk, RefundCurrencyMnt, RefundCurrencyMop, RefundCurrencyMru, RefundCurrencyMur, RefundCurrencyMvr, RefundCurrencyMwk, RefundCurrencyMxn, RefundCurrencyMyr, RefundCurrencyMzn, RefundCurrencyNad, RefundCurrencyNgn, RefundCurrencyNio, RefundCurrencyNok, RefundCurrencyNpr, RefundCurrencyNzd, RefundCurrencyOmr, RefundCurrencyPab, RefundCurrencyPen, RefundCurrencyPgk, RefundCurrencyPhp, RefundCurrencyPkr, RefundCurrencyPln, RefundCurrencyPyg, RefundCurrencyQar, RefundCurrencyRon, RefundCurrencyRsd, RefundCurrencyRub, RefundCurrencyRwf, RefundCurrencySar, RefundCurrencySbd, RefundCurrencyScr, RefundCurrencySek, RefundCurrencySgd, RefundCurrencyShp, RefundCurrencySle, RefundCurrencySll, RefundCurrencySos, RefundCurrencySrd, RefundCurrencySsp, RefundCurrencyStn, RefundCurrencySvc, RefundCurrencySzl, RefundCurrencyThb, RefundCurrencyTnd, RefundCurrencyTop, RefundCurrencyTry, RefundCurrencyTtd, RefundCurrencyTwd, RefundCurrencyTzs, RefundCurrencyUah, RefundCurrencyUgx, RefundCurrencyUsd, RefundCurrencyUyu, RefundCurrencyUzs, RefundCurrencyVes, RefundCurrencyVnd, RefundCurrencyVuv, RefundCurrencyWst, RefundCurrencyXaf, RefundCurrencyXcd, RefundCurrencyXof, RefundCurrencyXpf, RefundCurrencyYer, RefundCurrencyZar, RefundCurrencyZmw:
		return true
	}
	return false
}

type RefundNewParams struct {
	// The unique identifier of the payment to be refunded.
	PaymentID param.Field[string] `json:"payment_id,required"`
	// The amount to be refunded. Must be non-negative. Optional. Partial refunds are
	// currently disabled.
	Amount param.Field[int64] `json:"amount"`
	// The reason for the refund, if any. Maximum length is 3000 characters. Optional.
	Reason param.Field[string] `json:"reason"`
}

func (r RefundNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RefundListParams struct {
	// Get events after this created time
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get events created before this time
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by status
	Status param.Field[RefundListParamsStatus] `query:"status"`
}

// URLQuery serializes [RefundListParams]'s query parameters as `url.Values`.
func (r RefundListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status
type RefundListParamsStatus string

const (
	RefundListParamsStatusSucceeded RefundListParamsStatus = "succeeded"
	RefundListParamsStatusFailed    RefundListParamsStatus = "failed"
	RefundListParamsStatusPending   RefundListParamsStatus = "pending"
	RefundListParamsStatusReview    RefundListParamsStatus = "review"
)

func (r RefundListParamsStatus) IsKnown() bool {
	switch r {
	case RefundListParamsStatusSucceeded, RefundListParamsStatusFailed, RefundListParamsStatusPending, RefundListParamsStatusReview:
		return true
	}
	return false
}
