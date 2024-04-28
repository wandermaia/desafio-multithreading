# **Repositório para o Desafio de Multithreading da Pós de Go Expert**

Este repositório foi criado exclusivamente para hospedar o desenvolvimento do Desfio sobre Multithreading da **Pós Go Expert**, ministrado pela **Full Cycle**.


## Descrição do Desafio

 Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Informações sobre a Implementação

O desafio foi implementado para ser utilizado diretamente no shell, onde o CEP é passado como parâmetro. Abaixo segue um exemplo da execução:

```bash

wander@bsnote283:~$ go run main.go 01153000
CEP Informado: 01153000

O CEP 01153000 é válido.

Dados recebidos do site https://brasilapi.com.br: 

{
  "cep": "01153000",
  "state": "SP",
  "city": "São Paulo",
  "neighborhood": "Barra Funda",
  "street": "Rua Vitorino Carmilo",
  "service": "widenet"
}

wander@bsnote283:~$ 

```

Também há uma pasta de `test`, mas ela contém apenas as chamadas às APIs do https://brasilapi.com.br e do https://viacep.com.br. Essas chamadas foram utilizadas para testar essas APIs e verificar o retorno dos dados.

