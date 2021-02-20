package main

import "fmt"

const  englishPrefix = "hello  "
const spanishPrefix = "hala  "
const frenchPrefix = "Bonjour  "

func Hello(name, langueage string) string{
	if name == ""{
		name = "world"
	}
	return greetingPrefix(langueage) + name
}

func greetingPrefix(langueage string) (prefix string){
	switch langueage{
	case "spanish":
		prefix = spanishPrefix
	case "french":
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return prefix
}


func main(){
 fmt.Println("hello")
}