@echo off

set /p selectPlatform=please select build platform:[1]linux [2]win:(linux)


if "%selectPlatform%"=="" (
    set selectPlatform=1
)

set /p buildPackageName=please input package name:(xadmin_linux)

if "%buildPackageName%"=="" (
    set buildPackageName=xadmin_release
)

if %selectPlatform%==1 (
    echo start build linux: %buildPackageName%

    go env -w GOOS=linux

    del /f /s /q .\build\*.*
    rd /s /q .\build\


    if not exist build (
        mkdir build
    )


    if not exist build\config (
        mkdir build\config
    )
    if not exist build\public (
        mkdir build\config
    )


    cd ..\xadmin\
    go work sync
    go build -o %buildPackageName% ../xadmin/main.go

    xcopy .\%buildPackageName% ..\xbuild\build  /s /e /y
    del %buildPackageName%


    xcopy   .\config ..\xbuild\build\config  /s /y
    xcopy   .\public ..\xbuild\build\public  /s /y

    rd /s /q .\logs
    exit
)

if %selectPlatform%==2 (
    echo start build win: %buildPackageName%

    go env -w GOOS=windows
    go env
    del /f /s /q .\build\*.*
    rd /s /q .\build\

    if not exist build (
        mkdir build
    )

    if not exist build\config (
        mkdir build\config
    )
    if not exist build\public (
        mkdir build\config
    )


    cd ..\xadmin\
    go work sync
    go build -o %buildPackageName% ../xadmin/main.go

    echo f | xcopy .\%buildPackageName% ..\xbuild\build  /s /e /y
    del %buildPackageName%

    echo d |  xcopy   ".\config" "..\xbuild\build\config"  /s /e /y
    echo d |  xcopy   ".\public" "..\xbuild\build\public"  /s /e /y
    rd /s /q .\logs
    exit
)