package ctrls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

// StartTime @Tags 服务监控
var StartTime = time.Now()

func ServerInfo(c *gin.Context) {
	cpuNum := runtime.NumCPU() //核心数

	var cpuUsed float64 = 0  //用户使用率
	var cpuAvg5 float64 = 0  //CPU负载5
	var cpuAvg15 float64 = 0 //当前空闲率

	cpuInfo, err := cpu.Percent(time.Second, false)
	if err == nil {
		cpuUsed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuInfo[0]), 64)
	}

	loadInfo, err := load.Avg()
	if err == nil {
		cpuAvg5, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
		cpuAvg15, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
	}

	var memTotal uint64 = 0  //总内存
	var memUsed uint64 = 0   //总内存  := 0 //已用内存
	var memFree uint64 = 0   //剩余内存
	var memUsage float64 = 0 //使用率

	v, err := mem.VirtualMemory()
	if err == nil {
		memTotal = v.Total / 1024 / 1024
		memUsed = v.Used / 1024 / 1024
		memFree = memTotal - memUsed
		memUsage, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v.UsedPercent), 64)
	}

	var goTotal uint64 = 0  //go分配的总内存数
	var goUsed uint64 = 0   //go使用的内存数
	var goFree uint64 = 0   //go剩余的内存数
	var goUsage float64 = 0 //使用率

	var goMem runtime.MemStats
	runtime.ReadMemStats(&goMem)
	goUsed = goMem.Sys / 1024 / 1024
	goUsage = float64(goUsed) / float64(memTotal) * 100
	sysComputerIp := "" //服务器IP

	ip, err := GetLocalIP()
	if err == nil {
		sysComputerIp = ip
	}

	sysComputerName := "" //服务器名称
	sysOsName := ""       //操作系统
	sysOsArch := ""       //系统架构

	sysInfo, err := host.Info()

	if err == nil {
		sysComputerName = sysInfo.Hostname
		sysOsName = sysInfo.OS
		sysOsArch = sysInfo.KernelArch
	}

	goVersion := runtime.Version() // go环境版本

	goRunTime := time.Since(StartTime) //运行时长

	goHome := runtime.GOROOT() //安装路径
	goUserDir := ""            //项目路径

	curDir, err := os.Getwd()

	if err == nil {
		goUserDir = curDir
	}

	//服务器磁盘信息
	diskList := make([]disk.UsageStat, 0)
	diskInfo, err := disk.Partitions(true) //所有分区
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskDetail.Total = diskDetail.Total / 1024 / 1024
				diskDetail.Used = diskDetail.Used / 1024 / 1024
				diskDetail.Free = diskDetail.Free / 1024 / 1024
				diskList = append(diskList, *diskDetail)
			}
		}
	}
	res := map[string]interface{}{
		"cpuNum":          cpuNum,
		"cpuUsed":         cpuUsed,
		"cpuAvg5":         cpuAvg5,
		"cpuAvg15":        cpuAvg15,
		"memTotal":        memTotal,
		"goTotal":         goTotal,
		"memUsed":         memUsed,
		"goUsed":          goUsed,
		"memFree":         memFree,
		"goFree":          goFree,
		"memUsage":        memUsage,
		"goUsage":         goUsage,
		"sysComputerName": sysComputerName,
		"sysOsName":       sysOsName,
		"sysComputerIp":   sysComputerIp,
		"sysOsArch":       sysOsArch,
		"goVersion":       goVersion,
		"goStartTime":     StartTime,
		"goRunTime":       fmt.Sprintf("%.2f", goRunTime.Seconds()) + "s",
		"goHome":          goHome,
		"goUserDir":       goUserDir,
		"diskList":        diskList,
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

// GetLocalIP 服务端ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}
