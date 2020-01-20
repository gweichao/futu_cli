package libfunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"simplejson"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	randSeedFlag   uint32 = randSeedFlagLimit
	randSeedFlagts int64  = 0
	goOS           string = ""
)

const (
	CharTypeUnknown int = iota
	CharTypeUseUft8
	CharTypeUseGbk
)

func StringToInt1(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 || str == "" {
		return 0
	}

	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		prnLog.LogPrint(LOG_ERROR, 1, false, false, "", err)
		return -1 // ErrUndefined
	}
	return int(number)
}

// 一层 json结构
func FieldReset(body []byte, field string, target string) (result []byte) {

	js, err := simplejson.NewJson(body)
	if err != nil {
		return
	}

	_, ok := js.CheckGet(field)
	if !ok {
		return
	}
	js.Set(field, target)
	result, _ = js.MarshalJSON()
	return

}

func StringToInt64(str string) int64 {
	str = strings.TrimSpace(str)
	if len(str) == 0 || str == "" {
		return 0
	}

	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		prnLog.LogPrint(LOG_ERROR, 1, false, false, "", err)
		return -1
	}
	return number
}

// fn 1-5：指定位数；6及以上,同6位; <=0 没有秒后的小数点。
func NanoToTimeStr(deltaNano int64, fn int) string {
	//deltaNano = time.Now().UnixNano() - tsnano1
	var (
		secNum24h int64 = 86400
		secTotal  int64 = deltaNano / 1e9
		days      int64 = secTotal / secNum24h
		sec1Day   int64 = secTotal % secNum24h
		usec      int64 = (deltaNano % 1e9) / 1000

		h, m, s int64 = sec1Day / 3600, (sec1Day % 3600) / 60, sec1Day % 60
	)

	str1 := ""
	strD := ""

	if days > 0 && sec1Day+usec == 0 {
		return fmt.Sprintf("%dd", days)
	}

	if days > 0 {
		strD = fmt.Sprintf("%dd ", days)
	}

	if h > 0 {
		str1 += fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	} else if m > 0 {
		str1 += fmt.Sprintf("%02d:%02d", m, s)
	} else {
		str1 += fmt.Sprintf("%d", s)
	}
	if fn > 6 {
		fn = 6
	}
	if fn > 0 {
		fmts := ".%0" + IntToString(fn) + "d"

		// prnLog.Debugf("str1=%v, fmts=%v, usec=%v, fn=%v ", str1, fmts, usec, fn)

		str1 += fmt.Sprintf(fmts, usec/int64(math.Pow10((6-fn))))
	}

	if str1 == "0" {
		if len(strD) > 0 {
			return strD
		} else {
			return fmt.Sprintf("0.%01d", usec/int64(math.Pow10((6-1)))) // "0.0"
		}
	}

	str1 = strD + str1

	return str1
}

