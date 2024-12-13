// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/dodo-payments-go/internal/apijson"
	"github.com/stainless-sdks/dodo-payments-go/internal/apiquery"
	"github.com/stainless-sdks/dodo-payments-go/internal/param"
	"github.com/stainless-sdks/dodo-payments-go/internal/requestconfig"
	"github.com/stainless-sdks/dodo-payments-go/option"
)

// PayoutService contains methods and other services that help with interacting
// with the dodopayments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayoutService] method instead.
type PayoutService struct {
	Options []option.RequestOption
}

// NewPayoutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPayoutService(opts ...option.RequestOption) (r *PayoutService) {
	r = &PayoutService{}
	r.Options = opts
	return
}

func (r *PayoutService) List(ctx context.Context, query PayoutListParams, opts ...option.RequestOption) (res *PayoutListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PayoutListResponse struct {
	Items []PayoutListResponseItem `json:"items,required"`
	JSON  payoutListResponseJSON   `json:"-"`
}

// payoutListResponseJSON contains the JSON metadata for the struct
// [PayoutListResponse]
type payoutListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayoutListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutListResponseJSON) RawJSON() string {
	return r.raw
}

type PayoutListResponseItem struct {
	Amount            int64                           `json:"amount,required"`
	BusinessID        string                          `json:"business_id,required"`
	Chargebacks       int64                           `json:"chargebacks,required"`
	CreatedAt         time.Time                       `json:"created_at,required" format:"date-time"`
	Currency          PayoutListResponseItemsCurrency `json:"currency,required"`
	Fee               int64                           `json:"fee,required"`
	PaymentMethod     string                          `json:"payment_method,required"`
	PayoutID          string                          `json:"payout_id,required"`
	Refunds           int64                           `json:"refunds,required"`
	Status            PayoutListResponseItemsStatus   `json:"status,required"`
	Tax               int64                           `json:"tax,required"`
	UpdatedAt         time.Time                       `json:"updated_at,required" format:"date-time"`
	Name              string                          `json:"name,nullable"`
	PayoutDocumentURL string                          `json:"payout_document_url,nullable"`
	Remarks           string                          `json:"remarks,nullable"`
	JSON              payoutListResponseItemJSON      `json:"-"`
}

// payoutListResponseItemJSON contains the JSON metadata for the struct
// [PayoutListResponseItem]
type payoutListResponseItemJSON struct {
	Amount            apijson.Field
	BusinessID        apijson.Field
	Chargebacks       apijson.Field
	CreatedAt         apijson.Field
	Currency          apijson.Field
	Fee               apijson.Field
	PaymentMethod     apijson.Field
	PayoutID          apijson.Field
	Refunds           apijson.Field
	Status            apijson.Field
	Tax               apijson.Field
	UpdatedAt         apijson.Field
	Name              apijson.Field
	PayoutDocumentURL apijson.Field
	Remarks           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PayoutListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutListResponseItemJSON) RawJSON() string {
	return r.raw
}

type PayoutListResponseItemsCurrency string

