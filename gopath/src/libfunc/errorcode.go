package libfunc

var Errormap map[int]string = make(map[int]string, 0)

var (
	StatusInitCommunity = -1000 // 正在初始化数据库
)

var (
	Success     = 0
	thisSuccess = 0

	ErrResetByPeer = -120
	ErrNoSuchHost  = -121
	ErrRefused     = -122
	ErrFromNginx   = -123
	ErrDbVisitFail = -124 /// 数据库访问失败  ///
	ErrStatus404   = -125
	ErrStatus403   = -126
	ErrUndefined   = -200
	ErrData        = -201

	ErrUndefinedNoPrn = -333
)

func init() {
	var errormap_dna = map[int]string{
		0: "ok",

		ErrResetByPeer: "请求tcp连接时对端复位", // connection reset by peer
		ErrNoSuchHost:  "对端不存在,或本机断网",
		ErrRefused:     "对端拒绝访问,目标进程可能未运行",
		ErrFromNginx:   "nginx报错,目标进程可能未运行",
		ErrStatus404:   "对端报404错误(接口不存在)",
		ErrStatus403:   "对端报403错误(接口禁止)",
		ErrDbVisitFail: "数据库访问失败",
		ErrData:        "数据错误",

		ErrUndefined: "未定义",
	}

	for k, v := range errormap_dna {
		Errormap[k] = v
	}
}
