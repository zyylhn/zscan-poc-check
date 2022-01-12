# zscan-poc-check

暂时兼容xray v1版本的poc，根据实际使用需求看是否需要升级到v2

参考

```
https://github.com/shadow1ng/fscan
https://github.com/jjf012/gopoc
```

## 简单使用

查看所有内置poc

```
./zscan_poc_check -list       
```

查看指定内置poc

```
./zscan_poc_check -list -name weblogic
```

使用全部内置poc扫描目标

```
./zscan_poc_check -t http://192.168.101.104:7001 
```

使用指定poc扫描目标

```
./zscan_poc_check -t http://192.168.101.104:7001 -name weblogic
```

不使用内置poc使用外部poc

```
./zscan_poc_check -t http://192.168.101.104:7001 -path cmd/pocs
```

不使用内置poc使用外部poc的weblogic的poc

```
./zscan_poc_check -t http://192.168.101.104:7001 -path cmd/pocs -name weblogic
```

使用外部指定的某个poc

```
./zscan_poc_check -t http://192.168.101.104:7001 -path cmd/pocs/weblogic-cve-2017-10271.yml 
```