const (
	PayoutListResponseItemsCurrencyAed PayoutListResponseItemsCurrency = "AED"
	PayoutListResponseItemsCurrencyAll PayoutListResponseItemsCurrency = "ALL"
	PayoutListResponseItemsCurrencyAmd PayoutListResponseItemsCurrency = "AMD"
	PayoutListResponseItemsCurrencyAng PayoutListResponseItemsCurrency = "ANG"
	PayoutListResponseItemsCurrencyAoa PayoutListResponseItemsCurrency = "AOA"
	PayoutListResponseItemsCurrencyArs PayoutListResponseItemsCurrency = "ARS"
	PayoutListResponseItemsCurrencyAud PayoutListResponseItemsCurrency = "AUD"
	PayoutListResponseItemsCurrencyAwg PayoutListResponseItemsCurrency = "AWG"
	PayoutListResponseItemsCurrencyAzn PayoutListResponseItemsCurrency = "AZN"
	PayoutListResponseItemsCurrencyBam PayoutListResponseItemsCurrency = "BAM"
	PayoutListResponseItemsCurrencyBbd PayoutListResponseItemsCurrency = "BBD"
	PayoutListResponseItemsCurrencyBdt PayoutListResponseItemsCurrency = "BDT"
	PayoutListResponseItemsCurrencyBgn PayoutListResponseItemsCurrency = "BGN"
	PayoutListResponseItemsCurrencyBhd PayoutListResponseItemsCurrency = "BHD"
	PayoutListResponseItemsCurrencyBif PayoutListResponseItemsCurrency = "BIF"
	PayoutListResponseItemsCurrencyBmd PayoutListResponseItemsCurrency = "BMD"
	PayoutListResponseItemsCurrencyBnd PayoutListResponseItemsCurrency = "BND"
	PayoutListResponseItemsCurrencyBob PayoutListResponseItemsCurrency = "BOB"
	PayoutListResponseItemsCurrencyBrl PayoutListResponseItemsCurrency = "BRL"
	PayoutListResponseItemsCurrencyBsd PayoutListResponseItemsCurrency = "BSD"
	PayoutListResponseItemsCurrencyBwp PayoutListResponseItemsCurrency = "BWP"
	PayoutListResponseItemsCurrencyByn PayoutListResponseItemsCurrency = "BYN"
	PayoutListResponseItemsCurrencyBzd PayoutListResponseItemsCurrency = "BZD"
	PayoutListResponseItemsCurrencyCad PayoutListResponseItemsCurrency = "CAD"
	PayoutListResponseItemsCurrencyChf PayoutListResponseItemsCurrency = "CHF"
	PayoutListResponseItemsCurrencyClp PayoutListResponseItemsCurrency = "CLP"
	PayoutListResponseItemsCurrencyCny PayoutListResponseItemsCurrency = "CNY"
	PayoutListResponseItemsCurrencyCop PayoutListResponseItemsCurrency = "COP"
	PayoutListResponseItemsCurrencyCrc PayoutListResponseItemsCurrency = "CRC"
	PayoutListResponseItemsCurrencyCup PayoutListResponseItemsCurrency = "CUP"
	PayoutListResponseItemsCurrencyCve PayoutListResponseItemsCurrency = "CVE"
	PayoutListResponseItemsCurrencyCzk PayoutListResponseItemsCurrency = "CZK"
	PayoutListResponseItemsCurrencyDjf PayoutListResponseItemsCurrency = "DJF"
	PayoutListResponseItemsCurrencyDkk PayoutListResponseItemsCurrency = "DKK"
	PayoutListResponseItemsCurrencyDop PayoutListResponseItemsCurrency = "DOP"
	PayoutListResponseItemsCurrencyDzd PayoutListResponseItemsCurrency = "DZD"
	PayoutListResponseItemsCurrencyEgp PayoutListResponseItemsCurrency = "EGP"
	PayoutListResponseItemsCurrencyEtb PayoutListResponseItemsCurrency = "ETB"
	PayoutListResponseItemsCurrencyEur PayoutListResponseItemsCurrency = "EUR"
	PayoutListResponseItemsCurrencyFjd PayoutListResponseItemsCurrency = "FJD"
	PayoutListResponseItemsCurrencyFkp PayoutListResponseItemsCurrency = "FKP"
	PayoutListResponseItemsCurrencyGbp PayoutListResponseItemsCurrency = "GBP"
	PayoutListResponseItemsCurrencyGel PayoutListResponseItemsCurrency = "GEL"
	PayoutListResponseItemsCurrencyGhs PayoutListResponseItemsCurrency = "GHS"
	PayoutListResponseItemsCurrencyGip PayoutListResponseItemsCurrency = "GIP"
	PayoutListResponseItemsCurrencyGmd PayoutListResponseItemsCurrency = "GMD"
	PayoutListResponseItemsCurrencyGnf PayoutListResponseItemsCurrency = "GNF"
	PayoutListResponseItemsCurrencyGtq PayoutListResponseItemsCurrency = "GTQ"
	PayoutListResponseItemsCurrencyGyd PayoutListResponseItemsCurrency = "GYD"
	PayoutListResponseItemsCurrencyHkd PayoutListResponseItemsCurrency = "HKD"
	PayoutListResponseItemsCurrencyHnl PayoutListResponseItemsCurrency = "HNL"
	PayoutListResponseItemsCurrencyHrk PayoutListResponseItemsCurrency = "HRK"
	PayoutListResponseItemsCurrencyHtg PayoutListResponseItemsCurrency = "HTG"
	PayoutListResponseItemsCurrencyHuf PayoutListResponseItemsCurrency = "HUF"
	PayoutListResponseItemsCurrencyIdr PayoutListResponseItemsCurrency = "IDR"
	PayoutListResponseItemsCurrencyIls PayoutListResponseItemsCurrency = "ILS"
	PayoutListResponseItemsCurrencyInr PayoutListResponseItemsCurrency = "INR"
	PayoutListResponseItemsCurrencyIqd PayoutListResponseItemsCurrency = "IQD"
	PayoutListResponseItemsCurrencyJmd PayoutListResponseItemsCurrency = "JMD"
	PayoutListResponseItemsCurrencyJod PayoutListResponseItemsCurrency = "JOD"
	PayoutListResponseItemsCurrencyJpy PayoutListResponseItemsCurrency = "JPY"
	PayoutListResponseItemsCurrencyKes PayoutListResponseItemsCurrency = "KES"
	PayoutListResponseItemsCurrencyKgs PayoutListResponseItemsCurrency = "KGS"
	PayoutListResponseItemsCurrencyKhr PayoutListResponseItemsCurrency = "KHR"
	PayoutListResponseItemsCurrencyKmf PayoutListResponseItemsCurrency = "KMF"
	PayoutListResponseItemsCurrencyKrw PayoutListResponseItemsCurrency = "KRW"
	PayoutListResponseItemsCurrencyKwd PayoutListResponseItemsCurrency = "KWD"
	PayoutListResponseItemsCurrencyKyd PayoutListResponseItemsCurrency = "KYD"
	PayoutListResponseItemsCurrencyKzt PayoutListResponseItemsCurrency = "KZT"
	PayoutListResponseItemsCurrencyLak PayoutListResponseItemsCurrency = "LAK"
	PayoutListResponseItemsCurrencyLbp PayoutListResponseItemsCurrency = "LBP"
	PayoutListResponseItemsCurrencyLkr PayoutListResponseItemsCurrency = "LKR"
	PayoutListResponseItemsCurrencyLrd PayoutListResponseItemsCurrency = "LRD"
	PayoutListResponseItemsCurrencyLsl PayoutListResponseItemsCurrency = "LSL"
	PayoutListResponseItemsCurrencyLyd PayoutListResponseItemsCurrency = "LYD"
	PayoutListResponseItemsCurrencyMad PayoutListResponseItemsCurrency = "MAD"
	PayoutListResponseItemsCurrencyMdl PayoutListResponseItemsCurrency = "MDL"
	PayoutListResponseItemsCurrencyMga PayoutListResponseItemsCurrency = "MGA"
	PayoutListResponseItemsCurrencyMkd PayoutListResponseItemsCurrency = "MKD"
	PayoutListResponseItemsCurrencyMmk PayoutListResponseItemsCurrency = "MMK"
	PayoutListResponseItemsCurrencyMnt PayoutListResponseItemsCurrency = "MNT"
	PayoutListResponseItemsCurrencyMop PayoutListResponseItemsCurrency = "MOP"
	PayoutListResponseItemsCurrencyMru PayoutListResponseItemsCurrency = "MRU"
	PayoutListResponseItemsCurrencyMur PayoutListResponseItemsCurrency = "MUR"
	PayoutListResponseItemsCurrencyMvr PayoutListResponseItemsCurrency = "MVR"
	PayoutListResponseItemsCurrencyMwk PayoutListResponseItemsCurrency = "MWK"
	PayoutListResponseItemsCurrencyMxn PayoutListResponseItemsCurrency = "MXN"
	PayoutListResponseItemsCurrencyMyr PayoutListResponseItemsCurrency = "MYR"
	PayoutListResponseItemsCurrencyMzn PayoutListResponseItemsCurrency = "MZN"
	PayoutListResponseItemsCurrencyNad PayoutListResponseItemsCurrency = "NAD"
	PayoutListResponseItemsCurrencyNgn PayoutListResponseItemsCurrency = "NGN"
	PayoutListResponseItemsCurrencyNio PayoutListResponseItemsCurrency = "NIO"
	PayoutListResponseItemsCurrencyNok PayoutListResponseItemsCurrency = "NOK"
	PayoutListResponseItemsCurrencyNpr PayoutListResponseItemsCurrency = "NPR"
	PayoutListResponseItemsCurrencyNzd PayoutListResponseItemsCurrency = "NZD"
	PayoutListResponseItemsCurrencyOmr PayoutListResponseItemsCurrency = "OMR"
	PayoutListResponseItemsCurrencyPab PayoutListResponseItemsCurrency = "PAB"
	PayoutListResponseItemsCurrencyPen PayoutListResponseItemsCurrency = "PEN"
	PayoutListResponseItemsCurrencyPgk PayoutListResponseItemsCurrency = "PGK"
	PayoutListResponseItemsCurrencyPhp PayoutListResponseItemsCurrency = "PHP"
	PayoutListResponseItemsCurrencyPkr PayoutListResponseItemsCurrency = "PKR"
	PayoutListResponseItemsCurrencyPln PayoutListResponseItemsCurrency = "PLN"
	PayoutListResponseItemsCurrencyPyg PayoutListResponseItemsCurrency = "PYG"
	PayoutListResponseItemsCurrencyQar PayoutListResponseItemsCurrency = "QAR"
	PayoutListResponseItemsCurrencyRon PayoutListResponseItemsCurrency = "RON"
	PayoutListResponseItemsCurrencyRsd PayoutListResponseItemsCurrency = "RSD"
	PayoutListResponseItemsCurrencyRub PayoutListResponseItemsCurrency = "RUB"
	PayoutListResponseItemsCurrencyRwf PayoutListResponseItemsCurrency = "RWF"
	PayoutListResponseItemsCurrencySar PayoutListResponseItemsCurrency = "SAR"
	PayoutListResponseItemsCurrencySbd PayoutListResponseItemsCurrency = "SBD"
	PayoutListResponseItemsCurrencyScr PayoutListResponseItemsCurrency = "SCR"
	PayoutListResponseItemsCurrencySek PayoutListResponseItemsCurrency = "SEK"
	PayoutListResponseItemsCurrencySgd PayoutListResponseItemsCurrency = "SGD"
	PayoutListResponseItemsCurrencyShp PayoutListResponseItemsCurrency = "SHP"
	PayoutListResponseItemsCurrencySle PayoutListResponseItemsCurrency = "SLE"
	PayoutListResponseItemsCurrencySll PayoutListResponseItemsCurrency = "SLL"
	PayoutListResponseItemsCurrencySos PayoutListResponseItemsCurrency = "SOS"
	PayoutListResponseItemsCurrencySrd PayoutListResponseItemsCurrency = "SRD"
	PayoutListResponseItemsCurrencySsp PayoutListResponseItemsCurrency = "SSP"
	PayoutListResponseItemsCurrencyStn PayoutListResponseItemsCurrency = "STN"
	PayoutListResponseItemsCurrencySvc PayoutListResponseItemsCurrency = "SVC"
	PayoutListResponseItemsCurrencySzl PayoutListResponseItemsCurrency = "SZL"
	PayoutListResponseItemsCurrencyThb PayoutListResponseItemsCurrency = "THB"
	PayoutListResponseItemsCurrencyTnd PayoutListResponseItemsCurrency = "TND"
	PayoutListResponseItemsCurrencyTop PayoutListResponseItemsCurrency = "TOP"
	PayoutListResponseItemsCurrencyTry PayoutListResponseItemsCurrency = "TRY"
	PayoutListResponseItemsCurrencyTtd PayoutListResponseItemsCurrency = "TTD"
	PayoutListResponseItemsCurrencyTwd PayoutListResponseItemsCurrency = "TWD"
	PayoutListResponseItemsCurrencyTzs PayoutListResponseItemsCurrency = "TZS"
	PayoutListResponseItemsCurrencyUah PayoutListResponseItemsCurrency = "UAH"
	PayoutListResponseItemsCurrencyUgx PayoutListResponseItemsCurrency = "UGX"
	PayoutListResponseItemsCurrencyUsd PayoutListResponseItemsCurrency = "USD"
	PayoutListResponseItemsCurrencyUyu PayoutListResponseItemsCurrency = "UYU"
	PayoutListResponseItemsCurrencyUzs PayoutListResponseItemsCurrency = "UZS"
	PayoutListResponseItemsCurrencyVes PayoutListResponseItemsCurrency = "VES"
	PayoutListResponseItemsCurrencyVnd PayoutListResponseItemsCurrency = "VND"
	PayoutListResponseItemsCurrencyVuv PayoutListResponseItemsCurrency = "VUV"
	PayoutListResponseItemsCurrencyWst PayoutListResponseItemsCurrency = "WST"
	PayoutListResponseItemsCurrencyXaf PayoutListResponseItemsCurrency = "XAF"
	PayoutListResponseItemsCurrencyXcd PayoutListResponseItemsCurrency = "XCD"
	PayoutListResponseItemsCurrencyXof PayoutListResponseItemsCurrency = "XOF"
	PayoutListResponseItemsCurrencyXpf PayoutListResponseItemsCurrency = "XPF"
	PayoutListResponseItemsCurrencyYer PayoutListResponseItemsCurrency = "YER"
	PayoutListResponseItemsCurrencyZar PayoutListResponseItemsCurrency = "ZAR"
	PayoutListResponseItemsCurrencyZmw PayoutListResponseItemsCurrency = "ZMW"
)

