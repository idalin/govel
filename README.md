# Govel

用go语言写的网络小说下载阅读器，兼容 [阅读](https://github.com/gedoor/MyBookshelf) 的书源格式

----
## 代码结构
```
govel
    ├── config.json     # web 配置文件
    ├── controllers # web 版的contrller
    │   ├── book.go
    │   ├── book_routers.go
    │   ├── controllers.go
    │   └── websockets.go
    ├── gui     # gui版，qt的go绑定写的
    │   ├── images
    │   ├── main.go
    │   ├── main.qrc
    │   ├── qml
    │   └── search.go
    ├── main.go # web版的主程序入口
    ├── models  # 各种数据结构
    │   ├── book.go
    │   ├── chapter.go
    │   ├── rule.go # 规则的，实际没使用
    │   ├── shelf.go 
    │   ├── source.go 
    │   └── source_test.go
    ├── README.md
    ├── storage # 本地存储，准备做成interface，支持关系型数据库和本地文件
    │   ├── common.go
    │   ├── file.go
    │   └── file_test.go
    └── utils  # 一些通用工具
        ├── common.go
        ├── fetcher.go
        ├── fetcher_test.go
        └── parser.go # 规则解析，准备放到models里的，不过刚开始做，先整出个能用的来测试测试一下
```
----
## 安装使用
还没有正式完成，所以其实安装使用不了的。不过主要功能可以用了，通过 go test可以看到。 models/source_test.go 可以测试搜书等功能， storage/file_test.go可以实现本地缓存的测试。
当然，需要自备书源配置文件。