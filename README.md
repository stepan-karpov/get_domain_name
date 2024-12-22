`sudo docker build -t yandex-get-domain-name .`

`sudo docker run -d -p 8081:8080 --name yandex-get-domain-name-container yandex-get-domain-name`

`sudo docker ps -a`

`sudo docker stop yandex-get-domain-name-container`

`sudo docker rm yandex-get-domain-name-container`

`curl http://15.236.182.228:8081/get_domain_name -d "{\"ip\": \"87.250.251.140\"}"`