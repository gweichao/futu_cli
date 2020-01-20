package libfunc

const (
	dataSizeDefault int = 4096

	formatYMDHMSZone = "2006-01-02 15:04:05 -0700"
	formatYMDHMS     = "2006-01-02 15:04:05"
	formatDTimeNoDot = "20060102150405"
	formatYMD        = "2006-01-02"
	formatHMS        = "15:04:05"

	SecNum24h int64 = 86400
	SecNum04h int64 = 14400
	SecNum01h int64 = 3600

	NanoTf64 float64 = 1000000000.0
	NanoTi64 int64   = 1000000000

	FileSize10M = (1 << 20) * 10
	FileSize01M = (1 << 20)

	PrintTagKey = "printtag"
)

const (
	innerKnum int64 = 1 << 10
	innerMnum int64 = 1 << 20
	innerGnum int64 = 1 << 30
	innerTnum int64 = 1 << 40

	thisfileSizeLimit int64  = innerMnum // 调试文件大小限制
	randSeedFlagLimit uint32 = 1000000
)

const (
	// 0 ready, 1 do, 2 wait to ready
	ST_READY int = iota
	ST_DO
	ST_WAIT
)
