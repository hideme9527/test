package libtest

import (
	"github.com/peanut996/CloudflareWarpSpeedTest/task"
)

//export Test
func Test() string {
	task.InitRandSeed()
	task.InitHandshakePacket()

	pingData := task.NewWarping().Run().FilterDelay().FilterLossRate()
	// utils.ExportCsv(pingData)
	return pingData.Print()
}

//https://github.com/peanut996/CloudflareWarpSpeedTest
