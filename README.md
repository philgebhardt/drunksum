# drunksum

Checksums with Drunken Bishop text output. Largely modeled after [The drunken bishop: An analysis of the OpenSSH fingerprint visualization algorithm][1]

[1]: http://www.dirk-loss.de/sshvis/drunken_bishop.pdf


```shell
$ go build .
$ ./drunksum -a sha512 < drunksum
_                   _
      oBBo
     ooo oo
   o  o =o o
  o o  ==oo
 o  ooo oS
 o=o=o    o
 B =oo
 %ooE
 BX
_                   _

sha512: 43f40481820f09f144a8783b0296a864a59d5e013174e9e6d317962c0e72481a2fd68678cca2024b529dee303f58cb0c14d446d55359e70835107c0e1afbc34f
```

# why?

This project is not intended to provide anything beyond the joy of exploring random data visually. 
