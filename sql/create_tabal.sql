CREATE TABLE bootcamp.dbo.farm (
                                   id int IDENTITY(1,1) NOT NULL,
                                   character_id int NULL,
                                   check_point_x int NULL,
                                   check_point_y int NULL,
                                   plant_date datetimeoffset NULL,
                                   harvest_date datetimeoffset NULL,
                                   remaining_harvest int NULL,
                                   plant_dex_id int NULL,
                                   is_watered bit NULL,
                                   update_date datetimeoffset NULL,
                                   create_date datetimeoffset NULL
);


CREATE TABLE bootcamp.dbo.item (
                                   id int IDENTITY(1,1) NOT NULL,
                                   market_id int NULL,
                                   item_name varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                   price_per_unit int NULL,
                                   item_type varchar(10) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                   permit_currency int NULL,
                                   update_date datetimeoffset NULL,
                                   plant_description varchar(100) COLLATE Thai_CI_AS NULL
);

CREATE TABLE bootcamp.dbo.lottery (
                                      id int IDENTITY(1,1) NOT NULL,
                                      character_id int NULL,
                                      lottery_number varchar(3) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
                                      round_id int NOT NULL,
                                      update_date datetimeoffset NULL,
                                      create_date datetimeoffset NULL
);



CREATE TABLE bootcamp.dbo.[character] (
                                          id int IDENTITY(1,1) NOT NULL,
    login_id int NULL,
    gold int NULL,
    coin int NULL,
    gender int NULL,
    skin_id int NULL,
    hat_id int NULL,
    shirt_id int NULL,
    shoes_id int NULL,
    update_date datetimeoffset NULL,
    create_date datetimeoffset NULL
    );

CREATE TABLE bootcamp.dbo.inventory (
                                        id int IDENTITY(1,1) NOT NULL,
                                        character_id int NULL,
                                        item_id int NULL,
                                        quantity int NULL,
                                        update_date datetimeoffset NULL,
                                        create_date datetimeoffset NULL
);



CREATE TABLE bootcamp.dbo.plant_dex (
                                        id int IDENTITY(1,1) NOT NULL,
                                        item_id int NULL,
                                        state_id int NULL,
                                        state_name varchar(30) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                        hour_to_grow int NULL,
                                        plant_description varchar(255) COLLATE Thai_CI_AS NULL,
                                        plant_name varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                        plant_type int NULL,
                                        is_special_item bit NULL,
                                        update_date datetime2(7) NULL
);


CREATE TABLE bootcamp.dbo.login (
                                    id int IDENTITY(1,1) NOT NULL,
                                    login_uuid varchar(40) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                    username varchar(15) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                    password varchar(70) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                    email varchar(30) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                    register_date datetimeoffset NULL,
                                    last_login datetimeoffset NULL,
                                    update_date datetimeoffset NULL,
    [section] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
    exp_section datetimeoffset NULL,
    create_date datetimeoffset NULL
    );

CREATE TABLE bootcamp.dbo.buff (
                                   id int IDENTITY(1,1) NOT NULL,
                                   character_id int NULL,
                                   buff_name varchar(30) COLLATE Thai_CI_AS NULL,
                                   update_date datetimeoffset NULL,
                                   start_date datetimeoffset NULL,
                                   end_date datetimeoffset NULL,
                                   remaining int NULL,
                                   description varchar(255) COLLATE Thai_CI_AS NULL,
                                   value int NULL
);

CREATE TABLE bootcamp.dbo.market (
                                     id int IDENTITY(1,1) NOT NULL,
                                     market_name varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                     market_desc varchar(500) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
                                     update_date datetimeoffset NULL,
                                     CONSTRAINT PK__market__3213E83F627F0284 PRIMARY KEY (id)
);

CREATE TABLE bootcamp.dbo.item_pool (
                                        id int IDENTITY(1,1) NOT NULL,
                                        character_id int NULL,
                                        item_id int NULL,
                                        round_id int NULL,
                                        update_date datetimeoffset NULL,
                                        create_date datetimeoffset NULL
);
