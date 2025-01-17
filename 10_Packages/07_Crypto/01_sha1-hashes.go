package main

/*

SHA1 hashes are frequently used to compute short
identities for binary or text blobs.
For example, the git revision control system uses
SHA1s extensively to identify versioned files and directories.
*/
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "Golang is crazy langauge"
	// s := "Golang"

	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}