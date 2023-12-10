package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value         int
	rightParent   *Node
	leftNeighbour *Node
	leftChild     *Node
	rightChild    *Node
}

func parseNode(input string, leftNode *Node) Node {
	item, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("failed to parse item: %v", input)
	}

	return Node{
		value:         item,
		leftNeighbour: leftNode,
	}
}

func (n *Node) addLeftChild(leftNode *Node) {
	if leftNode == nil {
		return
	}
	newVal := n.value - leftNode.value

	n.leftChild = &Node{
		value:         newVal,
		rightParent:   n,
		leftNeighbour: leftNode.leftChild,
	}

	n.leftNeighbour.rightChild = n.leftChild
}
func (n *Node) getPrediction() (prediction int) {
	for n.leftChild != nil {
		n = n.leftChild
	}

	prediction = n.value
	for n.rightParent != nil {
		n = n.rightParent
		prediction += n.value
	}

	return
}

func (n *Node) predictPast() (prediction int) {
	for n.rightChild != nil {
		n = n.rightChild
	}

	for {
		prediction = n.value - prediction

		if n.rightParent == nil {
			break
		}

		n = n.rightParent.leftNeighbour
	}

	return
}

func addAllChildren(input []*Node) {
	allZero := true
	for _, node := range input {
		if node.value != 0 {
			allZero = false
		}
	}

	if allZero {
		return
	}

	var children []*Node
	for index, node := range input {
		if index == 0 {
			continue
		}

		node.addLeftChild(input[index-1])
		children = append(children, node.leftChild)
	}

	addAllChildren(children)
}

func printTree(input []*Node) {
	fmt.Println()
	var children []*Node
	for _, node := range input {
		fmt.Printf(" %v", node.value)
		if node.leftChild != nil {
			children = append(children, node.leftChild)
		}
	}

	if len(children) == 0 {
		return
	}

	printTree(children)
}

func TaskOne() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day9/input")
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err.Error())
		}
	}()

	if err != nil {
		log.Fatalf("failed to open file input1\nwith err: %v", err.Error())
	}

	fileScanner := bufio.NewScanner(inputFile)
	var lines [][]*Node

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		var prevNode *Node
		var nodes []*Node

		for _, tmp := range line {
			currNode := parseNode(tmp, prevNode)
			currNode.addLeftChild(prevNode)
			prevNode = &currNode

			nodes = append(nodes, &currNode)
		}

		lines = append(lines, nodes)
		addAllChildren(nodes)
	}

	sum := 0
	for _, line := range lines {
		prediction := line[len(line)-1].getPrediction()
		sum += prediction
	}

	fmt.Println("\n######################")
	fmt.Printf("\nResult: %v", sum)
}

func TaskTwo() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day9/input")
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err.Error())
		}
	}()

	if err != nil {
		log.Fatalf("failed to open file input1\nwith err: %v", err.Error())
	}

	fileScanner := bufio.NewScanner(inputFile)
	var lines [][]*Node

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		var prevNode *Node
		var nodes []*Node

		for _, tmp := range line {
			currNode := parseNode(tmp, prevNode)
			currNode.addLeftChild(prevNode)
			prevNode = &currNode

			nodes = append(nodes, &currNode)
		}

		lines = append(lines, nodes)
		addAllChildren(nodes)
	}

	past := 0
	for _, line := range lines {
		prediction := line[0].predictPast()
		past += prediction
	}

	fmt.Println("\n######################")
	fmt.Printf("\nResult: %v", past)
}
