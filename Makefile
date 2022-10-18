test:
	ENV=TEST go test ./...

run:
	ENV=DEV go run ./cmd/main.go

deploy:
	terraform -chdir=terraform apply

destroy:
	terraform -chdir=terraform destroy