func (r PayoutListResponseItemsCurrency) IsKnown() bool {
	switch r {
	case PayoutListResponseItemsCurrencyAed, PayoutListResponseItemsCurrencyAll, PayoutListResponseItemsCurrencyAmd, PayoutListResponseItemsCurrencyAng, PayoutListResponseItemsCurrencyAoa, PayoutListResponseItemsCurrencyArs, PayoutListResponseItemsCurrencyAud, PayoutListResponseItemsCurrencyAwg, PayoutListResponseItemsCurrencyAzn, PayoutListResponseItemsCurrencyBam, PayoutListResponseItemsCurrencyBbd, PayoutListResponseItemsCurrencyBdt, PayoutListResponseItemsCurrencyBgn, PayoutListResponseItemsCurrencyBhd, PayoutListResponseItemsCurrencyBif, PayoutListResponseItemsCurrencyBmd, PayoutListResponseItemsCurrencyBnd, PayoutListResponseItemsCurrencyBob, PayoutListResponseItemsCurrencyBrl, PayoutListResponseItemsCurrencyBsd, PayoutListResponseItemsCurrencyBwp, PayoutListResponseItemsCurrencyByn, PayoutListResponseItemsCurrencyBzd, PayoutListResponseItemsCurrencyCad, PayoutListResponseItemsCurrencyChf, PayoutListResponseItemsCurrencyClp, PayoutListResponseItemsCurrencyCny, PayoutListResponseItemsCurrencyCop, PayoutListResponseItemsCurrencyCrc, PayoutListResponseItemsCurrencyCup, PayoutListResponseItemsCurrencyCve, PayoutListResponseItemsCurrencyCzk, PayoutListResponseItemsCurrencyDjf, PayoutListResponseItemsCurrencyDkk, PayoutListResponseItemsCurrencyDop, PayoutListResponseItemsCurrencyDzd, PayoutListResponseItemsCurrencyEgp, PayoutListResponseItemsCurrencyEtb, PayoutListResponseItemsCurrencyEur, PayoutListResponseItemsCurrencyFjd, PayoutListResponseItemsCurrencyFkp, PayoutListResponseItemsCurrencyGbp, PayoutListResponseItemsCurrencyGel, PayoutListResponseItemsCurrencyGhs, PayoutListResponseItemsCurrencyGip, PayoutListResponseItemsCurrencyGmd, PayoutListResponseItemsCurrencyGnf, PayoutListResponseItemsCurrencyGtq, PayoutListResponseItemsCurrencyGyd, PayoutListResponseItemsCurrencyHkd, PayoutListResponseItemsCurrencyHnl, PayoutListResponseItemsCurrencyHrk, PayoutListResponseItemsCurrencyHtg, PayoutListResponseItemsCurrencyHuf, PayoutListResponseItemsCurrencyIdr, PayoutListResponseItemsCurrencyIls, PayoutListResponseItemsCurrencyInr, PayoutListResponseItemsCurrencyIqd, PayoutListResponseItemsCurrencyJmd, PayoutListResponseItemsCurrencyJod, PayoutListResponseItemsCurrencyJpy, PayoutListResponseItemsCurrencyKes, PayoutListResponseItemsCurrencyKgs, PayoutListResponseItemsCurrencyKhr, PayoutListResponseItemsCurrencyKmf, PayoutListResponseItemsCurrencyKrw, PayoutListResponseItemsCurrencyKwd, PayoutListResponseItemsCurrencyKyd, PayoutListResponseItemsCurrencyKzt, PayoutListResponseItemsCurrencyLak, PayoutListResponseItemsCurrencyLbp, PayoutListResponseItemsCurrencyLkr, PayoutListResponseItemsCurrencyLrd, PayoutListResponseItemsCurrencyLsl, PayoutListResponseItemsCurrencyLyd, PayoutListResponseItemsCurrencyMad, PayoutListResponseItemsCurrencyMdl, PayoutListResponseItemsCurrencyMga, PayoutListResponseItemsCurrencyMkd, PayoutListResponseItemsCurrencyMmk, PayoutListResponseItemsCurrencyMnt, PayoutListResponseItemsCurrencyMop, PayoutListResponseItemsCurrencyMru, PayoutListResponseItemsCurrencyMur, PayoutListResponseItemsCurrencyMvr, PayoutListResponseItemsCurrencyMwk, PayoutListResponseItemsCurrencyMxn, PayoutListResponseItemsCurrencyMyr, PayoutListResponseItemsCurrencyMzn, PayoutListResponseItemsCurrencyNad, PayoutListResponseItemsCurrencyNgn, PayoutListResponseItemsCurrencyNio, PayoutListResponseItemsCurrencyNok, PayoutListResponseItemsCurrencyNpr, PayoutListResponseItemsCurrencyNzd, PayoutListResponseItemsCurrencyOmr, PayoutListResponseItemsCurrencyPab, PayoutListResponseItemsCurrencyPen, PayoutListResponseItemsCurrencyPgk, PayoutListResponseItemsCurrencyPhp, PayoutListResponseItemsCurrencyPkr, PayoutListResponseItemsCurrencyPln, PayoutListResponseItemsCurrencyPyg, PayoutListResponseItemsCurrencyQar, PayoutListResponseItemsCurrencyRon, PayoutListResponseItemsCurrencyRsd, PayoutListResponseItemsCurrencyRub, PayoutListResponseItemsCurrencyRwf, PayoutListResponseItemsCurrencySar, PayoutListResponseItemsCurrencySbd, PayoutListResponseItemsCurrencyScr, PayoutListResponseItemsCurrencySek, PayoutListResponseItemsCurrencySgd, PayoutListResponseItemsCurrencyShp, PayoutListResponseItemsCurrencySle, PayoutListResponseItemsCurrencySll, PayoutListResponseItemsCurrencySos, PayoutListResponseItemsCurrencySrd, PayoutListResponseItemsCurrencySsp, PayoutListResponseItemsCurrencyStn, PayoutListResponseItemsCurrencySvc, PayoutListResponseItemsCurrencySzl, PayoutListResponseItemsCurrencyThb, PayoutListResponseItemsCurrencyTnd, PayoutListResponseItemsCurrencyTop, PayoutListResponseItemsCurrencyTry, PayoutListResponseItemsCurrencyTtd, PayoutListResponseItemsCurrencyTwd, PayoutListResponseItemsCurrencyTzs, PayoutListResponseItemsCurrencyUah, PayoutListResponseItemsCurrencyUgx, PayoutListResponseItemsCurrencyUsd, PayoutListResponseItemsCurrencyUyu, PayoutListResponseItemsCurrencyUzs, PayoutListResponseItemsCurrencyVes, PayoutListResponseItemsCurrencyVnd, PayoutListResponseItemsCurrencyVuv, PayoutListResponseItemsCurrencyWst, PayoutListResponseItemsCurrencyXaf, PayoutListResponseItemsCurrencyXcd, PayoutListResponseItemsCurrencyXof, PayoutListResponseItemsCurrencyXpf, PayoutListResponseItemsCurrencyYer, PayoutListResponseItemsCurrencyZar, PayoutListResponseItemsCurrencyZmw:
		return true
	}
	return false
}

type PayoutListResponseItemsStatus string

const (
	PayoutListResponseItemsStatusInProgress PayoutListResponseItemsStatus = "in_progress"
	PayoutListResponseItemsStatusFailed     PayoutListResponseItemsStatus = "failed"
	PayoutListResponseItemsStatusSuccess    PayoutListResponseItemsStatus = "success"
)

func (r PayoutListResponseItemsStatus) IsKnown() bool {
	switch r {
	case PayoutListResponseItemsStatusInProgress, PayoutListResponseItemsStatusFailed, PayoutListResponseItemsStatusSuccess:
		return true
	}
	return false
}

type PayoutListParams struct {
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [PayoutListParams]'s query parameters as `url.Values`.
func (r PayoutListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
