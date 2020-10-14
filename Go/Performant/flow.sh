#!/bin/bash
##Busiest 10(head default) processes
ps -eo pid,ppid,cmd,%mem,%cpu --sort=-%mem | head

sysdig -p"*%evt.time %proc.pid %proc.ppid %evt.dir %proc.exeline( evt.dir=< and evt.type=execve ) or evt.type=procexit"
##consoldited time
top -Sd1