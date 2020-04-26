create table user_inf (
  uid int primary key,
  username varchar(30) not null unique ,
  password varchar(16) not null,
  email varchar (50),
  phone varchar (11),
  gender int,
  birthday timestamp ,
  join_time date
);


create table article (
  article_id int primary key ,
  title varchar (30) not null ,
  content text ,
  article_time timestamp
);