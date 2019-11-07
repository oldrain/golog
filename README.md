# golog
A simple logger, small but enough


## Installation

> Use go get

    go get -u github.com/oldrain/golog

## Features
* Levels: off, fatal, error, warn, info, debug, trace, all
* Log to console or file
* Colored console
* File log support timer writing. Date and size rotate
* We can append something(e.g. uuid) to head or tail of the log message

## Usage

### New logger

> Init

```go
package main

import (
    "github.com/oldrain/golog"
)

func main() {
    config := new(golog.Config)
    config.SetLevel(golog.LevelInfo)
    config.SetPath("logs")
    config.SetRotate(golog.RotateDate)

    golog.SetLogMode(golog.ModeFile)
    golog.SetGlobalConfig(config)

    var moduleName = "app.user"
    logger := golog.GetLogger(moduleName)
    logger.Info("info...")
}

// 2019-01-01 00:00:00 [app.user] [INFO] [append...]
```

> Appending


```go
package main

import (
    "github.com/oldrain/golog"
)

func main() {
    var moduleName = "app.user"

    logger := golog.GetFileLogger(moduleName)
    logger.AppendHead("[before]")
    logger.AppendTail("[after]")
    // Erase(), EraseHead(), EraseTail

    logger.Info("append...")
}

// 2019-01-01 00:00:00 [app.user] [INFO] [before] [append...] [after]
```

### Console log

> Default logger

```go
package main

import (
    "github.com/oldrain/golog"
)

func main() {
    var moduleName = "app.user"
    logger := golog.GetConsoleLogger(moduleName)
    logger.Info("console info...")
}

// 2019-01-01 00:00:00 [app.user] [INFO] [console info...]
```

> or customize

```go
package main

import (
    "github.com/oldrain/golog"
)

func main() {
    var moduleName = "app.user"

    config := new(golog.Config)
    config.SetLevel(golog.LevelInfo)

    logger := golog.ConsoleLogger(moduleName, config)
    logger.Info("console info...")
}

// 2019-01-01 00:00:00 [app.user] [INFO] [console info...]
```

### File Log

> Rotate file size

```go
package main

import (
    "github.com/oldrain/golog"
)

func main() {
	var moduleName = "app.user"

    config := new(golog.Config)
    config.SetLevel(golog.LevelInfo)
    config.SetPath("logs")
    config.SetRotate(golog.RotateSize)
    config.SetRotateSize(10 * golog.MB)

    logger := golog.GetLogger(moduleName)
    logger.Info("file size...")
}

// 2019-01-01 00:00:00 [app.user] [INFO] [file size...]
```

> Time writing

```go
package main

import (
    "github.com/oldrain/golog"
)

func main() {
    var moduleName = "app.timer"

    config := new(golog.Config)
    config.SetLevel(golog.LevelInfo)
    config.SetPath("logs")
    config.SetRotate(golog.RotateDate)
    config.SetTimerWrite(true)

    logger := golog.FileLogger(moduleName, config)
    logger.Info("timer writing...")
}

// 2019-01-01 00:00:00 [app.user] [INFO] [timer writing...]
```
## Output
![output](https://github.com/oldrain/golog/raw/master/output.png)