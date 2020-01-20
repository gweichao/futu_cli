package libfunc

/*******************************************************************************
  功能描述:

  更改历史
  日期      		作者                    详情


*******************************************************************************/
// 加到 $GOPATH/src/devlog
import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	log "devlog"
)

//type LogLevel int

const (
	// !nashtsai! following level also match syslog.Priority value
	LOG_DEBUG int = iota
	LOG_INFO
	LOG_WARNING
	LOG_ERROR
	LOG_OFF
	LOG_UNKNOWN
)

/*
=== Debug ===
        这个级别最低。一般的来说，在系统实际运行过程中，一般都是不输出的。
        因此这个级别的信息，可以随意的使用，任何觉得有利于在调试时更详细的了解系统运行状态的东东，比如变量的值等等，都输出来看看也无妨。

=== Info ===
        这个应该用来反馈系统的当前状态给最终用户的，所以，在这里输出的信息，应该对最终用户具有实际意义，也就是最终用户要能够看得明白是什么意思才行。
        从某种角度上说，Info 输出的信息可以看作是软件产品的一部分（就像那些交互界面上的文字一样），所以需要谨慎对待，不可随便。

=== Warn、Error、Fatal ===
        警告、错误、严重错误，这三者应该都在系统运行时检测到了一个不正常的状态，大致区分：

        所谓警告，应该是这个时候进行一些修复性的工作，应该还可以把系统恢复到正常状态中来，系统应该可以继续运行下去。

        所谓错误，就是说可以进行一些修复性的工作，但无法确定系统会正常的工作下去，系统在以后的某个阶段，很可能会因为当前的这个问题，导致一个无法修复的错误（例如宕机），但也可能一直工作到停止也不出现严重问题。

        所谓Fatal，那就是相当严重的了，可以肯定这种错误已经无法修复，并且如果系统继续运行下去的话，可以肯定必然会越来越乱。这时候采取的最好的措施不是试图将系统状态恢复到正常，而是尽可能地保留系统有效数据并停止运行。
*/

var (
	localPrnLog  PrnLogStu
	localSysflag            = "libfv2"
	prnLog       *PrnLogStu = GetPrnLogInstance()
)

const (
	DEBUGName   = "DEBUG"
	INFOName    = "INFO"
	WARNINGName = "WARNING"
	ERRORName   = "ERROR"
)

var logPrintNameMAP map[int]string = map[int]string{
	LOG_DEBUG:   "D", // DEBUGName,
	LOG_INFO:    "I", // INFOName,
	LOG_WARNING: "W", // WARNINGName,
	LOG_ERROR:   "E", // ERRORName,
	LOG_OFF:     "OFF",
	LOG_UNKNOWN: "U", // "UNKNOWN",
}

const (
	formatLeader    string = "→→→"
	formatTag       string = "" //"tag="
	formatLongTime  string = "+"
	formatLeaderLen int    = len(formatLeader)

	DefaultLenPerLine int = 20 * 1024
)

type PrnLogStu struct {
	sysflag string

	logLevel   int
	dbLogLevel int

	useFulFmt      bool
	saveToFileOnly bool
	logPATHFile    string
	logTimes       uint32

	printtag   string
	tsNanobase int64

	LimitLenPerLine int
}

func GetPrnLogInstance() (inst *PrnLogStu) {
	inst = &localPrnLog
	return
}

func PrnLogGet(sysflagIn string) (inst *PrnLogStu) {
	inst = &localPrnLog
	if len(sysflagIn) > 0 {
		inst.sysflag = sysflagIn
	}
	return
}

func PrnLogNewGet(sysflagIn string) (inst *PrnLogStu) {
	inst = &PrnLogStu{}
	inst.Init(localSysflag, LOG_DEBUG, 0)
	if len(sysflagIn) > 0 {
		inst.sysflag = sysflagIn
	}
	return
}

