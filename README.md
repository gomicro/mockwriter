# Pen Name
A mock that implements the Writer interface for testing.

# Example

```go
import(
	"fmt"
	"io"
	"os"

	"github.com/gomicro/penname"
)

func main(){
	mockWrite := penname.New()
	mw := io.MultiWriter(os.Stdout, mockWrite)

	mw.Write("A random line to write")

	if strings.Contains( string(mockWrite.Written), "random" ){
		fmt.Println("Found a random")
	}
}
```
