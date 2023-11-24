package main

import (
	"fmt"
)

type Marble struct {
	value int
	prev  *Marble
	next  *Marble
}

func insert(head *Marble, marble int) *Marble {
	new_marble := &Marble{value: marble, prev: nil, next: nil}
	current_marble := head.next
	next_marble := current_marble.next
	current_marble.next = new_marble
	new_marble.prev = current_marble
	new_marble.next = next_marble
	next_marble.prev = new_marble
	return new_marble
}

func print(head *Marble) {
	current_head := head
	for {
		fmt.Println(head.value)
		head = head.next
		if head == current_head {
			break
		}
	}
}

func remove_7th(head *Marble) (*Marble, int) {
	for i := 0; i < 7; i++ {
		head = head.prev
	}
	value := head.value
	head = head.prev
	// fmt.Println("removing:", value)
	// fmt.Println("current node:", head.value)
	head.next = head.next.next
	head.prev = head
	return head.next, value
}

func max_score(players []int) int {
	max := 0
	for _, score := range players {
		if score > max {
			max = score
		}
	}
	return max
}

func solution(nplayers int, last_marble int) int {
	players := make([]int, nplayers)
	_ = players
	current_marble := 0
	current_player := 0

	head := &Marble{value: current_marble, prev: nil, next: nil}
	head.next = head
	head.prev = head

	for {
		current_marble += 1
		current_player = (current_player + 1) % nplayers
		if current_marble%23 == 0 {
			new_head, value := remove_7th(head)
			head = new_head
			players[current_player] += current_marble + value
		} else {
			head = insert(head, current_marble)
		}

		if current_marble == last_marble {
			break
		}
	}
	// print(head)
	return max_score(players)
}

func main() {
	fmt.Println("Part 1:", solution(446, 71522))   // 390592
	fmt.Println("Part 2:", solution(446, 7152200)) // 3277920293
}
