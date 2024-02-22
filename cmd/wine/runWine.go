package wine

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func RunWine(name string, app string) {
	winePath := GetWinePath() + "wine"
	bottlePath := viper.GetString("WINE_PATH") + "/bottles/" + name
	command := fmt.Sprintf(`MVK_CONFIG_RESUME_LOST_DEVICE=1 WINEPREFIX="%s" "%s" "%s"`, bottlePath, winePath, app)
	println(command)
	println(name)
	println(winePath)
	println(bottlePath)
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))

}

// exec('MVK_CONFIG_RESUME_LOST_DEVICE=1 WINEPREFIX="'.__DIR__.'"  ./wine/bin/wine  ./game/Divimera.exe', $output, $retval);
