CREATE TABLE "user_api" (
  "id" int PRIMARY KEY AUTOINCREMENT NOT NULL,
  "username" varchar,
  "password" varchar
);

CREATE TABLE "access_token" (
  "id" int KEY AUTOINCREMENT NOT NULL,
  "id_user_api" int,
  "token" text,
  "expired_at" datetime
);

CREATE TABLE "validation" (
  "id" int PRIMARY KEY AUTOINCREMENT NOT NULL,
  "id_user_api" int,
  "id_biometria" int,
  "hash" text
);

CREATE TABLE "biometria" (
  "id" int PRIMARY KEY AUTOINCREMENT NOT NULL,
  "rg" varchar,
  "cpf" int
);

CREATE TABLE "stack_frame" (
  "id" int PRIMARY KEY AUTOINCREMENT NOT NULL,
  "id_image" int,
  "id_biometria" int
);

CREATE TABLE "image" (
  "id" int PRIMARY KEY AUTOINCREMENT NOT NULL,
  "base64" text,
  "path" text
);

ALTER TABLE "access_token" ADD FOREIGN KEY ("id_user_api") REFERENCES "user_api" ("id");

ALTER TABLE "validation" ADD FOREIGN KEY ("id_user_api") REFERENCES "user_api" ("id");

ALTER TABLE "validation" ADD FOREIGN KEY ("id_biometria") REFERENCES "biometria" ("id");

ALTER TABLE "stack_frame" ADD FOREIGN KEY ("id_image") REFERENCES "image" ("id");

ALTER TABLE "stack_frame" ADD FOREIGN KEY ("id_biometria") REFERENCES "biometria" ("id");
