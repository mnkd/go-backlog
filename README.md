# go-backlog
Go library for accessing the Backlog API https://developer.nulab-inc.com/ja/docs/backlog/

Inspired by https://github.com/google/go-github

# Usage

```
import "github.com/mnkd/go-backlog/backlog"
```

Construct a new Backlog client. For example:

```go
client := backlog.NewClient(nil, space, apiKey)

// list all projects for youe Backlog space
projects, _, err := client.Projects.ListAll()
```

See also [examples](./examples)

# Test
```
$ go test ./backlog
```

# License
This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
