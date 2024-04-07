package common

import (
	"fmt"
	"io"
	"os"
)

func AnnounceChapter(w io.Writer, chapter int, name string) {
	if w == nil {
		w = os.Stdout
	}
	fmt.Fprintf(w, "Chapter %d: %s\n", chapter, name)
}
