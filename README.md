# Go_rc4
## 由Go编写的RC4 cli工具，《计算机网安安全原理》的实验二的源码。本次算法及其简单。

## 如何使用
``` bash
PS C:\Users\SU15VTE\Go_rc4> .\Go_RC4.exe encrypt --key "SU15VTE" --text "Hello World!"
encrypt called
en called
Key:    SU15VTE
text:   Hello World!
ciphertext: 4dc1dacdad813f8f50a9b9c7
PS C:\Users\SU15VTE\Go_rc4> .\Go_RC4.exe decrypt --key "SU15VTE" --text "4dc1dacdad813f8f50a9b9c7"
decrypt called
de called
Key:    SU15VTE
text:   4dc1dacdad813f8f50a9b9c7
plaintext: Hello World!
```
