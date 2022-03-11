//用户表
create table users(
id      int primary key  auto_increment,
name    varchar(20) not null,
sex     char(1) default '0',
belong  char(1) default '0'
) charset=utf8

//管理员表
create



//商品分类表
create table goodclassify(
id      int primary key  auto_increment,
name    varchar(10)  UNIQUE,
supclass  varchar(10) NOT NULL DEFAULT "一级分类",
classnum int NOT NULL DEFAULT 1,
shopnum int DEFAULT 0
)charset=utf8

//商品表

create table goods(
id      int primary key  auto_increment,
cid     int,
name    varchar(10)  UNIQUE,
minmoney float DEFAULT 0,
maxmoney float DEFAULT 0,
briefintro text,
foreign key(cid) references goodclassify(id)
)charset=utf8

//商品图片表
create table goodsimg(
gid int,
url varchar(50),
num int default 0,
foreign key(gid) references goods(id)
)charset=utf8