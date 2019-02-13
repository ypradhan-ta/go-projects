package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		fmt.Println("returning", err)
		return
	}

	// file read
	// fmt.Println(contents)
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

// author: ypradhan
// function to calculate CPU usage

func getCPUUsage() () {

	fileContent, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	normalCC := uint64(0)
	sysCC := uint64(0)
	idleCC := uint64(0)

	lines := strings.Split(string(fileContent), "\n")
	for _, line := range(lines) {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i<numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				fmt.Println(i, val)

				if i == 1 {
					normalCC = val
				}
				if i == 3 {
					sysCC = val
				}
				if i == 4 {
					idleCC = val
				}
			}
		}

		fmt.Println("normal: %v, sys: %v, idle: %v", normalCC, sysCC, idleCC)
		percentUsage := float64(normalCC + sysCC)*100/float64(normalCC + sysCC + idleCC)
		fmt.Println("Usage ", percentUsage, "%")
		return
	}
	return
}

func main() {
	//idle0, total0 := getCPUSample()
	//time.Sleep(3 * time.Second)
	//idle1, total1 := getCPUSample()

	//idleTicks := float64(idle1 - idle0)
	//totalTicks := float64(total1 - total0)
	//cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	//fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)

	//_, _ = getCPUSample()
	getCPUUsage()
	fmt.Println("Done")
}
