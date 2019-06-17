# e
error handling utilities

## Usage

```bash
go get github.com/reusee/e/v2
```

```go
import (
  "github.com/reusee/e/v2"
)

var (
  makeErr = e.Default.WithStack().WithName("program")
  checkErr, handleErr = e.New(makeErr)
  // or name them me, ce, he to save key strokes.
)
```

```go
func CopyFile(src, dst string) (err error) {
  defer he(&err, "copy %s to %s", src, dst)

  r, err := os.Open(src)
  ce(err, "open %s", src)
  defer r.Close() 

  w, err := os.Create(dst)
  ce(err, "create %s", dst)
  defer func() {
    w.Close()
    if err != nil {
      os.Remove(dst)
    }
  }()

  _, err = io.Copy(w, r)
  ce(err, "copy")

  ce(w.Close(), "close writer")

  return nil
}
```

