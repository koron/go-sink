# sink - put errors into sink at first, then evaluate later.

    go get -u github.com/koron/go-sink

```go
import "github.com/koron/go-sink"

errs := &sink.Errors{}

// Put an error.
errs.Put(err)

// Check errors are stored or not.
if errs.Has() {
  // Get first occured error.
  return errs.First()
}

for _, err := range errs.All() {
  // do something for each errors.
}

// Get a concatenated error.
err := errs.Error()
fmt.Println("something wrong: ", err.Error())
```

## LICENSE

MIT license.  See LICENSE.
