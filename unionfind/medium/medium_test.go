package medium

import (
	"testing"
)

func Test_numIslands(t *testing.T)  {
	t.Log(numIslands([][]byte{
	{'1','1','1','1','0'},
	{'1','1','0','1','0'},
	{'1','1','0','0','0'},
	{'0','0','0','0','0'},
	}))

	t.Log(numIslands([][]byte{
	{'1','1','0','0','0'},
	{'1','1','0','0','0'},
	{'0','0','1','0','0'},
	{'0','0','0','1','1'},
	}))
}
