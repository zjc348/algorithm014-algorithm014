func findCircleNum(M [][]int) int {
    ll := len(M)
    if ll == 0{
        return 0
    }
    ufind := NewUnionFind(ll)
    for i := 0;i < ll;i++{
        for j := i+1;j < ll;j++{
            if M[i][j] == 1{
                ufind.Union(i, j)
            }
        }
    }
    return ufind.count
}

type UnionFind struct{
    count int
    parent []int
}

func NewUnionFind(n int) *UnionFind{
    parent := make([]int, n)
    for i := 0;i < n;i++{
        parent[i] = i
    }
    return &UnionFind{
        count: n,
        parent: parent,
    }
}

func (u *UnionFind) Find(i int)int{
    root := i
    for u.parent[root] != root{
        root = u.parent[root]
    }
    for u.parent[i] != i{
        i, u.parent[i] = u.parent[i], root
    }
    return root
}

func (u *UnionFind) Union(i, j int){
    pi := u.Find(i)
    pj := u.Find(j)
    if pi != pj{
        u.parent[pi] = pj
        u.count--
    }
}