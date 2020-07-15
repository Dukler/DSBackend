@ECHO OFF
:: Assign all Path variables
SET CGO_ENABLED=0
SET GOOS=windows
set output=C:\Users\iarwa\Workspace\Servers\DSBackend\
set input=C:\Users\iarwa\Workspace\Go\DSBackend\

go build -a -installsuffix cgo -o server.exe .

ECHO build finished
ECHO copying executable
xcopy "%input%server.exe" %output% /K /D /H /Y
ECHO copying Apps
xcopy "%input%Apps" "%output%Apps" /E /H /C /I
ECHO copying tokens
xcopy "%input%tokens" "%output%tokens" /E /H /C /I
ECHO done.