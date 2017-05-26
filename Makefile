db:
	@docker run -d -p 3306:3306 --name "xml2db_db" imega/mysql
	@docker run --rm \
		-v $(CURDIR)/sql:/sql \
		--link xml2db_db:s \
		imega/mysql-client \
    	mysql --host=s -e "source /sql/schema.sql"

clean:
	@-docker stop xml2db_db
	@-docker rm -fv xml2db_db
