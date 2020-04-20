/*
 Navicat Premium Data Transfer

 Source Server         : mac_local
 Source Server Type    : PostgreSQL
 Source Server Version : 120002
 Source Host           : localhost:5432
 Source Catalog        : test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 120002
 File Encoding         : 65001

 Date: 20/04/2020 10:52:31
*/


-- ----------------------------
-- Sequence structure for user_uid_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_uid_seq";
CREATE SEQUENCE "public"."user_uid_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_uid_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for topic_classes
-- ----------------------------
DROP TABLE IF EXISTS "public"."topic_classes";
CREATE TABLE "public"."topic_classes" (
  "class_id" int4,
  "class_name" text COLLATE "pg_catalog"."default",
  "class_remark" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."topic_classes" OWNER TO "postgres";

-- ----------------------------
-- Records of topic_classes
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for topics
-- ----------------------------
DROP TABLE IF EXISTS "public"."topics";
CREATE TABLE "public"."topics" (
  "topic_id" int4 NOT NULL,
  "topic_title" text COLLATE "pg_catalog"."default",
  "topic_short_title" text COLLATE "pg_catalog"."default",
  "user_ip" text COLLATE "pg_catalog"."default",
  "topic_score" int4,
  "topic_url" text COLLATE "pg_catalog"."default",
  "topic_date" timestamptz(6)
)
;
ALTER TABLE "public"."topics" OWNER TO "postgres";

-- ----------------------------
-- Records of topics
-- ----------------------------
BEGIN;
INSERT INTO "public"."topics" VALUES (1, 'title 1', 's title 1', '127.0.0.1', 8, 'baidu.com 1', '2020-04-17 18:02:48+08');
INSERT INTO "public"."topics" VALUES (2, 'title 2', 's title 2', '127.0.0.1', 9, 'baidu.com 2', '2020-04-17 18:02:48+08');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "uid" int4 NOT NULL DEFAULT nextval('user_uid_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default",
  "age" int4,
  "birthday" timestamptz(6),
  "user_name" text COLLATE "pg_catalog"."default" NOT NULL,
  "pass_word" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."user" OWNER TO "postgres";

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_uid_seq"
OWNED BY "public"."user"."uid";
SELECT setval('"public"."user_uid_seq"', 2, false);

-- ----------------------------
-- Primary Key structure for table topics
-- ----------------------------
ALTER TABLE "public"."topics" ADD CONSTRAINT "topics_pkey" PRIMARY KEY ("topic_id");

-- ----------------------------
-- Uniques structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_user_name_key" UNIQUE ("user_name");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("uid");
