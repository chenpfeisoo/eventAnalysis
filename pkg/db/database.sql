create  table  event_info(
    event_name    varchar(32),
    image_name varchar(32),
    pullimage_time float
    ) ;
INSERT INTO event_info (event_name,image_name,pullimage_time)
VALUES ("test","nginx",5);

curl -H "Content-Type:application/json" -X POST -d '{"event_name":"qq","image_name":"qq","pullimage_time":5}' 'http://localhost:9999/show'
