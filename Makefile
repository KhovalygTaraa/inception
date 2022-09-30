DCKER_COMPOSE = $(shell command -v docker-compose)
H=\033[0;36mHELP\033[0m

all:
ifndef $(DOCKER_COMPOSE)
	@docker-compose -v
else
	echo "install docker-compose"
endif
	touch Hello

up:
	@docker-compose -f srcs/docker-compose.yaml up -d --build

down:
	@docker-compose -f srcs/docker-compose.yaml down
	@rm -rf ../data/mariadb/*

system_prune:
	@docker system prune -af --volumes

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
	@echo "||                    $(H)                        ||"
	@echo "||================================================||"
	@echo "||                                                ||"
	@echo "||  1)upload - upload project to the remote host  ||"
	@echo "||  2)update - update project on the remote host  ||"
	@echo "||  3)delete - delete project on the remote host  ||"
	@echo "||  4)remote_build  - run make on remote host     ||"
	@echo "||                                                ||"
	@echo "||================================================||"
	@echo
.PHONY: upload update remote_build help all
