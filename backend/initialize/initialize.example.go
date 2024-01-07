package initialize

import "os"

func InitExample() {
	os.Setenv("POLYGON_MAINNET_URL", "")
	os.Setenv("CREDIT_CONTRACT", "")
}
