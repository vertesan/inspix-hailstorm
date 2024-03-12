package crypto

import (
  "bytes"
  "crypto/aes"
  "crypto/cipher"
  "errors"
  "io"
  "log"
)

// PKCS7 errors.
var (
  // ErrInvalidBlockSize indicates hash blocksize <= 0.
  ErrInvalidBlockSize = errors.New("invalid blocksize")

  // ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
  ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

  // ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
  ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

func Decrypt(key []byte, iv []byte, reader io.Reader, writer io.Writer) {
  ciphertext, err := io.ReadAll(reader)
  if err != nil {
    panic(err)
  }

  block, err := aes.NewCipher(key)
  if err != nil {
    panic(err)
  }

  mode := cipher.NewCBCDecrypter(block, iv)

  // CryptBlocks can work in-place if the two arguments are the same.
  plainBytes := make([]byte, len(ciphertext))
  mode.CryptBlocks(plainBytes, ciphertext)

  // If the original plaintext lengths are not a multiple of the block
  // size, padding would have to be added when encrypting, which would be
  // removed at this point. For an example, see
  // https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
  // critical to note that ciphertexts must be authenticated (i.e. by
  // using crypto/hmac) before being decrypted in order to avoid creating
  // a padding oracle.

  plainBytes, err = pkcs7Unpad(plainBytes, 128)
  if err != nil {
    panic(err)
  }

  n, err := io.Copy(writer, bytes.NewBuffer(plainBytes))
  if err != nil {
    panic(err)
  }
  if n <= 0 {
    log.Panicf("%q bytes of data has been written.", n)
  }
}

// pkcs7Pad right-pads the given byte slice with 1 to n bytes, where
// n is the block size. The size of the result is x times n, where x
// is at least 1.
func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
  if blocksize <= 0 {
    return nil, ErrInvalidBlockSize
  }
  if len(b) == 0 {
    return nil, ErrInvalidPKCS7Data
  }
  blocksize = blocksize / 8
  n := blocksize - (len(b) % blocksize)
  pb := make([]byte, len(b)+n)
  copy(pb, b)
  copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
  return pb, nil
}

// pkcs7Unpad validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
  if blocksize <= 0 {
    panic(ErrInvalidBlockSize)
  }
  if len(b) == 0 {
    panic(ErrInvalidPKCS7Data)
  }
  blocksize = blocksize / 8
  if len(b)%blocksize != 0 {
    panic(ErrInvalidPKCS7Padding)
  }
  c := b[len(b)-1]
  n := int(c)
  if n == 0 || n > len(b) {
    panic(ErrInvalidPKCS7Padding)
  }
  for i := 0; i < n; i++ {
    if b[len(b)-n+i] != c {
      panic(ErrInvalidPKCS7Padding)
    }
  }
  return b[:len(b)-n], nil
}
