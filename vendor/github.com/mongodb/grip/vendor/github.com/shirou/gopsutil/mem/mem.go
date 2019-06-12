package mem

import (
	"encoding/json"

	"github.com/shirou/gopsutil/internal/common"
)

var invoke common.Invoker = common.Invoke{}

// Memory usage statistics. Total, Available and Used contain numbers of bytes
// for human consumption.
//
// The other fields in this struct contain kernel specific values.
type VirtualMemoryStat struct {
	// Total amount of RAM on this system
	Total uint64 `json:"total" bson:"total"`

	// RAM available for programs to allocate
	//
	// This value is computed from the kernel specific values.
	Available uint64 `json:"available" bson:"available"`

	// RAM used by programs
	//
	// This value is computed from the kernel specific values.
	Used uint64 `json:"used" bson:"used"`

	// Percentage of RAM used by programs
	//
	// This value is computed from the kernel specific values.
	UsedPercent float64 `json:"usedPercent,omitempty" bson:"usedPercent"`

	// This is the kernel's notion of free memory; RAM chips whose bits nobody
	// cares about the value of right now. For a human consumable number,
	// Available is what you really want.
	Free uint64 `json:"free" bson:"free"`

	// OS X / BSD specific numbers:
	// http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/
	Active   uint64 `json:"active" bson:"active"`
	Inactive uint64 `json:"inactive" bson:"inactive"`
	Wired    uint64 `json:"wired" bson:"wired"`

	// FreeBSD specific numbers:
	// https://reviews.freebsd.org/D8467
	Laundry uint64 `json:"laundry" bson:"laundry"`

	// Linux specific numbers
	// https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
	// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
	// https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
	Buffers        uint64 `json:"buffers" bson:"buffers"`
	Cached         uint64 `json:"cached" bson:"cached"`
	Writeback      uint64 `json:"writeback" bson:"writeback"`
	Dirty          uint64 `json:"dirty" bson:"dirty"`
	WritebackTmp   uint64 `json:"writebacktmp" bson:"writebacktmp"`
	Shared         uint64 `json:"shared" bson:"shared"`
	Slab           uint64 `json:"slab" bson:"slab"`
	SReclaimable   uint64 `json:"sreclaimable" bson:"sreclaimable"`
	PageTables     uint64 `json:"pagetables" bson:"pagetables"`
	SwapCached     uint64 `json:"swapcached" bson:"swapcached"`
	CommitLimit    uint64 `json:"commitlimit" bson:"commitlimit"`
	CommittedAS    uint64 `json:"committedas" bson:"committedas"`
	HighTotal      uint64 `json:"hightotal" bson:"hightotal"`
	HighFree       uint64 `json:"highfree" bson:"highfree"`
	LowTotal       uint64 `json:"lowtotal" bson:"lowtotal"`
	LowFree        uint64 `json:"lowfree" bson:"lowfree"`
	SwapTotal      uint64 `json:"swaptotal" bson:"swaptotal"`
	SwapFree       uint64 `json:"swapfree" bson:"swapfree"`
	Mapped         uint64 `json:"mapped" bson:"mapped"`
	VMallocTotal   uint64 `json:"vmalloctotal" bson:"vmalloctotal"`
	VMallocUsed    uint64 `json:"vmallocused" bson:"vmallocused"`
	VMallocChunk   uint64 `json:"vmallocchunk" bson:"vmallocchunk"`
	HugePagesTotal uint64 `json:"hugepagestotal" bson:"hugepagestotal"`
	HugePagesFree  uint64 `json:"hugepagesfree" bson:"hugepagesfree"`
	HugePageSize   uint64 `json:"hugepagesize" bson:"hugepagesize"`
}

type SwapMemoryStat struct {
	Total       uint64  `json:"total" bson:"total"`
	Used        uint64  `json:"used" bson:"used"`
	Free        uint64  `json:"free" bson:"free"`
	UsedPercent float64 `json:"usedPercent,omitempty" bson:"usedPercent"`
	Sin         uint64  `json:"sin" bson:"sin"`
	Sout        uint64  `json:"sout" bson:"sout"`
	PgIn        uint64  `json:"pgin" bson:"pgin"`
	PgOut       uint64  `json:"pgout" bson:"pgout"`
	PgFault     uint64  `json:"pgfault" bson:"pgfault"`
}

func (m VirtualMemoryStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func (m SwapMemoryStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
