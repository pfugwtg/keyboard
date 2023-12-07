# 防止某键盘误触发

新买的键盘用了几个月，一个问题越来越明显，就是误触发，按一次，重复响应多次，随时随地敲错特别烦。

就算我按得再轻也没用。

这个问题用 win 自带的 筛选键功能无法解决。

然而键盘超出了包退期，要自己花运费寄过去修QAQ！

所以我写了这个小程序，屏蔽重复触发。

希望能帮到不小心买到类似键盘的你。

解压后，双击里面的 keyboard.exe 运行。

可以设置为开机启动，这样以后就不需要手动双击了。

下载地址：https://github.com/garfeng/keyboard/releases/latest

## 编译

Windows 环境下编译时，请确保系统已安装如下工具：
- [Go](https://go.dev/dl/)：我使用的版本是 1.18，其它版本未测试
- [MinGW](https://www.mingw-w64.org/)：包含了 gcc、g++ 等工具
- [make](https://gnuwin32.sourceforge.net/packages/make.htm)（可选）

编译指令：
```bash
make
#
# 若未安装 make 工具，也可直接执行 Makefile 中的指令：
#     go build -ldflags "-w -s -H windowsgui -extldflags --static"
#
```

编译完成后，会在当前路径生成一个 exe 可执行文件。

## 日志

* 2022年9月28日：加系统托盘图标

* 2018年6月：基本功能完成