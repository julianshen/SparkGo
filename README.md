SparkGo
=======
Example usage
-------
````go
package main

import (
   spark "github.com/julianshen/SparkGo"
   "time"
)


func main() {
    core := spark.SparkCore{"YOUR_ACCESS_TOKEN","DEVICE_ID"}
 
    c := time.Tick(1 * time.Second)
    b := true
    for _ = range c {
        core.DigitalWrite("D7", b)
        b=!b
    }
}
````