package tools

func Banner() {
	banner := `
 __  __                _  __
|  \/  | __ _ _ __ ___| |/ /
| |\/| |/ _' | '__/ __| ' /
| |  | | (_| | |  \__ \ . \
|_|  |_|\__,_|_|  |___/_|\_\
`
	print(banner)
}
func Flag() {
	Banner()
	print("-u url\n")
	print("-f 文件\n")
	print("-nd 不使用DNSLOG\n")
	print("-exp get_webshell\n")

}
