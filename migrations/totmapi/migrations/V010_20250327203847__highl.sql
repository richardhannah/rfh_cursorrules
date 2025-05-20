create table if not exists totm.shop_stock
(
    shopid varchar(36) not null,
    name                text not null,
    description         text not null,
    encumbrance         int not null,
    unit                varchar(2) not null,
    quantity_available  int not null,
    category            text not null
);
