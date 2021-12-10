package hash

import (
	"crypto/sha1"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	//虚拟节点数量
	DefaultVirtualSpots = 200
)

//节点信息
type node struct {
	nodeKey   string //节点的key
	spotValue uint32 //节点的槽位值
}

type nodesArray []node

func (p nodesArray) Len() int           { return len(p) }
func (p nodesArray) Less(i, j int) bool { return p[i].spotValue < p[j].spotValue }
func (p nodesArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p nodesArray) Sort()              { sort.Sort(p) }

// 带权重hash环
type HashRing struct {
	virualSpots int            //虚拟节点
	nodes       nodesArray     //节点数组(槽位)
	weights     map[string]int //节点信息对应其权重
	mu          sync.RWMutex
}

// 初始化Hash环, vSpots为虚拟节点数量; vSpots小于0, 则使用默认虚拟节点数200
func NewHashRing(vSpots int) *HashRing {
	spots := DefaultVirtualSpots
	if vSpots > 0 {
		spots = vSpots
	}

	h := &HashRing{
		virualSpots: spots,
		weights:     make(map[string]int),
	}
	return h
}

// 用于向hash环中添加节点信息nodes
func (h *HashRing) AddNodes(nodeWeight map[string]int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for nodeKey, w := range nodeWeight {
		h.weights[nodeKey] = w
	}
	h.generate()
}

// 向hash环中添加一个节点信息
func (h *HashRing) AddNode(nodeKey string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.weights[nodeKey] = weight
	h.generate()
}

// 根据节点的key(nodeKey)从hash环上删除
func (h *HashRing) RemoveNode(nodeKey string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.weights, nodeKey)
	h.generate()
}

// 更新hash环上某个节点的信息
func (h *HashRing) UpdateNode(nodeKey string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.weights[nodeKey] = weight
	h.generate()
}

// 重新计算并生成hash环中的nodes
func (h *HashRing) generate() {
	var totalW int
	for _, w := range h.weights {
		totalW += w
	}

	totalVirtualSpots := h.virualSpots * len(h.weights) //总的虚拟节点(槽位)
	h.nodes = nodesArray{}

	for nodeKey, w := range h.weights {
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirtualSpots))) //计算出该节点的槽位数
		for i := 1; i <= spots; i++ {                                                       //计算出该节点的hash值并添加到nodes中
			hash := sha1.New()
			hash.Write([]byte(nodeKey + ":" + strconv.Itoa(i)))
			hashBytes := hash.Sum(nil)
			n := node{
				nodeKey:   nodeKey,
				spotValue: genValue(hashBytes[6:10]),
			}
			h.nodes = append(h.nodes, n)
			hash.Reset()
		}
	}
	h.nodes.Sort() //对nodes进行排序
}

// 将bs转成uint32
func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}

// 根据key获取对应的节点名称
func (h *HashRing) GetNode(s string) string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.nodes) == 0 {
		return ""
	}

	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	//获取nodes中第一个槽位(节点)的值大于该key对应的hash值的槽位(节点)索引
	i := sort.Search(len(h.nodes), func(i int) bool { return h.nodes[i].spotValue >= genValue(hashBytes) })

	if i == len(h.nodes) {
		i = 0
	}
	return h.nodes[i].nodeKey
}
