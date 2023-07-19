package tools

import (
	"flag"
	"fmt"
)

func Console() {
	urlPtr := flag.String("u", "", "URL")
	filePtr := flag.String("f", "", "File")
	nodns := flag.Bool("nd", false, "No dns validation is used")
	expPtr := flag.Bool("exp", false, "Execute ExpWebshell")
	flag.Parse()
	url := *urlPtr
	file := *filePtr
	exp := *expPtr
	nd := *nodns
	// 如果没有提供 -f 或 -u 参数，则不执行 POC，并直接返回
	if url == "" && file == "" && !exp {
		return
	}
	if exp {
		if url != "" {
			ExpWebshell(UrlHandler(url))
		}

		if file != "" {
			lines, err := ReadFile(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, line := range lines {
				ExpWebshell(UrlHandler(line))
			}
		}
	} else {

		if url != "" {
			if nd {
				FileExp(UrlHandler(url))
			} else {
				POC(UrlHandler(url))
			}

		}

		if file != "" {
			lines, err := ReadFile(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, line := range lines {
				if nd {
					FileExp(UrlHandler(line))
				} else {
					POC(UrlHandler(line))
				}

			}
		}
	}

}
