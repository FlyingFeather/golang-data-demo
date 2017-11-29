# golang-data-demo
golang 构建数据服务
使用xorm
## 1.下载Docker
1.
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/1.png)

2.
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/2.png)

3.
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/3.png)

4.
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/4.png)

5.
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/5.png)

6.
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/6.png)

7.
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/7.png)


## 2.安装数据库
### 2.1.下载镜像
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/8.png)

### 2.2 启动 mysql 作为主机服务
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/9.png)

### 2.3 启动 mysql client 访问服务器

![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/10.png)


- sudo docker run -it --rm --net host mysql:5.7 "sh" 启动了容器内 sh 进程。

        -it 等价于 -i -t ，表示使用当前 stdin 和 stdout 作为该进程的 io
        --rm , 当该进程结束时，自动清理该容器的文件系统空间
        --net host，表示容器使用当前主机的网络环境
        参数1 参数2，分别是镜像和在镜像中执行的命令

- # 表示你处于容器的超级管理员的 shell
- mysql -h127.0.0.1 -P3306 -uroot -proot mysql 客户端的命令

![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/11.png)

## 3.创建数据库test
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/12.png)

## 4.构建数据服务测试结果
基本代码与课件基本一致
### 4.1.测试web服务
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/13-qidong.png)

#### 添加数据
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/13.png)
- -d 参数是 POST 到服务器的数据

#### 查询数据
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/14.png)

### 4.2.ab测试结果
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/15.png)

### 4.2.curl测试结果
![](https://github.com/FlyingFeather/golang-data-demo/blob/master/screenshot/16.png)








