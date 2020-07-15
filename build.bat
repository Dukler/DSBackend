@ECHO OFF
:: Assign all Path variables
SET CGO_ENABLED=0
SET GOOS=windows
set output=C:\Users\iarwa\Workspace\Servers\DSBackend\
set input=C:\Users\iarwa\Workspace\Go\DSBackend\

go build -a -installsuffix cgo -o server.exe .