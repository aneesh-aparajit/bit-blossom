package main

type BloomFilter struct {
	filter    []uint8
	functions *HashingFunctions
	size      int
}

func NewBloomFilter(size, numHashFunctions int) *BloomFilter {
	/*
		Instead of having an array of size `size`, we split the array into uint8 elements grouped by 8 bits.
		So, let's say we want to set the index of 12, then in that case, we'll first go to the element 12/8, and then
		go to the 12%8'th bit in that index and set that bit.
	*/
	return &BloomFilter{
		filter:    make([]uint8, size),
		size:      size,
		functions: NewHashingFunctions(numHashFunctions),
	}
}

func (b *BloomFilter) setBitForHashValue(hashValue int) {
	idx := hashValue % b.size
	b.filter[idx/8] |= (1 << (idx % 8))
}

func (b *BloomFilter) getBitForHashValue(hashValue int) bool {
	idx := hashValue % b.size
	return b.filter[idx/8]>>(idx%8)&1 == 1
}

func (b *BloomFilter) Add(key string) {
	hashes := b.functions.GetHashValuesForKey([]byte(key))
	for i := 0; i < b.functions.Size(); i++ {
		b.setBitForHashValue(hashes[i])
	}
}

func (b *BloomFilter) Exists(key string) bool {
	hashes := b.functions.GetHashValuesForKey([]byte(key))
	exists := true
	for _, value := range hashes {
		exists = exists && b.getBitForHashValue(value)
	}
	return exists
}
