package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetSoonestPossibleDeparture(ts, bus int) int {
    rem := ts % bus
    if rem == 0 {
        return ts
    }
    return bus - rem
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    ts, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    busesStr := strings.Split(scanner.Text(), ",")
    var buses []int
    for _, busStr := range busesStr {
        bus_num, err := strconv.Atoi(busStr)
        if err == nil {
            buses = append(buses, bus_num)
        }
    }

    // fmt.Println("ts:", ts)
    // for _, bus := range buses {
    //     fmt.Println(bus, GetSoonestPossibleDeparture(ts, bus))
    // }

    min_ts := GetSoonestPossibleDeparture(ts, buses[0])
    min_bus := buses[0]
    for i := 1; i < len(buses); i++ {
        dep := GetSoonestPossibleDeparture(ts, buses[i])
        if dep < min_ts {
            min_ts = dep
            min_bus = buses[i]
        }
    }
    fmt.Println("[1]", min_ts * min_bus)
}


