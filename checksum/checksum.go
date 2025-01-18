package checksum

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
)

func NewChecksum(algorithm string, r *bufio.Reader) ([]byte, error) {
	var h hash.Hash
	switch algorithm {
	case "sha1":
		h = sha1.New()

	case "sha256":
		h = sha256.New()

	case "sha512":
		h = sha512.New()

	case "md5":
		h = md5.New()
	default:
	}

	if _, err := io.Copy(h, r); err != nil {
		return nil, fmt.Errorf("read: %v", err)
	}
	return h.Sum(nil), nil
}
