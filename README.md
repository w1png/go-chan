# ⚠️ Caution: this is a WIP

# go-chan
go-chan is a [4chan api](https://github.com/4chan/4chan-API) wrapper for golang.

## Examples
### Getting all boards
```go
package main

import (
    "fmt"

    gochan "github.com/w1png/go-chan"
)

func main() {
    // []types.Board
    boards, _ := gochan.GetBoards()

	for _, board := range boards {
		fmt.Println(board)
	}
}
```

### Getting pages and threads
#### Pages
```go
package main

import (
    "fmt"

    gochan "github.com/w1png/go-chan"
)

func main() {
	board, _ := gochan.GetBoard("g")
	pages, _ := board.GetPages()

	fmt.Println(pages)
}

```

#### Threads
```go
package main

import (
    "fmt"

    gochan "github.com/w1png/go-chan"
)

func main() {
	board, _ := gochan.GetBoard("g")
	threads, _ := board.GetAllThreads()

	fmt.Println(threads)
}
```

