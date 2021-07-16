#scrip เพิ่มข้อมูลในdatabase
ภาษาที่ใช้คือ sql
--ในการจะทดสอบบฟังก์ชั่นต่างๆอย่างน้อย  table item,item_pool,lottery,plant_dex,market ใน database ต้องมีข้อมูล


#scrip เสกของ ของ inventory ให้ characterID มีครบทุก item(ในปัจจุบัน) 
```
DECLARE  @characterID INT
DECLARE  @quantity INT
SET @characterID = 51
SET @quantity = 99
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 1, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 2, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 3, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 4, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 5, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 6, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 7, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 8, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 9, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 10, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 11, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 12, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 13, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 14, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 15, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 16, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 17, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 18, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 19, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 20, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 21, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 22, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 23, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 24, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 25, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 26, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 27, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 28, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
INSERT INTO inventory
(character_id, item_id, quantity, update_date, create_date)
VALUES(@characterID, 29, @quantity, N'2021-07-04 18:37:05 +00:00', N'0001-01-01 00:00:00 +00:00');
```
#scrip เสกตัง  ใน table character

```
DECLARE  @characterID INT
DECLARE  @gold INT
DECLARE  @coin INT
SET @characterID = 72 //ใส่ characterIdที่จะเสกตัง
SET @gold = 2000
SET @coin = 2000
UPDATE bootcamp.dbo.[character]
SET gold=@gold, coin=@coin
WHERE id =@characterID
 ```

#scrip ลดวัน ใน table farm


```
DECLARE  @characterID INT
DECLARE     @date_back INT
SET @characterID = 72 //ใส่ characterIdที่จะลดวันปลูก ใน farm
SET @date_back = 1 // ใส่ จำนวนวันที่จะลด
UPDATE bootcamp.dbo.farm
SET plant_date = CAST( GETDATE()-@date_back AS Date )
WHERE character_id = @characterID
```