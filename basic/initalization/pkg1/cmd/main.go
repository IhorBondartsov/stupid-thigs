package pkg1

import (
	"github.com/IhorBondartsov/stupid-things/basic/initalization/pkg2/cmd"
	"log"
)

const NAME_PKG = "Const pkg1"

var  namepkg= "Var pkg1"

func init(){
	log.Println("pkg1: Init().", NAME_PKG, namepkg)
}

func main(){
	log.Println("pkg1: main().")
	cmd.PKG2()
}

func PKG1 ()  {
	log.Println("func PKG1")
}