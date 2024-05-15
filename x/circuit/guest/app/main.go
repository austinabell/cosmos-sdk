package main

/*
#include <stdint.h>
#cgo LDFLAGS: -L./out/platform/riscv32im-risc0-zkvm-elf/release -lzkvm_platform
void sys_halt(uint8_t exit_code, uint32_t* initial_sha_state);
void sys_write(uint32_t fd, uint8_t* byte_ptr, int len);
void sys_sha_buffer(uint32_t* out_state, uint32_t* in_state, uint8_t* buf, uint32_t count);
*/
import "C"

import (
	"context"
	"encoding/binary"
	"time"
	"unsafe"

	appmanager "cosmossdk.io/core/app"
	"cosmossdk.io/x/circuit/app/deps"
)

var app = deps.NewStf()

func main() {
	_, _, _ = app.DeliverBlock(context.Background(), &appmanager.BlockRequest[deps.Tx]{
		Height:            0,
		Time:              time.Time{},
		Hash:              nil,
		ChainId:           "",
		AppHash:           nil,
		Txs:               nil,
		ConsensusMessages: nil,
	},
		nil, // TODO: state
	)

	var assumptionsDigestState = [8]uint32{0, 0, 0, 0, 0, 0, 0, 0}

	// TODO replace this output with program specific output
	outputBytes := [4]byte{0,1,2,3}

	// No output, but finalize still calls sha on empty bytes it seems
	journalDigest := shaBuffer(sha256InitState(), outputBytes[:])

	C.sys_write(3 /* journal */, &outputBytes[0], C.int(len(outputBytes)))

	output := taggedStruct("risc0.Output", [][8]uint32{journalDigest, assumptionsDigestState})

	C.sys_halt(0, &output[0])
}

func shaBuffer(initialState [8]uint32, bytes []byte) [8]uint32 {
	var outState [8]uint32
	// TODO doesn't really handle alignment
	padLen := computeU32sNeeded(len(bytes))

	// Create buffer of u32s to ensure word aligned
	padBuf := make([]uint32, padLen)
	if !(len(bytes) <= len(padBuf)*4) {
		panic("bytes too long")
	}
	padBufu8 := unsafe.Slice((*byte)(unsafe.Pointer(&padBuf[0])), 4*len(padBuf))

	// Copy whole bytes
	length := int(min(len(bytes), padLen*4))
	copy(padBufu8[:length], bytes[:length])

	// Add END marker since this is always with a trailer
	padBufu8[length] = 0x80

	// Add trailer with number of bits written. This needs to be big endian.
	bitsTrailer := 8 * uint32(len(bytes))

	// Swap bits to BE
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, bitsTrailer)
	bitsTrailer = uint32(binary.LittleEndian.Uint32(b))

	padBuf[padLen-1] = bitsTrailer

	numBlocks := padLen / 16
	C.sys_sha_buffer(&outState[0], &initialState[0], &padBufu8[0], uint32(numBlocks))
	return outState
}

func taggedStruct(tag string, digests [][8]uint32) [8]uint32 {
	tagDigest := shaBuffer(sha256InitState(), []byte(tag))

	// Create buffer for all bytes to hash
	allBytes := make([]byte, 0, len(tagDigest)*4+len(digests)*32+2)

	for _, word := range tagDigest {
		allBytes = binary.LittleEndian.AppendUint32(allBytes, word)
	}

	for _, digest := range digests {
		for _, word := range digest {
			allBytes = binary.LittleEndian.AppendUint32(allBytes, word)
		}
	}

	// TODO doesn't handle extending with data, not needed for this use case yet

	allBytes = binary.LittleEndian.AppendUint16(allBytes, uint16(len(digests)))

	return shaBuffer(sha256InitState(), allBytes)
}

func sha256InitState() [8]uint32 {
	// BE conversion of the sha256 init state
	return [8]uint32{
		1743128938,
		2242799547,
		1928556092,
		989155237,
		2136084049,
		2355627419,
		2883158815,
		432922715,
	}
}



func computeU32sNeeded(lenBytes int) int {
	const WordSize = 4
	const BlockWords = 16
	// Add one byte for end marker
	nWords := alignUp(lenBytes+1, WordSize) / WordSize
	// Add two words for length at end (even though we only
	// use one of them, being a 32-bit architecture)
	nWords += 2

	return alignUp(nWords, BlockWords)
}

func alignUp(addr, al int) int {
	return (addr + al - 1) & ^(al - 1)
}
