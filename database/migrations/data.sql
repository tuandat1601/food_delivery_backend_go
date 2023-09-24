CREATE TABLE "Food"(
    "id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "image" TEXT NOT NULL,
    "price" INTEGER NOT NULL,
    "menuId" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" BIGINT NOT NULL
);
ALTER TABLE
    "Food" ADD PRIMARY KEY("id");
CREATE TABLE "OrderItem"(
    "id" BIGINT NOT NULL,
    "order_id" BIGINT NOT NULL,
    "food_id" BIGINT NOT NULL,
    "quantity" INTEGER NOT NULL,
    "total_price" DOUBLE PRECISION NOT NULL
);
ALTER TABLE
    "OrderItem" ADD PRIMARY KEY("id");
CREATE TABLE "users"(
    "id"SERIAL PRIMARY KEY ,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NOT NULL,
    "is_deleted" BOOLEAN DEFAULT false,
    "created_at"TIMESTAMPTZ DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ DEFAULT NOW()
);
ALTER TABLE
    "User" ADD PRIMARY KEY("id");
CREATE TABLE "restaurants"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "user_id" BIGINT NOT NULL,
    "day_of_week" VARCHAR(255) NOT NULL,
    "opening_time" VARCHAR(255) NOT NULL,
    "closing_time" VARCHAR(255) NOT NULL,
    "is_deleted" BOOLEAN DEFAULT false,
    "created_at"TIMESTAMPTZ DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ DEFAULT NOW()
);
ALTER TABLE
    "Restaurants" ADD PRIMARY KEY("id");
CREATE TABLE "Menus"(
    "id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "restaurantID" BIGINT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_ad" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "Menus" ADD PRIMARY KEY("id");
CREATE TABLE "Orders"(
    "id" BIGINT NOT NULL,
    "customer_id" BIGINT NOT NULL,
    "order_time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "Orders" ADD PRIMARY KEY("id");
ALTER TABLE
    "Menus" ADD CONSTRAINT "menus_restaurantid_foreign" FOREIGN KEY("restaurantID") REFERENCES "Restaurants"("id");
ALTER TABLE
    "Orders" ADD CONSTRAINT "orders_id_foreign" FOREIGN KEY("id") REFERENCES "User"("id");
ALTER TABLE
    "OrderItem" ADD CONSTRAINT "orderitem_food_id_foreign" FOREIGN KEY("food_id") REFERENCES "Food"("id");
ALTER TABLE
    "OrderItem" ADD CONSTRAINT "orderitem_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "Orders"("id");
ALTER TABLE
    "Food" ADD CONSTRAINT "food_menuid_foreign" FOREIGN KEY("menuId") REFERENCES "Menus"("id");