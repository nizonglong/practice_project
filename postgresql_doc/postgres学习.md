GIN框架 + psql 和 redis

postgresql数据库，官网的文档、索引、数据库优化，安装、同步备份等



## postgresql学习

### docker 安装postgresql

安装最新版本：`docker pull postgres`

安装指定版本：`docker pull postgres:9.4`，这里把9.4换成指定版本即可

### docker 启动postgresql

`docker run --name psql_study -e POSTGRES_PASSWORD=password -p 54321:5432 -d postgres:9.4 `

> run，创建并运行一个容器；
--name，指定创建的容器的名字；
-e POSTGRES_PASSWORD=password，设置环境变量，指定数据库的登录口令为password；
-p 54321:5432，端口映射将容器的5432端口映射到外部机器的54321端口；
-d postgres:9.4，指定使用postgres:9.4作为镜像。

然后输入`docker ps`查看当前活动进程

### 链接docker中的postgresql数据库

1. `docker exec -it cid 命令`，例如`docker exec -it d6t5987ru7 psql -U postgres`

这里就可以对postgresql进行命令行操作

2. 使用pgAdmin4进行链接

`docker port cnane`使用这个命令查看指定容器名称对应的端口号映射

例如`docker port psql_study`查看psql容器对应的端口号映射

> 5432/tcp -> 0.0.0.0:54321

这里是psql_study容器将 5432 端口映射到 Docker 主机的 54321 端口上

### 保存docker中对postgresql的修改

`docker commit -a "nizonglong" -m "add user root, make root admin_user" d70695442ee8 psql_study`

命令的意思是：将d70695442ee8提交为一个新的镜像，保存了对此镜像的所有更改，添加了作者信息以及备注信息

-a:修改者信息 

-m:注释、说明 紧跟着当前操作的容器id 最后是要生成的新的镜像名称 



PS：不能重复使用同一条提交命令，会产生错误，因为生成的id不同

### docker安装pgAdmin4

docker 下载pgAdmin4: `docker pull dpage/pgadmin4`

启动pgAdmin4，以HTTP方式启动Docker

- PGADMIN_DEFAULT_EMAIL：这个可以作为你登录Pgadmin4的用户名
- PGADMIN_DEFAULT_PASSWORD：登录Pgadmin4的默认密码

```
docker run -p 80:80 \
-e "PGADMIN_DEFAULT_EMAIL=user@domain.com" \
-e "PGADMIN_DEFAULT_PASSWORD=SuperSecret" \
-d dpage/pgadmin4
```

然后在浏览器里输入`localhost:80`就可以进入pgAdmin4的web管理界面

连接的话需要使用`docker inspect cid`查看容器的信息，其中IPAddress是容器的ip，填入host框即可，默认用户名是postgres，密码则是创建postgresql容器的时候填入的信息。

## postgresql基础命令

### 特殊注意

1. 若需要插入字符串: this is li's book，但是由于插入数据库时候使用的是`'`因此可以使用`两个'`代替一个`'`。即：`this is li''s book`
2. 注释：使用`--`，示例：`-- 这是普通的单行注释`

多行注释：使用

```
/*
 * 这是
 * 多行
 * 注释
 */
```

### 基础知识

#### 约束

1. 检查约束check

> 一个检查约束是最普通的约束类型。它允许我们指定一个特定列中的值必须要满足一个布尔表达式。例如，为了要求正值的产品价格，我们可以使用：
>
> CREATE TABLE products (
>
> product_no integer, 
>
> name text,
>
> price numeric CHECK (price > 0)  
>
> );

也可以给约束取名：`price numeric CONSTRAINT positive_price CHECK (price > 0)`

一个检查约束也可以引用多个列：`CHECK (price > discounted_price)`这种约束写在最后，又称表约束。当然，表约束也可以取名：`CONSTRAINT valid_discount CHECK (price > discounted_price)`

2. 非空约束 NOT NULL

```sql
CREATE TABLE products (
product_no integer NOT NULL,
name text NOT NULL,
price numeric
);
```

3. 唯一约束

```sql
CREATE TABLE products (
product_no integer UNIQUE,
name text,
price numeric
);
```

也可以有表约束：`UNIQUE (product_no)`

也可以组合约束：`UNIQUE (a, c)`

当然，命名也是可以的：`product_no integer CONSTRAINT must_be_different UNIQUE,`

4. 主键约束

```sql
CREATE TABLE products (
product_no integer PRIMARY KEY,
name text,
price numeric
);
```

主键也可以包含多于一个列，其语法和唯一约束相似：`PRIMARY KEY (a, c)`

5. 

### console cmd控制台命令

- \h：查看SQL命令的解释，比如\h select。
- \?：查看psql命令列表。
- \l：列出所有数据库。
- \c [database_name]：连接其他数据库。
- \d：列出当前数据库的所有表格。
- \d [table_name]：列出某一张表格的结构。
- \du：列出所有用户。
- \e：打开文本编辑器。
- \conninfo：列出当前数据库和连接的信息。

### CURD基本命令

```sql
# 创建新表
CREATE TABLE user_tbl(name VARCHAR(20), signup_date DATE);

