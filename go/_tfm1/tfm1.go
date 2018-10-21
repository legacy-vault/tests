package main

import (
	"encoding/base64"
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/twofish"
)

const BytesInblocks = 3 * 16 // 3 Blocks of 16 Bytes = 48 Bytes
const BlockSize = 16         // 16 Bytes
const KeySize = 32           // 32 Bytes

const InputMaxLen = 64 - 1 // Maximum Length of Text String that can be encoded
// 64 Symbols = 48 Bytes (4 Symbols = 3 Bytes), 1 for Length.

const InputMinLen = 8

const AllowedSymbols string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789*."

// Base64 (0..63): "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
// Allowed Characters are similar to Base64 with one Exception:
// last Characters are "*" and "." instead of "+" and "/".

//------------------------------------------------------------------------------

type Encoder struct {
	Key   []byte
	BG    []byte
	Block []byte
}

//------------------------------------------------------------------------------

func (enc Encoder) NewEncoder(key []byte) (e *Encoder) {

	var i int

	e = new(Encoder)

	if len(key) != KeySize {
		log.Println("Bad Key Length.") //
		return nil
	}

	// Save Key
	e.Key = make([]byte, KeySize)
	copy(e.Key, key)

	// Create BG
	rand.Seed(time.Now().UnixNano())
	enc.BG = make([]byte, BytesInblocks)
	for i = 0; i < BytesInblocks; i++ {
		enc.BG[i] = uint8(rand.Int())
	}

	// Init Block
	e.Block = make([]byte, BytesInblocks)

	return e
}

//------------------------------------------------------------------------------

func (enc Encoder) Encode(input, encoded []byte) (ok bool) {

	var li int = len(input)
	var le int = len(encoded)

	if (li > InputMaxLen) || (li < InputMinLen) {
		log.Println("Input has wrong Length.") //
		return false
	}

	if le != BytesInblocks {
		log.Println("Receiver has bad Length.") //
		return false
	}

	ok = enc.prepareBlocks(input)
	if !ok {
		log.Println("Preparation failed.") //
		return false
	}

	ok = enc.encryptBlocks(input, encoded)
	if !ok {
		log.Println("Encoding failed.") //
		return false
	}

	return true
}

//------------------------------------------------------------------------------

func (enc Encoder) Decode(input, decoded []byte) (ok bool) {

	var li int = len(input)
	var ld int = len(decoded)
	var tmp []byte = make([]byte, BytesInblocks)

	if (li != BytesInblocks) || (ld != BytesInblocks) {
		log.Println("Bad Length specified.") //
		return false
	}

	ok = enc.decryptBlocks(input, tmp)
	if !ok {
		log.Println("Decoding failed.") //
		return false
	}

	enc.getString(tmp, decoded)

	return true
}

//------------------------------------------------------------------------------

func (enc Encoder) prepareBlocks(src []byte) (ok bool) {

	// Reads a String, converts it into an Array of Bytes and lays it upon the
	// background Array of Bytes. It is recommended that Number of symbols in
	// String is multiple (divisible by) Four (4). Length of a String must be
	// not more than 64 allowed Symbols.

	var l, m, i, j, n, data_len int
	var srcB, data []byte
	var srcS string
	var b1, b2, b3, b4, tb byte
	var err error

	b1 = []byte("*")[0]
	b2 = []byte("+")[0]
	b3 = []byte(".")[0]
	b4 = []byte("/")[0]

	l = len(src)
	m = len(AllowedSymbols)
	ok = false

	// Check allowed Symbols and convert into Base64
	for i = 0; i < l; i++ {
		ok = false
		for j = 0; j < m; j++ {
			if srcB[i] == AllowedSymbols[j] {
				ok = true
			}
		}
		if !ok {
			log.Println("Bad Symbol found at pos #", i+1) //
			return ok
		}
	}

	// Replace two Symbols to make it Base64 compatible
	for i = 0; i < l; i++ {
		if srcB[i] == b1 { // "*" -> "+"
			srcB[i] = b2
		}
		if srcB[i] == b3 { // "." -> "/"
			srcB[i] = b4
		}
	}

	// Convert into String
	srcS = string(srcB)

	// Push Length
	srcS = string(AllowedSymbols[l]) + srcS

	// Update Length
	l = len(srcS)

	// Add '+++' Ending if needed
	if l%4 == 1 {

		srcS = srcS + "+++"

	} else if l%4 == 2 {

		srcS = srcS + "++"

	} else if l%4 == 3 {

		srcS = srcS + "+"

	}

	// Base64 String -> []byte
	data, err = base64.StdEncoding.DecodeString(srcS)
	if err != nil {
		log.Println("Conversion Base64 to []byte failed.", err) //
		return false
	}

	// Fill DST with BG
	n = copy(enc.Block[0:BytesInblocks], enc.BG[0:BytesInblocks])
	if n != len(enc.BG) {
		log.Println("Copy BG failed.") //
		return false
	}

	// Add converted SRC to BG
	data_len = len(data)
	n = copy((enc.Block)[0:l], data[0:data_len])
	if n != data_len {
		log.Println("Copy Data failed.") //
		return false
	}

	// Rotate
	i = 0
	for j = 0; j < 16; j++ {
		if j%3 == 1 {

			tb = (enc.Block)[16*(i+2)+j]
			enc.Block[16*(i+2)+j] = enc.Block[16*(i+1)+j]
			enc.Block[16*(i+1)+j] = enc.Block[16*(i)+j]
			enc.Block[16*(i)+j] = tb

		} else if j%3 == 2 {

			tb = enc.Block[16*(i+2)+j]
			enc.Block[16*(i+2)+j] = enc.Block[16*(i)+j]
			enc.Block[16*(i)+j] = enc.Block[16*(i+1)+j]
			enc.Block[16*(i+1)+j] = tb

		}
	}

	return ok
}

