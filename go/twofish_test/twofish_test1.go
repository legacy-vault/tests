package tfm1

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/twofish"
)

//------------------------------------------------------------------------------

type tfm1Encoder struct {
	key []byte
	bg  []byte
}

//------------------------------------------------------------------------------

func (enc tfm1Encoder) init() {

	const blockSize = 16
	const nMax = 3 * blockSize // 3 Blocks = 48 Bytes

	//
}

//------------------------------------------------------------------------------

func main() {

	const blockSize = 16
	const dataSize = 3 * blockSize // 3 Blocks = 48 Bytes

	var main_key []byte
	var s, s2 string
	var res bool
	var salt, block, encoded, decoded []byte

	rand.Seed(time.Now().UnixNano())

	salt = make([]byte, dataSize)
	salt = *tfm1_createBg()
	fmt.Println("salt=", salt) //

	s = "HelloHowAreYou"
	fmt.Println("\ns=", []byte(s)) //

	main_key = []byte("Crazy is just a Test of crazy Fu")
	fmt.Println("\nKey in Main =", main_key) //

	block = make([]byte, dataSize)
	encoded = make([]byte, dataSize)
	decoded = make([]byte, dataSize)

	res = tfm1_prepare3Blocks(&s, &block, &salt)
	if !res {
		fmt.Println("\nError in Prepare.") //
		return
	}

	res = tfm1_encrypt3Blocks(&block, &encoded, main_key)
	if !res {
		fmt.Println("\nError in Encrypt.") //
		return
	}
	fmt.Println("\nencoded=", encoded) //

	tfm1_decrypt3Blocks(&encoded, &decoded, main_key)
	if !res {
		fmt.Println("\nError in Decrypt.") //
		return
	}
	fmt.Println("\ndecoded=", decoded) //

	tfm1_getString(&decoded, &s2)
	fmt.Println("\ns2=", s2) //

}

//------------------------------------------------------------------------------

func tfm1_createBg() (bg *[]byte) {

	var i int
	var arr []byte

	arr = make([]byte, nMax)

	for i = 0; i < nMax; i++ {
		arr[i] = uint8(rand.Int())
	}

	return &arr
}

//------------------------------------------------------------------------------

func tfm1_prepare3Blocks(src *string, block, bg *[]byte) (ok bool) {

	// Reads a String, converts it into an Array of Bytes and lays it upon the
	// background Array of Bytes. It is recommended that Number of symbols in
	// String is multiple (divisible by) Four (4). Length of a String must be
	// not more than 64 allowed Symbols.

	const blockSize = 16
	const dataSize = 3 * blockSize // 3 Blocks = 48 Bytes

	const strMaxlen = 64 - 1 // 64 Symbols = 48 Bytes (4 Symbols = 3 Bytes), 1 for Length.
	const strMinlen = 8
	const allowedSymbols string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789*."
	// Base64 (0..63): "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// Allowed Characters are similar to Base64 with one Exception:
	// last Characters are " " and "." instead of "+" and "/".
	var l, m, i, j, n, data_len int
	var srcB, data []byte
	var srcS string
	var b1, b2, b3, b4, tb byte
	var err error

	b1 = []byte("*")[0]
	b2 = []byte("+")[0]
	b3 = []byte(".")[0]
	b4 = []byte("/")[0]

	l = len(*src)
	m = len(allowedSymbols)
	ok = false

	if (l > strMaxlen) || (l < strMinlen) {
		log.Println("SRC has wrong Length.") //
		return false
	}

	if len(*bg) != dataSize {
		log.Println("BG has wrong Length.") //
		return false
	}

	srcB = []byte(*src)

	if l != len(srcB) {
		log.Println("String has bad Characters.") //
		return ok
	}

	// Check allowed Symbols and convert into Base64
	for i = 0; i < l; i++ {
		ok = false
		for j = 0; j < m; j++ {
			if srcB[i] == allowedSymbols[j] {
				ok = true
			}
		}
		if !ok {
			log.Println("Bad Symbol found at pos #", i+1) //
			return ok
		}
	}

	for i = 0; i < l; i++ {
		if srcB[i] == b1 { // "*" -> "+"
			srcB[i] = b2
		}
		if srcB[i] == b3 { // "." -> "/"
			srcB[i] = b4
		}
	}

	srcS = string(srcB)

	// Push Length
	srcS = string(allowedSymbols[l]) + srcS

	// Update Length
	l = len(srcS)

	if l%4 == 1 {

		srcS = srcS + "+++"

	} else if l%4 == 2 {

		srcS = srcS + "++"

	} else if l%4 == 3 {

		srcS = srcS + "+"

	}

	// Base64 string to []byte
	data, err = base64.StdEncoding.DecodeString(srcS)
	if err != nil {
		log.Println("Conversion Base64 to []byte failed.", err) //
		return false
	}

	// Fill DST with BG
	n = copy((*block)[0:dataSize], (*bg)[0:dataSize])
	if n != len(*bg) {
		log.Println("Copy BG failed.") //
		return false
	}

	// Add converted SRC to BG
	data_len = len(data)
	n = copy((*block)[0:l], data[0:data_len])
	if n != data_len {
		log.Println("Copy Data failed.") //
		return false
	}

	i = 0
	for j = 0; j < 16; j++ {
		if j%3 == 1 {

			tb = (*block)[16*(i+2)+j]
			(*block)[16*(i+2)+j] = (*block)[16*(i+1)+j]
			(*block)[16*(i+1)+j] = (*block)[16*(i)+j]
			(*block)[16*(i)+j] = tb

		} else if j%3 == 2 {

			tb = (*block)[16*(i+2)+j]
			(*block)[16*(i+2)+j] = (*block)[16*(i)+j]
			(*block)[16*(i)+j] = (*block)[16*(i+1)+j]
			(*block)[16*(i+1)+j] = tb

		}
	}

	return ok
}

