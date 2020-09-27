var fathers = make(map[int]int)
func numIslands(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}

	nc := len(grid[0])

	// 初始化fahters
	for i := 0; i < nr * nc; i++ {
		fathers[i] = i
	}
	
	// 循环列表，将相邻的单独岛屿合并
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			// 如果是海洋  不处理
			if grid[i][j] == '0' {
				continue
			}
			// 转换当前二维坐标 => 一维坐标
			cur_pos := i * nc + j
			// 右边的坐标是否是陆地
			if j+1 < nc && grid[i][j+1] == '1' {
				// 合并这两个坐标
				union(cur_pos, cur_pos+1)
			}
			// 下边的坐标是否是陆地
			if i+1 < nr && grid[i+1][j] == '1' {
				union(cur_pos, cur_pos+nc)
			}
		}
	}
	
	// 计算岛屿数量
	island := 0
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' && fathers[i*nc+j] == i*nc+j {
				island++
			}
		}
	}
	return island
}



func find(i int) int {
	bak := i
	if fathers[i] == i {
		return i
	}
	
	for fathers[i] != i {
		i = fathers[i]
	}
	
	for bak != i {
		tmp := fathers[bak]
		fathers[bak] = i
		bak = tmp
	}
	
	return i
}

func union(i, j int)  {
	f1 := find(i)
	f2 := find(j)
	if f1 != f2 {
		fathers[f1] = f2
	}
}
