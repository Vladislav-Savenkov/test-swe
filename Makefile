.PHONY: test clean logs

test:
	docker-compose up --build --abort-on-container-exit

clean:
	rm -rf test_logs
	docker-compose down --volumes --remove-orphans

logs:
	cat test_logs/test_output.txt
