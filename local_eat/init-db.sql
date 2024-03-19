create table USER
(USERNAME varchar(20) not null,
 PASSWORD varchar(64) not null, -- SHA-256 produit un hash de 64 caractères hexadécimaux
 MAIL varchar(50) not null,
 primary key (USERNAME));
 
 create table ADDRESS
(ID char(5) not null,
 STREET varchar(30) not null,
 STREET_NUMBER int not null,
 BOX_NUMBER varchar(5) null,
 POSTCODE int not null,
 LOCALITY varchar(20) not null,
 COUNTRY varchar(20) not null,
 primary key (ID));
 
create table PRODUCER
(ID char(7) not null,
 USERNAME varchar(20) not null,
 FIRSTNAME char(20) not null,
 LASTNAME char(20) not null,
 ADDRESS char(5) not null,
 PHONE_NUM char(10) not null,
 primary key (ID),
 foreign key (USERNAME) references USER(USERNAME),
 foreign key (ADDRESS) references ADDRESS(ID));
 
create table COMPANY
(COMPANY_NAME varchar(20) not null,
 PASSWORD varchar(64) not null, -- SHA-256 produit un hash de 64 caractères hexadécimaux
 ALIAS varchar(50) not null,
 ADDRESS char(5) not null,
 MAIL varchar(50) not null,
 PHONE_NUM char(10) not null,
 VAT_NUM char(12) not null,
 DESCRIPTION longtext null,
 primary key (COMPANY_NAME),
 foreign key (ADDRESS) references ADDRESS(ID));
  
create table REL_COMP_PROD
(PRODUCER_ID char(7) not null,
 COMPANY_NAME varchar(20) not null,
 primary key (PRODUCER_ID, COMPANY_NAME),
 foreign key (PRODUCER_ID) references PRODUCER(ID),
 foreign key (COMPANY_NAME) references COMPANY(COMPANY_NAME));