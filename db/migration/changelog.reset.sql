--liquibase formatted sql

--changeset Ecomm:reset runInTransaction:false runAlways:true validCheckSum:any
DROP DATABASE IF EXISTS ecommerce;
CREATE DATABASE ecommerce;

