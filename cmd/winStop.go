package main

import (
    "os/exec"
    "fmt"
    "strconv"
    "golang.org/x/sys/windows"	
)

const processEntrySize = 568

func processID(name string) (uint32, error) {
   h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
   if e != nil { return 0, e }
   p := windows.ProcessEntry32{Size: processEntrySize}
   for {
      e := windows.Process32Next(h, &p)
      if e != nil { return 0, e }
      if windows.UTF16ToString(p.ExeFile[:]) == name {
          return p.ProcessID, nil
      }
   }
   return 0, fmt.Errorf("not found")
}

func main() {

		n, e := processID("TeamViewer.exe")
		if e != nil {
		  println("Couldn't found this file.")
		}

		pID := int(n)
	
		println("Killing process.")
    kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(pID))
    err := kill.Run()
    if err != nil {
        println("Error killing process.")
        return
    }
    
    println("process was killed.")
}