# GO (Programming Language)

Notes:

```PowerShell
# run main program direct from the file
go run main.go

# Compiles the program into a executable (.exe in Windows)
go build main.go
./main.exe
```

Go was created by a team from Google. Go is a language that has the feel dynamic languages such as JavaScript and PHP and the performance and efficiency of strongly typed languages as C++ and Java.

- compiled
- memory safaty
- channel-based concorrency

Go is written in text files and then compiled to machine code packaged into a single, standalone executable file. The result is self-contained meaning it doesn't need nothing more to run. Go programs can be compiled to Windows, Linux, macOS and Android.
- Write once and run it anywhere.

Go has a statically typed and type-safe memory model with a garbage collector.
- Go has same level of productivity as dynamically typed languages without giving up perfomance and efficiency.
- Go is designed to take advantage of multiple CPU cores.

This code below is our package declaration. If you want to run the code directly, you'll need to name it main. If you don't name it main, then you can use it as a library and import it into other Go code.

```Go
package main
```

Importing code from packages
```Go
// Import extra functionality from packages
import (
  "errors"
  "fmt"
  "log"
  "math/rand"
  "strconv"
  "time"
)
```

