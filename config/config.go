package config

import(
	"os"
	"github.com/spf13/viper"
	"log"
)

//实例化config,这次是指定目录了
func NewConfig(cnfName,cnfType string)*viper.Viper{
	viper := viper.New()
	viper.SetConfigName(cnfName)
	viper.SetConfigType(cnfType)

	env := "dev"
	regin_mod := os.Getenv("REGIN_RUNMODE")
	if regin_mod == "master"{
		env = "master"
	}
	if regin_mod == "test"{
		env = "test"
	}
	cnfDirPath := "./config/"+env
	
	viper.AddConfigPath(cnfDirPath)
	if err := viper.ReadInConfig();err != nil {
		log.Println(err)
	}
	return viper
}