package main
import (
	"fmt"
)

func isValidSudoku(arr [9][9]string ) bool {
	var jcolumn	[9]map[string]struct{}
	var xBox	[9]map[string]struct{}

	for i := 0; i < 9; i++{
		iRow := make(map[string]struct{})
		jcolumn[i] = make(map[string]struct{})
		for j := 0; j < 9; j++{
			ii := int(i / 3) * 3 + int(j / 3)
			if xBox[ii] == nil {
				xBox[ii] = make(map[string]struct{})
			}
			if arr[i][j] == "." {
				continue
			}
			
			if _, ok := iRow[arr[i][j]]; ok {
				return false
			}
			if _, ok := jcolumn[i][arr[i][j]]; ok {
				return false
			}
			if _, ok := xBox[ii][arr[i][j]]; ok {
				return false
			}
			iRow[arr[i][j]] = struct{}{}
			jcolumn[i][arr[i][j]] = struct{}{}
			xBox[ii][arr[i][j]] = struct{}{}
		}	
	}

	return true
}

func main() {
	// https://leetcode.cn/problems/valid-sudoku/
	//
	arr := [9][9]string{
		[9]string{"5","3",".",".","7",".",".",".","."},
		[9]string{"6",".",".","1","9","5",".",".","."},
		[9]string{".","9","8",".",".",".",".","6","."},
		[9]string{"8",".",".",".","6",".",".",".","3"},
		[9]string{"4",".",".","8",".","3",".",".","1"},
		[9]string{"7",".",".",".","2",".",".",".","6"},
		[9]string{".","6",".",".",".",".","2","8","."},
		[9]string{".",".",".","4","1","9",".",".","5"},
		[9]string{".",".",".",".","8",".",".","7","9"},
	}
	fmt.Println(isValidSudoku(arr))
	fmt.Println("xxxx")
}
