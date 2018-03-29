## Gossh

Gossh is a remote client that makes up for the inability of terminal ssh to save passwords.

## Why use golang?

Because golang does not require other dependencies or runtime environments, it only need the terminal

## Download
[https://github.com/i183/gossh/releases](https://github.com/i183/gossh/releases)

## Commands
### add
Add a remote server
```
$ gossh add <SERVER_NAME> <USERNAME> <IP> <PORT> <PASSWORD>
```
For example
````
$ gossh add serverName root@127.0.0.1 22 password
````

#### conn
Connect to the remote  server
```
$ gossh conn <SERVER_NAME>
```
For example
```
$ gossh conn serverName
```

#### ls
Show the remote server list
```
$ gossh ls
```

#### rm
Remove a remote server
```
$ gossh rm <SERVER_NAME>
```
For example
```
$ gossh rm serverName
```