func (this *PrnLogStu) SetLogLevel(logLevel, DbLogLevel string) {

	if len(DbLogLevel) == 0 {
		DbLogLevel = logLevel
	}

	logLevel = strings.TrimSpace(strings.ToUpper(logLevel))
	DbLogLevel = strings.TrimSpace(strings.ToUpper(DbLogLevel))

	switch logLevel {
	case DEBUGName:
		this.logLevel = LOG_DEBUG
	case INFOName:
		this.logLevel = LOG_INFO
	case WARNINGName:
		this.logLevel = LOG_WARNING
	case ERRORName:
		this.logLevel = LOG_ERROR
	}

	switch DbLogLevel {
	case DEBUGName:
		this.dbLogLevel = LOG_DEBUG
	case INFOName:
		this.dbLogLevel = LOG_INFO
	case WARNINGName:
		this.dbLogLevel = LOG_WARNING
	case ERRORName:
		this.dbLogLevel = LOG_ERROR
	}
}

/**
直接输出到log的公共buffer，不做位置补齐 */
func (this *PrnLogStu) LogDirect(s string) {
	if len(s) > 0 {
		log.PrintDirect(s)
	}
}

/**
一般打印 e.g.
	 [25 06:58:25]|ERROR|svcproxy|entrystruct.go|Init()|84|url= /main,/smartcom/hotel/info/query
	 日期 时间	   等级	程序名		文件名		 函数	行号 内容
参数
	logtype 等级，libf.	LOG_DEBUG(0)，LOG_INFO(1),LOG_WARNING,LOG_ERROR
	DeltaLevel 选择显示位置，0当前位置，1 父级，2 祖，...
	prnToStrOnly 是否仅输出到返回字符串，一般用于统一控制打印，由 LogDirect 直接输出
	doShort 短格式
	format 格式字符串，如果空串则有特定意义、代表无格式，类似 fmt.Println()

	如果有内置的tag，自动调用LogPrintTag;
	用法1, 用公共的a1.PrnLog：
		a1.PrnLog.LogPrint(libf.LOG_ERROR, 0, false, false, "InitByJson error, data=%v", string(data))
	用法2，拷贝一个实例、设置tag：
		this.PrnLog = *a1.PrnLog
		this.PrnLog.TagTsSet(this.PrintTag, this.StartTsNano)
		// 自动加tag打印
		this.PrnLog.LogPrint(libf.LOG_ERROR, prnLevel, false, false, " strconv.Atoi error %v", err)
*/
func (this *PrnLogStu) LogPrint(logtype int, DeltaLevel int, prnToStrOnly, doShort bool, format string, v ...interface{}) (s_out string) {

	if logtype < this.logLevel {
		return
	}

	// 用内置tag，如果为空则自动用非tag函数
	// 需要clone一个类实例然后 TagTsSet()
	if len(this.printtag) > 0 {
		s_out = this.logPrintTag(logtype, DeltaLevel+1, prnToStrOnly, doShort,
			this.printtag, this.tsNanobase, format, v...)
		return
	}

	s_out = this.outputf(logtype, DeltaLevel, prnToStrOnly, doShort, format, format, v...)

	return
}

func (this *PrnLogStu) Errorf(fmt string, args ...interface{}) {
	this.LogPrint(LOG_ERROR, 1, false, false, fmt, args...)
	return
}
func (this *PrnLogStu) Infof(fmt string, args ...interface{}) {
	this.LogPrint(LOG_INFO, 1, false, false, fmt, args...)
	return
}
func (this *PrnLogStu) Debugf(fmt string, args ...interface{}) {
	this.LogPrint(LOG_DEBUG, 1, false, false, fmt, args...)
	return
}
func (this *PrnLogStu) Warningf(fmt string, args ...interface{}) {
	this.LogPrint(LOG_WARNING, 1, false, false, fmt, args...)
	return
}

