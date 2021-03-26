package route

import (
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	// API v1
	ApiV1(e)
	// API v2
	// APIのVersionを追加する場合はvX.goの形式でファイルを新規作成して
	// 関数を呼び出すようにしてください
	// ApiV2(e)
}
