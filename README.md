# Distributed Calculator of Arithmetic Expressions

## Run project (Powershell method)
1. Install or Update Git
2. Get a copy of the project source code
```shell
git clone https://github.com/IvanNovCode/Distributed-calculator-of-arithmetic-expressions.git
```
3. Run agent
```sh
go run .\cmd\agent\main.go
```
4. Run orchestrator
```sh
go run .\cmd\orchestrator\main.go
```

## Request example
Create a new expression
```sh
curl -X POST --data-urlencode "expression=2+2*2" http://localhost:8080/addexpression
```
Get a list of all expressions
```sh
curl http://localhost:8080/getexpressions
```
Clear all previous expressions
```sh
curl http://localhost:8080/clearexpressions
```
Setting the Expression Settings
```sh
curl -X POST --data-urlencode "setting=+200" http://localhost:8080/setsetting
curl -X POST --data-urlencode "setting=-200" http://localhost:8080/setsetting
curl -X POST --data-urlencode "setting=*200" http://localhost:8080/setsetting
curl -X POST --data-urlencode "setting=/200" http://localhost:8080/setsetting
```
Get current settings
```sh
curl http://localhost:8080/getsettings
```