"Go é uma linguagem de programação de código aberto que facilita a criação de softwares simples, confiáveis e eficientes."

Criada em 2007 pro Robert Griesemer, Rob Pike e Ken Thompson. <br>
Go nasceu por canta da complexidade de códigos de soluções do Google. <br>
Convenhamos que não é só o Google que tem problemas com complexidade de soluções. <br>
Então... Go não é só uma linguagem para o Google. <br>
A linguagem escalou e a comunidade ganhou força.

##### Pontos fortes
- Tipagem forte e estática;
- Inferência de tipo;
- Compilada;
- GC;
- Concorrência;
- Rápida;
- Simples; e 
- Até mesmo divertida.

Go é uma linguagem que podemos considerar como "multiparadigma".

- O.O.;
- Procedural; e
- Funcional.

##### CI/CD
CI: Continuos Integration <br>
CD: Continuous Delivery

##### Qualidade de código
- Testes;
- Lint; e
- FMT.

##### Build
- Dockerização

##### Antipatterns
- Implantar software manualmente;
- Implantar em um ambiente similiar ao de produção apenas quando o desenvolvimento esiver completo.; e
- Vantagens: Reduzir erros, estresse, flexibilidade de implantação e rollback mais simples.

##### Variáveis de Ambiente
Adicionar em Path:
```
C:\Program Files\Go\bin
```

Criar uma nova:
```
GO_HOME
```

```
C:\Program Files\Go
```

##### Iniciando o projeto
```golang
go mod init github.com/lucasrmagalhaes/planejamento_financeiro-golang
```

##### Declarando variáveis em GO
```golang
var var1 = ""
```

```golang
var2 := ""
```

```golang
var var3 string
```

##### POST
```json
[
    {
        "title": "Salário",
        "amount": 1200,
        "type": 0,
        "created_at": "2022-01-16T10:20:00"
    }
]
```

##### Go Lint
```golang
go get -u golang.org/x/lint/golint
```

Verifica os diretórios onde possui teste(s)
```golang
go test ./...
```

Verifica os problemas
```golang
golint ./...
```

##### Go Fmt
```golang
go fmt ./...
```

##### Prometheus
```golang
go get github.com/prometheus/client_golang/prometheus
```

##### CircleCI
[Acesso](https://circleci.com/vcs-authorize/)

##### Docker
[Download](https://www.docker.com/products/docker-desktop) <br>
[WSL 2](https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi)

```
make build-image
```

##### Endpoints
```
http://localhost:8080/transactions
```

```
http://localhost:8080/transactions/create
```

```
http://localhost:8080/health
```

```
http://localhost:8080/metrics
```