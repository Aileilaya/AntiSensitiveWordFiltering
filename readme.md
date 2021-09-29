# 反敏感词屏蔽工具

自动对文章进行修改，从而绕过敏感词屏蔽机制(此工具仅针对菠萝包轻小说的屏蔽机制，可能不适用于其他高级屏蔽机制)

例

|        | 修改前 | 修改后 |
| ------ | ------ | ------ |
| 屏蔽前 | 榨干   | 榨^干  |
| 屏蔽后 | **     | 榨^干  |

## 参数说明

| 选项 | 说明                   |
| ---- | ---------------------- |
| -i   | 指定要修改的文件或目录 |
| -s   | 指定保存结果的目录     |

注意本程序仅处理文本文件(后缀为`.txt`)

## 配置文件说明

| 键         | 值                          |
| ---------- | --------------------------- |
| Version    | 配置文件的版本号            |
| Dictionary | 屏蔽词列表 用户可以自行修改 |
| Separator  | 用于插^入屏蔽词中间的分隔符 |

## 使用方法举例

### 处理单个文件

`.\` 表示当前目录

```
.\shatang.exe -i .\test\测试.txt -o .\out\
正在处理 D:\novel\shatang\test\测试.txt
```

也可以使用绝对路径

```
.\shatang.exe -m 1 -i D:\novel\shatang\test\测试.txt -o D:\novel\shatang\out
正在处理 D:\novel\shatang\test\测试.txt
```

### 处理目录下的所有文件

```
.\shatang.exe -i .\test\ -o .\out\
正在处理 D:\novel\shatang\test\测试 - 副本.txt
正在处理 D:\novel\shatang\test\测试.txt
```

也可以使用绝对路径

```
.\shatang.exe -i D:\novel\shatang\test -o D:\novel\shatang\out
正在处理 D:\novel\shatang\test\测试 - 副本.txt
正在处理 D:\novel\shatang\test\测试.txt
```

## 如何修改屏蔽词字典

修改目录下的`config.json`，在`Dictionary`键中添加新屏蔽词

务必按照格式添加

## 如何修改插入字符

修改目录下的`config.json`，`separator`键的内容即为隔断字符串

## 如何打开终端

1. 如果你是windows10 用户 推荐使用 应用商店中的windows terminal
2. 如果你想使用cmd 同时按下Windows徽标键+R 在输入框内输入 cmd
   1. 关于cmd切换工作目录的方法请自行搜索答案
3. 如果你想使用powershell，在文件资源管理器中 按住shift键 鼠标右键在菜单中选择"在此处打开powershell窗口"