func IntToString(i int) string {
	return Int64ToString(int64(i))
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

//允许跨域访问 调试使用 不含 content-type
func SetHeader(w http.ResponseWriter) {

	w.Header().Add("Access-Control-Allow-Origin", "*")  //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "*") //header的类型
	w.Header().Add("Access-Control-Allow-Headers", "accept")
	w.Header().Add("Access-Control-Allow-Headers", `filename, Content-Type, Content-Range, Content-Disposition`)
	w.Header().Add("Access-Control-Allow-Headers", `Content-Description,x-requested-with`)

	w.Header().Add("Access-Control-Allow-Methods", "POST")    //header的类型
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS") //header的类型

}

var pfDirList []string = []string{
	"./",
	"./runtime/",
	"../runtime/",
	"../../runtime/",
	"./config/",
	"../config/",
	"../../config/",
	"./output/",
	"../output/",
	"../../output/",
	"./key/",
	"../key/",
	"../../key/",
	"./sql/",
	"../sql/",
	"../../sql/",
}

// 优先按 fnamelist 顺序查找，如果找不到则取 __file_list.txt 里面的文件名称，每行一个文件名
func FindPathFileInNameList(fnamelist []string) (pathfile1 string) {
	filename1 := ""
	for _, fn1 := range fnamelist {
		filename1 = FindPathFile(fn1)
		if FileExist(filename1) {
			pathfile1 = filename1
			return
		}
	}
	filename1 = FindPathFile("__file_list.txt")
	if !FileExist(filename1) {
		return
	}
	//prnLog.Debugf("filename1==%v==", filename1)
	fnamelist = strings.Split(GetStringFromFile(filename1), "\n")
	for _, fn1 := range fnamelist {
		fn1 = strings.TrimSpace(fn1)
		if len(fn1) == 0 {
			continue
		}
		if strings.HasPrefix(fn1, "#") {
			continue
		}
		//prnLog.Debugf("==%v==", fn1)
		filename1 = FindPathFile(fn1)
		if FileExist(filename1) {
			pathfile1 = filename1
			return
		}
	}

	return
}
func FindPathFile(filename string) (pathfile string) {
	/* 顺序
	1. 当前文件
	2. 如果空则取缺省文件
	3. 依次加 ./  ./config/ ../config/ ../../config/ ./output/ ../output/ ../../output/ ./key/ ../key/ ../../key/
	   存在则返回 路径+文件名 */
	filename = strings.TrimSpace(filename)
	pathfile = filename

	if len(filename) == 0 {
		return
	}
	if FileExist(filename) {
		pathfile, _ = filepath.Abs(filename)
		return
	}

	for _, path1 := range pfDirList {
		pfile1 := path1 + filename
		if FileExist(pfile1) {
			pathfile, _ = filepath.Abs(pfile1)
			return
		}
	}

	return
}
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
func StructToJsonString(v interface{}) (resp string, ret int) {

	resp = ""
	ret = thisSuccess

	if v == nil {
		return
	}

	if value1, ok1 := v.(string); ok1 {
		resp = value1
		return
	} else if value2, ok2 := v.([]byte); ok2 {
		resp = string(value2)
		return
	}

	jsonBody, err := json.Marshal(v)
	if err != nil {
		ret = ErrUndefined
		prnLog.LogPrint(LOG_ERROR, 1, false, false, "", "json Marshal error:", err)
		return
	}

	resp = string(jsonBody)

	return
}
func StructToJsonStringOne(v interface{}) string {
	resp, _ := StructToJsonString(v)
	return resp

}

//return rand number in [0,max)
func GetRand(max int) int {
	if randSeedFlag >= randSeedFlagLimit && time.Now().Unix()-randSeedFlagts > 3600 {
		//fmt.Println("=================rand.Seed")
		randSeedFlag = 0
		randSeedFlagts = time.Now().Unix()
		rand.Seed(time.Now().UnixNano())
	}
	randSeedFlag++
	result := rand.Intn(max)
	return result
}

// 生成随机字符串, lenth 返回字符串长度； bbase 表示是最大顺序号；
/*10:	纯数字；					16: 十六进制；
  36:	所有数字、小写字符；		62: 所有数字、大小写字符；
  82:	所有数字、大小写、符号字符；	63: 所有数字、大写字符；
  64:	大写字符；				65:	小写字符；
  66:	大小写字符；				16: 如果取值为其他，则使用默认值16 */
func GetRandStr(lenth int, bbase int) string {
	var result string

	switch bbase {
	case 10, 16, 36, 62, 82:
		for i := 0; i < lenth; i++ {
			result += nummap[GetRand(bbase)]
		}
	case 63: // 0..9, 36..61
		bbase = 36
		for i := 0; i < lenth; i++ {
			seq := GetRand(bbase)
			if 10 <= seq && seq <= 35 {
				seq += 26
			}
			result += nummap[seq]
		}
	case 64: // 36..61
		bbase = 26
		for i := 0; i < lenth; i++ {
			seq := GetRand(bbase)
			seq += 36
			result += nummap[seq]
		}
	case 65: // 10..35
		bbase = 26
		for i := 0; i < lenth; i++ {
			seq := GetRand(bbase)
			seq += 10
			result += nummap[seq]
		}
	case 66: // 10..61
		bbase = 52
		for i := 0; i < lenth; i++ {
			seq := GetRand(bbase)
			seq += 10
			result += nummap[seq]
		}
	default:
		bbase = 16
		for i := 0; i < lenth; i++ {
			result += nummap[GetRand(bbase)]
		}
	}

	return result
}

var nummap = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z",
	62: ".", // `
	63: "~",
	64: "!",
	65: "@",
	66: "#",
	67: "$",
	68: "%",
	69: "^",
	70: "&",
	71: "*",
	72: "(",
	73: ")",
	74: "_",
	75: "+",
	76: "=",
	77: "-",
	78: "/",
	79: ";",
	80: "*", // :
	81: ",", // '
}

