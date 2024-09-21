i should use load.io

and they will request to add there txt file, u should add it in public and also in cmd/main.go

run 
`go build -o outPut ./cmd/main.go`

and them

`./outPut`

the performance are stunning
while Laravel with Swoole or FrankenPHP were crashing after 10seconds of loads with high CPU usage in all cores which is understandable

### With 10k per Minutes
<img src="./htop10kperMinutes.png" />
<img src="./with10kperMinutes.png" />

### With 10k per Second
<img src="./htop10kperSecond.png" />
<img src="./with10kperSecond.png" />
