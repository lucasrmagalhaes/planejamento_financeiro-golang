build-image:
	docker build -t lucasrmagalhaes/planejamento_financeiro-golang .

run-app:
	docker-compose -f devops/app.yml up -d

lint:
	golint ./...
	go fmt -n ./...

unit_test:
	go test ./...