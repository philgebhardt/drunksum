# drunksum

Checksums with Drunken Bishop text output. Largely modeled after [The drunken bishop: An analysis of the OpenSSH fingerprint visualization algorithm][1]

[1]: http://www.dirk-loss.de/sshvis/drunken_bishop.pdf


```shell
$ go build .
$ ./drunksum -a md5 < drunksum
_                   _


           =
        X # =
       E S   =
          = =
           =


_                   _

md5: fdbea08fbded2b874a1e78700562d03a
```

# why?

This project is not intended to provide anything beyond the joy of exploring random data visually. 
