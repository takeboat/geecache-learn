package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int            // 虚拟节点倍数
	keys     []int          // sorted 哈希环
	hashMap  map[int]string // 虚拟节点和真实节点的映射表
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash) // 将哈希值添加到环上
			m.hashMap[hash] = key         // 建立虚拟节点和真实节点的映射
		}
	}
	sort.Ints(m.keys) // 对哈希值进行排序
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	// 计算key的哈希值
	hash := int(m.hash([]byte(key)))
	// 在哈希环中找到顺时针方向上比它大的第一个节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// 这里的sort函数如果没有找到 会返回 len(m.keys)
	// 所以下面获取对应的真实节点需要mod len(m.keys)
	// 这属于是环形取值
	// 通过hashMap获取对应的真实节点
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
