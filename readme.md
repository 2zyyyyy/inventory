## 待办清单项目

****

### 项目背景

使用 golang + gin + gorm开发的一个练手项目。通过此项目可以初步掌握 gin web框架开发应用以及 gorm 的使用。（学习链接：https://github.com/Q1mi/bubble，感谢七米老师在哔哩哔哩的视频教程）

### 功能介绍

- 项目结构

  ```shell
  $ tree  
  ├── config											// 配置文件，如数据库配置等
  │   ├── app.yaml
  │   ├── main.go
  │   └── test.yaml
  ├── controller									// 控制层，根据请求路由执行不同的动作
  │   └── controller.go
  ├── dao													// 数据访问对象(data access object，DAO)
  │   └── mysql.go
  ├── go.mod
  ├── go.sum
  ├── main.go											// 项目main函数
  ├── models											// 模型层，存放对应数据库模型及其操作
  │   └── todo.go
  ├── readme.md
  ├── routers											// 路由注册，控制前端请求的 URL 及对应的 controller 层的动作
  │   └── routers.go
  ├── setting											// 读取 config 层的 yaml 配置文件数据(未完成)
  │   └── setting.go
  ├── static											// 前端打包的静态文件
  │   ├── css
  │   │   ├── app.8eeeaf31.css
  │   │   └── chunk-vendors.57db8905.css
  │   ├── fonts
  │   │   ├── element-icons.535877f5.woff
  │   │   └── element-icons.732389de.ttf
  │   └── js
  │       ├── app.007f9690.js
  │       └── chunk-vendors.ddcb6f91.js
  ├── templates										// 模板文件
  │   ├── favicon.ico
  │   └── index.html
  └── tmp
      └── main
  12 directories, 21 files
  ```

  ![image-20220504211922447](https://tva1.sinaimg.cn/large/e6c9d24ely1h1wor26i7sj20nk0yi0tz.jpg)

- 功能演示

  1. 主页

     在根目录启动本项目 

     > $ go run main.go  (建议安装 air 热加载框架，更新后无需再次构建。启动直接使用 air 命令即可)

     访问设置的本机地址（127.0.0.1:8888）。

     ![image-20220504212128423](https://tva1.sinaimg.cn/large/e6c9d24ely1h1wot9y175j211u0u0abb.jpg)

  2. 添加
     ![image-20220504212229241](https://tva1.sinaimg.cn/large/e6c9d24ely1h1woubdzs1j214c0s6abs.jpg)
  3. 完成
     ![image-20220504212252721](https://tva1.sinaimg.cn/large/e6c9d24ely1h1wous92rdj215g0swtai.jpg)
  4. 删除
     ![image-20220504212317052](https://tva1.sinaimg.cn/large/e6c9d24ely1h1wov5nvy6j214i0oota7.jpg)
  5. 恢复
     ![image-20220504212337963](https://tva1.sinaimg.cn/large/e6c9d24ely1h1woviekd6j21560oiabl.jpg)

- 数据查看

  我们可以查看本地的 mysql 数据库来查看对应的数据落库情况。

  ```mysql
  mysql> mysql> select * from todos;
  +----+-----------------+--------+
  | id | title           | status |
  +----+-----------------+--------+
  |  1 | basketball      |      0 |
  |  4 | 告别 996        |      0 |
  |  5 | 万里大帅比      |      0 |
  +----+-----------------+--------+
  3 rows in set (0.00 sec)
  ```

- 终端查看访问记录

  ![image-20220504212646561](https://tva1.sinaimg.cn/large/e6c9d24ely1h1woysz1b4j21bo0sgtfv.jpg)