# keva

DynamoDB is great! This makes it greater for me with my lambada functions...

### Installing

`go get lusikas.com/keva`



### Usage

First go create DynameDB table "tablename" with "key" (string) as primary
partition key.

And then just write code like:

```go
package main

import (
	"lusikas.com/keva"
)

func main() {
	kv := keva.New("tablename")
	kv.Set("blah", "aaaaaaaaaaaaa")
	kv.Set("blah", 123.444)
	kv.Set("blah", []string{"a", "b", "c"})
	_ = kv.GetSlice("blah")[1]
}
```
