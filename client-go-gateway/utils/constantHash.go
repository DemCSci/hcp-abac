package utils

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

//声明切片类型
type units []uint32

//返回切片长度
func (x units) Len() int {
	return len(x)
}

//比较两个值的大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

//切片中值交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

//当hash环没有数据时，提示错误
var emptyErr = errors.New("hash 环为空")

//创建结构体，保存一致性hash信息
type Consistent struct {
	//hash环，key为hash值，值存放的是节点信息
	circle map[uint32]string
	//已经排序的节点hash切片
	sortedHashes units
	//虚拟节点个数，用来增加hash的平衡性
	virtualNode int
	//读写锁
	sync.RWMutex
}

func NewConsistent() *Consistent {
	return &Consistent{
		//初始化变量
		circle: make(map[uint32]string),
		//设置虚拟节点个数
		virtualNode: 20,
	}
}

//自动生成key值
func (c *Consistent) generateKey(element string, index int) string {
	//副本key生成逻辑
	return element + strconv.Itoa(index)
}

//获取hash位置
func (c *Consistent) hashKey(key string) uint32 {
	if len(key) < 64 {
		var srcatch [64]byte
		copy(srcatch[:], key)

		//使用IEEE多项式返回数据的CRC-32校验和
		return crc32.ChecksumIEEE(srcatch[:len(key)])
	}

	return crc32.ChecksumIEEE([]byte(key))
}

//更新排序，方便查找
func (c *Consistent) updateSortedHashes() {
	hash := c.sortedHashes[:0]

	//判断切片容量是否过大，如果过大则重置
	if cap(c.sortedHashes)/(c.virtualNode*4) > len(c.circle) {
		hash = nil
	}

	//添加hash
	for k := range c.circle {
		hash = append(hash, k)
	}

	//排序
	sort.Sort(hash)

	c.sortedHashes = hash
}

//向hash环添加节点
func (c *Consistent) Add(element string) {
	c.Lock()
	defer c.Unlock()
	c.add(element)
}

func (c *Consistent) add(element string) {
	//生成虚拟节点
	var (
		i    int
		hash uint32
		key  string
	)

	//循环虚拟节点，设置副本
	for i = 0; i < c.virtualNode; i++ {
		key = c.generateKey(element, i)
		hash = c.hashKey(key)
		c.circle[hash] = element
	}

	//更新排序
	c.updateSortedHashes()
}

func (c *Consistent) remove(element string) {
	//生成虚拟节点
	var (
		i    int
		hash uint32
		key  string
	)

	//循环虚拟节点，设置副本
	for i = 0; i < c.virtualNode; i++ {
		key = c.generateKey(element, i)
		hash = c.hashKey(key)
		delete(c.circle, hash)
	}

	//更新排序
	c.updateSortedHashes()
}

//删除节点
func (c *Consistent) Remove(element string) {
	c.Lock()
	defer c.Unlock()
	c.remove(element)
}

//顺时针查找最近的节点
func (c *Consistent) search(key uint32) int {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}

	i := sort.Search(len(c.sortedHashes), f)

	if i >= len(c.sortedHashes) {
		i = 0
	}

	return i
}

//根据数据表示，获取对应的服务器节点信息
func (c *Consistent) Ger(name string) (string, error) {
	c.Lock()
	defer c.Unlock()

	if len(c.circle) == 0 {
		return "", emptyErr
	}

	key := c.hashKey(name)
	i := c.search(key)

	return c.circle[c.sortedHashes[i]], nil
}
