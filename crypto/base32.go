package crypto

import "encoding/base32"

var (
  Base32Encoder = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)
)
