# FileDeduplication		|		文件去重

使用 Go 语言实现的文件去重，单线程
内存占用，越到后面可能占用越大，扫描**777个文件**到后期内存占用达到 **100~150 MB**
刚入门 Go，可能有很多不规范的、冗余的代码，接受指正

# 实现思路
遍历文件下所有文件，计算文件 MD5值，然后压入 map，map 结构大概是这样的
```
  key    :  value
------------------
文件路径 : 文件MD5值
```
然后根据当前 **计算的 MD5值** 在 map 中寻找，如果找到了，就判定文件有重复
将当前 重复文件路径，写入到 **重复文件清单**

删除重复文件：
按行读取 **重复文件清单**，然后执行删除操作



# 安装
下载 `fileCheck.go` 文件，然后安装 `go` 环境进行编译

执行以下命令
```
go build fileCheck.go
```

编译完成你将得到一个 **二进制文件**
- **Windows 平台** 将得到一个名为 `fileCheck.exe` 二进制文件
- **Linux、Android 平台** 将得到一个名为 `fileCheck` 二进制文件
- ...

# 使用
> 这里以 **Windows** 平台为例

```
E:\awesomeProject>fileCheck -h
Usage of file_check:
  -de string
        需要删除重复文件的清单
  -dn string
        需要检查重复文件的路径，如 A:\CloudMusic，路径最后面建议不要带 \ 或 / (default "./")
```

## 查询重复文件
> 基本语法：`fileCheck -dn "需要查重的文件夹"`

实例：
```
fileCheck -dn "A:\公共音乐"
```

查询完成后，如果**没有重复文件**，将显示如下信息：
```
E:\awesomeProject>fileCheck -dn "A:\公共音乐"
92a20dcf766db2f9ecb2a54740b4dbb8        A:\公共音乐/马頔 - 南山南.MP3
f5ffdf3d7d446ae0f163ace933dfd2a4        A:\公共音乐/马頔 - 皆非.mp3
......
067968ad7e563925eb462e2bd0e515c2        A:\公共音乐/齐一 - 这个年纪.flac
9fa20140df5805b14fb605c5cc1f58bb        A:\公共音乐/내가 설렐 수 있게 - Apink.flac
扫描完成，共扫描 16 文件，有 0 文件重复，耗时：0s
```

查询完成后，如果**有重复文件**，将显示如下信息：
```
E:\awesomeProject>fileCheck -dn "A:\公共音乐"
47617b7b37e5e006b490f405b7e55a72        A:\公共音乐/颜人中 - 有可能的夜晚.flac
fd2609dc9ad3c0ad66a6da8e0f074b87        A:\公共音乐/魔鬼中的天使 - 田馥甄.flac
  文件名：A:\公共音乐/马頔 - 南山南.MP3
重复路径：A:\公共音乐/马頔 - 南山南 (1).MP3

f5ffdf3d7d446ae0f163ace933dfd2a4        A:\公共音乐/马頔 - 皆非.mp3
......
067968ad7e563925eb462e2bd0e515c2        A:\公共音乐/齐一 - 这个年纪.flac
9fa20140df5805b14fb605c5cc1f58bb        A:\公共音乐/내가 설렐 수 있게 - Apink.flac
扫描完成，共扫描 777 文件，有 1 文件重复，耗时：395s
重复文件清单路径：A:\公共音乐/1581495121_delete_list.txt
```

## 删除重复文件
> 自己手动检查 `重复文件清单路径` 后，如果**有问题**的，可以自己修改路径，格式是**一行一个文件绝对路径**

> 基本语法：`fileCheck -de "重复文件清单路径"`

```
E:\awesomeProject>fileCheck -de "A:\公共音乐\1581495121_delete_list.txt"
[删除成功]A:\公共音乐/G.E.M. 邓紫棋 - 喜欢你 [mqms2].mp3
[删除成功]A:\公共音乐/马頔 - 南山南.MP3
[删除成功]A:\公共音乐/Gamper & Dadoni Ember Island - Creep (1).mp3
[删除成功]A:\公共音乐/Groove Coverage - Far Away from Home (1).MP3
[删除成功]A:\公共音乐/JC - 说散就散.mp3
[删除成功]A:\公共音乐/Taylor Swift - Look What You Made Me Do (1).MP3
[删除成功]A:\公共音乐/Vicetone Cozi Zuehlsdorff - Nevada (feat. Cozi Zuehlsdorff) (1).MP3
[删除成功]A:\公共音乐/Vicetone Cozi Zuehlsdorff - Nevada (feat. Cozi Zuehlsdorff).MP3
[删除成功]A:\公共音乐/Wiz Khalifa Charlie Puth - See You Again (feat. Charlie Puth).mp3
[删除成功]A:\公共音乐/周冬雨 - 不完美女孩 [mqms2].mp3
[删除成功]A:\公共音乐/好妹妹 - 往事只能回味.MP3
[删除成功]A:\公共音乐/好妹妹 - 月.MP3
[删除成功]A:\公共音乐/岑宁儿 - 追光者 [mqms2].mp3
[删除成功]A:\公共音乐/徐佳莹 - 湫兮如风 [mqms2].mp3
[删除成功]A:\公共音乐/李易峰 - 年少有你 [mqms2].mp3
[删除成功]A:\公共音乐/杨宗纬 - 越过山丘 [mqms2].mp3
[删除成功]A:\公共音乐/Ryan.B AY楊佬叁 - 再也没有.mp3
[删除成功]A:\公共音乐/赵英俊 - 守候 [mqms2].mp3
[删除成功]A:\公共音乐/赵雷 - 成都 [mqms2].mp3
[删除成功]A:\公共音乐/是阿涵阿 - 过客.mp3
[删除成功]A:\公共音乐/是阿涵阿 王冕 - 讨厌你.mp3
[删除成功]A:\公共音乐/陈雪凝 - 我唯一青春里的路人.MP3
[删除成功]A:\公共音乐/鲤 仙灵女巫果妹 - Russ - 早安少女 (Psycho remix)（鲤   台灯家的果妹 remix） (1).mp3
```
