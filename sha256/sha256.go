package sha256

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

func GenerateSha256Hash(r io.Reader) ([]byte, error) {
	br := bufio.NewReader(r)
	data := make([]byte, 0)

	for {
		b, err := br.ReadByte()
		if err != nil {
			log.Printf("Error reading %s: %v\n", r, err)
			break
		}
		data = append(data, b)
	}

	h := sha256.New()

	n, err := h.Write(data)
	if err != nil {
		log.Printf("Error writing to hash: %v\n", err)
		return nil, err
	}
	fmt.Printf("wrote %d bytes to hash\n", n)
	return h.Sum(nil), nil
}

func SaveHashOutput(w io.Writer, hash []byte) error {
	bw := bufio.NewWriter(w)

	n, err := bw.Write(hash)
	if err != nil && err != io.EOF {
		log.Printf("Error writing hash &v: %v\n", hash, err)
		return err
	}
	fmt.Printf("%d bytes written\n", n)
	err = bw.Flush()
	if err != nil && err != io.EOF {
		log.Printf("Error flushing write buffer: %v\n", err)
		return err
	}
	return nil
}
