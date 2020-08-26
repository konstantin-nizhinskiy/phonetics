PHONETICS
=========

Implemented algorithms:

* [Metaphone](http://en.wikipedia.org/wiki/Metaphone)
* [Metaphone-ru](http://rsdn.org/forum/db/906963.hot)

### Example
```go

 phonetics.EncodeMetaphone("Wat")//WT
 phonetics.EncodeMetaphone("What")//WT
 phonetics.EncodeMetaphone("Неженськый")//НИЖИНСК7
 phonetics.EncodeMetaphone("Ніженский")//НИЖИНСК7
 phonetics.EncodeMetaphone("Неженский")//НИЖИНСК7

```
More information can be found in the [godocs](http://godoc.org/github.com/konstantin-nizhinskiy/phonetics)
