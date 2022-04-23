/*
 * +===============================================
 * | Author:        Ahmad Foroughi <pyhp2017@gmail.com>
 * |
 * | Creation Date: 23-04-2022
 * |
 * | File Name:     p10.go
 * +===============================================
 */

package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row_nums := [10]int{}
		col_nums := [10]int{}
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				row_nums[int(board[j][i]-'0')]++
			}
			if board[i][j] != '.' {
				col_nums[int(board[i][j]-'0')]++
			}
		}
		for j := 0; j < 10; j++ {
			if row_nums[j] > 1 || col_nums[j] > 1 {
				return false
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			box_nums := [10]int{}
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if board[k][l] != '.' {
						box_nums[int(board[k][l]-'0')]++
					}
				}
			}
			for k := 0; k < 10; k++ {
				if box_nums[k] > 1 {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]string{}
	for i := 0; i < 9; i++ {
		input := make([]string, 9)
		for j := 0; j < 9; j++ {
			fmt.Scanf("%s", &input[j])
		}
		board = append(board, input)
	}

	converted_board := [][]byte{}
	for i := 0; i < 9; i++ {
		input := make([]byte, 9)
		for j := 0; j < 9; j++ {
			input[j] = byte(board[i][j][0])
		}
		converted_board = append(converted_board, input)
	}

	if isValidSudoku(converted_board) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
