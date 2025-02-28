package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// 綁定（Bind）並驗證（Validate）請求參數
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	// 綁定請求參數
	err := c.ShouldBind(v)
	if err != nil {
		fmt.Println("🔴 1 BindAndValid 發生錯誤:", err)
		// 取得 gin.Context 中的翻譯器
		v := c.Value("trans")
		trans, _ := v.(ut.Translator) // 將 trans 轉換成 ut.Translator（通用翻譯介面）
		// 判斷錯誤類型是否為 ValidationErrors
		verrs, ok := err.(val.ValidationErrors) // err.(val.ValidationErrors)：檢查 err 是否為 驗證錯誤類型，如果不是，就直接返回錯誤。
		if !ok {
			fmt.Println("🔴 2 BindAndValid 發生錯誤:", err)
			return false, errs
		}

		// 翻譯錯誤訊息 (verrs.Translate(trans) 會把錯誤訊息翻譯成可讀的格式)
		for key, value := range verrs.Translate(trans) {
			// 把這些錯誤存進 errs 陣列，稍後返回給調用者
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		// 如果驗證失敗，回傳 false, errs
		return false, errs
	}

	// 如果驗證成功，回傳 true, nil
	return true, nil
}
