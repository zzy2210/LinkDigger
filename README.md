# LinkDigger
Golang开发的网页链接爬虫
- 支持基于域名的结果显示
- 支持深度爬取与单独爬取
- 支持结果输出至文件

# USAGE
`-h`

打印帮助页面

`-u`
指定目标（需指明协议） 如：https://www.jd.com

`-d`

使用深度爬取，每爬取到一属于目标域名下的地址则启用goroutine对其进行爬取。

使用深度爬取时会自动将输出至域名.txt文件内。

# RUN
`go run LinkDigger.go -u https://www.jd.com -d`

## v1.0功能
- 基于域名的链接挖掘
- 多线程
- txt标准输出
