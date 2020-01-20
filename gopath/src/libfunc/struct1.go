package libfunc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type CheckTsStu struct {
	tsMap map[string]time.Time

	init bool

	sync.RWMutex
}

func (this *CheckTsStu) Set(token string, ts1 time.Time) {
	if !this.init {
		this.tsMap = make(map[string]time.Time, 0)
		this.init = true
	}

	this.Lock()
	this.tsMap[token] = ts1
	this.Unlock()

	return
}

// true print, false not print
func (this *CheckTsStu) Check(token string, tsec int) (retb bool) {
	if !this.init {
		this.tsMap = make(map[string]time.Time, 0)
		this.init = true
	}

	this.RLock()
	ts1, ok := this.tsMap[token]
	this.RUnlock()

	if !ok {
		retb = true

		this.Lock()
		this.tsMap[token] = time.Now().Add(time.Second * time.Duration(tsec))
		this.Unlock()

		return
	}

	if time.Now().UnixNano() >= ts1.UnixNano() {
		retb = true

		this.Lock()
		this.tsMap[token] = time.Now().Add(time.Second * time.Duration(tsec))
		this.Unlock()

		return
	}

	return
}

//loopback结构体
type LoopbackReqStu struct {
	Id      string        `json:"id,omitempty"`      //随机6位字符串
	LastRtt string        `json:"lastrtt,omitempty"` //上次往返时间
	NextHop NextHopReqStu `json:"nexthop,omitempty"` //
}

// header 目前不用
func (this *LoopbackReqStu) set(body []byte, header http.Header) (addmsg string, ret int) {
	retp := ParseJsonByStruct(body, this, 0)
	if retp != Success {
		ret = retp
		addmsg = "[Loopback]" + " 提取body入参失败 "
		return
	}

	return
}

// header 目前不用
func (this *LoopbackReqStu) setHttpReq(req *http.Request) (addmsg string, ret int) {
	if req == nil {
		return
	}

	// req.Body注意只能读一次
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		addmsg = err.Error()
		ret = ErrUndefined
		return
	}

	retp := ParseJsonByStruct(body, this, 0)
	if retp != Success {
		ret = retp
		addmsg = "[Loopback]" + " 提取body入参失败 "
		return
	}

	return
}

// body , header 同时为nil时使用 req
func (this *LoopbackReqStu) Execute(body []byte, header http.Header,
	req *http.Request) (loopbackResp LoopbackRespStu, addmsg string, ret int) {

	if body == nil && header == nil {
		addmsg, ret = this.setHttpReq(req)
	} else {
		addmsg, ret = this.set(body, header)
	}

	loopbackResp.Error = Success
	loopbackResp.Msg = "ok"
	loopbackResp.Id = this.Id

	if ret != Success {
		return
	}

	return
}

type NextHopReqStu struct {
	Method  string `json:"method,omitempty"`
	Fullurl string `json:"fullurl,omitempty"`
	Header  string `json:"header,omitempty"`
	Body    string `json:"body,omitempty"`
}

type LoopbackRespStu struct {
	Error   int            `json:"error,omitempty"`
	Msg     string         `json:"msg,omitempty"`
	Id      string         `json:"id,omitempty"`
	NextHop NextHopRespStu `json:"nexthop,omitempty"`
}

type NextHopRespStu struct {
	Error int    `json:"error,omitempty"`
	Msg   string `json:"msg,omitempty"`
	Rtt   string `json:"rtt,omitempty"`
}

// skiplevel >=0 +upper,<0 not print
func ParseJsonByStruct(body []byte, v interface{}, skiplevel int) int {
	if v == nil {
		return 0
	}
	// 2016-10-12
	if len(body) == 0 {
		body = []byte("{}")
		// if skiplevel >= 0 {
		// 	prnLog.Debugf("", "income body is null, changed to {}")
		// }
	}

	if err := json.Unmarshal(body, v); err != nil {
		if skiplevel >= 0 {
			prnLog.LogPrint(LOG_ERROR, skiplevel, false, false, "err=%v,body=%v",
				err, string(body))
		}
		return ErrData
	}
	return thisSuccess
}
