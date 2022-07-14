//go:generate stringer -type ErrCode -linecomment
package enum

type ErrCode int64 // 错误码

const (
	ERR_CODE_OK             ErrCode = iota // OK
	ERR_CODE_INVALID_PARAMS                // 无效参数
	ERR_CODE_TIMEOUT                       // 超时
)