/* LevelNum 接口沿路的层级数量，0仅本级，1 父级+本级，2 祖+父+本
e.g. libf_test.go|P3.342,P2.338,P1.334| this is test P1 */
func (this *PrnLogStu) CErrorf(levelNum int, fmt string, args ...interface{}) {
	this.outputItf(LOG_ERROR, levelNum, 2, false, false, fmt, args...)
	return
}
func (this *PrnLogStu) CInfof(levelNum int, fmt string, args ...interface{}) {
	this.outputItf(LOG_INFO, levelNum, 2, false, false, fmt, args...)
	return
}
func (this *PrnLogStu) CDebugf(levelNum int, fmt string, args ...interface{}) {
	this.outputItf(LOG_DEBUG, levelNum, 2, false, false, fmt, args...)
	return
}
func (this *PrnLogStu) CWarningf(levelNum int, fmt string, args ...interface{}) {
	this.outputItf(LOG_WARNING, levelNum, 2, false, false, fmt, args...)
	return
}

/**
一般打印 e.g.
	 [25 08:32:21]|DEBUG|svcproxy|hotelwxmplogin.go|A3Prepare()|170|tag=RpFCKE,0.596|from=ali,uid=2088502911522641
	 日期 时间	   等级	程序名		文件名			函数			行号 标签       相对起点的时间		内容
参数
	logtype 等级，libf.	LOG_DEBUG(0)，LOG_INFO(1),LOG_WARNING,LOG_ERROR
	DeltaLevel 选择显示位置，0当前位置，1 父级，2 祖，...
	prnToStrOnly 是否仅输出到返回字符串，一般用于统一控制打印，由 LogDirect 直接输出
	doShort 短格式
	printtag 标签
	tsNanobase 起点系统时间的纳秒值
	format 格式字符串，如果空串则有特定意义、代表无格式，类似 fmt.Println()
*/
func (this *PrnLogStu) logPrintTag(logtype int, DeltaLevel int, prnToStrOnly, doShort bool,
	printtag string, tsNanobase int64, format string, v ...interface{}) (s_out string) {

	if logtype < this.logLevel {
		return
	}

	format0 := format
	// 构造 tag --> format
	var dur int64
	str1 := ""
	if tsNanobase > 0 {
		tag := formatTag
		dur = time.Now().UnixNano() - tsNanobase
		secNum := int(dur / 1e9)
		if secNum >= 3 {
			secNum = 3
		}
		if secNum > 0 {
			tag += strings.Repeat(formatLongTime, secNum)
		}

		str1 += tag + printtag + "," + NanoToTimeStr(dur, 3) + "| "
	} else {
		tag := formatTag
		str1 += tag + printtag + ",-." + "| "
	}

	format = str1 + format

	s_out = this.outputf(logtype, DeltaLevel, prnToStrOnly, doShort, format0, format, v...)

	return
}

// format0 用于判断是否空
func (this *PrnLogStu) outputf(logtype int, DeltaLevel int,
	prnToStrOnly, doShort bool, format0, format string, v ...interface{}) string {

	if this.LimitLenPerLine == 0 {
		this.LimitLenPerLine = DefaultLenPerLine
	}

	var s_out string

	if logtype < this.logLevel {
		return ""
	}

	isLogPrintTag := format0 != format
	spaceStr := " "
	if isLogPrintTag {
		spaceStr = ""
	}
	rtnStr := "\n"
	if prnToStrOnly {
		rtnStr = ""
	}

	logTypeName, ok := logPrintNameMAP[logtype]
	if !ok {
		this.LogDirect("no this log type=" + IntToString(logtype) + "\n")
		return ""
	}

	DeltaLevel++
	fileName, funcName, line := callerInfoGet(1 + DeltaLevel)

	if !doShort {

		if format0 == "" {
			s_out = log.Sprintf("|%v|%v|%v|%v()|%v|%v%v",
				logTypeName, this.sysflag, fileName, funcName, line, spaceStr,
				strings.TrimLeft(format+fmt.Sprintln(v...), " "))
		} else {
			s_out = log.Sprintf("|%v|%v|%v|%v()|%v|%v%v",
				logTypeName, this.sysflag, fileName, funcName, line, spaceStr,
				strings.TrimLeft(fmt.Sprintf(format, v...)+rtnStr, " "))
		}

	} else {
		if format0 == "" {
			s_out = log.Sprintf("|%v|%v|%v|%v%v",
				logTypeName, fileName, line, spaceStr,
				strings.TrimLeft(format+fmt.Sprintln(v...), " "))
		} else {
			s_out = log.Sprintf("|%v|%v|%v|%v%v",
				logTypeName, fileName, line, spaceStr,
				strings.TrimLeft(fmt.Sprintf(format, v...)+rtnStr, " "))
		}
	}

	this.incLogTimes()

	if this.LimitLenPerLine > 0 && len(s_out) > this.LimitLenPerLine {
		s_out = s_out[:this.LimitLenPerLine] + "(trim...)"
	}

	if !prnToStrOnly {
		log.PrintDirect(s_out)
		return ""
	}

	return s_out
}