func GetStringFromFile(pathfile string) string {
	allFilebytes, err := ioutil.ReadFile(pathfile)
	if err != nil {
		prnLog.Errorf("", "open file error:", err)
		return ""
	}
	return string(allFilebytes)
}

// 获取http body数据
func GetPostData(req *http.Request) (body []byte, ret int) {
	if req == nil {
		return
	}
	var err error
	if body, err = ioutil.ReadAll(req.Body); err != nil {
		prnLog.Errorf("read body error=%v", err)
		ret = ErrUndefined
		return
	}
	ret = Success
	return
}
func HttpBodyGet(req *http.Request) (body []byte, ret int) {
	if req == nil {
		return
	}
	var err error
	if body, err = ioutil.ReadAll(req.Body); err != nil {
		prnLog.Errorf("read body error=%v", err)
		ret = ErrUndefined
		return
	}
	ret = Success
	return
}

func ParseJsonByStructMsg(body []byte, v interface{}) (msg string, ret int) {
	if v == nil {
		return
	}
	// 2016-10-12
	ret = thisSuccess
	msg = ""

	if len(body) == 0 {
		body = []byte("{}")
	}

	if err := json.Unmarshal(body, v); err != nil {
		msg = err.Error()
		ret = ErrData
		return
	}
	return
}

func GetRemoteIP2(req *http.Request, nginxIp string) string {

	if req == nil {
		return ""
	}

	// req.RemoteAddr 为协议自带，不可篡改
	ip1 := (AddrT(req.RemoteAddr)).RemoteIP() // 类似 RemoteAddr=10.10.30.102:37336

	return ip1
}

type AddrT string

func (ad AddrT) RemoteIP() string {
	s := strings.Split(string(ad), ":")
	if len(s) != 0 {
		return s[0]
	}
	return ""
}
func CreateAndPutTofile(pfile string, printstr string) error {
	fo1, err1 := os.Create(pfile)
	if err1 != nil {
		//prnLog.Errorf("create file error:%v\n", err1)
		return err1
	}

	defer func() {
		//打开文件出错处理
		if fo1 != nil {
			fo1.Close()
		}
	}()

	fo1.WriteString(printstr)

	fo1.Sync()
	return nil
}

