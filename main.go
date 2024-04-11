package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type laptopData struct {
	id        int
	buyPrice  int
	sellPrice int
}
type node struct {
	up        *node
	left      *node
	right     *node
	laptop    laptopData
	profit    int
	maxProfit int
}

func (n *node) buildTree(laptops []laptopData) *node {
	if len(laptops) == 0 {
		return nil
	}

	top := new(node)
	middleIndex := len(laptops) / 2
	top.laptop = laptops[middleIndex]
	top.profit = top.laptop.sellPrice - top.laptop.buyPrice
	top.maxProfit = top.profit

	top.left = n.buildTree(laptops[:middleIndex])
	top.right = n.buildTree(laptops[middleIndex+1:])

	if top.right != nil {
		top.right.up = top
		top.maxProfit = max(top.maxProfit, top.right.maxProfit)
	}
	if top.left != nil {
		top.left.up = top
		top.maxProfit = max(top.maxProfit, top.left.maxProfit)
	}

	return top
}

type studentController struct {
	laptopsTree *node
	capital     int
}

func (c *studentController) buildLaptopsTree(laptops []laptopData) {
	c.laptopsTree = c.laptopsTree.buildTree(laptops)
}

func (c *studentController) buyLaptop(top *node) {
	if top == nil {
		return
	}
	c.capital += top.profit
	top.profit = 0

	tempTop := top
	for tempTop != nil {
		tempTop.maxProfit = tempTop.profit
		if tempTop.left != nil {
			tempTop.maxProfit = max(tempTop.maxProfit, tempTop.left.maxProfit)
		}
		if tempTop.right != nil {
			tempTop.maxProfit = max(tempTop.maxProfit, tempTop.right.maxProfit)
		}
		tempTop = tempTop.up
	}
}

func (c *studentController) getMostProfitLaptop(top *node) *node {

	if top == nil || top.maxProfit == 0 {
		return nil
	}
	if top.left != nil && top.maxProfit == top.left.maxProfit {
		return c.getMostProfitLaptop(top.left)
	}
	if top.right != nil && top.maxProfit == top.right.maxProfit {
		return c.getMostProfitLaptop(top.right)
	}
	return top
}

func (c *studentController) maxProfitLaptop(laptop1, laptop2 *node) *node {
	if laptop1 != nil && laptop2 != nil {
		if laptop1.profit > laptop2.profit {
			return laptop1
		}
		return laptop2
	}
	if laptop1 != nil {
		return laptop1
	}
	return laptop2
}

func (c *studentController) buyMostProfitLaptopWithCapital(top *node) *node {
	if top == nil {
		return nil
	}

	if top.laptop.buyPrice > c.capital {
		return c.buyMostProfitLaptopWithCapital(top.left)
	}

	maxProfitLaptop := c.maxProfitLaptop(c.getMostProfitLaptop(top.left), c.buyMostProfitLaptopWithCapital(top.right))
	if top.profit != 0 {
		return c.maxProfitLaptop(maxProfitLaptop, top)
	}
	return maxProfitLaptop
}

func printTree(top *node) {
	if top == nil {
		return
	}

	printTree(top.left)
	fmt.Println(top)
	printTree(top.right)
}

func getReader(filePath string) (*csv.Reader, *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return csv.NewReader(file), file
}

func getRecord(reader *csv.Reader) ([]string, bool) {
	record, err := reader.Read()
	if err == io.EOF {
		return record, false
	}
	if err != nil {
		panic(err)
	}
	if len(record) != 2 {
		panic(errors.New("incorrect record"))
	}
	return record, true
}

func handleRecord(mas *[]laptopData, record []string) {

	var laptop laptopData
	var err error

	laptop.id = len(*mas)
	laptop.buyPrice, err = strconv.Atoi(record[0])
	laptop.sellPrice, err = strconv.Atoi(record[1])

	if err != nil {
		return
	}
	*mas = append(*mas, laptop)
}

func readLaptopsFromCSV(path string) []laptopData {
	reader, file := getReader(path)
	defer file.Close()

	//Read header of CSV file
	record, _ := getRecord(reader)

	var result []laptopData

	breakFlag := true
	for breakFlag {
		record, breakFlag = getRecord(reader)
		if breakFlag {
			handleRecord(&result, record)
		}
	}
	return result
}

func main() {
	laptops := readLaptopsFromCSV("laptops.csv")
	sort.Slice(laptops, func(i, j int) bool {
		return laptops[i].buyPrice < laptops[j].buyPrice
	})
	var student studentController
	student.buildLaptopsTree(laptops)
	student.capital = 9
	//printTree(student.laptopsTree)

	for i := 0; i < 10; i++ {
		a := student.buyMostProfitLaptopWithCapital(student.laptopsTree)
		if a != nil {
			student.buyLaptop(a)
		}
		//fmt.Println("///////////////////////////")
		//printTree(student.laptopsTree)
		fmt.Println(a)
		//fmt.Println(student.capital)
		//fmt.Println("///////////////////////////")
	}

}
