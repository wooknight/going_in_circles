package main
//$env:GODEBUG="gctrace=1" - for powershell
//GODEBUG=gctrace=1 go run -- for mac and linux
import (
	"fmt"
	"log"
	// "log/syslog"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("---------------------------------------------------")
}
func main() {
	var mem runtime.MemStats
	// syslog1, err := syslog.Dial("", "", syslog.LOG_ALERT|syslog.LOG_MAIL, "GC_Coll")
	// if err != nil {
	// 	log.Panic("Could not connect to syslog")
	// }
	// log.SetOutput(syslog1)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	printStats(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
			log.Println("Failed to allocate 5000000 bytes of memory")
		}
	}
	printStats(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed")
			log.Println("Failed to allocate 10000000 bytes of memory")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
	log.Println("All Good men come to the aid of the party")
}
