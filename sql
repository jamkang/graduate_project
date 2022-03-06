//用户表
create table users(
id      int primary key  auto_increment,
name    varchar(20) not null,
sex     char(1) default '0',
belong  char(1) default '0'
) charset=utf8

//管理员表
create