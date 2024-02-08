// Code generated by fastssz. DO NOT EDIT.
// Hash: 953d511e4b9ab9601054ad0869d034020890d4dc921edfd28c9e2ff8785fe6ee
package capella

import (
	ssz "github.com/prysmaticlabs/fastssz"
	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	v1 "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"
)

// MarshalSSZ ssz marshals the BeaconKitBlockCapella object
func (b *BeaconKitBlockCapella) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconKitBlockCapella object to a target array
func (b *BeaconKitBlockCapella) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(44)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Offset (1) 'Body'
	dst = ssz.WriteOffset(dst, offset)
	if b.Body == nil {
		b.Body = new(BeaconKitBlockBodyCapella)
	}
	offset += b.Body.SizeSSZ()

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	dst = append(dst, b.PayloadValue...)

	// Field (1) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconKitBlockCapella object
func (b *BeaconKitBlockCapella) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 44 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Slot'
	b.Slot = github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Offset (1) 'Body'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 44 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'PayloadValue'
	if cap(b.PayloadValue) == 0 {
		b.PayloadValue = make([]byte, 0, len(buf[12:44]))
	}
	b.PayloadValue = append(b.PayloadValue, buf[12:44]...)

	// Field (1) 'Body'
	{
		buf = tail[o1:]
		if b.Body == nil {
			b.Body = new(BeaconKitBlockBodyCapella)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconKitBlockCapella object
func (b *BeaconKitBlockCapella) SizeSSZ() (size int) {
	size = 44

	// Field (1) 'Body'
	if b.Body == nil {
		b.Body = new(BeaconKitBlockBodyCapella)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconKitBlockCapella object
func (b *BeaconKitBlockCapella) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconKitBlockCapella object with a hasher
func (b *BeaconKitBlockCapella) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (1) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	hh.PutBytes(b.PayloadValue)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BeaconKitBlockBodyCapella object
func (b *BeaconKitBlockBodyCapella) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconKitBlockBodyCapella object to a target array
func (b *BeaconKitBlockBodyCapella) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(132)

	// Field (0) 'RandaoReveal'
	if size := len(b.RandaoReveal); size != 96 {
		err = ssz.ErrBytesLengthFn("--.RandaoReveal", size, 96)
		return
	}
	dst = append(dst, b.RandaoReveal...)

	// Field (1) 'Graffiti'
	if size := len(b.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Graffiti", size, 32)
		return
	}
	dst = append(dst, b.Graffiti...)

	// Offset (2) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(v1.ExecutionPayloadCapella)
	}
	offset += b.ExecutionPayload.SizeSSZ()

	// Field (2) 'ExecutionPayload'
	if dst, err = b.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconKitBlockBodyCapella object
func (b *BeaconKitBlockBodyCapella) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 132 {
		return ssz.ErrSize
	}

	tail := buf
	var o2 uint64

	// Field (0) 'RandaoReveal'
	if cap(b.RandaoReveal) == 0 {
		b.RandaoReveal = make([]byte, 0, len(buf[0:96]))
	}
	b.RandaoReveal = append(b.RandaoReveal, buf[0:96]...)

	// Field (1) 'Graffiti'
	if cap(b.Graffiti) == 0 {
		b.Graffiti = make([]byte, 0, len(buf[96:128]))
	}
	b.Graffiti = append(b.Graffiti, buf[96:128]...)

	// Offset (2) 'ExecutionPayload'
	if o2 = ssz.ReadOffset(buf[128:132]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 132 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'ExecutionPayload'
	{
		buf = tail[o2:]
		if b.ExecutionPayload == nil {
			b.ExecutionPayload = new(v1.ExecutionPayloadCapella)
		}
		if err = b.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconKitBlockBodyCapella object
func (b *BeaconKitBlockBodyCapella) SizeSSZ() (size int) {
	size = 132

	// Field (2) 'ExecutionPayload'
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(v1.ExecutionPayloadCapella)
	}
	size += b.ExecutionPayload.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconKitBlockBodyCapella object
func (b *BeaconKitBlockBodyCapella) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconKitBlockBodyCapella object with a hasher
func (b *BeaconKitBlockBodyCapella) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	if size := len(b.RandaoReveal); size != 96 {
		err = ssz.ErrBytesLengthFn("--.RandaoReveal", size, 96)
		return
	}
	hh.PutBytes(b.RandaoReveal)

	// Field (1) 'Graffiti'
	if size := len(b.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Graffiti", size, 32)
		return
	}
	hh.PutBytes(b.Graffiti)

	// Field (2) 'ExecutionPayload'
	if err = b.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