/**
一般打印 e.g.
	 [25 07:03:52]|DEBUG|svcproxy|server.go|HandlerSmartcom.155,Execute.307,hotelwxmplogin.go:B2_Cache.246()|sessioncache01={"id":"910b0f5bb89a
	 日期 时间	   等级	程序名		文件名	函数调用顺序、及行号 												 内容
参数
	logtype 等级，libf.	LOG_DEBUG(0)，LOG_INFO(1),LOG_WARNING,LOG_ERROR
	DeltaLevel 选择显示位置，0当前位置，1 父级，2 祖，...
	LevelNum 接口沿路的层级数量，0仅本级，1 父级+本级，2 祖+父+本
	prnToStrOnly 是否仅输出到返回字符串，一般用于统一控制打印，由 LogDirect 直接输出
	doShort 短格式
	format 格式字符串，如果空串则有特定意义、代表无格式，类似 fmt.Println()
*/
func (this *PrnLogStu) LogPrintItf(logtype int, LevelNum int, prnToStrOnly, doShort bool, format string, v ...interface{}) string {
	return this.outputItf(logtype, LevelNum, 2, prnToStrOnly, doShort, format, v...)
}
func (this *PrnLogStu) outputItf(logtype int, LevelNum, DeltaLevel int,
	prnToStrOnly, doShort bool, format string, v ...interface{}) string {

	if this.LimitLenPerLine == 0 {
		this.LimitLenPerLine = DefaultLenPerLine
	}

	var s_out string

	if logtype < this.logLevel {
		return ""
	}

	logTypeName, ok := logPrintNameMAP[logtype]
	if !ok {
		this.LogDirect("no this log type=" + IntToString(logtype) + "\n")
		return ""
	}

	fileName, funcName, line := callerInfoGet(LevelNum + DeltaLevel) // callerInfoGet(DeltaLevel + 2)
	funcName += "." + IntToString(line)
	for i2 := LevelNum - 1; i2 >= 0; i2-- {
		fileName2, funcName2, line2 := callerInfoGet(i2 + DeltaLevel)

		str1 := ","
		if i2 == 1 && fileName2 != fileName {
			str1 += fileName2 + ":"
		}
		funcName += str1 + funcName2 + "." + IntToString(line2)
	}

	if !doShort {

		if format == "" {
			s_out = log.Sprintf("|%v|%v|%v| %v", logTypeName, fileName, funcName,
				strings.TrimLeft(fmt.Sprintln(v...), " "))
		} else {
			s_out = log.Sprintf("|%v|%v|%v| %v", logTypeName, fileName, funcName,
				strings.TrimLeft(fmt.Sprintf(format, v...)+"\n", " "))
		}

	} else {
		if format == "" {
			s_out = log.Sprintf("|%v|%v|%v| %v", logTypeName, fileName, line,
				strings.TrimLeft(fmt.Sprintln(v...), " "))
		} else {
			s_out = log.Sprintf("|%v|%v|%v| %v", logTypeName, fileName, line,
				strings.TrimLeft(fmt.Sprintf(format, v...)+"\n", " "))
		}
	}

	this.incLogTimes()

	if this.LimitLenPerLine > 0 && len(s_out) > this.LimitLenPerLine {
		s_out = s_out[:this.LimitLenPerLine] + "(trim...)"
	}

	if !prnToStrOnly {
		log.PrintDirect(s_out)
		return ""
	}

	return s_out
}

