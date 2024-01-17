package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()
	config, err := LoadConfig("../")
	if err != nil {
		t.Error("Failed to load config")
	}
	expectedConfig := Config{
		DBDriver:             config.DBDriver,
		DBSource:             config.DBSource,
		HTTPServerAddress:    config.HTTPServerAddress,
		GRPCServerAddress:    config.GRPCServerAddress,
		AccessTokenDuration:  15 * time.Minute,
		TokenSymmetricKey:    config.TokenSymmetricKey,
		RefreshTokenDuration: 24 * time.Hour,
		Enviroment:           config.Enviroment,
	}

	t.Logf(config.DBDriver)

	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, expectedConfig, config)
}
