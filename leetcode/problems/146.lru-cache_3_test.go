package problems

import "testing"

func Test_LRUCache2(t *testing.T) {

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if r1 := cache.Get(1); r1 != 1 { // returns 1
		t.Fatalf("should be 1")
	}

	cache.Put(3, 3)                   // evicts key 2
	if r2 := cache.Get(2); r2 != -1 { // returns -1 (not found)
		t.Fatalf("should be -1")
	}

	cache.Put(4, 4)                   // evicts key 1
	if r1 := cache.Get(1); r1 != -1 { // returns -1 (not found)
		t.Fatalf("should be -1")
	}
	if r3 := cache.Get(3); r3 != 3 { // returns 3
		t.Fatalf("should be 3")
	}

	if r4 := cache.Get(4); r4 != 4 { // returns 4
		t.Fatalf("should be 4")
	}
}

func Test_LRUCache3(t *testing.T) {

	cache := Constructor(3)
	cache.Put(1, 1)                  //1
	cache.Put(2, 2)                  //2 1
	cache.Put(3, 3)                  //3 2 1
	cache.Put(4, 4)                  //4 3 2
	if r4 := cache.Get(4); r4 != 4 { // 4 3 2
		t.Fatalf("should be 1")
	}
	if r3 := cache.Get(3); r3 != 3 { // 3 4 2
		t.Fatalf("should be 3")
	}
	if r2 := cache.Get(2); r2 != 2 { // 2 3 4
		t.Fatalf("should be 2")
	}
	if r1 := cache.Get(1); r1 != -1 { // 2 3 4
		t.Fatalf("should be -1")
	}

	cache.Put(5, 5)                   // 5 2 3
	if r1 := cache.Get(1); r1 != -1 { // 5 3 2
		t.Fatalf("should be -1")
	}
	if r2 := cache.Get(2); r2 != 2 { // 2 5 3
		t.Fatalf("should be 2")
	}
	if r3 := cache.Get(3); r3 != 3 { // 3 2 5
		t.Fatalf("should be 3")
	}
	if r4 := cache.Get(4); r4 != -1 { // 3 2 5
		t.Fatalf("should be -1")
	}
	if r5 := cache.Get(5); r5 != 5 { // 5 3 2
		t.Fatalf("should be 5")
	}

	//["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	//[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
}
