package main

import (
	
	"fmt"
	"os"
	"strings"
	"reverse/functions"
)

func main() {
	args := os.Args[1:]


if len(args)==1 && !strings.HasPrefix(args[0], "--reverse="){
	input:=args[0]
	banner:="standard"
	result:=reverse.AsciiNormal(input,banner)
	fmt.Print(result)
	return
}else if len(args)==1 && strings.HasPrefix(args[0], "--reverse="){
	reverse.Reverse(args)
	fmt.Println("reverse worked ")

}else if len(args)==2{

	if (strings.HasPrefix(args[0], "--color=")&&strings.HasPrefix(args[1], "--output="))||(strings.HasPrefix(args[1], "--color=")&&strings.HasPrefix(args[0], "--output="))||
	(strings.HasPrefix(args[0], "--color=")&&strings.HasPrefix(args[1], "--reverse="))||(strings.HasPrefix(args[1], "--color=")&&strings.HasPrefix(args[0], "--reverse="))||
	(strings.HasPrefix(args[0], "--output=")&&strings.HasPrefix(args[1], "--reverse="))||(strings.HasPrefix(args[1], "--output=")&&strings.HasPrefix(args[0], "--reverse=")){
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName> ")
	
	} else if strings.HasPrefix(args[0], "--color="){

	result := reverse.AsciiColor(args)
	fmt.Print(result)
    return

}else if strings.HasPrefix(args[0], "--output=") {

	reverse.AsciiOutput(args)

}else if args[1]=="standard"||args[1]=="thinkertoy"||args[1]=="shadow"{
	input:=args[0]
    banner:=args[1]
    result:=reverse.AsciiFs(input,banner)
    fmt.Print(result)
    return
}else {
	fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName> ")

    return
}


}else if len(args)==3{

	if (strings.HasPrefix(args[0], "--color=")&&strings.HasPrefix(args[1], "--output="))||(strings.HasPrefix(args[1], "--color=")&&strings.HasPrefix(args[0], "--output=")){
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something'")
	
	} else if strings.HasPrefix(args[0], "--color="){
		result := reverse.AsciiColor(args)
		fmt.Print(result)
		return
	}else if strings.HasPrefix(args[0], "--output=") {

		reverse.AsciiOutput(args)
	
	}else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something' ")
		return
	}

}else if len(args)==4{

	if (strings.HasPrefix(args[0], "--color=")&&strings.HasPrefix(args[1], "--output="))||(strings.HasPrefix(args[1], "--color=")&&strings.HasPrefix(args[0], "--output=")){
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something'")
	
	} else if strings.HasPrefix(args[0], "--color="){
		result := reverse.AsciiColor(args)
		fmt.Print(result)
		return
	}else{
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something' ")
		return
	}
}else {
	fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something'")

		return
	}
}
