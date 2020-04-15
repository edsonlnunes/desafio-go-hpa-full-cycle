# Desafio Go e HPA - Curso MS

* Processo de CI
* Processo de criar imagem e public no Container Registry
* Processo de publicar imagem no docker hub
* Processo de publicar no Kubernetes
* Processo de Auto Scaling (HPA)


# Link image projeto go
https://hub.docker.com/repository/docker/edsonnunes/go-hpa


# Executando o cenário do HPA
  * Executar o apply dos yamls (go-deployment, go-hpa, busybox-pod)
  * Entrar no terminal do pod busybox (kubectl exec -it busybox /bin/sh)
  * Efetuar requisição ao endereço do service go-hpa (while true; do wget -q -O- http://34.67.128.13:8000/; done;)