# Quick Distributed Brute Force

A simple tool for pentesters looking to automate their tests

## Flags
-b (500) Set amount of parallel requests

-o (out.log) Set the output file for criteria results

-p (7575) Set port for inter-node communication

## Config Structure

```yaml
request:
  method: <GET | POST | PUT | ...> # any valid http method
  url: http://test/$prodid$
  params:                          # keys describe the param name
    username: Test
    password: lol
    id: "$id$"                     # define variables between '$'
  headers:
    cookie: "sessionid=ab3235bac5"
  body: '{"user":"test", "bro":55}'

params: # Different types of variables (only one can be used)
  prodid:
    type: RANGE
    from: 1
    to: 10000
  id:
    type: DICT
    dict: ['a', 'b', 'c']
  bro:
    type: FILE              
    file: dict.txt

criteria: 
  type: <STOP | LOG>
  response:
    status: 200
    headers: # Key value pairs
      cookie: "sesion=sa4d65541a"
    body: 'sjdhaksjdh' # Regex supported

helpers: # Other instances that can help split the work
  - 192.168.1.1
  - 123.45.67.8
```