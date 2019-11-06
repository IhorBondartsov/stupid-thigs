package cmd

import "log"

const NAME_PKG = "Const pkg3"

var  namepkg= "Var pkg3"

func init(){
	log.Println("pkg3: Init().", NAME_PKG, namepkg)
}

func main(){
	log.Println("pkg3: main().")
}

func PKG3 ()  {
	log.Println("func PKG3")

}