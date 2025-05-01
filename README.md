# stress-test
Atividade Pós-GoExpert Stress test

## Descrição

Comand line criado para uma atividade da pós-graduação GoExpert, o objetivo é ter um programa que faça o envio de requests e exiba um relatório final com dados da execução, tendo como opção paralelizar as requests.

## Testando

```bash
docker build --no-cache -t stress:latest .
docker run stress --url=http://teste.com.br --requests=100 --concurrency=2
```

Também pode ser executado dessa forma:

```bash
docker run stress -u http://teste.com.br -r 100 -c 2
```

Uma observação, caso o --requests não seja informado, o programa assumirá 1 como valor default, o mesmo se aplica ao parâmetro --concurrency.


https://github.com/user-attachments/assets/c1b7a457-fbea-462d-bffb-c615d84195b5
