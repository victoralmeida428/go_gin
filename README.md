# Estrutura API

## Necessário

- Arquivo .env na pasta ./src

| Name    | Type    |
|---------|---------|
| PORT    | Integer |
| DB_PASS | String  |
| DB_HOST | String  |
| DB_PASS | String  |
| DB_USER | String  |
| DB_NAME | String  |
| DB_PORT | Integer |
| DB_JWT_SECRET_KEYPORT | string |
| GIN_MODE | release / debug |
| ALLOWED_ORIGIN | string |


---
## Migrate
- Criar Arquivo
```bash
migrate create -seq -ext=.sql -dir=./migrations nome_da_etapa
 ```
- Migrar
```bash
migrate -path=./src/migrations -database=postgres://user:password@host:port/dbname up
``` 
---
## Iniciar
```bash
go run .
```
---
## Documentação Swagger
/swagger/index.html
---
## Controller
Pasta onde controla as ações de cada endpoint.<br>
Exemplo:
```go
package controller

import (
	"api/src/utils"
	"net/http"
)

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusOK
	defer c.cfg.Logger.PrintDebug(r, status, nil)

	dados, err := c.repository.Grupo.FindAll()

	if err != nil {
		status = http.StatusInternalServerError
		c.cfg.Error.ErrorReponse(w, r, status, err.Error())
		return
	}
	if err := utils.WriteJSON(w, dados, status, nil); err != nil {
		status = http.StatusInternalServerError
		c.cfg.Error.ErrorReponse(w, r, status, err.Error())
		return
	}
}

```
---
## Build
### Win
```bash
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o app.exe ./main.go 
```
### Linux
```bash
go build -ldflags "-s -w" ./main.go 
```