//------------------------------------------------------------------------------

func (enc Encoder) encryptBlocks(src, dst []byte) (ok bool) {

	var cypher *twofish.Cipher
	var err error
	var key, a, b, c, d, e, f []byte
	var i, r int
	var tb byte

	// Local Key
	key = make([]byte, KeySize)
	copy(key, enc.Key) // Local key will be modified

	// Prepare Data
	cypher, err = twofish.NewCipher(key)
	if err != nil {
		log.Println("Error creating Cypher.", err)
		return false
	}
	a = make([]byte, BlockSize)
	b = make([]byte, BlockSize)
	c = make([]byte, BlockSize)
	d = make([]byte, BlockSize)
	e = make([]byte, BlockSize)
	f = make([]byte, BlockSize)

	// Read Input
	copy(a, src[0:BlockSize])
	copy(b, src[BlockSize:BlockSize*2])
	copy(c, src[BlockSize*2:BlockSize*3])

	// Pass I
	cypher.Encrypt(d, a)
	cypher.Encrypt(e, b)
	cypher.Encrypt(f, c)

	// Pass II
	for r = 0; r < KeySize-1; r++ {

		// Rotate Key
		tb = key[0]
		for i = 1; i < KeySize; i++ {
			key[i-1] = key[i]
		}
		key[KeySize-1] = tb

		// Set Key
		cypher, err = twofish.NewCipher(key)
		if err != nil {
			log.Println("Error creating Cypher.", err)
			return false
		}

		// Set Src
		copy(a, d)
		copy(b, e)
		copy(c, f)

		// Encode
		cypher.Encrypt(d, a)
		cypher.Encrypt(e, b)
		cypher.Encrypt(f, c)
	}

	// Output
	copy(dst[0:BlockSize], d)
	copy(dst[BlockSize:BlockSize*2], e)
	copy(dst[BlockSize*2:BlockSize*3], f)

	return true
}

//------------------------------------------------------------------------------

func (enc Encoder) decryptBlocks(src, dst []byte) (ok bool) {

	var cypher *twofish.Cipher
	var err error
	var key, a, b, c, d, e, f []byte
	var tb byte
	var i, r int

	// Local Key
	key = make([]byte, KeySize)
	copy(key, enc.Key) // Local key will be modified

	// Prepare Data
	a = make([]byte, BlockSize)
	b = make([]byte, BlockSize)
	c = make([]byte, BlockSize)
	d = make([]byte, BlockSize)
	e = make([]byte, BlockSize)
	f = make([]byte, BlockSize)

	// Read Input
	copy(d, src[0:BlockSize])
	copy(e, src[BlockSize:BlockSize*2])
	copy(f, src[BlockSize*2:BlockSize*3])

	// Anti-Pass II
	for r = 0; r < KeySize-1; r++ {

		// Anti-Rotate Key
		tb = key[KeySize-1]
		for i = KeySize - 1; i > 0; i-- {
			key[i] = key[i-1]
		}
		key[0] = tb

		// Set Key
		cypher, err = twofish.NewCipher(key)
		if err != nil {
			log.Println("Error creating Cypher.", err)
			return false
		}

		// Decode
		cypher.Decrypt(a, d)
		cypher.Decrypt(b, e)
		cypher.Decrypt(c, f)

		// Set
		copy(d, a)
		copy(e, b)
		copy(f, c)
	}

	// Anti-Rotate Key to return to original
	tb = key[KeySize-1]
	for i = KeySize - 1; i > 0; i-- {
		key[i] = key[i-1]
	}
	key[0] = tb

	cypher, err = twofish.NewCipher(key)
	if err != nil {
		log.Println("Error creating Cypher.", err)
		return false
	}

	// Anti-Pass I
	cypher.Decrypt(a, d)
	cypher.Decrypt(b, e)
	cypher.Decrypt(c, f)

	// Output
	copy(dst[0:BlockSize], a)
	copy(dst[BlockSize:BlockSize*2], b)
	copy(dst[BlockSize*2:BlockSize*3], c)

	return true
}

//------------------------------------------------------------------------------

func (enc Encoder) getString(input, out []byte) {

	var i, j, l int
	var tb byte
	var tmpStr string

	i = 0
	for j = 0; j < 16; j++ {
		if j%3 == 2 {

			tb = input[16*(i+2)+j]
			input[16*(i+2)+j] = input[16*(i+1)+j]
			input[16*(i+1)+j] = input[16*(i)+j]
			input[16*(i)+j] = tb

		} else if j%3 == 1 {

			tb = input[16*(i+2)+j]
			input[16*(i+2)+j] = input[16*(i)+j]
			input[16*(i)+j] = input[16*(i+1)+j]
			input[16*(i+1)+j] = tb

		}
	}

	// []byte -> Base64 string
	tmpStr = base64.StdEncoding.EncodeToString(input[0:BytesInblocks])
	l = strings.Index(AllowedSymbols, string(tmpStr[0]))
	out = []byte(tmpStr[1 : l+1])
}
