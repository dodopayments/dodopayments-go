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
	"github.com/stainless-sdks/dodo-payments-go/packages/pagination"
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

func (r *PayoutService) List(ctx context.Context, query PayoutListParams, opts ...option.RequestOption) (res *pagination.PageNumberPage[PayoutListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "payouts"
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

func (r *PayoutService) ListAutoPaging(ctx context.Context, query PayoutListParams, opts ...option.RequestOption) *pagination.PageNumberPageAutoPager[PayoutListResponse] {
	return pagination.NewPageNumberPageAutoPager(r.List(ctx, query, opts...))
}

type PayoutListResponse struct {
	Amount            int64                      `json:"amount,required"`
	BusinessID        string                     `json:"business_id,required"`
	Chargebacks       int64                      `json:"chargebacks,required"`
	CreatedAt         time.Time                  `json:"created_at,required" format:"date-time"`
	Currency          PayoutListResponseCurrency `json:"currency,required"`
	Fee               int64                      `json:"fee,required"`
	PaymentMethod     string                     `json:"payment_method,required"`
	PayoutID          string                     `json:"payout_id,required"`
	Refunds           int64                      `json:"refunds,required"`
	Status            PayoutListResponseStatus   `json:"status,required"`
	Tax               int64                      `json:"tax,required"`
	UpdatedAt         time.Time                  `json:"updated_at,required" format:"date-time"`
	Name              string                     `json:"name,nullable"`
	PayoutDocumentURL string                     `json:"payout_document_url,nullable"`
	Remarks           string                     `json:"remarks,nullable"`
	JSON              payoutListResponseJSON     `json:"-"`
}

// payoutListResponseJSON contains the JSON metadata for the struct
// [PayoutListResponse]
type payoutListResponseJSON struct {
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

func (r *PayoutListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutListResponseJSON) RawJSON() string {
	return r.raw
}

type PayoutListResponseCurrency string

const (
	PayoutListResponseCurrencyAed PayoutListResponseCurrency = "AED"
	PayoutListResponseCurrencyAll PayoutListResponseCurrency = "ALL"
	PayoutListResponseCurrencyAmd PayoutListResponseCurrency = "AMD"
	PayoutListResponseCurrencyAng PayoutListResponseCurrency = "ANG"
	PayoutListResponseCurrencyAoa PayoutListResponseCurrency = "AOA"
	PayoutListResponseCurrencyArs PayoutListResponseCurrency = "ARS"
	PayoutListResponseCurrencyAud PayoutListResponseCurrency = "AUD"
	PayoutListResponseCurrencyAwg PayoutListResponseCurrency = "AWG"
	PayoutListResponseCurrencyAzn PayoutListResponseCurrency = "AZN"
	PayoutListResponseCurrencyBam PayoutListResponseCurrency = "BAM"
	PayoutListResponseCurrencyBbd PayoutListResponseCurrency = "BBD"
	PayoutListResponseCurrencyBdt PayoutListResponseCurrency = "BDT"
	PayoutListResponseCurrencyBgn PayoutListResponseCurrency = "BGN"
	PayoutListResponseCurrencyBhd PayoutListResponseCurrency = "BHD"
	PayoutListResponseCurrencyBif PayoutListResponseCurrency = "BIF"
	PayoutListResponseCurrencyBmd PayoutListResponseCurrency = "BMD"
	PayoutListResponseCurrencyBnd PayoutListResponseCurrency = "BND"
	PayoutListResponseCurrencyBob PayoutListResponseCurrency = "BOB"
	PayoutListResponseCurrencyBrl PayoutListResponseCurrency = "BRL"
	PayoutListResponseCurrencyBsd PayoutListResponseCurrency = "BSD"
	PayoutListResponseCurrencyBwp PayoutListResponseCurrency = "BWP"
	PayoutListResponseCurrencyByn PayoutListResponseCurrency = "BYN"
	PayoutListResponseCurrencyBzd PayoutListResponseCurrency = "BZD"
	PayoutListResponseCurrencyCad PayoutListResponseCurrency = "CAD"
	PayoutListResponseCurrencyChf PayoutListResponseCurrency = "CHF"
	PayoutListResponseCurrencyClp PayoutListResponseCurrency = "CLP"
	PayoutListResponseCurrencyCny PayoutListResponseCurrency = "CNY"
	PayoutListResponseCurrencyCop PayoutListResponseCurrency = "COP"
	PayoutListResponseCurrencyCrc PayoutListResponseCurrency = "CRC"
	PayoutListResponseCurrencyCup PayoutListResponseCurrency = "CUP"
	PayoutListResponseCurrencyCve PayoutListResponseCurrency = "CVE"
	PayoutListResponseCurrencyCzk PayoutListResponseCurrency = "CZK"
	PayoutListResponseCurrencyDjf PayoutListResponseCurrency = "DJF"
	PayoutListResponseCurrencyDkk PayoutListResponseCurrency = "DKK"
	PayoutListResponseCurrencyDop PayoutListResponseCurrency = "DOP"
	PayoutListResponseCurrencyDzd PayoutListResponseCurrency = "DZD"
	PayoutListResponseCurrencyEgp PayoutListResponseCurrency = "EGP"
	PayoutListResponseCurrencyEtb PayoutListResponseCurrency = "ETB"
	PayoutListResponseCurrencyEur PayoutListResponseCurrency = "EUR"
	PayoutListResponseCurrencyFjd PayoutListResponseCurrency = "FJD"
	PayoutListResponseCurrencyFkp PayoutListResponseCurrency = "FKP"
	PayoutListResponseCurrencyGbp PayoutListResponseCurrency = "GBP"
	PayoutListResponseCurrencyGel PayoutListResponseCurrency = "GEL"
	PayoutListResponseCurrencyGhs PayoutListResponseCurrency = "GHS"
	PayoutListResponseCurrencyGip PayoutListResponseCurrency = "GIP"
	PayoutListResponseCurrencyGmd PayoutListResponseCurrency = "GMD"
	PayoutListResponseCurrencyGnf PayoutListResponseCurrency = "GNF"
	PayoutListResponseCurrencyGtq PayoutListResponseCurrency = "GTQ"
	PayoutListResponseCurrencyGyd PayoutListResponseCurrency = "GYD"
	PayoutListResponseCurrencyHkd PayoutListResponseCurrency = "HKD"
	PayoutListResponseCurrencyHnl PayoutListResponseCurrency = "HNL"
	PayoutListResponseCurrencyHrk PayoutListResponseCurrency = "HRK"
	PayoutListResponseCurrencyHtg PayoutListResponseCurrency = "HTG"
	PayoutListResponseCurrencyHuf PayoutListResponseCurrency = "HUF"
	PayoutListResponseCurrencyIdr PayoutListResponseCurrency = "IDR"
	PayoutListResponseCurrencyIls PayoutListResponseCurrency = "ILS"
	PayoutListResponseCurrencyInr PayoutListResponseCurrency = "INR"
	PayoutListResponseCurrencyIqd PayoutListResponseCurrency = "IQD"
	PayoutListResponseCurrencyJmd PayoutListResponseCurrency = "JMD"
	PayoutListResponseCurrencyJod PayoutListResponseCurrency = "JOD"
	PayoutListResponseCurrencyJpy PayoutListResponseCurrency = "JPY"
	PayoutListResponseCurrencyKes PayoutListResponseCurrency = "KES"
	PayoutListResponseCurrencyKgs PayoutListResponseCurrency = "KGS"
	PayoutListResponseCurrencyKhr PayoutListResponseCurrency = "KHR"
	PayoutListResponseCurrencyKmf PayoutListResponseCurrency = "KMF"
	PayoutListResponseCurrencyKrw PayoutListResponseCurrency = "KRW"
	PayoutListResponseCurrencyKwd PayoutListResponseCurrency = "KWD"
	PayoutListResponseCurrencyKyd PayoutListResponseCurrency = "KYD"
	PayoutListResponseCurrencyKzt PayoutListResponseCurrency = "KZT"
	PayoutListResponseCurrencyLak PayoutListResponseCurrency = "LAK"
	PayoutListResponseCurrencyLbp PayoutListResponseCurrency = "LBP"
	PayoutListResponseCurrencyLkr PayoutListResponseCurrency = "LKR"
	PayoutListResponseCurrencyLrd PayoutListResponseCurrency = "LRD"
	PayoutListResponseCurrencyLsl PayoutListResponseCurrency = "LSL"
	PayoutListResponseCurrencyLyd PayoutListResponseCurrency = "LYD"
	PayoutListResponseCurrencyMad PayoutListResponseCurrency = "MAD"
	PayoutListResponseCurrencyMdl PayoutListResponseCurrency = "MDL"
	PayoutListResponseCurrencyMga PayoutListResponseCurrency = "MGA"
	PayoutListResponseCurrencyMkd PayoutListResponseCurrency = "MKD"
	PayoutListResponseCurrencyMmk PayoutListResponseCurrency = "MMK"
	PayoutListResponseCurrencyMnt PayoutListResponseCurrency = "MNT"
	PayoutListResponseCurrencyMop PayoutListResponseCurrency = "MOP"
	PayoutListResponseCurrencyMru PayoutListResponseCurrency = "MRU"
	PayoutListResponseCurrencyMur PayoutListResponseCurrency = "MUR"
	PayoutListResponseCurrencyMvr PayoutListResponseCurrency = "MVR"
	PayoutListResponseCurrencyMwk PayoutListResponseCurrency = "MWK"
	PayoutListResponseCurrencyMxn PayoutListResponseCurrency = "MXN"
	PayoutListResponseCurrencyMyr PayoutListResponseCurrency = "MYR"
	PayoutListResponseCurrencyMzn PayoutListResponseCurrency = "MZN"
	PayoutListResponseCurrencyNad PayoutListResponseCurrency = "NAD"
	PayoutListResponseCurrencyNgn PayoutListResponseCurrency = "NGN"
	PayoutListResponseCurrencyNio PayoutListResponseCurrency = "NIO"
	PayoutListResponseCurrencyNok PayoutListResponseCurrency = "NOK"
	PayoutListResponseCurrencyNpr PayoutListResponseCurrency = "NPR"
	PayoutListResponseCurrencyNzd PayoutListResponseCurrency = "NZD"
	PayoutListResponseCurrencyOmr PayoutListResponseCurrency = "OMR"
	PayoutListResponseCurrencyPab PayoutListResponseCurrency = "PAB"
	PayoutListResponseCurrencyPen PayoutListResponseCurrency = "PEN"
	PayoutListResponseCurrencyPgk PayoutListResponseCurrency = "PGK"
	PayoutListResponseCurrencyPhp PayoutListResponseCurrency = "PHP"
	PayoutListResponseCurrencyPkr PayoutListResponseCurrency = "PKR"
	PayoutListResponseCurrencyPln PayoutListResponseCurrency = "PLN"
	PayoutListResponseCurrencyPyg PayoutListResponseCurrency = "PYG"
	PayoutListResponseCurrencyQar PayoutListResponseCurrency = "QAR"
	PayoutListResponseCurrencyRon PayoutListResponseCurrency = "RON"
	PayoutListResponseCurrencyRsd PayoutListResponseCurrency = "RSD"
	PayoutListResponseCurrencyRub PayoutListResponseCurrency = "RUB"
	PayoutListResponseCurrencyRwf PayoutListResponseCurrency = "RWF"
	PayoutListResponseCurrencySar PayoutListResponseCurrency = "SAR"
	PayoutListResponseCurrencySbd PayoutListResponseCurrency = "SBD"
	PayoutListResponseCurrencyScr PayoutListResponseCurrency = "SCR"
	PayoutListResponseCurrencySek PayoutListResponseCurrency = "SEK"
	PayoutListResponseCurrencySgd PayoutListResponseCurrency = "SGD"
	PayoutListResponseCurrencyShp PayoutListResponseCurrency = "SHP"
	PayoutListResponseCurrencySle PayoutListResponseCurrency = "SLE"
	PayoutListResponseCurrencySll PayoutListResponseCurrency = "SLL"
	PayoutListResponseCurrencySos PayoutListResponseCurrency = "SOS"
	PayoutListResponseCurrencySrd PayoutListResponseCurrency = "SRD"
	PayoutListResponseCurrencySsp PayoutListResponseCurrency = "SSP"
	PayoutListResponseCurrencyStn PayoutListResponseCurrency = "STN"
	PayoutListResponseCurrencySvc PayoutListResponseCurrency = "SVC"
	PayoutListResponseCurrencySzl PayoutListResponseCurrency = "SZL"
	PayoutListResponseCurrencyThb PayoutListResponseCurrency = "THB"
	PayoutListResponseCurrencyTnd PayoutListResponseCurrency = "TND"
	PayoutListResponseCurrencyTop PayoutListResponseCurrency = "TOP"
	PayoutListResponseCurrencyTry PayoutListResponseCurrency = "TRY"
	PayoutListResponseCurrencyTtd PayoutListResponseCurrency = "TTD"
	PayoutListResponseCurrencyTwd PayoutListResponseCurrency = "TWD"
	PayoutListResponseCurrencyTzs PayoutListResponseCurrency = "TZS"
	PayoutListResponseCurrencyUah PayoutListResponseCurrency = "UAH"
	PayoutListResponseCurrencyUgx PayoutListResponseCurrency = "UGX"
	PayoutListResponseCurrencyUsd PayoutListResponseCurrency = "USD"
	PayoutListResponseCurrencyUyu PayoutListResponseCurrency = "UYU"
	PayoutListResponseCurrencyUzs PayoutListResponseCurrency = "UZS"
	PayoutListResponseCurrencyVes PayoutListResponseCurrency = "VES"
	PayoutListResponseCurrencyVnd PayoutListResponseCurrency = "VND"
	PayoutListResponseCurrencyVuv PayoutListResponseCurrency = "VUV"
	PayoutListResponseCurrencyWst PayoutListResponseCurrency = "WST"
	PayoutListResponseCurrencyXaf PayoutListResponseCurrency = "XAF"
	PayoutListResponseCurrencyXcd PayoutListResponseCurrency = "XCD"
	PayoutListResponseCurrencyXof PayoutListResponseCurrency = "XOF"
	PayoutListResponseCurrencyXpf PayoutListResponseCurrency = "XPF"
	PayoutListResponseCurrencyYer PayoutListResponseCurrency = "YER"
	PayoutListResponseCurrencyZar PayoutListResponseCurrency = "ZAR"
	PayoutListResponseCurrencyZmw PayoutListResponseCurrency = "ZMW"
)

func (r PayoutListResponseCurrency) IsKnown() bool {
	switch r {
	case PayoutListResponseCurrencyAed, PayoutListResponseCurrencyAll, PayoutListResponseCurrencyAmd, PayoutListResponseCurrencyAng, PayoutListResponseCurrencyAoa, PayoutListResponseCurrencyArs, PayoutListResponseCurrencyAud, PayoutListResponseCurrencyAwg, PayoutListResponseCurrencyAzn, PayoutListResponseCurrencyBam, PayoutListResponseCurrencyBbd, PayoutListResponseCurrencyBdt, PayoutListResponseCurrencyBgn, PayoutListResponseCurrencyBhd, PayoutListResponseCurrencyBif, PayoutListResponseCurrencyBmd, PayoutListResponseCurrencyBnd, PayoutListResponseCurrencyBob, PayoutListResponseCurrencyBrl, PayoutListResponseCurrencyBsd, PayoutListResponseCurrencyBwp, PayoutListResponseCurrencyByn, PayoutListResponseCurrencyBzd, PayoutListResponseCurrencyCad, PayoutListResponseCurrencyChf, PayoutListResponseCurrencyClp, PayoutListResponseCurrencyCny, PayoutListResponseCurrencyCop, PayoutListResponseCurrencyCrc, PayoutListResponseCurrencyCup, PayoutListResponseCurrencyCve, PayoutListResponseCurrencyCzk, PayoutListResponseCurrencyDjf, PayoutListResponseCurrencyDkk, PayoutListResponseCurrencyDop, PayoutListResponseCurrencyDzd, PayoutListResponseCurrencyEgp, PayoutListResponseCurrencyEtb, PayoutListResponseCurrencyEur, PayoutListResponseCurrencyFjd, PayoutListResponseCurrencyFkp, PayoutListResponseCurrencyGbp, PayoutListResponseCurrencyGel, PayoutListResponseCurrencyGhs, PayoutListResponseCurrencyGip, PayoutListResponseCurrencyGmd, PayoutListResponseCurrencyGnf, PayoutListResponseCurrencyGtq, PayoutListResponseCurrencyGyd, PayoutListResponseCurrencyHkd, PayoutListResponseCurrencyHnl, PayoutListResponseCurrencyHrk, PayoutListResponseCurrencyHtg, PayoutListResponseCurrencyHuf, PayoutListResponseCurrencyIdr, PayoutListResponseCurrencyIls, PayoutListResponseCurrencyInr, PayoutListResponseCurrencyIqd, PayoutListResponseCurrencyJmd, PayoutListResponseCurrencyJod, PayoutListResponseCurrencyJpy, PayoutListResponseCurrencyKes, PayoutListResponseCurrencyKgs, PayoutListResponseCurrencyKhr, PayoutListResponseCurrencyKmf, PayoutListResponseCurrencyKrw, PayoutListResponseCurrencyKwd, PayoutListResponseCurrencyKyd, PayoutListResponseCurrencyKzt, PayoutListResponseCurrencyLak, PayoutListResponseCurrencyLbp, PayoutListResponseCurrencyLkr, PayoutListResponseCurrencyLrd, PayoutListResponseCurrencyLsl, PayoutListResponseCurrencyLyd, PayoutListResponseCurrencyMad, PayoutListResponseCurrencyMdl, PayoutListResponseCurrencyMga, PayoutListResponseCurrencyMkd, PayoutListResponseCurrencyMmk, PayoutListResponseCurrencyMnt, PayoutListResponseCurrencyMop, PayoutListResponseCurrencyMru, PayoutListResponseCurrencyMur, PayoutListResponseCurrencyMvr, PayoutListResponseCurrencyMwk, PayoutListResponseCurrencyMxn, PayoutListResponseCurrencyMyr, PayoutListResponseCurrencyMzn, PayoutListResponseCurrencyNad, PayoutListResponseCurrencyNgn, PayoutListResponseCurrencyNio, PayoutListResponseCurrencyNok, PayoutListResponseCurrencyNpr, PayoutListResponseCurrencyNzd, PayoutListResponseCurrencyOmr, PayoutListResponseCurrencyPab, PayoutListResponseCurrencyPen, PayoutListResponseCurrencyPgk, PayoutListResponseCurrencyPhp, PayoutListResponseCurrencyPkr, PayoutListResponseCurrencyPln, PayoutListResponseCurrencyPyg, PayoutListResponseCurrencyQar, PayoutListResponseCurrencyRon, PayoutListResponseCurrencyRsd, PayoutListResponseCurrencyRub, PayoutListResponseCurrencyRwf, PayoutListResponseCurrencySar, PayoutListResponseCurrencySbd, PayoutListResponseCurrencyScr, PayoutListResponseCurrencySek, PayoutListResponseCurrencySgd, PayoutListResponseCurrencyShp, PayoutListResponseCurrencySle, PayoutListResponseCurrencySll, PayoutListResponseCurrencySos, PayoutListResponseCurrencySrd, PayoutListResponseCurrencySsp, PayoutListResponseCurrencyStn, PayoutListResponseCurrencySvc, PayoutListResponseCurrencySzl, PayoutListResponseCurrencyThb, PayoutListResponseCurrencyTnd, PayoutListResponseCurrencyTop, PayoutListResponseCurrencyTry, PayoutListResponseCurrencyTtd, PayoutListResponseCurrencyTwd, PayoutListResponseCurrencyTzs, PayoutListResponseCurrencyUah, PayoutListResponseCurrencyUgx, PayoutListResponseCurrencyUsd, PayoutListResponseCurrencyUyu, PayoutListResponseCurrencyUzs, PayoutListResponseCurrencyVes, PayoutListResponseCurrencyVnd, PayoutListResponseCurrencyVuv, PayoutListResponseCurrencyWst, PayoutListResponseCurrencyXaf, PayoutListResponseCurrencyXcd, PayoutListResponseCurrencyXof, PayoutListResponseCurrencyXpf, PayoutListResponseCurrencyYer, PayoutListResponseCurrencyZar, PayoutListResponseCurrencyZmw:
		return true
	}
	return false
}

type PayoutListResponseStatus string

const (
	PayoutListResponseStatusInProgress PayoutListResponseStatus = "in_progress"
	PayoutListResponseStatusFailed     PayoutListResponseStatus = "failed"
	PayoutListResponseStatusSuccess    PayoutListResponseStatus = "success"
)

func (r PayoutListResponseStatus) IsKnown() bool {
	switch r {
	case PayoutListResponseStatusInProgress, PayoutListResponseStatusFailed, PayoutListResponseStatusSuccess:
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
