package examples

//
// Example of using buffered FIFO with LSM6DSL Gyroscope (3 axes)
//

import (
    "fmt"

    "github.com/brendan-ta/goiio"
)

const (
    iiodhost = "10.7.0.27"
    iiodport = "30431"
)

func convertToLeS16(msb uint16, lsb uint16) int16 {
    n := msb<<8 + lsb
    return int16(n)
}

func main() {
    c, err := goiio.New(fmt.Sprintf("%s:%s", iiodhost, iiodport))
    if err != nil {
        panic(err)
    }

    defer c.Close()

    device_name := "lsm6dsl_gyro"

    err = c.CreateBuffer(device_name, 24, "00000007")
    if err != nil {
        panic(err)
    }

    k := 0
    for k <= 10 {
        val, e := c.DumpBuffer(device_name, 96)
        if e != nil {
            fmt.Println("dump error")
            panic(e)
        }
        for ; len(val) >= 2; {
            v := convertToLeS16(uint16(val[1]), uint16(val[0]))
            fmt.Printf("%d\n", v)
            val = val[2:]
        }
        k++
    }

    err = c.DestroyBuffer(device_name)
    if err != nil {
        panic(err)
    }
}
