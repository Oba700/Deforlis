# License:
Licensed under [STV License](https://github.com/Oba700/STV_License/blob/main/README.md)

# What?
Pet-project web-server

# Why?
To discover C++ as a tool and ~~flex~~ slay publically.

# Project maturity grade
It doesn't work at this point

# Development
 - Install GCC
 - Run `g++ main.cpp`
 - create config file in ini with roughly such content
```
[binding]
bindingName=unpriv1
IPv4addr=0.0.0.0
IPv4port=8080

[binding]
bindingName=unpriv2
IPv4addr=0.0.0.0
IPv4port=8081

[binding]
bindingName=unpriv3
IPv4addr=0.0.0.0
IPv4port=8082
```
 - Run `./a.out --configfile=yoa.ini`
 - Enjoy, go down the spiral, never learn