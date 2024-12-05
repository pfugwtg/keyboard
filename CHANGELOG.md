v1.1 版本
【新增】
- 增加日志打印到文件功能；
- 增加按键映射功能；

【改进】
- 完善菜单栏；
- README.md 增加编译相关描述；
- 过滤重复按键时，只对字母和数字键生效；#https://github.com/garfeng/keyboard/issues/2
- 默认过滤判定时间修改

【BugFix】
- Makefile 编译指令改为静态链接，防止找不到库的情况； #https://github.com/garfeng/keyboard/issues/3
