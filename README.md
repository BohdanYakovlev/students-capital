# Students Capital
[Task #8](https://docs.google.com/document/d/1agfSo-UU_ONLyN5OUOQl_wIqhULtC5w9_fxyUCT9tNQ/edit?usp=sharing) in Distributed Lab Challenge

This command-line interface (CLI) application helps a student maximize their capital during the summer by buying, fixing, and selling laptops. The program is designed to select the most profitable combination of laptops within the given constraints to increase the student's capital.

## How to Run the Code

To run this console app, follow these steps:

1. Clone this repository to your local machine:

```
git clone <repository_url>
```

2. Navigate to the directory containing the Go code:

```
cd students-capital
```

3. Ensure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).

4. Execute the Go code by running the following command:

```
go run main.go
```

5.Follow the on-screen instructions to provide the required parameters.

6.Once the program finishes executing, it will display the capital at the end of the summer based on the selected laptops and their gains.

## Input Parameters:
  - Limit on the number of laptops purchased.
  - Initial capital.
  - Laptops count.
  - Array of gains for each laptop.
  - Array of prices for each laptop.

## Algorithm
The application utilizes a modified binary search tree data structure to efficiently select the most profitable combination of laptops while considering the given constraints of the student's capital and the number of laptops to buy.
  - Building the Tree: The program builds a binary search tree where each node represents a laptop. Laptops are sorted based on their buy prices.
  - Buying Laptops: The application traverses the tree to select the most profitable laptops within the given capital constraint. It iteratively updates the student's capital and set gains of laptop to zero to ignore it in next iteration.
  - Maximizing Profit: During the buying process, the program keeps track of the maximum gains achievable at each node. This information is used to efficiently select the most profitable laptops to buy.

## Conclusion
The Student’s Capital CLI application provides a robust solution for maximizing a student's capital during the summer by efficiently buying, fixing, and selling laptops. By leveraging a modified binary search tree algorithm, the program optimally selects the most profitable combination of laptops while adhering to the constraints of the student's capital and the number of laptops to buy. Feel free to reach out if you have any questions or encounter any issues while running the code.
