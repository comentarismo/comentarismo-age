Comentarismo Age API

# A API to determine Age on names

Age demographics for 2016

* 18-24  -> 1998 to 1992 
* 25-34  -> 1991 to 1982 
* 35-44  -> 1981 to 1972 
* 45-54  -> 1971 to 1962 
* 55-64  -> 1961 to 1952 


Inside the redis bayes the following is applied:

```
* condition: (now().year - 18) to (now().year - 24) ->eg: 1998 to 1992 -->map: bayes:18_24
* condition: (now().year - 25) to (now().year - 34) ->eg: 1991 to 1982 -->map: bayes:25_34
* condition: (now().year - 35) to (now().year - 44) ->eg: 1981 to 1972 -->map: bayes:35_44
* condition: (now().year - 45) to (now().year - 54) ->eg: 1971 to 1962 -->map: bayes:45_54
* condition: (now().year - 55) to (now().year - 64) ->eg: 1961 to 1952 -->map: bayes:55_64
```


Running defaults & debug mode & learn age names
```
$ AGE_DEBUG=true LEARNAGE=true PORT=3004 godep go run main.go
```

# Options
```
AGE_DEBUG, if true will debug all log entries for age detection (optional)

LEARNAGE, if true will learn words for each available lang
REDIS_HOST, ip addr of the redis instance to be used (required) -> defaults to g7-host
REDIS_PORT, port number of the redis instance to be used (required) defaults to 6379
REDIS_PASSWORD, password for this instance to be used (optional)
```

