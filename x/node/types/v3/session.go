package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Session) GetDownloadBytes() sdkmath.Int { return m.DownloadBytes }
func (m *Session) GetUploadBytes() sdkmath.Int   { return m.UploadBytes }

func (m *Session) SetDownloadBytes(v sdkmath.Int) { m.DownloadBytes = v }
func (m *Session) SetDuration(v time.Duration)    { m.Duration = v }
func (m *Session) SetInactiveAt(v time.Time)      { m.InactiveAt = v }
func (m *Session) SetStatusAt(v time.Time)        { m.StatusAt = v }
func (m *Session) SetStatus(v v1base.Status)      { m.Status = v }
func (m *Session) SetUploadBytes(v sdkmath.Int)   { m.UploadBytes = v }

func (m *Session) paymentAmountForBytes() sdk.Coin {
	decPrice := m.Price.Amount.ToLegacyDec()
	bytePrice := decPrice.QuoInt(base.Gigabyte)
	totalBytes := m.DownloadBytes.Add(m.UploadBytes)
	amount := bytePrice.MulInt(totalBytes).Ceil().TruncateInt()

	return sdk.NewCoin(m.Price.Denom, amount)
}

func (m *Session) paymentAmountForDuration() sdk.Coin {
	decPrice := m.Price.Amount.ToLegacyDec()
	nsPrice := decPrice.QuoInt64(time.Hour.Nanoseconds())
	nsDuration := m.Duration.Nanoseconds()
	amount := nsPrice.MulInt64(nsDuration).Ceil().TruncateInt()

	return sdk.NewCoin(m.Price.Denom, amount)
}

func (m *Session) PaymentAmount() sdk.Coin {
	if !m.MaxBytes.IsZero() {
		return m.paymentAmountForBytes()
	}
	if m.MaxDuration != 0 {
		return m.paymentAmountForDuration()
	}

	return sdk.NewCoin(m.Price.Denom, sdk.ZeroInt())
}

func (m *Session) RefundAmount() sdk.Coin {
	payment := m.PaymentAmount()
	if payment.IsGTE(m.Deposit) {
		payment = m.Deposit
	}

	return m.Deposit.Sub(payment)
}
