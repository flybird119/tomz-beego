tomz
====

golang web app


在环境变量中设置根目录 $GOPATH,在 $GOPATH 下新建 src 目录，用于存放web app,
整个项目结构如下：


      $GOPATH                         $GOPATH 根目录
        src                           $GOPATH/src 目录
          github.com/astaxie/beego    依赖的第三方工具->>Beego源码目录
          ...
          tomz                App 根目录
              main.go         Main 包文件
              controllers     App 控制器目录
              models          App 模型目录
              views           模板/HTML文件
              conf            配置文件目录
                  app.conf    主要配置文件
              static          静态文件目录
                  css         CSS 文件目录
                  js          Javascript 文件目录
                  img         图片文件目录
            
            

第三方依赖：

  tomz依赖第三方工具beego作为开发框架，

  通过 go get github.com/astaxie/beego 下载工具

  更多信息请查看beego的官网：http://beego.me/quickstart
