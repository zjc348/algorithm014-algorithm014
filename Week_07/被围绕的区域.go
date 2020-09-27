func solve(board [][]byte) {
	if len(board) <= 2 {
		return
	}

	var (
		m, n         = len(board), len(board[0])
		dummy, ufind = m * n, NewUnionFind(m*n + 1)
		position     func(i, j int) int
	)
	position = func(i, j int) int {
		return i*n + j
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i == 0 || j == 0 || i == m-1 || j == n-1 {
					ufind.Union(position(i, j), dummy)
				}
				if i > 0 && board[i-1][j] == 'O' {
					ufind.Union(position(i, j), position(i-1, j))
				}
				if j > 0 && board[i][j-1] == 'O' {
					ufind.Union(position(i, j), position(i, j-1))
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !ufind.IsConnected(position(i, j), dummy) {
				board[i][j] = 'X'
			}
		}
	}
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
	}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != root {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u *UnionFind) IsConnected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}
