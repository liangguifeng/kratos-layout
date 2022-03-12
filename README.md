# kratos-layout

## 简介

基于kratos封装的基础微服务架构

## 使用方法

1. 新建项目

```shell
git clone git@github.com:liangguifeng/kratos-layout.git kratos-layout-demo
# 如果觉得慢的话可以使用gitee
git clone git@gitee.com:liangguifeng/kratos-layout.git kratos-layout-demo
```

2. 配置环境变量

```shell
# 当前项目所属环境
GO_ENV=dev;
# nacos地址(使用nacos作为配置中心、服务注册、服务发现)
NACOS_ADDRESS=127.0.0.1;
# nacos端口
NACOS_ENDPOINT=8848;
# nacos命名空间
NACOS_NAMESPACE_ID=98745969-8801-47f3-8d65-7ba9bb40f858
```