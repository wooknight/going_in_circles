package main

import (
	"fmt"
	"os"
)

func main() {
	//list all processes

	//list all interesting attributes
	{
		envs := os.Environ()
		for _, val := range envs {
			fmt.Printf("Before %v\n", val)
		}
	}
	fmt.Println(os.ExpandEnv("$CAS_USERNAME lives in ${HOME}."))
	os.Clearenv()
	{
		envs := os.Environ()
		for _, val := range envs {
			fmt.Printf("After %v\n", val)
		}
	}
	nam, _ := os.Executable()
	fmt.Println("Process Name %v", nam)
	fmt.Println(os.ExpandEnv("$CAS_USERNAME lives in ${HOME}."))
	{
		grps, _ := os.Getgroups()
		for _, val := range grps {
			fmt.Printf("in Group %v\n", val)
		}
	}
	fmt.Println("Page size - ", os.Getpagesize())
	fmt.Println("Process Id - ", os.Getpid())
	fmt.Println("Process Parent ID- ", os.Getppid())
	fmt.Println("my ID- ", os.Getuid())
	hstname, _ := os.Hostname()
	fmt.Println("my hostname- ", hstname)
	myfil, err := os.Open("/Users/ramesh/go/src/going_in_circles/go/thread")
	if err != nil {
		fmt.Println(err)
	}
	defer myfil.Close()
	// 	{
	// 	fileinfo,err := myfil.Readdir(0)
	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// 	for i,val := range fileinfo{
	// 		fmt.Printf("%d File info - %v\n\n",i,val)
	// 	}
	// }
	{
		fileinfo, err := myfil.Readdirnames(0)
		if err != nil {
			fmt.Println(err)
		}
		for i, val := range fileinfo {
			fmt.Printf("%d File names - %v\n\n", i, val)
		}
	}
}
