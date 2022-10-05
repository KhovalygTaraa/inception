DCKER_COMPOSE = $(shell command -v docker-compose)
H=\033[0;36mHELP\033[0m
VOLUMES_LIST=$(shell command docker volume ls -q)

.PHONY: upload update remote_build help all up down total_delete remote_down remote_up remote_system_prune
all: up
ifndef $(DOCKER_COMPOSE)
	@docker-compose -v
else
	echo "install docker-compose"
endif

up:
	@mkdir -p $(HOME)/data/wordpress
	@mkdir -p $(HOME)/data/adminer
	@mkdir -p $(HOME)/data/mariadb
	@cp $(HOME)/inception/srcs/requirements/tools/daemon.json /etc/docker/daemon.json
	@docker-compose --env-file srcs/.env -f srcs/docker-compose.yaml up -d --build

down:
	@docker-compose --env-file srcs/.env -f srcs/docker-compose.yaml down

total_down:
	@docker-compose --env-file srcs/.env -f srcs/docker-compose.yaml down
	@docker volume rm $(VOLUMES_LIST)
	@docker system prune -af --volumes
	@rm -rf $(HOME)/data/mariadb/*
	@rm -rf $(HOME)/data/wordpress/*
	@rm -rf $(HOME)/data/adminer/*

upload:
	@ansible-playbook -i srcs/requirements/tools/ansible/inventory srcs/requirements/tools/ansible/playbook.yaml -t upload

update:	
	@ansible-playbook -i srcs/requirements/tools/ansible/inventory srcs/requirements/tools/ansible/playbook.yaml -t update

remote_system_prune:
	@ansible-playbook -i srcs/requirements/tools/ansible/inventory srcs/requirements/tools/ansible/playbook.yaml -t remote_system_prune

remote_down:
	@ansible-playbook -i srcs/requirements/tools/ansible/inventory srcs/requirements/tools/ansible/playbook.yaml -t remote_down

remote_up:
	@ansible-playbook -i srcs/requirements/tools/ansible/inventory srcs/requirements/tools/ansible/playbook.yaml -t build
help:
	@echo "||                    $(H)                                     ||"
	@echo "||=============================================================||"
	@echo "||                                                             ||"
	@echo "||  1)upload - upload project to the remote host               ||"
	@echo "||  2)update - update project on the remote host               ||"
	@echo "||  3)remote_system_prune - delete project on the remote host  ||"
	@echo "||  4)remote_up - up projec on remote host                     ||"
	@echo "||  5)remote_down - down project on the remote host            ||"
	@echo "||=============================================================||"
	@echo
