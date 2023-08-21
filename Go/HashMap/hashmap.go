package HashMap

// Basic Node for Hashmap
type Node struct {
	Key   string
	Value string
}

// Mutable Hashmap
type HashMap struct {
	Size   uint64
	Bucket []Node
}
