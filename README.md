# Trabalho 2 de Criptografia

## Como usar

### Compilação

```sh
mage
```

Caso não tenha mage, você pode fazer `go run mage.go`.

Caso queira instalar mage, você pode usar o target ensureMage:
```sh
go run mage.go ensureMage
```

O binário gerado será salvo em `./bin`.

### Uso

Para detalhes sobre como usar o programa, execute:

```sh
./bin/t2 -h
```

#### Exemplos

-   `echo hello | ./bin/t2 enc 7 187`

    Criptografa a mensagem "hello" usando a chave pública {7, 187}.

-   `echo s1QwMJuv | ./bin/t2 dec 7 187`

    Descriptografa a mensagem "s1QwMJuv" usando a chave pública {7, 187}.
