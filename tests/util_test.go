package bank

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func requireErrorIsNoPanic(t *testing.T, expected error, f func() error) {
	t.Helper()
	var err error
	require.NotPanics(t, func() {
		err = f()
	})
	require.ErrorIs(t, err, expected)

}

func requireBalanceErrorIsNoPanic(t *testing.T, expected error, f func() (int64, error)) {
	t.Helper()
	var err error
	require.NotPanics(t, func() {
		_, err = f()
	})
	require.ErrorIs(t, err, expected)

}

func requireErrorContainsNoPanic(t *testing.T, expected string, f func() error) {
	t.Helper()

	var err error
	require.NotPanics(t, func() {
		err = f()
	})
	require.Error(t, err)
	require.ErrorContains(t, err, expected)
}

func requireBalanceErrorContainsNoPanic(t *testing.T, expected string, f func() (int64, error)) {
	t.Helper()

	var err error
	require.NotPanics(t, func() {
		_, err = f()
	})
	require.Error(t, err)
	require.ErrorContains(t, err, expected)
}
