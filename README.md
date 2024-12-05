# 防止某键盘误触发

## 背景

​	我的键盘问题容易误触发，即，按一次，重复响应多次，随时随地敲错特别烦。网上找了一下，发现有人也有类似的烦恼，但用起来感觉有 Bug，而且频率不适合自己，于是我就 Fork 了原作者的仓库，自己修改使用。原作者仓库链接：

https://github.com/garfeng/keyboard



## 主要功能

我根据自己的需求，实现了如下功能：

- 限制用户按键输入频率（原始需求）；
- 映射快捷键 （后加的）：将 Ctrl + C/V/F 分别映射到 F2/F3/4 （这几个按键我基本不用）；



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



## 使用指南

### 1 如何使用？

运行生成的 exe 程序即可。

### 2 如何开机自启动？

- Win + R
- 输入 `shell:startup`，点击“确定”后会打开文件夹
- 在该文件夹下，创建本 exe 的快捷方式



## 日志

* 2024年12月4日：增加解盘按钮映射

* 2022年9月28日：增加系统托盘图标

* 2018年6月：基本功能完成