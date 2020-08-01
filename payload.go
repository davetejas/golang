// We have 3 structs: Cpu, Mem and Disk.
// There are predefined ‘fetch’ functions that return instances of each.

// Complete the fetchPayload() method to return a Payload. Optimize your implementation to be as performant as possible.


package main

import (
	"fmt"
)

func main() {
	x := fetchPayload()
	fmt.Println("CPU:",x.CpuMetric)
	fmt.Println("MEM:",x.MemMetric)
	fmt.Println("DISK:",x.DiskMetric)
}

type Cpu struct {
	Pct float64
}

type Mem struct {
	UsedBytes int64
}

type Disk struct {
	UsedBytes int64
}


//predefined method
func fetchCPUMetric() Cpu {
	//querys system and returns Cpu struct
	//dummy return val
	return Cpu {
		Pct: float64(2),
	}
}

//predefined method
func fetchMemMetric() Mem {
	//querys system and returns Mem struct
	//dummy return val
	return Mem {
		UsedBytes: 7500,
	}
}

//predefined method
func fetchDiskMetric() Disk {
	//querys system and returns Disk struct
	//dummy return val
	return Disk{
		UsedBytes: 199,
	}
}

type Payload struct {
	CpuMetric Cpu
	MemMetric Mem
	DiskMetric Disk
}


func fetchPayload() Payload {
	//TODO complete this. write the most optimized implementation
	
	var dump Payload
	
	dump.CpuMetric = fetchCPUMetric()
	dump.MemMetric = fetchMemMetric()
	dump.DiskMetric = fetchDiskMetric()
	
	return dump 
}
