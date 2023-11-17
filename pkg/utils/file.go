package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
)

func encodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func compress(s []byte) []byte {

	zipbuf := bytes.Buffer{}
	zipped := gzip.NewWriter(&zipbuf)
	zipped.Write(s)
	zipped.Close()
	fmt.Println("compressed size (bytes): ", len(zipbuf.Bytes()))
	return zipbuf.Bytes()
}

func Decompress(s []byte) []byte {

	rdr, _ := gzip.NewReader(bytes.NewReader(s))
	data, err := io.ReadAll(rdr)
	if err != nil {
		log.Fatal(err)
	}
	rdr.Close()
	fmt.Println("uncompressed size (bytes): ", len(data))
	return data
}

func WriteToFile(p interface{}, file string) {
	dataOut := encodeToBytes(p)
	dataOut = compress(dataOut)

	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}

	f.Write(dataOut)
}

func ReadFromFile(path string) []byte {

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func DecodeToTransaction(s []byte, target interface{}) {

	//trans := transactions.Transaction{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&target)
	if err != nil {
		log.Fatal(err)
	}
}
