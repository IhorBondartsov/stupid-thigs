package cmd

import (
	"github.com/IhorBondartsov/stupid-things/basic/initalization/pkg3/cmd"
	"log"
)

const NAME_PKG = "Const pkg2"

var  namepkg= "Var pkg2"

func init(){
	log.Println("pkg2: Init().", NAME_PKG, namepkg)
}

func main(){
	log.Println("pkg2: main().")
	cmd.PKG3()
}

func PKG2 ()  {
	log.Println("func PKG2")
}