/* flag: <=0 覆盖，>=1 附加,代表保存历史文件的数量,最多50;
fileSizeLimit 字节数，追加时如果已有文件大小超该值，则改为覆盖。0表示默认1M，-1 不限制
对于附加方式，当第一个文件到达上限时，将其改名 pfile.n 再重建，保证内容连续性，其中n=flag代表历史文件数
比如flag=3，会依次保存 pfile pfile.1 pfile.2 pfile.3
*/
func PutToFile(pfile string, printstr string, flag int, fileSizeLimit int64) error {

	var (
		debugflag = false
	)

	if debugflag {
		prnLog.Debugf("pfile=%v, flag=%v, Limit=%v",
			pfile, flag, fileSizeLimit)
	}

	if flag <= 0 {
		err1 := CreateAndPutTofile(pfile, printstr)
		return err1
	}
	if flag > 50 {
		flag = 50
	}

	if fileSizeLimit == 0 {
		fileSizeLimit = thisfileSizeLimit
	}

	// 以只写的模式，打开文件
	fo2, err2 := os.OpenFile(pfile, os.O_WRONLY, 0644)
	//
	if err2 != nil &&
		(strings.Contains(err2.Error(), "cannot find the file specified") ||
			strings.Contains(err2.Error(), "no such file")) {
		err1 := CreateAndPutTofile(pfile, printstr)
		return err1
	}
	defer func() {
		//打开文件出错处理
		if fo2 != nil {
			fo2.Close()
		}
	}()

	if err2 != nil {
		//prnLog.LogPrint(LOG_ERROR, 1, false, false, " file create failed. err: %v\n", err2.Error())
		return err2
	} else {
		// 查找文件末尾的偏移量
		n, _ := fo2.Seek(0, os.SEEK_END)

		if fileSizeLimit > 0 && n > fileSizeLimit {

			if debugflag {
				prnLog.Debugf("file size=%v ", n)
			}

			fo2.Close()
			fo2 = nil

			os.Remove(pfile + "." + IntToString(flag))
			for seq := flag; seq >= 1; seq-- {

				left := seq - 1
				leftPfile := pfile + "." + IntToString(left)
				if left == 0 {
					leftPfile = pfile
				}

				right := seq
				rightPfile := pfile + "." + IntToString(right)

				if debugflag {
					prnLog.Debugf("rename, %v,%v ", leftPfile, rightPfile)
				}

				os.Rename(leftPfile, rightPfile)
			}

			err1 := CreateAndPutTofile(pfile, printstr)
			return err1
		}
		// 从末尾的偏移量开始写入内容
		_, err2 = fo2.WriteAt([]byte(printstr), n)
	}
	return err2

}

// 判断obj是否在tarGet中，tarGet支持的类型arrary,slice,map
func Contain(obj interface{}, tarGet interface{}) bool {
	if obj == nil || tarGet == nil {
		return false
	}
	tarGetValue := reflect.ValueOf(tarGet)
	// ? tarGet.(type)

	switch reflect.TypeOf(tarGet).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < tarGetValue.Len(); i++ {
			if tarGetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if tarGetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

// 可用 strings.TrimSpace(str)，不要使用 strings.Trim(str, " ")
func StringTrim3(str string) string {
	var s1, e2 int = 0, len(str) - 1
	if e2 < 0 {
		return str

	}
	for ; e2 >= 0; e2-- {
		if str[e2] == '\n' || str[e2] == '\r' || str[e2] == '\t' || str[e2] == ' ' {
			continue
		}
		break
	}
	for ; s1 <= e2; s1++ {
		if str[s1] == '\n' || str[s1] == '\r' || str[s1] == '\t' || str[s1] == ' ' {
			continue
		}
		break
	}

	return str[s1 : e2+1]
}

func CharUtf8ToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}
	return string(dst[:nDst]), nil
}
func CharGbkToUTF8(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}
	return string(dst[:nDst]), nil
}
func CharBytesToGBK(bytesIn []byte) (string, error) {
	dst := make([]byte, len(bytesIn)*2)
	tr := simplifiedchinese.GB18030.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(bytesIn), true)
	if err != nil {
		return string(bytesIn), err
	}
	return string(dst[:nDst]), nil
}
func CharBytesToUTF8(bytesIn []byte) (string, error) {
	dst := make([]byte, len(bytesIn)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(bytesIn), true)
	if err != nil {
		return string(bytesIn), err

	}
	return string(dst[:nDst]), nil
}
