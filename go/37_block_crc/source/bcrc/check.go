// check.go.

// Block CRC Check.

// Date: 2018-09-22.

// Notes.
//
// 1. Block's Double CheckSum Error may mean one of the following:
//	-	Both Block CheckSums are good, but Data File Block is corrupted;
//	-	Data File Block is good, but both Block CheckSums are corrupted;
//	-	Everything is corrupted: Data File Block and Block's both CheckSums.
//	If further Blocks contain Sliding Block CheckSum Error, then this often
// 	means the Data File Block Corruption. If there are no further Sliding Block
// 	CheckSum Errors, then it often means that Data File Block is good and the
// 	Error may be in the CheckSum.
//
// 2. Single Current Block's CheckSum Error, if it is not followed by Errors
// 	in further Blocks, often means corrupted Current Block's CheckSum.
//
// 3. Single Sliding Block's CheckSum Error, if it is not followed by Errors
// 	in further Blocks, often means corrupted Sliding Block's CheckSum.

package bcrc

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

const MsgBlockCRCMismatchDouble = "Double CRC Mismatch in Block #%v.\r\n"
const MsgBlockCRCMismatchCurrentBlock = "CRC Mismatch in Current Block #%v.\r\n"
const MsgBlockCRCMismatchSlidingBlock = "CRC Mismatch in Sliding Block #%v.\r\n"

// Creates Block CRC.
// 'dfp' = Data File Path.
// 'sfp' = Sum File Path.
func BCRCCheck(dfp, sfp string) error {

	var blockNumber uint64
	var blockSize uint64
	var blockSizeBA []byte
	var bytesRead int
	var bytesReadSumFileTotal uint64
	var bytesReadDataFileTotal uint64
	var checkSumMismatchCurrentBlock bool
	var checkSumMismatchSlidingBlock bool
	var currentBlock []byte
	var crcTable *crc32.Table
	var dataFile *os.File
	var err error
	var expectedSumOfCurrentBlock uint32
	var expectedSumOfCurrentBlockBA []byte
	var expectedSumOfSlidingBlock uint32
	var expectedSumOfSlidingBlockBA []byte
	var i uint64
	var mismatchIsFound bool
	var readSumOfCurrentBlock uint32
	var readSumOfCurrentBlockBA []byte
	var readSumOfSlidingBlock uint32
	var readSumOfSlidingBlockBA []byte
	var slidingTotalBlock []byte // "Running Total" Block.
	var sumFile *os.File

	// Open the Data File.
	dataFile, err = os.Open(dfp)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	// Open the Sum File.
	sumFile, err = os.Open(sfp)
	if err != nil {
		return err
	}
	defer sumFile.Close()

	// Read Block Size from the Sum File.
	blockSizeBA = make([]byte, 8)
	bytesRead, err = sumFile.Read(blockSizeBA)
	if err != nil {
		return err
	}
	bytesReadSumFileTotal = bytesReadSumFileTotal + uint64(bytesRead)
	blockSize = binary.BigEndian.Uint64(blockSizeBA)

	// Preparations.

	// 1. Buffer for current Block.
	currentBlock = make([]byte, blockSize)
	slidingTotalBlock = make([]byte, 0)
	blockNumber = 0
	mismatchIsFound = false

	// 2. CRC Table and CRC Sum Holders.
	crcTable = crc32.IEEETable
	readSumOfCurrentBlockBA = make([]byte, 4)
	readSumOfSlidingBlockBA = make([]byte, 4)
	expectedSumOfCurrentBlockBA = make([]byte, 4)
	expectedSumOfSlidingBlockBA = make([]byte, 4)

	// Read all CheckSums from Sum File.
	for true {

		// Read current Block CheckSum.
		blockNumber++
		bytesRead, err = sumFile.Read(readSumOfCurrentBlockBA)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		bytesReadSumFileTotal = bytesReadSumFileTotal + uint64(bytesRead)
		readSumOfCurrentBlock = binary.BigEndian.Uint32(readSumOfCurrentBlockBA)

		// Read sliding Block CheckSum.
		bytesRead, err = sumFile.Read(readSumOfSlidingBlockBA)
		if err != nil {
			return err
		}
		bytesReadSumFileTotal = bytesReadSumFileTotal + uint64(bytesRead)
		readSumOfSlidingBlock = binary.BigEndian.Uint32(readSumOfSlidingBlockBA)

		// Read current Block from Data File.
		bytesRead, err = dataFile.Read(currentBlock)
		bytesReadDataFileTotal = bytesReadDataFileTotal + uint64(bytesRead)
		if err != nil {
			return err
		}

		// Update sliding Total Block.
		slidingTotalBlock = append(slidingTotalBlock, currentBlock...)

		// No Errors. Process a Block.
		expectedSumOfCurrentBlock = crc32.Checksum(currentBlock, crcTable)
		binary.BigEndian.PutUint32(
			expectedSumOfCurrentBlockBA,
			expectedSumOfCurrentBlock,
		)
		expectedSumOfSlidingBlock = crc32.Checksum(slidingTotalBlock, crcTable)
		binary.BigEndian.PutUint32(
			expectedSumOfSlidingBlockBA,
			expectedSumOfSlidingBlock,
		)

		// Compare Checksums.
		checkSumMismatchCurrentBlock =
			readSumOfCurrentBlock != expectedSumOfCurrentBlock
		checkSumMismatchSlidingBlock =
			readSumOfSlidingBlock != expectedSumOfSlidingBlock

		if checkSumMismatchCurrentBlock || checkSumMismatchSlidingBlock {

			// At least One Mismatch has been found.
			mismatchIsFound = true
			if checkSumMismatchCurrentBlock == checkSumMismatchSlidingBlock {

				// Report.
				fmt.Printf(MsgBlockCRCMismatchDouble, blockNumber)

			} else if checkSumMismatchCurrentBlock {

				// Report.
				fmt.Printf(MsgBlockCRCMismatchCurrentBlock, blockNumber)

			} else {

				// Report.
				fmt.Printf(MsgBlockCRCMismatchSlidingBlock, blockNumber)
			}
		}

		// Clear the Buffer before new Data arrives.
		for i = 0; i < blockSize; i++ {
			currentBlock[i] = 0
		}
	}

	// Summary.
	if mismatchIsFound {
		fmt.Println("Mismatch(es) has been found!")
	} else {
		fmt.Println("Good. CheckSums match the Data.")
	}

	return nil
}
