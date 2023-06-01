

CREATE TABLE `customer` (
  `id` VARCHAR(40) NOT NULL,
  `address` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY(`id`)
);

INSERT INTO `customer`
(id, address, city, state)
VALUES('M1-ABDI-11', 'Taman Sari, Kec. Setu', 'Bekasi', 'Jawa Barat');
INSERT INTO `customer`
(id, address, city, state)
VALUES('M1-PUMS-11', 'Gunung Sindur', 'Bogor', 'Jawa Barat');
INSERT INTO `customer`
(id, address, city, state)
VALUES('M1-STQI-11', 'Kel. Rorotan- Cilincing', 'Jakarta Utara', 'DKI Jakarta');
INSERT INTO `customer`
(id, address, city, state)
VALUES('M1-SUGP-11', 'Gn. Putri Kec. Gn. Putri', 'Bekasi', 'Jawa Barat');
INSERT INTO `customer`
(id, address, city, state)
VALUES('M1-SURA-11', 'Padurenan, Mustika Jaya', 'Bekasi', 'Jawa Barat');


CREATE TABLE `supplier` (
  `id` VARCHAR(40) NOT NULL,
  `address` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY(`id`)
);

INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-KPS-1', 'Kelurahan Klender, Kecamatan Duren Sawit', 'Jakarta Timur', 'DKI Jakarta');
INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-HSC-1', 'Mangga Dua Sel., Kecamatan Sawah Besar', 'Jakarta Pusat', 'DKI Jakarta');
INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-PSB-1', 'Tubagus Angke', 'Jakarta', 'DKI Jkarata');
INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-SUM-1', 'Kedung Waringin, Kec. Tanah Sereal', 'Kota Bogor', 'Jawa Barat');
INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-ISB-1', 'Penjaringan', 'Jakarta Utara', 'DKI Jakarta');
INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-FIX-1', 'Godangdia', 'Jakarta Pusat', 'DKI Jakarta');
INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-SAM-1', 'Bantar Gebang', 'Kota Bekasi', 'Jawa Barat');
INSERT INTO `supplier`
(id, address, city, state)
VALUES('S1-SSC-1', 'Kec. Sawah Besar', 'Jakarta Pusat', 'DKI Jakarta');

CREATE TABLE `logistics` (
  `id` VARCHAR(40) NOT NULL,
  `type` VARCHAR(40) NOT NULL,
  `capacity` INT NOT NULL,
  `cost` BIGINT NOT NULL,
  `city` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY(`id`)
);


INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('fe9deaaa-446f-4877-9d29-79035cbcbeb7','Fuso',8,1600000,'Jakarta Timur','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('244e038b-5c61-4f32-89c2-7f7a677d8a34','Fuso',8,1600000,'Jakarta Barat','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('d00a77c1-9508-4631-a5fc-f85e13cc775d','Fuso',8,1600000,'Jakarta Utara','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('00b79dc4-884a-49bb-a7c9-3e25d50cb099','Fuso',8,1600000,'Jakarta Selatan','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('febdcc23-bd87-4b55-b2af-6d85e02d46de','Fuso',8,1800000,'Bekasi','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('04cef342-688d-49db-9b3e-5287a487d405','Fuso',8,1800000,'Depok','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('66ed623b-7f49-48fd-949c-b3288da3c609','Fuso',8,1800000,'Cikarang','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('56aa2de1-818a-4beb-b6cc-65205a9695bc','Fuso',8,1900000,'Gunung Sindur','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('0ba7844d-507d-478f-aea5-eb64bbe57cf1','Fuso',8,1800000,'Tangerang','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('c4260add-3675-47b8-a8ce-04a87056e011','Tronton',22,2500000,'Jakarta Timur','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('2df50801-56be-4438-8479-06b38f7f4775','Tronton',22,2500000,'Jakarta Barat','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('d9698db9-d683-4cc7-a49c-c121cc0eea87','Tronton',22,2500000,'Jakarta Utara','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('68dde84c-b15a-4daf-afb7-95b0ed4056b6','Tronton',22,2500000,'Jakarta Selatan','DKI Jakarta');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('576b622f-b746-4ded-a19f-7b3df3d18911','Tronton',22,3450000,'Bekasi','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('1cd8abc9-2180-4e3e-b38a-fba08d6e4931','Tronton',22,3450000,'Depok','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('31ea5ccc-1f99-42b1-b0da-b56e948dcb1f','Tronton',22,3450000,'Cikarang','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('88ffc7ff-85dc-4590-a474-349b9cff027b','Tronton',22,3600000,'Gunung Sindur','Jawa Barat');
INSERT INTO `logistics`
(id, type, capacity, cost, city, state)
VALUES('31c10fdb-a180-428b-a4d1-f991a5507481','Tronton',22,3450000,'Tangerang','Jawa Barat');

CREATE TABLE `product_pricelist` (
  `id` VARCHAR(40) NOT NULL,
  `supplier_id` VARCHAR(40) NOT NULL,
  `sku` VARCHAR(40) NOT NULL,
  `price` BIGINT NOT NULL,
  `stock` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY(`id`)
);


INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('46b3e61e-ff95-4659-b092-73ab75ed1b4e', 'S1-FIX-1', 'PLT-SPHC1000', '5596216.21621622', '1');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('b1fc655d-44f6-4fe2-bdcf-1eea418e7ccb', 'S1-FIX-1', 'PLT-SPHC0400', '1789729.72972973', '3');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('0a1ce3ac-30f0-4594-91f7-b7e9664e8abb', 'S1-FIX-1', 'PLT-SPHC0500', '2390540.54054054', '5');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('6bc2ea83-be89-49a3-a780-f42d2d390cc4', 'S1-FIX-1', 'PLT-SPHC0600', '3008783.78378378', '4');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('0b83e419-4d59-49c3-8e34-492a715953d9', 'S1-FIX-1', 'PLT-SPHC1200-GG', '3640765.76576577', '11');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('e1e0e0cf-bcc2-40b0-868f-8c9ecfac861e', 'S1-FIX-1', 'SIK-100100-IBB', '1066572.97297297', '30');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('2f220d6b-de18-4db9-8ddf-588f95be0fba', 'S1-FIX-1', 'SIK-120120-IBB', '1540090.09009009', '5');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('bcf23706-a706-4aac-b042-873bdcbc2c0f', 'S1-FIX-1', 'SIK-040040-IBB', '166129.72972973', '51');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('771c89cb-0a59-4d38-9883-a0ddb487a77a', 'S1-FIX-1', 'SIK-050050-IBB', '259720.720720721', '72');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('5e1669a7-6af9-4a5b-bc99-1ffc46ce3a5e', 'S1-FIX-1', 'SIK-060060-IBB', '372075.675675676', '8');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('f7304589-1dd5-4d56-8f86-c3483ee4d73e', 'S1-FIX-1', 'SIK-070070-IBB', '506627.027027027', '8');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('eefb162c-6b32-4b04-b8a1-1e27d61d44e1', 'S1-HSC-1', 'PIP-SCH404', '1559459.45945946', '3');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('b209c621-defb-4442-9bca-653e0516d2cd', 'S1-HSC-1', 'PIP-SCH4060', '2727027.02702703', '3');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('bd547ab2-4cf6-4394-bc9e-6a674263287a', 'S1-HSC-1', 'PIP-SCH4080', '4163063.06306306', '2');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('de86e42b-9d4e-47a0-82fc-766c00998990', 'S1-HSC-1', 'PLT-SPHC1200-GG', '3544144.14414414', '11');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('49d2331b-8191-4bab-a856-2625a563d8dc', 'S1-HSC-1', 'SIK-100100-IBB', '990990.990990991', '30');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('e2ffd305-b48c-422b-86d1-73425e4bd87f', 'S1-HSC-1', 'SIK-120120-IBB', '1447747.74774775', '30');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('61326338-aef4-4380-a3c2-868066622caf', 'S1-HSC-1', 'SIK-040040-IBB', '143694', '51');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('99adecfe-db05-4447-83b1-21aa9ac87790', 'S1-HSC-1', 'SIK-060060-IBB', '324775', '112');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('3ae60fce-d76d-43f1-966c-a43ca00ada1d', 'S1-HSC-1', 'SIK-070070-IBB', '443243', '24');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('84f05cfd-e3b1-473e-a5db-91fc0a15f45e', 'S1-HSC-1', 'SIK-080080-IBB', '633783.783783784', '24');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('cc82c8b4-792d-4cf3-bc6b-1c634efc2c7f', 'S1-HSC-1', 'SIK-090090-KS', '804504.504504504', '52');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('66194372-38d9-4766-9c01-d78a392fc85a', 'S1-HSC-1', 'UNP-200', '2266667', '6');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('662930dd-fcd4-4226-9228-38882231375a', 'S1-HSC-1', 'PLT-SPHC0155', '518919', '50');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('6e816ee9-26db-4e85-a944-e354060bd03b', 'S1-HSC-1', 'PLT-SPHC0180', '582702.702702703', '50');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('701c46db-bd4b-41b4-aaf7-268f82a1f023', 'S1-HSC-1', 'UNP-120', '1968468.46846847', '10');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('378f9983-7756-4f06-b7c3-08596fc9bad8', 'S1-ISB-1', 'UNP-200', '4387387.38738739', '6');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('bb6631dd-a858-4ccc-b553-a88e02f69990', 'S1-ISB-1', 'PLT-BRDS0230', '828828.828828829', '100');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('0ccd99a2-168c-49f5-956f-757ecbc89b7b', 'S1-ISB-1', 'PLT-SPHC0120', '400900.900900901', '170');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('d9a90bb5-6d9a-44e1-b492-eebfc5df7b7e', 'S1-KPS-1', 'PLT-SPHC0150', '504504.504504504', '20');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('93ca4f0a-9c1a-4180-81a6-80b86910f0be', 'S1-KPS-1', 'PLT-SPHC0200', '627027.027027027', '15');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('bc8d55d9-1758-4023-a83c-00f9231a5566', 'S1-PSB-1', 'PLT-SPHC0026', '788288.288288288', '10');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('c6813d8d-e510-4745-8d53-af787b0296cd', 'S1-PSB-1', 'PLT-SPHC0280', '855855.855855856', '20');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('de8d5ac2-0f06-4508-b129-e43831a62d8b', 'S1-PSB-1', 'SIK-100100-IBB', '1576767.56756757', '30');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('7e165d19-7316-41cb-99d3-9bae2cc84b6b', 'S1-PSB-1', 'SIK-120120-IBB', '2366674.77477477', '12');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('d0cf3630-6982-48c2-b894-cb4e4695f17d', 'S1-SAM-1', 'SIK-040040-IBB', '151577.477477477', '51');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('eec03c22-468e-410e-8c09-b32ae7dc115f', 'S1-SAM-1', 'SIK-050050-IBB', '475662.162162162', '72');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('f67e5140-4db7-40b1-bd8a-dce5ef572451', 'S1-SSC-1', 'SIK-060060-IBB', '687067.567567568', '112');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('9be365df-5360-485f-84c3-251077ad0f8f', 'S1-SSC-1', 'SIK-070070-IBB', '468262.162162162', '124');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('80f6e4b1-86be-444b-b267-b3b8831536c5', 'S1-SSC-1', 'SIK-080080-IBB', '1233198.1981982', '20');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('9cd0cabb-120e-4e9c-98ff-094ad678d59d', 'S1-SSC-1', 'WFL-300-GG', '7040000.', '8');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('bb9b4ec2-1f0f-411b-8720-e6f2ab6bbbe1', 'S1-SSC-1', 'PLT-KPL0100', '9864865', '2');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('5f03c5be-6a7d-4081-b2d5-a5234bed4ad6', 'S1-SSC-1', 'SIK-050050-IBB', '189189', '72');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('6e7fe4c3-d2b8-4d12-b82d-525a528735ab', 'S1-SSC-1', 'SIK-120120-IBB', '1261261.26126126', '5');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('b5e7d432-d810-4111-b5ad-79608bfada1d', 'S1-SSC-1', 'SIK-040040-IBB', '144144.144144144', '51');

INSERT INTO `product_pricelist`
(id, supplier_id, sku, price, stock)
VALUES('bd3b4859-080d-4aef-af6b-2ab4d29bd1c3', 'S1-SUM-1', 'SIK-050050-IBB', '189189.189189189', '72');


CREATE TABLE `purchase_order` (
  `id` VARCHAR(40) NOT NULL,
  `customer_id` VARCHAR(40) NOT NULL,
  `sku_id` VARCHAR(40) NOT NULL,
  `order_date` TIMESTAMP NOT NULL,
  `order_qty` BIGINT NOT NULL,
  `order_unit` varchar(255) NOT NULL,
  `unit_selling_price` BIGINT NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY(`id`)
);

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('1e79cf9b-dccc-43c7-b1bb-4e889a93bd37','M1-PUMS-11','SIK-040040-IBB','2022-12-20','142','Batang','143694');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('ce7d1ecd-c64c-45cc-9624-1ddd984c302d','M1-PUMS-11','SIK-060060-IBB','2022-12-20','112','Batang','324775');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('f62686dc-8cab-4c40-a3c4-ce61930511f8','M1-PUMS-11','SIK-070070-IBB','2022-12-20','24','Batang','443243');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('ba22ac12-901b-476f-bb52-958371e66347','M1-PUMS-11','SIK-080080-IBB','2022-12-20','24','Batang','581982');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('63ad5ebb-5c10-4afb-9c7c-f07019ba86c5','M1-PUMS-11','SIK-090090-KS','2022-12-20','52','Batang','738739');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('199a55d0-f907-421a-96c1-e663d66bd0db','M1-PUMS-11','SIK-100100-IBB','2022-12-20','90','Batang','910360');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('2fbb0f30-31b6-490b-9617-b44f851095d6','M1-STQI-11','PLT-BRDS0230','2023-01-17','100','Lembar','902832');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('7b98ccc7-e88a-4177-b110-f52dd5c6a38f','M1-STQI-11','PLT-SPHC0155','2023-01-17','200','Lembar','502978');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('3023d2fb-8704-4708-824e-d6b28cbaa865','M1-STQI-11','PLT-SPHC0180','2023-01-17','50','Lembar','576828');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('4b4fbf3b-448e-4c9d-be60-7f50c0c6baae','M1-STQI-11','PLT-SPHC0026','2023-01-19','10','Lembar','831644');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('68b5c1ac-e531-4a3d-96a3-2772dc37e2b2','M1-STQI-11','PLT-SPHC0280','2023-01-19','20','Lembar','902927');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('dd7a7a9e-4a2c-4ba7-9a84-c0eb2fafdb6d','M1-PUMS-11','PIP-SCH404','2023-01-20','3','Batang','1557900');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('a9149455-62bc-43de-88c9-8156223e1398','M1-PUMS-11','PIP-SCH4060','2023-01-20','3','Batang','2724300');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('aeee5714-14db-4b9c-a931-7d03df8bc1fd','M1-PUMS-11','PIP-SCH4080','2023-01-20','2','Batang','4158900');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('4c1d14fc-1db0-49f1-a22a-613c44bb7a0f','M1-PUMS-11','PLT-SPHC0400','2023-01-20','3','Lembar','1782900');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('71a353d1-eb54-4fc1-8dfc-b43d7dcf50ba','M1-PUMS-11','PLT-SPHC0500','2023-01-20','5','Lembar','2364570');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('0d6c926a-cd0c-4b47-ac41-b89da387fe35','M1-PUMS-11','PLT-SPHC0600','2023-01-20','4','Lembar','2946285');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('bdefda2c-21fe-434f-9663-bb8932ffc98d','M1-PUMS-11','PLT-SPHC1000','2023-01-20','1','Lembar','4766940');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('f9962ea8-04e5-4df6-9368-f439e6cd75ae','M1-PUMS-11','PLT-SPHC1200-GG','2023-01-20','11','Lembar','3540600');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('b11852be-17b6-41a7-b9bf-0447a752a17b','M1-PUMS-11','PLT-SPHC1200-GG','2023-01-20','14','Lembar','3540600');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('74a43b49-913b-4c7d-9e10-9fe004cb27ed','M1-PUMS-11','SIK-040040-IBB','2023-01-20','51','Batang','156600');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('8151fd38-707e-48cc-a53d-b84371ea1774','M1-PUMS-11','SIK-050050-IBB','2023-01-20','72','Batang','246600');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('dda3c2e5-f530-4b74-8082-6371a4b60309','M1-PUMS-11','SIK-060060-IBB','2023-01-20','8','Batang','353700');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('4bb3c7fc-f802-419d-b136-ce6b7173a92e','M1-PUMS-11','SIK-070070-IBB','2023-01-20','8','Batang','482400');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('7fba2d36-db0e-469d-bf2b-a7bc0344c701','M1-PUMS-11','SIK-080080-IBB','2023-01-20','10','Batang','633150');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('8e25e571-f88e-4782-b36a-0908f3207284','M1-PUMS-11','SIK-090090-KS','2023-01-20','17','Batang','803700');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('80a1dd6b-dc57-445f-82a7-8dbfcd249517','M1-PUMS-11','SIK-100100-IBB','2023-01-20','1','Batang','990000');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('da59629b-85f0-4144-a973-c854886fe1b1','M1-PUMS-11','SIK-100100-IBB','2023-01-20','18','Batang','990000');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('b0064beb-9dce-47dd-a1de-45e07a5fbfbe','M1-PUMS-11','SIK-120120-IBB','2023-01-20','5','Batang','1446300');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('570452ab-8cf6-49fc-9859-4b5a72e9d874','M1-PUMS-11','SIK-120120-IBB','2023-01-20','20','Batang','1446300');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('004d9eb8-0c88-4af5-99eb-8a64e8e2c0e7','M1-STQI-11','PLT-SPHC0280','2023-01-20','20','Lembar','902927');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('51a2980a-47b4-4300-8d99-e8feda5aec20','M1-SUGP-11','SIK-040040-IBB','2023-01-27','30','Batang','162973');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('98595c78-0641-4af0-a728-b0f4b27d68d5','M1-SUGP-11','SIK-050050-IBB','2023-01-27','60','Batang','514147');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('57199cbb-ad4d-4881-88f9-4edc99696d06','M1-SUGP-11','SIK-060060-IBB','2023-01-27','50','Batang','737809');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('c5d5df9d-cdd7-4969-ae2c-6f9cecd242eb','M1-SUGP-11','SIK-070070-IBB','2023-01-27','120','Batang','503237');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('d640282c-7793-47a2-9306-3a0f3557d36d','M1-SUGP-11','SIK-080080-IBB','2023-01-27','15','Batang','1322680');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('7edacab2-d419-452e-9303-b07a62c4198d','M1-SUGP-11','SIK-100100-IBB','2023-01-27','30','Batang','1674972');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('f787d98f-acf5-40b3-83c6-7139fde83a26','M1-SUGP-11','SIK-120120-IBB','2023-01-27','11','Batang','2542785');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('d6f11b0c-e982-49c3-8635-0177887926e8','M1-SUGP-11','WFL-300-GG','2023-01-27','8','Batang','7706558');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('40674506-b785-475f-968b-dcad3f686b31','M1-ABDI-11','UNP-200','2023-02-01','20','Batang','2162162');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('e9436210-2a18-4715-bebe-9a298fcc3751','M1-SURA-11','PLT-SPHC0120','2023-02-02','170','Lembar','432072');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('4519d78b-6c23-4bb1-a328-a685a374bb0c','M1-SURA-11','PLT-SPHC0150','2023-02-02','20','Lembar','531261');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('f9d0bf2e-3c62-47f7-b79b-10414aa5f042','M1-SURA-11','PLT-SPHC0200','2023-02-02','15','Lembar','634775');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('b8ce2567-6bc8-4489-9b28-d95e9e09d7b5','M1-STQI-11','PLT-SPHC0155','2023-02-03','50','Lembar','502978');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('50f69bd4-fc51-4a41-a369-5a53325ee0e1','M1-ABDI-11','UNP-120','2023-02-06','6','Batang','1968468.46846847');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('53869daa-3cf9-4805-862c-bd051a7a79e5','M1-ABDI-11','UNP-200','2023-02-06','10','Batang','4533333.33333333');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('979ee9e7-2578-4435-a276-04f3b1667e8c','M1-SUGP-11','PLT-KPL0100','2023-02-06','2','Batang','10333445');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('4f640801-5003-40b9-8e5d-00034f549133','M1-SUGP-11','SIK-040040-IBB','2023-02-06','100','Batang','150990');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('135db044-c64e-4353-818d-e3339a05f760','M1-SUGP-11','SIK-050050-IBB','2023-02-06','30','Batang','198175');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('5f3920df-00e2-4b99-8356-e10e10f0653a','M1-SUGP-11','SIK-120120-IBB','2023-02-06','3','Batang','1259831');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('6e209a28-d942-45f4-a735-ee7e8af8f030','M1-SUGP-11','SIK-120120-IBB','2023-02-06','22','Batang','1321171');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('335e1e62-9d90-423a-a624-a5d27352ab67','M1-SUGP-11','PLT-KPL0100','2023-02-23','2','Lembar','10333445');

INSERT INTO `purchase_order`
(id, customer_id, sku_id, order_date, order_qty, order_unit,unit_selling_price)
VALUES('dc5922ba-9062-46ed-b047-880fc810dd0a','M1-SUGP-11','SIK-050050-IBB','2023-02-23','30','Batang','198175');
