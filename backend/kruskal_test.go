package main

import "testing"

func isSameSpanningTree(ST1 []MinimumSpanningTree, ST2 []MinimumSpanningTree) bool {
	if (len(ST1) != len(ST2)) {
		return false
	} else {
		var i int
		for i = 0; i < len(ST1); i++ {
			if (!isSame(ST1[i].p1, ST2[i].p1) || !isSame(ST1[i].p2, ST2[i].p2)) {
				return false
			}
		}
		return true
	}
}

func TestKruskal(t *testing.T) {
	var i int
	list_point := []Point{}
	for i = 0; i < 3; i++ {
		var point Point

		point.x = i + 1
		point.y = i + 2
		list_point = append(list_point, point)
	}
	var point Point
	point.x = 2
	point.y = 4
	list_point = append(list_point, point)

	var result []MinimumSpanningTree
	var MST MinimumSpanningTree
	for i = 0; i < 3; i++ {
		var p1 Point
		var p2 Point
		if (i == 0) {
			p1.x = 2
			p1.y = 3
			p2.x = 2
			p2.y = 4
			MST.p1 = p1
			MST.p2 = p2
		} else if (i == 1) {
			p1.x = 3
			p1.y = 4
			p2.x = 2
			p2.y = 4
			MST.p1 = p1
			MST.p2 = p2
		} else {
			p1.x = 1
			p1.y = 2
			p2.x = 2
			p2.y = 3
			MST.p1 = p1
			MST.p2 = p2
		}
		result = append(result, MST)
	}

	if (!isSameSpanningTree(KruskalAlgorithm(list_point), result)) {
		t.Fatalf("Ada Algoritma Kruskal yang salah!")
	}
}