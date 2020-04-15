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



指定数据持久化目录启动方式

```
docker run --name psql-study -e POSTGRES_PASSWORD=123456 -p 54321:5432 -v /Users/zonst/docker_data/study_psql:/var/lib/postgresql/data -d postgres:latest
```



这样，创建容器的时候就会自动在主机上保存数据获得初始化，以后启动数据可以挂载这个数据卷

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



持久化启动

```
docker run --name pgdmin4-study -p 80:80 -e "PGADMIN_DEFAULT_EMAIL=nizonglong@163.com" -e "PGADMIN_DEFAULT_PASSWORD=123456" -v /Users/zonst/docker_data/study_pgadmin4:/var/lib/pgadmin -d dpage/pgadmin4
```



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

9. 列操作【新增，修改，删除】

新增列：`alter table iuser add column email varchar(50);`

删除列：`alter table iuser drop column email;`

修改列数据类型：`alter table iuser alter column sex type int  `

重命名：`alter table iuser rename column sex to gender;`

10. 索引相关

> 索引的类型：
>
> - UNIQUE(唯一索引)：不可以出现相同的值，可以有NULL值
> - INDEX(普通索引)：允许出现相同的索引内容
> - PROMARY KEY(主键索引)：不允许出现相同的值
> - fulltext index(全文索引)：可以针对值中的某个单词，但效率确实不敢恭维
> - 组合索引：实质上是将多个字段建到一个索引里，列值的组合必须唯一

PostgreSQL提供了多种索引类型： B-tree、Hash、GiST、SP-GiST 、GIN 和 BRIN。每一种索引类型使用了 一种不同的算法来适应不同类型的查询。默认情况下， CREATE INDEX命令创建适合于大部分情况的B-tree 索引。

创建一个id和name的索引：`create index id_name_idx on iuser (id,name);`

11. 并发控制





## go + postgresql

### 连接示例

使用go链接postgresql示例【官网】

```go
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	…
}
```

You can also connect to a database using a URL. For example:

```go
connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
db, err := sql.Open("postgres", connStr)
```

### CRUD示例

#### 1. select

```go
func sqlSelect() {
	//查询数据
	rows, err := db.Query("SELECT * FROM iuser")
	checkErr(err)

	fmt.Println("-----------")
	for rows.Next() {
		var id int
		var name sql.NullString
		var age sql.NullInt64
		var birthday sql.NullString
		var motto sql.NullString
		var email sql.NullString
		err = rows.Scan(&id, &name, &age, &birthday, &motto, &email)
		checkErr(err)
		fmt.Printf("%d, %s, %d, %s, %s, %s -----------\n", id, name, age, birthday, motto, email)
	}
}
```

查询可以使用多行查询`db.Query`或者单行查询`db.QueryRow`

此处使用多行查询，然后遍历rows获得数据

#### 2. insert

```go
func sqlInsert() {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO iuser(id,name,age) VALUES($1,$2,$3) RETURNING id")
	checkErr(err)

	res, err := stmt.Exec(4, "李4", 23)
	res, err = stmt.Exec(5, "李5", 25)
	//这里的三个参数就是对应上面的$1,$2,$3了

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
```

#### 3. update

```go
func sqlUpdate() {
	//更新数据
	stmt, err := db.Prepare("update iuser set motto=$1 where id=$2")
	checkErr(err)

	res, err := stmt.Exec("li 5 motto", 5)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
```

#### 4. delete

```go
func sqlDelete() {
	//删除数据
	stmt, err := db.Prepare("delete from iuser where id=$1")
	checkErr(err)

	res, err := stmt.Exec(4)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
```

#### 小结

查：`db.Query(sql)`或者`db.QueryRow(sql)`

rows, err := db.Query(sql)然后使用rows.Next()来逐行查询

增删改：`db.Prepare(sql)`

stmt, err := db.Prepare(sql)      预编译sql

Stmt.Exec(params)     执行sql

### 问题与解决

1. 若数据库对应的列为null，则无法读取相应的数据，会报错

```go
// 这个email列在数据库中为null
var email string
```

```go
// 在scan的时候会报错
err = rows.Scan(&id, &name, &age, &birthday, &motto, &email)
```

```verilog
panic: sql: Scan error on column index 4: unsupported Scan, storing driver.Value type <nil> into type *string
```

解决：使用可空数据类型，在sql对应的库里。这里可以使用`sql.NullString`来接收email

```go
5, {李5 %!s(bool=true)}, {25 %!d(bool=true)}, { %!s(bool=false)}, {li 5 motto %!s(bool=true)}, { %!s(bool=false)} 
```

若使用可空数据类型，为空的数据列打印出来如上所示。

## go + redis

### docker 启动 redis

```
docker run --name redis-study -p 6379:6379 -v /Users/zonst/docker_data/study_redis/data:/data -v /Users/zonst/docker_data/study_redis/redis.conf:/etc/redis/redis.conf -d redis:latest redis-server /etc/redis/redis.conf --appendonly yes
```

-p 6379:6379:把容器内的6379端口映射到宿主机6379端口
-v /root/redis/redis.conf:/etc/redis/redis.conf：把宿主机配置好的redis.conf放到容器内的这个位置中
-v /root/redis/data:/data：把redis持久化的数据在宿主机内显示，做数据备份
redis-server /etc/redis/redis.conf：这个是关键配置，让redis不是无配置启动，而是按照这个redis.conf的配置启动
–appendonly yes：redis启动后开启日志



进入redis容器: `docker exec -it b70f3592136f /bin/bash`

进入后执行`redis-cli`若设有密码则输入密码，若无密码则进入了redis指令界面



### go 连接 redis 与基本示例

基本示例

```go
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "test", "go-redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	val, err := redis.String(c.Do("GET", "test"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get test: %v \n", val)
	}
}
```

获取所有key-value

```go
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 打印redis所有k-v
func printValues(keys [][]byte) {
	for _, v := range keys { // 忽略 index
		val, _ := redis.String(c.Do("GET", v))
		fmt.Printf("key %s : %s\n", v, val)
	}
}

var c, err = redis.Dial("tcp", "127.0.0.1:6379")

func main() {

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// 获取所有的key
	keys, err := redis.ByteSlices(c.Do("keys", "*"))
	if err != nil {
		fmt.Println("redis get keys failed:", err)
	} else {
		printValues(keys)
	}
}
```

### 使用go-redis库

```go
package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
```

#### 问题与解决

1. 由于系统环境为go1.9因此一些最新的功能无法支持导致编译出问题，会报错说类似`cannot find package "github.com/go-redis/redis/v7/internal`

解决：在go-redis开启终端，使用`git log`查看日志，然后使用较早版本的代码就可以，使用`git checkout version-number`切换版本即可，此处我使用的是`git checkout 0fdd200bc73d0033e6742edd51979bc81cda2d52`