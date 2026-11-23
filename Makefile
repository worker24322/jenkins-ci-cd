.PHONY: run test up down jenkins-up jenkins-down jenkins-password

# Chạy thẳng bằng Go, không cần Docker.
run:
	go run .

test:
	go vet ./... && go test ./...

# Chạy bằng Docker.
up:
	docker compose up --build -d

down:
	docker compose down

# Jenkins tại http://localhost:8080
jenkins-up:
	docker compose -f jenkins/docker-compose.yml up --build -d

jenkins-down:
	docker compose -f jenkins/docker-compose.yml down

# In mật khẩu dùng cho lần mở Jenkins đầu tiên.
jenkins-password:
	docker compose -f jenkins/docker-compose.yml exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
