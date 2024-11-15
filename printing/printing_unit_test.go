package printing_test

import (
	"fmt"
	"testing"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/printing"

	"github.com/stretchr/testify/require"
)

func TestHandlerTag(t *testing.T) {
	output := printing.HandlerTag("<C10>ajdhaf")
	require.Equal(t, output, "C10")
	output = printing.HandlerTag("<C> asdwoeiur")
	require.Equal(t, output, "C")
	output = printing.HandlerTag("<F> jkadshfl")
	require.Equal(t, output, "F")
}

func TestHandlerBar(t *testing.T) {
	output := printing.HandlerBar("-")
	require.Equal(t, output, printing.Dash)
	output = printing.HandlerBar("=")
	require.Equal(t, output, printing.DoubleDash)
	output = printing.HandlerBar(" ")
	require.Equal(t, output, printing.Space)
}

func TestHandlerHeader(t *testing.T) {
	tag, body, height, width, err := printing.HandlerHeader("<C10>ajdhaf")
	require.NoError(t, err)
	require.Equal(t, "C", tag)
	require.Equal(t, "ajdhaf", body)
	require.Equal(t, 1, height)
	require.Equal(t, 0, width)

	tag, body, height, width, err = printing.HandlerHeader("<C>ajdhaf")
	require.NoError(t, err)
	require.Equal(t, "C", tag)
	require.Equal(t, "ajdhaf", body)
	require.Equal(t, -1, height)
	require.Equal(t, -1, width)

	_, _, _, _, err = printing.HandlerHeader("<Cxy>ajdhaf")
	require.EqualError(t, err, "invalid number")
}

func TestHandleOrder(t *testing.T) {
	output, err := printing.HandleOrder(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}

func TestHandleDiscount(t *testing.T) {
	output, err := printing.HandleDiscount(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}

func TestHandlePayment(t *testing.T) {
	output, err := printing.HandlePayment(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}

func TestHandleService(t *testing.T) {
	output, err := printing.HandleService(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}

func TestHandleTax(t *testing.T) {
	output, err := printing.HandleTax(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}

func TestHandleEntity(t *testing.T) {
	output, err := printing.HandleEntity(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}

func TestHandleLayout(t *testing.T) {
	output, _, err := printing.HandleLayout(printing.LayoutTest)
	require.NoError(t, err)
	fmt.Println(output)
}
