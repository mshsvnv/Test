create_db_command := psql -U postgres -d tests -f ./sql/create.sql
delete_db_command := psql -U postgres -d tests -f ./sql/drop.sql

create_db:
	$(create_db_command)

delete_db:
	$(delete_db_command)

test:
	rm -rf ../allure-results
	export ALLURE_OUTPUT_PATH="/c/Users/mshsv/OneDrive/Рабочий стол/studying/7_sem/Test" && go test -v ./...