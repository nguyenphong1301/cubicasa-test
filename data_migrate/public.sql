/*
 Navicat Premium Data Transfer

 Source Server         : Ama-Subscription-Local
 Source Server Type    : PostgreSQL
 Source Server Version : 130000
 Source Host           : 127.0.0.1:5432
 Source Catalog        : cubicasa_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130000
 File Encoding         : 65001

 Date: 07/01/2021 14:03:49
*/


-- ----------------------------
-- Sequence structure for hubs_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."hubs_id_seq";
CREATE SEQUENCE "public"."hubs_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for teams_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."teams_id_seq";
CREATE SEQUENCE "public"."teams_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for hubs
-- ----------------------------
DROP TABLE IF EXISTS "public"."hubs";
CREATE TABLE "public"."hubs" (
  "id" int8 NOT NULL DEFAULT nextval('hubs_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default" NOT NULL,
  "geo_location" text COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;

-- ----------------------------
-- Table structure for teams
-- ----------------------------
DROP TABLE IF EXISTS "public"."teams";
CREATE TABLE "public"."teams" (
  "id" int8 NOT NULL DEFAULT nextval('teams_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default" NOT NULL,
  "team_type" text COLLATE "pg_catalog"."default" NOT NULL,
  "hub_id" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "role" text COLLATE "pg_catalog"."default",
  "email" text COLLATE "pg_catalog"."default" NOT NULL,
  "team_id" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."hubs_id_seq"
OWNED BY "public"."hubs"."id";
SELECT setval('"public"."hubs_id_seq"', 2, false);
ALTER SEQUENCE "public"."teams_id_seq"
OWNED BY "public"."teams"."id";
SELECT setval('"public"."teams_id_seq"', 2, true);
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 4, true);

-- ----------------------------
-- Indexes structure for table hubs
-- ----------------------------
CREATE INDEX "idx_hubs_name" ON "public"."hubs" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table hubs
-- ----------------------------
ALTER TABLE "public"."hubs" ADD CONSTRAINT "hubs_name_key" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table hubs
-- ----------------------------
ALTER TABLE "public"."hubs" ADD CONSTRAINT "hubs_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table teams
-- ----------------------------
CREATE INDEX "idx_teams_name" ON "public"."teams" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_teams_team_type" ON "public"."teams" USING btree (
  "team_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table teams
-- ----------------------------
ALTER TABLE "public"."teams" ADD CONSTRAINT "teams_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE INDEX "idx_users_email" ON "public"."users" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_email_key" UNIQUE ("email");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
