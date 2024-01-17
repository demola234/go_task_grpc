package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReturnsStringOfLengthN(t *testing.T) {
	// Arrange
	n := 10

	// Act
	result := RandomString(n)

	// Assert
	require.Equal(t, n, len(result))
}

func TestReturnsRandomEmail(t *testing.T) {
	// Arrange
	// Act
	result := RandomEmail()

	// Assert
	require.NotEmpty(t, result)
	require.Contains(t, result, "@")
	require.Contains(t, result, ".")
}

func TestReturnsRandomOwner(t *testing.T) {
	// Arrange
	// Act
	result := RandomOwner()

	// Assert
	require.NotEmpty(t, result)
}

func TestSplitS(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	result := SplitStrings(s)

	// Assert
	require.NotEmpty(t, result)
	require.Equal(t, len(s), len(result))
}

func TestReturnsRandomMoney(t *testing.T) {
	// Arrange
	// Act
	result := RandomMoney()

	// Assert
	require.NotEmpty(t, result)
}
