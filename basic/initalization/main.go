package main

import (
	"github.com/IhorBondartsov/stupid-things/basic/initalization/pkg1/cmd"
	"log"
)

const NAME_PKG = "Const main(pkg0)"

var  namepkg= "Var main(pkg0)"

func init(){
	log.Println("init pkg (pkg0)", NAME_PKG, namepkg)
}

func main() {
	log.Println("main pkg (pkg0)" )
	pkg1.PKG1()
}