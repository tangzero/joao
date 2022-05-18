# Joao
simple password recovery tool

## Identify

```sh
joao identify file.txt
```

or

```sh
echo hash-to-identify | joao identify
```

### Example output

```
20824005085df04078c1596a2ad00ea6 [MD5]
f9b2d62d0078783d0063aaf4731f3b3727327e08 [SHA1]
dcb66983cd7205493d3157b64789b59b [MD5]
f9fe0661e262cc5f27211e82e286e61d30540709 [SHA1]
```

## Decrypt

```sh
joao decrypt file.txt
```

or

```sh
echo hash | joao decrypt
```

### Example output

```
20824005085df04078c1596a2ad00ea6 [MD5] = casita
f9b2d62d0078783d0063aaf4731f3b3727327e08 [SHA1] = Panqueque
f9fe0661e262cc5f27211e82e286e61d30540709 [SHA1] = cuchito
```

### Toggle Case Attack

the Toggle Case Attack can be used to decrypt/recovery/crack some difficult hashes.
The `toggle` is used to specifiy how many letters from the beggining of the word should be toggled.

```sh
joao decrypt --toggle 1 file.txt
```

This will try to match  `banana` and `Banana` words, for example.

```sh
joao decrypt --toggle 2 file.txt
```

and this one will try `banana`, `Banana`, `bAnana` and `BAnana` .



 
  