
# go-cli-log

  Colored / keyed CLI logging for Go command-line programs.

```
       install : package express@3.2.1
         fetch : tarball for express
        unpack : tarball to node_modules/express
        unpack : express/Readme.md
        unpack : express/lib/application.js
        unpack : express/lib/request.js
        unpack : express/lib/response.js
          warn : duplicate express@3.x package found
         error : something exploded
```

  View the [docs](https://godoc.org/github.com/visionmedia/go-cli-log).

## Example

```go
package main

import "github.com/visionmedia/go-cli-log"
import "errors"
import "fmt"

func main() {
  fmt.Println()
  log.Verbose = true
  log.Info("install", "package %s@%s", "express", "3.2.1")
  log.Debug("fetch", "tarball for express")
  log.Info("unpack", "tarball to node_modules/express")
  log.Debug("unpack", "express/Readme.md")
  log.Debug("unpack", "express/lib/application.js")
  log.Debug("unpack", "express/lib/request.js")
  log.Debug("unpack", "express/lib/response.js")
  log.Warn("duplicate %s package found", "express@3.x")
  log.Error(errors.New("something exploded"))
  fmt.Println()
}
```

# License

 MIT