# 插入数据
INSERT INTO user_tbl(name, signup_date) VALUES('张三', '2013-12-22');

# 选择记录
SELECT * FROM user_tbl;

# 更新数据
UPDATE user_tbl set name = '李四' WHERE name = '张三';

# 删除记录
DELETE FROM user_tbl WHERE name = '李四' ;

# 添加栏位
ALTER TABLE user_tbl ADD email VARCHAR(40);

# 更新结构
ALTER TABLE user_tbl ALTER COLUMN signup_date SET NOT NULL;

# 更名栏位
ALTER TABLE user_tbl RENAME COLUMN signup_date TO signup;

# 删除栏位
ALTER TABLE user_tbl DROP COLUMN email;

# 表格更名
ALTER TABLE user_tbl RENAME TO backup_tbl;

# 删除表格
DROP TABLE IF EXISTS backup_tbl;
```

#### 基础练习

1. 创建数据库study `create database study;`，创建数据库test `create database test;`
2. 删除数据库test：`drop database test;`
3. 切换到sutdy数据库：`\c study;`
4. 创建表 iuser{id, name, age, birthday, motto}, motto是座右铭

```sql
create table iuser (
id bigint,
name varchar(20),
age int,
birthday date,
motto varchar(100) );
```

5. iuser表新增数据

> 1, '李1', 18, '2006-01-02', 'li 1 motto'
>
> 2, '李2', 19, '2005-02-03', 'li 2"s motto'
>
> 3, '李3', 20, '2004-03-04', 'li 3 motto'

```sql
insert into iuser values (1, '李1', 18, '2006-01-02', 'li 1 motto');
insert into iuser values (2, '李2', 19, '2005-02-03', 'li 2"s motto');
insert into iuser values (3, '李3', 20, '2004-03-04', 'li 3 motto');
```

6. 修改iuser表的id为2的motto: li 2 motto

`update iuser set motto='li 2 motto' where id=2;`

7. 新建表irole {rid, name, desc}

```sql
create table irole (
	rid smallint primary key,
	name varchar(20),
	descr varchar(30)
);
```

8. 将iuser表的id设置为主键

`alter table iuser add constraint iuser_pkey primary key (id);`

删除主键

`alter table iuser drop constraint iuser_pkey`

> --删除主键
> alter table 表名 drop constraint 主键名
>
> --添加主键
> alter table 表名 add constraint 主键名 primary key(字段名1,字段名2……)
>
> --添加非聚集索引的主键
> alter table 表名 add constraint 主键名 primary key NONCLUSTERED(字段名1,字段名2……)
>
> 新建约束：
> ALTER TABLE [表名] ADD CONSTRAINT 约束名 CHECK ([约束字段] <= \'2000-1-1\')
>
> 删除约束：
> ALTER TABLE [表名] DROP CONSTRAINT 约束名
>
> 新建默认值
> ALTER TABLE [表名] ADD CONSTRAINT 默认值名 DEFAULT \'51WINDOWS.NET\' FOR [字段名]
>
> 删除默认值
> ALTER TABLE [表名] DROP CONSTRAINT 默认值名

9. 索引相关

> 索引的类型：
>
> - UNIQUE(唯一索引)：不可以出现相同的值，可以有NULL值
> - INDEX(普通索引)：允许出现相同的索引内容
> - PROMARY KEY(主键索引)：不允许出现相同的值
> - fulltext index(全文索引)：可以针对值中的某个单词，但效率确实不敢恭维
> - 组合索引：实质上是将多个字段建到一个索引里，列值的组合必须唯一

