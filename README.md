# gotrees
`gotrees` is a simple implementation of an [AVL Tree](https://en.wikipedia.org/wiki/AVL_tree), written in Go.  This implementation provides basic primitives for building and operating on tree sets and tree maps, using [Go Generics](https://go.dev/blog/intro-generics).

## `TreeSet[K Comparable]` Usage
To create a new tree set containing `int`'s:
```
treeSet := gotrees.NewTreeSet[int]()
treeSet.AddItem(1)
treeSet.AddItem(4)
...
```

To test value membership of this tree set:
```
var hasValue bool
hasValue = treeSet.Contains(1) // true
hasValue = treeSet.Contains(2) // false
```

## `TreeMap[K Comparable, V any]` Usage
To create a new tree map of `int`, `string` key-values:
```
treeMap := gotrees.NewTreeMap[int, string]()
treeMap.AddItem(1, "A")
treeMap.AddItem(4, "D")
...
```

To retrieve a value for a certain key:
```
var value *KeyValueNode[K, V]
value = treeMap.Find(1) // "A"
value = treeMap.Find(0) // nil
```
