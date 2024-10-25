INSERT INTO mst_supliyer (id_supliyer, name_supliyer, balance) VALUES
(uuid_generate_v4(), 'All Operator', 1000000);

INSERT INTO mst_product (id_product, name_provider, nominal, price, id_supliyer) VALUES
(uuid_generate_v4(), 'Indosat', 10000, 10500, 'supliyer_uuid_value_here'),
(uuid_generate_v4(), 'Telkomsel', 10000, 10900, 'supliyer_uuid_value_here'),
(uuid_generate_v4(), 'Xl', 10000, 10700, 'supliyer_uuid_value_here'),
(uuid_generate_v4(), 'Tri', 10000, 10400, 'supliyer_uuid_value_here');