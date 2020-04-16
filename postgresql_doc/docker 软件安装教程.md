# docker 软件安装教程

## 数据库

### mysql

### redis

### postgresql

```
// pull
docker pull postgres

// run
docker run --name psql_study -e POSTGRES_PASSWORD=password -p 54321:5432 -d postgres
```

指定数据持久化目录启动方式

```
docker run --name psql-study -e POSTGRES_PASSWORD=123456 -p 54321:5432 -v /Users/zonst/docker_data/study_psql:/var/lib/postgresql/data -d postgres:latest
```

## web服务器

### tomcat

## 工具

### 接口工具

#### postwoman

```
#pull
docker pull liyasthomas/postwoman

#run
docker run -p 3000:3000 liyasthomas/postwoman:latest

#build
docker build -t postwoman:latest
```

数据持久化启动

```
docker run -p 3000:3000 -v /Users/zonst/docker_data/study_postwoman:/app liyasthomas/postwoman:latest
```

