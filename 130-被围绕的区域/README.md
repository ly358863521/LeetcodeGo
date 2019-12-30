

DFS递归

```go
func solve(board [][]byte)  {
    if board==nil ||len(board)==0{
        return 
    }
    m := len(board)
    n := len(board[0])
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            var isedge bool
            isedge = (i==0||j==0||i==m-1||j==n-1)
            if(isedge&&board[i][j]=='O'){
                dfs(board,i,j)
            }
        }
    }
    for i:=0;i<m;i++{
        for j :=0;j<n;j++{
            if board[i][j]=='O'{
                board[i][j]='X'
            }
            if board[i][j]=='#'{
                board[i][j]='O'
            }
        }
    }
}

func dfs(board [][]byte,i,j int){
    if (i<0 || j<0 || i>=len(board) || j>=len(board[0]) || board[i][j]=='X' || board[i][j]=='#'){
        return
    }
    board[i][j]='#'
    dfs(board,i-1,j)
    dfs(board,i+1,j)
    dfs(board,i,j-1)
    dfs(board,i,j+1)
}
```



并查集

