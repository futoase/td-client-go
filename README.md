td-client-go
============

TreasureData client for golang.

Referenced of [td-client-node](https://github.com/treasure-data/td-client-node).

How to install
--------------

```
> go get github.com/futoase/td-client-go
```

Sample
------

```go
package main

import (
	"fmt"
	"github.com/futoase/td-client-go"
)

func main() {
	options := td.NewOptions("Treasure data api key")
	client := td.NewTreasureData(options)
	d := client.ListDatabases()
	for i := 0; i < len(d.Databases); i++ {
	  row = d.Database[i]
		fmt.Printf("table: %v\n count: %v\n created_at: %v\n updated_at: %v\n", row.Name, row.Count, row.Created_At, row.Updated_At)
	}
}
```

**result**

```
> go build main.go
> ./main
table: applications
 count: 4
 created_at: 2013-08-16 21:11:22 UTC
 updated_at: 2013-08-16 21:11:22 UTC
table: debug
 count: 5
 created_at: 2013-09-04 08:43:59 UTC
 updated_at: 2013-09-04 08:43:59 UTC
table: fluent-test
 count: 209
 created_at: 2013-07-12 10:29:02 UTC
 updated_at: 2013-07-12 10:29:02 UTC
table: api-test
 count: 7
 created_at: 2013-07-05 02:31:50 UTC
 updated_at: 2013-07-05 02:31:50 UTC
```

LICENSE
-------

MIT.


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/futoase/td-client-go/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

