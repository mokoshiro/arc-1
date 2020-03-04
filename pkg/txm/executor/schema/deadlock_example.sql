start transaction;

insert into peer('peer_id', 'addr', 'credential', 'location') VALUES ('a', 'xxxx', 'yyyy', ST_GeomFromText('POINT(125.0 35.0)'));

commit;

update peer set location = ST_GeomFromText(POINT(111.0 11.0)) where peer_id = 'a';