//------------------------------------------------------------------------------

func tfm1_encrypt3Blocks(src, dst *[]byte, cypher_key []byte) (ok bool) {

	const blockSize = 16
	const dataSize = 3 * blockSize // 3 Blocks = 48 Bytes
	const hbs = blockSize / 2
	const qbs = hbs / 2
	const keySize = 32

	var cypher *twofish.Cipher
	var err error
	var key, a, b, c, d, e, f []byte
	var i, r int
	var tb byte

	key = make([]byte, len(cypher_key))
	copy(key, cypher_key)

	if (len(*src) != dataSize) || (len(*dst) != dataSize) || (len(key) != keySize) {
		log.Println("Wrong Lengths Given.") //
		return false
	}

	// Prepare Data
	cypher, err = twofish.NewCipher(key)
	if err != nil {
		log.Println("Error creating Cypher.", err)
		return false
	}
	a = make([]byte, blockSize)
	b = make([]byte, blockSize)
	c = make([]byte, blockSize)
	copy(a, (*src)[0:blockSize])
	copy(b, (*src)[blockSize:blockSize*2])
	copy(c, (*src)[blockSize*2:blockSize*3])

	// Pass I
	d = make([]byte, blockSize)
	e = make([]byte, blockSize)
	f = make([]byte, blockSize)
	cypher.Encrypt(d, a)
	cypher.Encrypt(e, b)
	cypher.Encrypt(f, c)

	for r = 0; r < keySize-1; r++ {

		// Rotate Key
		tb = key[0]
		for i = 1; i < keySize; i++ {
			key[i-1] = key[i]
		}
		key[keySize-1] = tb
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
	copy((*dst)[0:blockSize], d)
	copy((*dst)[blockSize:blockSize*2], e)
	copy((*dst)[blockSize*2:blockSize*3], f)

	return true
}

//------------------------------------------------------------------------------

func tfm1_decrypt3Blocks(src, dst *[]byte, decypher_key []byte) (ok bool) {

	const blockSize = 16
	const dataSize = 3 * blockSize // 3 Blocks = 48 Bytes
	const hbs = blockSize / 2
	const qbs = hbs / 2
	const keySize = 32

	var cypher *twofish.Cipher
	var err error
	var key, a, b, c, d, e, f []byte
	var tb byte
	var i, r int

	key = make([]byte, len(decypher_key))
	copy(key, decypher_key)

	if (len(*src) != dataSize) || (len(*dst) != dataSize) || (len(key) != keySize) {
		log.Println("Wrong Lengths Given.") //
		return false
	}

	// Prepare Data
	a = make([]byte, blockSize)
	b = make([]byte, blockSize)
	c = make([]byte, blockSize)
	d = make([]byte, blockSize)
	e = make([]byte, blockSize)
	f = make([]byte, blockSize)

	copy(d, (*src)[0:blockSize])
	copy(e, (*src)[blockSize:blockSize*2])
	copy(f, (*src)[blockSize*2:blockSize*3])

	for r = 0; r < keySize-1; r++ {

		// Anti-Rotate Key
		tb = key[keySize-1]
		for i = keySize - 1; i > 0; i-- {
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
	tb = key[keySize-1]
	for i = keySize - 1; i > 0; i-- {
		key[i] = key[i-1]
	}
	key[0] = tb

	cypher, err = twofish.NewCipher(key)
	if err != nil {
		log.Println("Error creating Cypher.", err)
		return false
	}

	// Pass Anti-I
	cypher.Decrypt(a, d)
	cypher.Decrypt(b, e)
	cypher.Decrypt(c, f)

	// Output
	copy((*dst)[0:blockSize], a)
	copy((*dst)[blockSize:blockSize*2], b)
	copy((*dst)[blockSize*2:blockSize*3], c)

	return true
}

//------------------------------------------------------------------------------

func tfm1_getString(data *[]byte, str *string) {

	const blockSize = 16
	const dataSize = 3 * blockSize // 3 Blocks = 48 Bytes
	const allowedSymbols string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789*."

	var i, j, l int
	var tb byte
	var tmpStr string

	i = 0
	for j = 0; j < 16; j++ {
		if j%3 == 2 {

			tb = (*data)[16*(i+2)+j]
			(*data)[16*(i+2)+j] = (*data)[16*(i+1)+j]
			(*data)[16*(i+1)+j] = (*data)[16*(i)+j]
			(*data)[16*(i)+j] = tb

		} else if j%3 == 1 {

			tb = (*data)[16*(i+2)+j]
			(*data)[16*(i+2)+j] = (*data)[16*(i)+j]
			(*data)[16*(i)+j] = (*data)[16*(i+1)+j]
			(*data)[16*(i+1)+j] = tb

		}
	}

	// []byte -> Base64 string
	tmpStr = base64.StdEncoding.EncodeToString((*data)[0:dataSize])
	l = strings.Index(allowedSymbols, string(tmpStr[0]))
	*str = tmpStr[1 : l+1]
}
