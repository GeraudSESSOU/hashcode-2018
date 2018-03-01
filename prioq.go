package main

// Invariant: both children are bigger

type prioq struct {
	bintree []*Car
}

func (pq *prioq) Add(car *Car) {
	pq.bintree = append(pq.bintree, car)

	// Rebalance tree to respect invariant
	var i = len(pq.bintree) - 1
	var p = (i - 1) / 2
	for p >= 0 && pq.bintree[p].Arrival > pq.bintree[i].Arrival {
		pq.bintree[p], pq.bintree[i] = pq.bintree[i], pq.bintree[p]
		i = p
		p = (i - 1) / 2
	}
}

func (pq *prioq) Pop() *Car {
	if len(pq.bintree) == 0 {
		panic("Trying to remove from empty queue")
	}

	if len(pq.bintree) == 1 {
		elem := pq.bintree[0]
		pq.bintree = pq.bintree[:0]
		return elem
	}

	elem := pq.bintree[0]
	// Put last element at root
	pq.bintree[0] = pq.bintree[len(pq.bintree)-1]
	// Remove last element
	pq.bintree = pq.bintree[:len(pq.bintree)-1]

	//        1                  9
	//    10     9	         10     12
	//  11 12   13 14  ->  11 12   13 14
	// 12

	// Rebalance tree to respect invariant
	len := len(pq.bintree)
	i, left, right := 0, 0, 0
	for {
		left = 2*i + 1
		right = 2*i + 2
		if left < len && right < len { // Two children
			if pq.bintree[left].Arrival <= pq.bintree[right].Arrival {
				if pq.bintree[i].Arrival <= pq.bintree[left].Arrival {
					break // Inferior to both children
				} else {
					pq.bintree[i], pq.bintree[left] = pq.bintree[left], pq.bintree[i]
					i = left
				}
			} else {
				if pq.bintree[i].Arrival <= pq.bintree[right].Arrival {
					break // Inferior to both children
				} else {
					pq.bintree[i], pq.bintree[right] = pq.bintree[right], pq.bintree[i]
					i = right
				}
			}
		} else if left < len { // One child (left)
			if pq.bintree[i].Arrival <= pq.bintree[left].Arrival {
				break // Inferior to only child
			}
			pq.bintree[i], pq.bintree[left] = pq.bintree[left], pq.bintree[i]
			i = left
		} else { // No child
			break
		}

	}

	return elem
}

func (pq *prioq) empty() bool {
	return len(pq.bintree) == 0
}
