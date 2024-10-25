// SPDX-FileCopyrightText: Copyright 2024 The Minder Authors
// SPDX-License-Identifier: Apache-2.0

package server

// CryptoConfig is the configuration for the crypto engine
type CryptoConfig struct {
	KeyStore KeyStoreConfig `mapstructure:"keystore"`
	Default  DefaultCrypto  `mapstructure:"default"`
	Fallback FallbackCrypto `mapstructure:"fallback"`
}

// KeyStoreConfig specifies the type of keystore to use and its configuration
// If we support multiple types of keystore (e.g. local, aws secrets, vault, etc.)
// it is intended that there will be one field for each type of keystore config,
// and that the `Type` field specifies which one to use
type KeyStoreConfig struct {
	Type  string              `mapstructure:"type" default:"local"`
	Local LocalKeyStoreConfig `mapstructure:"local"`
}

// DefaultCrypto defines the default crypto to be used for new data
type DefaultCrypto struct {
	// `token_key_passphrase` is the filename generated by `make bootstrap`
	KeyID string `mapstructure:"key_id"`
}

// FallbackCrypto defines the optional key and algorithm which can be used for
// decrypting old secrets.
// This is used for rotating keys or algorithms.
type FallbackCrypto struct {
	KeyID string `mapstructure:"key_id"`
}

// LocalKeyStoreConfig contains configuration for the local file keystore
type LocalKeyStoreConfig struct {
	// `./.ssh/` is the directory generated by `make bootstrap`
	KeyDir string `mapstructure:"key_dir" default:"./.ssh/"`
}
