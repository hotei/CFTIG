<center>
# CFTIG - Computer Forensics Toolkit In Go
</center>


## OVERVIEW

## What Is It?

### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/CFTIG```

Will copy github.com/hotei/CFTIG to the first entry of your $GOPATH

or if go is not installed yet :

> ```cd DestinationDirectory```

> ```git clone https://github.com/hotei/CFTIG.git```

### Features

### Limitations

* <font color="red">Some programs may open the raw device - such as /dev/sda - 
you must be root to run this kind of program.</font>

### Usage

Typical usage is :

```program | tee results.txt```

### BUGS

### To-Do

* Essential:
 * TBD
* Nice:
 * TBD
* Nice but no immediate need:
 * TBD

### Change Log
* 2014-03-xx Started

### Resources

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for program] [3]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/CFTIG "github.com/hotei/CFTIG"

Comments can be sent to <hotei1352@gmail.com> or to user "hotei" at github.com.
License is BSD-two-clause, in file "LICENSE"

License
-------
The 'CFTIG' go package/program is distributed under the Simplified BSD License:

> Copyright (c) 2014 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Documentation (c) 2014 David Rook 

// EOF README-CFTIG.md  (this is a markdown document and tested OK with blackfriday)