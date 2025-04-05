package g-util

import "fmt"

type SparseArray struct {
	Row int
	Col int
	Val int
}

func main() {

	chess := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 22, 0, 0, 0, 0, 15, 0, 0},
		{0, 0, -6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, -15, 0, 0, 0, 0, 0},
	}

	sparseArray := sliceToSparseArray(chess)
	fmt.Printf("%#v\n", sparseArray)

	slice := sparseToSlice(sparseArray)
	//打印slice
	for _, row := range slice {
		for _, col := range row {
			fmt.Printf("%d\t", col)
		}
		fmt.Println()
	}
}

func sliceToSparseArray(slice [][]int) []SparseArray {
	sparseArray := make([]SparseArray, 0)
	row := len(slice)
	col := len(slice[0])

	effectiveDataCount := 0

	for rowIndex, row := range slice {
		for colIndex, col := range row {
			if col == 0 {
				continue
			}
			effectiveDataCount++
			sparseArray = append(sparseArray, SparseArray{Row: rowIndex, Col: colIndex, Val: col})
		}
	}

	finalResult := make([]SparseArray, 0)
	finalResult = append(finalResult, SparseArray{Row: row, Col: col, Val: effectiveDataCount})
	finalResult = append(finalResult, sparseArray...)

	return finalResult
}

// 稀疏数组转二维数组
func sparseToSlice(sparse []SparseArray) [][]int {
	// 由于稀疏数组的特性，第一行的数据可以确定二维数组的行列
	var (
		rows  = sparse[0].Row
		col   = sparse[0].Col
		slice [][]int
	)

	// 初始化二维数组
	for i := 0; i < rows; i++ {
		colData := make([]int, col)
		slice = append(slice, colData)
	}

	// 为二维数组的有效数据赋值
	for i, data := range sparse {
		if i == 0 {
			continue
		}
		slice[data.Row][data.Col] = data.Val
	}

	return slice

}
