INSERT INTO "crud_sql"."companies" ("id", "name", "cnpj", "created_at", "updated_at", "deleted_at")
VALUES (DEFAULT, 'C1', '1', DEFAULT, DEFAULT, DEFAULT)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."companies" ("id", "name", "cnpj", "created_at", "updated_at", "deleted_at")
VALUES (DEFAULT, 'C2', '2', DEFAULT, DEFAULT, DEFAULT)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."users" ("id", "name", "email", "rg", "cpf", "created_at", "updated_at", "deleted_at")
VALUES (DEFAULT, 'Rodrigo 3', '1', '1', '1', DEFAULT, DEFAULT, DEFAULT)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."users" ("id", "name", "email", "rg", "cpf", "created_at", "updated_at", "deleted_at")
VALUES (DEFAULT, 'Rodrigo 3', '2', '2', '2', DEFAULT, DEFAULT, DEFAULT)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."users" ("id", "name", "email", "rg", "cpf", "created_at", "updated_at", "deleted_at")
VALUES (DEFAULT, 'Rodrigo 3', '3', '3', '3', DEFAULT, DEFAULT, DEFAULT)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."jobs" ("id",user_id, company_id, month_salary, hours_per_day)
VALUES (DEFAULT, '1', '1', 1, 1)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."jobs" ("id",user_id, company_id, month_salary, hours_per_day)
VALUES (DEFAULT, '1', '2', 1, 1)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."jobs" ("id",user_id, company_id, month_salary, hours_per_day)
VALUES (DEFAULT, '2', '1', 1, 1)
RETURNING "id", "updated_at", "deleted_at";

INSERT INTO "crud_sql"."jobs" ("id",user_id, company_id, month_salary, hours_per_day)
VALUES (DEFAULT, '2', '2', 1, 1)
RETURNING "id", "updated_at", "deleted_at";