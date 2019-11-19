# Summary

luniq - short non-sequential url-friendly unique id generator written on pure Go (version >= 1.10). You can safely use luniq on more than one machine because it's contain hostname hash. Package is optimized for multithreaded access. Moreover, it precomputes a pool of values ​​to reduce responce time.

# Install

```
go get github.com/belfinor/luniq
```

# Make new uniq

Process is very simple. You need only call *Next* method:

```go
package main

import (
  "fmt"

  "github.com/belfinor/luniq"
)

func main() {

  for i := 0 ; i < 10 ; i++ {
    fmt.Println( luniq.Next() )
  }
}
```

This code prints output like this:

```
f590291092005a35f5dd440c4aa157f22e61ad22459700cb2
f0af8737755e73b175dd440c4aa157f22e61b1c744704a50d
f269745a6aa3e2f5c5dd440c4aa157f22e61c2c87cb4a3aa7
f2f2ec5e1f9f2e31a5dd440c4aa157f22e61d372f8e364e9f
f61c2a0ab8f236c225dd440c4aa157f22e61e423470fd19bf
f0dae5b81de12a3dd5dd440c4aa157f22e61f4d8b5b964fc8
f3529bfdf62c7c3cb5dd440c4aa157f22e62058afdfc3a788
f79aec50fae6e56c75dd440c4aa157f22e621611619688ddd
f434584e8a20465375dd440c4aa157f22e622664bdf90c72a
f7112aa68239618395dd440c4aa157f22e6236b533f253ffb
```

# Verify uniq

Moreover, luniq allows you to check uniq values:

```go
package main

import (
  "fmt"

  "github.com/belfinor/luniq"
)

func main() {

  uniq := luniq.Next()

  // lite check if 2nd param is false
  if luniq.Check(uniq, true) {
    fmt.Println( uniq, "valid")
  }

}

```
