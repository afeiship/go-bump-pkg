# docker-mysql
> Docker for mysq@57.

## installation
```shell
# 安装
git clone git@github.com:afeiship/docker-mysql.git
cd docker-mysql

# 配置 - 修改密码等
npm run config:init

# 启动 
npm run start-d

# to docker test mysql
docker exec -it mysqllocal bash
mysql -uroot -p -h 127.0.0.1 -P 3306
```

## command
```shell
# login
mysql -uroot -p -h 127.0.0.1 -P 3306
# password: 123456
```
