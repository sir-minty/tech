Makeing keys for SSL mode

	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem

Connecting to MySQL that has been brought up by docker-compose

	# List the processes brought up by docker compose and you should see your db name
	$ docker-compose ps
	    Name                 Command             State            Ports
	----------------------------------------------------------------------------
	tech_data_1    /bin/sh                       Exit 0
	tech_mysql_1   docker-entrypoint.sh mysqld   Up       3306/tcp
	tech_web_1     app                           Up       0.0.0.0:3000->3000/tcp

	# take that and run it through docker
	$ docker exec -it tech_mysql_1 mysql -uroot -p