func callerInfoGet(level int) (fileName, funcName string, line int) {

	pc, pathfile, lineInt, _ := runtime.Caller(1 + level)

	fileName = func(pathfileIn string) string {
		fileNameOut := pathfileIn
		for i := len(pathfileIn) - 1; i > 0; i-- {
			if pathfileIn[i] == '/' {
				fileNameOut = pathfileIn[i+1:]
				break
			}
		}
		return fileNameOut
	}(pathfile)

	// 行号
	line = lineInt

	funcNameList := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	func(funcNameListIn []string) {
		len1 := len(funcNameListIn)
		if len1 == 0 {
			return
		}
		funcName = funcNameListIn[len1-1]

		for s := len1 - 1; s >= 0; s-- {
			if isCheckNum09(funcNameListIn[s]) {
				continue
			}
			if strings.HasPrefix(funcNameListIn[s], "func") {
				s1 := strings.Replace(funcNameListIn[s], "func", "", 1)
				if isCheckNum09(s1) {
					continue
				}
			}
			funcName = funcNameListIn[s]
			break
		}
	}(funcNameList)

	return
}

//数字类型字符串：0-9
var checkNum09Map = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func isCheckNum09(strIn string) (retb bool) {
	retb = true
	lenStr := len(strIn)
	for i := 0; i < lenStr; i++ {
		r := strIn[i]
		if _, ok := checkNum09Map[r]; !ok {
			retb = false
			return
		}
	}
	return
}

func (this *PrnLogStu) logger() *os.File {
	if !this.saveToFileOnly {
		return nil
	}
	log.Printf(" 仅打印到文件=%v, 适时刷新到文件、非实时。\n\n", this.logPATHFile)

	f, err := os.OpenFile(this.logPATHFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)

	return f
}

func (this *PrnLogStu) SysflagGet() (sysflag string) {
	sysflag = this.sysflag
	return
}

func (this *PrnLogStu) SysflagSet(SysflagIn string) {
	this.sysflag = SysflagIn
}

func (this *PrnLogStu) TagTsSet(printtagIn string, tsNanobaseIn int64) {
	this.printtag = printtagIn
	this.tsNanobase = tsNanobaseIn
}
func (this *PrnLogStu) TagTsGet() (printtag string, tsNanobase int64) {
	printtag = this.printtag
	tsNanobase = this.tsNanobase
	return
}
func (this *PrnLogStu) TagGet() (printtag string) {
	printtag = this.printtag
	return
}

func (this *PrnLogStu) Init(sysflag string, loglevel, limitLenPerLine int) {

	if len(sysflag) == 0 {
		sysflag = localSysflag
	}
	if limitLenPerLine <= 0 {
		limitLenPerLine = DefaultLenPerLine
	}

	this.sysflag = sysflag
	this.logLevel = loglevel
	this.dbLogLevel = loglevel
	this.useFulFmt = true
	this.saveToFileOnly = false
	this.logPATHFile = ""
	this.logTimes = 0

	this.LimitLenPerLine = limitLenPerLine

	// log.Ldd Ldate Lmmdd Ldd
	log.SetFlags(log.Lmmdd | log.Ltime /* | log.Lmicroseconds*/)

	this.logger()

}

func init() {
	localPrnLog.Init("", LOG_DEBUG, 0)
}

func (this *PrnLogStu) incLogTimes() {
}

// 在 header 检查是否存在 SysflagCheck，不存在则添加 SysflagIn 。SysflagCheck 空则取 SysflagIn
// SysflagIn 必选，如可以这样用 PrintTagSetToHeader("[sldjf]","sldjf",header)
func PrintTagSetToHeader(SysflagIn, SysflagCheck string, header http.Header) {

	if len(SysflagIn) == 0 {
		return
	}

	if len(SysflagCheck) == 0 {
		SysflagCheck = SysflagIn
	}
	tags := header.Get(PrintTagKey)
	if strings.Contains(tags, SysflagCheck) {
		return
	}
	header.Set(PrintTagKey, SysflagIn)

	return
}
