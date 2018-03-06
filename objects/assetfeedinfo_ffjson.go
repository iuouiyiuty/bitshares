// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetfeedinfo.go

package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetFeedInfo) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AssetFeedInfo) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"maintenance_collateral_ratio":`)
	fflib.FormatBits2(buf, uint64(j.MaintenanceCollateralRatio), 10, false)
	buf.WriteString(`,"maximum_short_squeeze_ratio":`)
	fflib.FormatBits2(buf, uint64(j.MaximumShortSqueezeRatio), 10, false)
	/* Struct fall back. type=objects.Price kind=struct */
	buf.WriteString(`,"settlement_price":`)
	err = buf.Encode(&j.SettlementPrice)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.Price kind=struct */
	buf.WriteString(`,"core_exchange_rate":`)
	err = buf.Encode(&j.CoreExchangeRate)
	if err != nil {
		return err
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAssetFeedInfobase = iota
	ffjtAssetFeedInfonosuchkey

	ffjtAssetFeedInfoMaintenanceCollateralRatio

	ffjtAssetFeedInfoMaximumShortSqueezeRatio

	ffjtAssetFeedInfoSettlementPrice

	ffjtAssetFeedInfoCoreExchangeRate
)

var ffjKeyAssetFeedInfoMaintenanceCollateralRatio = []byte("maintenance_collateral_ratio")

var ffjKeyAssetFeedInfoMaximumShortSqueezeRatio = []byte("maximum_short_squeeze_ratio")

var ffjKeyAssetFeedInfoSettlementPrice = []byte("settlement_price")

var ffjKeyAssetFeedInfoCoreExchangeRate = []byte("core_exchange_rate")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetFeedInfo) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetFeedInfo) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetFeedInfobase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAssetFeedInfonosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyAssetFeedInfoCoreExchangeRate, kn) {
						currentKey = ffjtAssetFeedInfoCoreExchangeRate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyAssetFeedInfoMaintenanceCollateralRatio, kn) {
						currentKey = ffjtAssetFeedInfoMaintenanceCollateralRatio
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetFeedInfoMaximumShortSqueezeRatio, kn) {
						currentKey = ffjtAssetFeedInfoMaximumShortSqueezeRatio
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyAssetFeedInfoSettlementPrice, kn) {
						currentKey = ffjtAssetFeedInfoSettlementPrice
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffjKeyAssetFeedInfoCoreExchangeRate, kn) {
					currentKey = ffjtAssetFeedInfoCoreExchangeRate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetFeedInfoSettlementPrice, kn) {
					currentKey = ffjtAssetFeedInfoSettlementPrice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetFeedInfoMaximumShortSqueezeRatio, kn) {
					currentKey = ffjtAssetFeedInfoMaximumShortSqueezeRatio
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAssetFeedInfoMaintenanceCollateralRatio, kn) {
					currentKey = ffjtAssetFeedInfoMaintenanceCollateralRatio
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAssetFeedInfonosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAssetFeedInfoMaintenanceCollateralRatio:
					goto handle_MaintenanceCollateralRatio

				case ffjtAssetFeedInfoMaximumShortSqueezeRatio:
					goto handle_MaximumShortSqueezeRatio

				case ffjtAssetFeedInfoSettlementPrice:
					goto handle_SettlementPrice

				case ffjtAssetFeedInfoCoreExchangeRate:
					goto handle_CoreExchangeRate

				case ffjtAssetFeedInfonosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_MaintenanceCollateralRatio:

	/* handler: j.MaintenanceCollateralRatio type=objects.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MaintenanceCollateralRatio.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaximumShortSqueezeRatio:

	/* handler: j.MaximumShortSqueezeRatio type=objects.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MaximumShortSqueezeRatio.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SettlementPrice:

	/* handler: j.SettlementPrice type=objects.Price kind=struct quoted=false*/

	{
		/* Falling back. type=objects.Price kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.SettlementPrice)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CoreExchangeRate:

	/* handler: j.CoreExchangeRate type=objects.Price kind=struct quoted=false*/

	{
		/* Falling back. type=objects.Price kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.CoreExchangeRate)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
