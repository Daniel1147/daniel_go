package dp

// import "fmt"

var board [][]int
var wayList [][]int

func knightProbability(N int, K int, r int, c int) float64 {
    initBoard(N)
    board[r][c] = 1
    boardIterate(N, K)
    sum := sumBoard(N)
    total := 1.0

    for i := 0; i < K; i++ {
        total *= 8.0
    }

    result := sum / total

    return result
}

func initBoard(n int) {
    board = make([][]int, n)
    for i := 0; i < n; i++{
        board[i] = make([]int, n)
    }

    wayList = [][]int{
        {1,  2},
        {-1, 2},
        {1,  -2},
        {-1, -2},
        {2,  1},
        {-2, 1},
        {2,  -1},
        {-2, -1},
    }
}

func boardIterate(n int, k int) {
    for j := 0; j < k; j++ {
        newBoard := make([][]int, n)
        for i := 0; i < n; i++{
            newBoard[i] = make([]int, n)
        }

        for i := range board {
            for j, _ := range board[i] {
                addBoard(newBoard, n, board[i][j], i, j)
            }
        }
        board = newBoard
    }
}

func addBoard(newBoard [][]int, n int, num int, r int, c int) {
    for _, way := range wayList {
        tmp_r := r + way[0]
        tmp_c := c + way[1]
        if (tmp_r > -1 && tmp_c > -1 && tmp_r < n && tmp_c < n) {
            newBoard[tmp_r][tmp_c] += num
        }
    }
}

func sumBoard(n int) float64 {
    sum := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++{
            sum += board[i][j]
        }
    }

    return float64(sum)
}
