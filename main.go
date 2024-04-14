package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

type laptopData struct {
	id       int
	buyPrice int
	gains    int
}
type node struct {
	up     *node
	left   *node
	right  *node
	laptop laptopData
	//profit    int
	maxGainLaptop *node
}

func (n *node) maxGain(first *node, second *node) *node {
	if first.laptop.gains > second.laptop.gains {
		return first
	}
	return second
}

func (n *node) buildTree(laptops []laptopData) *node {
	if len(laptops) == 0 {
		return nil
	}

	top := new(node)
	middleIndex := len(laptops) / 2
	top.laptop = laptops[middleIndex]
	top.maxGainLaptop = top

	top.left = n.buildTree(laptops[:middleIndex])
	top.right = n.buildTree(laptops[middleIndex+1:])

	if top.right != nil {
		top.right.up = top
		top.maxGainLaptop = n.maxGain(top.maxGainLaptop, top.right.maxGainLaptop)
	}
	if top.left != nil {
		top.left.up = top
		top.maxGainLaptop = n.maxGain(top.maxGainLaptop, top.left.maxGainLaptop)
	}

	return top
}

type studentController struct {
	laptopsTree *node
	capital     int
	laptopLimit int
}

func (c *studentController) buildLaptopsTree(laptops []laptopData) {
	c.laptopsTree = c.laptopsTree.buildTree(laptops)
}

func (c *studentController) buyLaptop(top *node) {
	if top == nil {
		return
	}
	c.capital += top.laptop.gains
	top.laptop.gains = 0

	tempTop := top
	for tempTop != nil {
		tempTop.maxGainLaptop = tempTop
		if tempTop.left != nil {
			tempTop.maxGainLaptop = tempTop.maxGain(tempTop.maxGainLaptop, tempTop.left.maxGainLaptop)
		}
		if tempTop.right != nil {
			tempTop.maxGainLaptop = tempTop.maxGain(tempTop.maxGainLaptop, tempTop.right.maxGainLaptop)
		}
		tempTop = tempTop.up
	}
}

func (c *studentController) maxProfitLaptop(laptop1, laptop2 *node) *node {
	if laptop1 != nil && laptop2 != nil {
		if laptop1.laptop.gains > laptop2.laptop.gains {
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
	if top.left == nil {
		return top.maxGainLaptop
	}

	maxProfitLaptop := c.maxProfitLaptop(top.left.maxGainLaptop, c.buyMostProfitLaptopWithCapital(top.right))
	if top.laptop.gains != 0 {
		return c.maxProfitLaptop(maxProfitLaptop, top)
	}
	return maxProfitLaptop
}

func (c *studentController) getResult() []laptopData {

	var result []laptopData

	for i := 0; i < c.laptopLimit; i++ {
		laptopToBuy := c.buyMostProfitLaptopWithCapital(c.laptopsTree)
		if laptopToBuy != nil {
			c.buyLaptop(laptopToBuy)
			result = append(result, laptopToBuy.laptop)
		}
	}
	return result
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
	laptop.gains, err = strconv.Atoi(record[1])

	if err != nil {
		return
	}

	if laptop.buyPrice <= 0 || laptop.gains <= 0 {
		return
	}

	*mas = append(*mas, laptop)
}

func readLaptopsFromCSV(path string) []laptopData {
	reader, file := getReader(path)

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
	file.Close()

	sort.Slice(result, func(i, j int) bool {
		return result[i].buyPrice < result[j].buyPrice
	})

	return result
}

func megaTest() {
	path := "laptops.csv"
	for {
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}

		writer := csv.NewWriter(file)
		for i := 0; i <= 1000; i++ {
			r := rand.Intn(1000) + 1
			writer.Write([]string{strconv.Itoa(r), strconv.Itoa(r + rand.Intn(100))})
			writer.Flush()
		}
		n := rand.Intn(50)
		fmt.Println("||||||||||||||||||||||||||||||||||||||||||||")
		laptops := readLaptopsFromCSV("laptops.csv")
		sort.Slice(laptops, func(i, j int) bool {
			return laptops[i].buyPrice < laptops[j].buyPrice
		})
		var student studentController
		student.buildLaptopsTree(laptops)
		student.capital = 9
		//printTree(student.laptopsTree)

		for i := 0; i < n; i++ {
			a := student.buyMostProfitLaptopWithCapital(student.laptopsTree)
			if a != nil {
				student.buyLaptop(a)
				fmt.Println(a.laptop)
			}
			//fmt.Println("///////////////////////////")
			//printTree(student.laptopsTree)
			//fmt.Println(student.capital)
			//fmt.Println("///////////////////////////")
		}
		file.Close()
		os.Remove(path)
	}
}

func getConsoleArray(arrayLen int) []int {

	var array []int

	for i := 0; i < arrayLen; i++ {
		var gains int
		_, err := fmt.Scan(&gains)
		if err != nil {
			fmt.Println("Can not read data")
			break
		}

		array = append(array, gains)
	}

	return array
}

func getLaptop(buyPrice int, gains int) laptopData {
	var res laptopData
	res.buyPrice = buyPrice
	res.gains = gains
	return res
}

func getLaptopsArray(buyPriceArray []int, gainsArray []int) []laptopData {
	var res []laptopData
	if len(buyPriceArray) != len(gainsArray) {
		return res
	}

	for i := 0; i < len(buyPriceArray); i++ {
		laptop := getLaptop(buyPriceArray[i], gainsArray[i])
		laptop.id = i
		res = append(res, laptop)
	}

	return res
}

func getParams() studentController {
	var res studentController

	fmt.Println("Enter laptops limit(N):")
	_, err := fmt.Scan(&res.laptopLimit)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter the capital(C):")
	_, err = fmt.Scan(&res.capital)
	if err != nil {
		panic(err)
	}

	var laptopsCount int
	fmt.Println("Enter laptops count(K)")
	_, err = fmt.Scan(&laptopsCount)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter laptops gains:")
	gains := getConsoleArray(laptopsCount)

	fmt.Println("Enter laptop buy prices:")
	buyPrices := getConsoleArray(laptopsCount)

	laptops := getLaptopsArray(buyPrices, gains)
	res.buildLaptopsTree(laptops)

	return res
}

func printResult(totalCapital int) {
	fmt.Printf("Capital at the end of the summer: %d\n", totalCapital)
}

func main() {
	student := getParams()
	student.getResult()
	printResult(student.capital)
}
