package bloom_filter

import "github.com/go-redis/redis"

type BloomFilter struct {
	cli       *redis.Client
	hashFuncs []hashFunc
}

func NewBloomFilter(cli *redis.Client) BloomFilter {
	return BloomFilter{
		cli:       cli,
		hashFuncs: []hashFunc{BKDRHash, SDBMHash, DJBHash},
	}
}

type hashFunc func(str string) int64

func (b *BloomFilter) Add(key, value string) error {
	for _, f := range b.hashFuncs {
		offset := f(value)
		intCmd := b.cli.SetBit(key, offset, 1)
		if intCmd.Err() != nil {
			return intCmd.Err()
		}
	}
	return nil
}

func (b *BloomFilter) Contains(key, value string) bool {
	for _, f := range b.hashFuncs {
		offset := f(value)
		intCmd := b.cli.GetBit(key, offset)
		if intCmd.Val() != 1 {
			return false
		}
	}
	return true
}

func BKDRHash(str string) int64 {
	seed := int64(131) // 31 131 1313 13131 131313 etc..
	hash := int64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int64(str[i])
	}
	return hash & 0x7FFFFFFF
}

func SDBMHash(str string) int64 {
	hash := int64(0)
	for i := 0; i < len(str); i++ {
		hash = int64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash & 0x7FFFFFFF
}

func DJBHash(str string) int64 {
	hash := int64(0)
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + int64(str[i])
	}
	return hash & 0x7FFFFFFF
}
