package main
import "fmt"
func main(){
	fmt.Print(tictactoe([]string{"OX ","OX ","O  "}))
}
//位运算
//func tictactoe(board []string) string {
//	isFull := true
//
//	XDiaWin := false
//	XDiaWinP := false
//
//	ODiaWin := false
//	ODiaWinP := false
//	for i:=0;i<len(board);i++{
//		XRowWin := false
//		XColWin := false
//		ORowWin := false
//		OColWin := false
//		for j:=0;j<len(board);j++{
//			if isFull && board[i][j]==' '{
//				isFull = false
//			}
//			if(!XRowWin){
//				if 'X' ^ board[i][j] == 0{
//					XRowWin = false
//				}else{
//					XRowWin = true
//				}
//			}
//			if(!XColWin){
//				if 'X' ^ board[j][i] == 0{
//					XColWin = false
//				}else{
//					XColWin = true
//				}
//			}
//			if(!ORowWin){
//				if 'O' ^ board[i][j] == 0{
//					ORowWin = false
//				}else{
//					ORowWin = true
//				}
//			}
//			if(!OColWin){
//				if 'O' ^ board[j][i] == 0{
//					OColWin = false
//				}else{
//					OColWin = true
//				}
//			}
//
//		}
//		if !XRowWin || !XColWin{return "X"}
//		if !ORowWin || !OColWin{return "O"}
//		if(!XDiaWin){
//			if 'X' ^ board[i][i] == 0{
//				XDiaWin = false
//			}else{
//				XDiaWin = true
//			}
//		}
//		if(!XDiaWinP){
//			if 'X' ^ board[i][len(board)-i-1] == 0{
//				XDiaWinP = false
//			}else{
//				XDiaWinP = true
//			}
//		}
//		if(!ODiaWin){
//			if 'O' ^ board[i][i] == 0{
//				ODiaWin = false
//			}else{
//				ODiaWin = true
//			}
//		}
//		if(!ODiaWinP){
//			if 'O' ^ board[i][len(board)-i-1] == 0{
//				ODiaWinP = false
//			}else{
//				ODiaWinP = true
//			}
//		}
//	}
//	if !XDiaWin || !XDiaWinP{return "X"}
//	if !ODiaWin || !ODiaWinP{return "O"}
//	if isFull{
//		return "Draw"
//	}
//	return "Pending"
//}


//计算每行每列和每个对角线的字符个数是否等于N
func tictactoe(board []string) string {
	xDiaCount := 0
	xDiaCountP := 0
	oDiaCount := 0
	oDiaCountP := 0
	isFull := true
	n := len(board)
	for i:=0;i<n;i++{
		if board[i][i] == 'X'{
			xDiaCount++
		}
		if board[i][i] == 'O'{
			oDiaCount++
		}
		if board[i][n-i-1] == 'X'{
			xDiaCountP++
		}
		if board[i][n-i-1] == 'O'{
			oDiaCountP++
		}
		xRowCount := 0
		xColCount := 0
		oRowCount := 0
		oColCount := 0
		for j:=0;j<n;j++{
			if board[i][j] == 'X'{
				xRowCount++
			}
			if board[i][j] == 'O'{
				oRowCount++
			}
			if board[j][i] == 'X'{
				xColCount++
			}
			if board[j][i] == 'O'{
				oColCount++
			}
			if isFull && board[i][j] == ' '{
				isFull = false
			}
		}
		if xRowCount == n || xColCount == n{
			return "X"
		}
		if oRowCount == n || oColCount == n{
			return "O"
		}
	}
	if xDiaCount == n || xDiaCountP == n{
		return "X"
	}
	if oDiaCount == n || oDiaCountP == n{
		return "O"
	}
	if isFull{
		return "Draw"
	}
	return "Pending"
}