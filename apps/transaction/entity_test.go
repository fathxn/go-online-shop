package transaction

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetSubTotal(t *testing.T) {
	var trx = Transaction{
		ProductPrice: 10_000,
		Amount:       10,
	}
	expected := uint(100_000)

	trx.SetSubTotal()

	require.Equal(t, expected, trx.SubTotal)
}

func TestGrandTotal(t *testing.T) {
	t.Run("without platform fee", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}
		expected := uint(100_000)

		trx.SetSubTotal()
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})

	t.Run("without set sub total first", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}
		expected := uint(100_000)
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})

	t.Run("with platform fee", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
			PlatformFee:  1_000,
		}
		expected := uint(101_000)

		trx.SetSubTotal()
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})
}
