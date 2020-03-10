# Input/Output

# Beginning
```
netstat -r | wsl.exe -d Ubuntu grep --color -E "Active Routes|$"
```

```
netstat -r | Select-String Active Routes
```

```
netstat -r | findstr.exe /C:"Active Routes"
```

## Linux
```
./palindrome.exe | cowsay
```

```
echo -n "Hello WSLConf!" | ./palindrome.exe | cowsay
```

## Windows Powershell
```
wsl.exe -d Ubuntu ./palindrome | wsl.exe -d Ubuntu cowsay
```

```
Write-Output "Hello WSLConf!" | wsl.exe -d Ubuntu ./palindrome
```

```
Write-Output "Hello WSLConf!" | wsl.exe -d Ubuntu bash -c "./palindrome | cowsay"
```

```
Write-Output "Hello WSLConf!" | wsl.exe -d Ubuntu ./palindrome | wsl.exe -d Ubuntu cowsay
```

# Socat

## 1
```
nc -l localhost 1939
```

```
nc localhost 1939
```

## 2
```
socat TCP-LISTEN:1939 STDIO
```

```
nc localhost 1939
```

## 3

```
socat UNIX-LISTEN:test.sock,fork EXEC:/usr/games/cowsay
```

```
nc -q 0 -U test.sock
```

## 4
```
socat UNIX-LISTEN:test.sock,fork EXEC:palindrome.exe
```

```
nc -q 0 -U test.sock | cowsay
```

Optional
```
echo -n "Hello WSLConf!" | nc -q 0 -U test.sock | cowsay
```

# Debug
```
fuser 1939/tcp
fuser -k 1939/tcp
```