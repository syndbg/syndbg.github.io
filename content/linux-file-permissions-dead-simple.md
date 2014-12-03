Date: 2014-11-29
Title: Linux: File Permissions, dead simple
Tags:
Slug: linux-file-permissions-dead-simple



## Intro

As every new Linux user may wonder, what do the file permissions means. What does file permissions as 777, 644 or 200 mean? How do I calculate that?


In this article, I'll try to explain this, __dead simple__.



## The bits and bats

So file permissions as the name suggest, decide if you can **read**, **write** and **execute** the given file. But what do the numbers and single letter characters mean, you might ask? They're represented in bits.

Each permission bit adds a number to its total:

* 4 = Read (R)
* 2 = Write (W)
* 1 = Execute (X)

The possible permissions scenarios are:

* 7 (R+W+X) - all permissions. Very dangerous unless you know what you're doing.
* 6 (R+W) - common scenario for non-bash scripts that must not be executed. For example Python scripts are like that.
* 5 (R+X) - useful for bash scripts that must be immutable, yet executable.
* 4 (R) - when you want your Python, Ruby scripts to work and remain immutable.
* 3 (W+X) - I honestly can't think of an usage with these permission bits.
* 2 (W) - useful if you have multi-accounts that you want to log into one file, but not know what the others wrote.
* 1 (X) - meaningless on non-binary files, since the shell needs to read the file to know what to do before executing it. Run binary files just fine with `./foo`



## Practical

Let's create a blank file and see its permissions with `ls -l` (just `l` in zsh).


```bash
$ touch afile
$ l afile
-rw-rw-r-- 1 syndbg syndbg 0 дек  4 00:39 afile
```

`-rw-rw-r--` are the permissions, but wait, why are they so many?

They're exactly 10 characters. The first one is the file type (I'll explain it in a bit) and the other 9 are separated in 3 triads. The first one is __u__ser (file owner) permissions, __g__roup (members from owner's group) permissions and last - __o__ther (world) permissions.


This means the file is readable and writeable for the owner and group, but only readable for others.


### The file type

**The most common ones**:

* `-` is a normal file,
* `d` is a directory,
* `l` is a symbolic link,



**The interesting ones:**

Following the Unix mantra, everything is a file, even a device (actually it's the device's interface).

* `s` is a socket device (unix socket),
* `b` is a block device (disk),
* `c` is a character device (stream),
* `p` is a pipe device (named pipe),



### Changing permissions

... is really easy! There are two ways, the first one is using the numbers, which I find easier to understand and the second is using characters.


In both cases we'll use `chmod`. So let's change our example `afile`'s permissions.

Currently they are `-rw-rw-r--`. Let's start with the first way, using numbers.


```bash
$ l afile
-rw-rw-r-- 1 syndbg syndbg 0 дек  4 00:39 afile
$ chmod 644 afile
$ l afile
-rw-r--r-- 1 syndbg syndbg 0 дек  4 00:39 afile
```


And the second way requires more commands understanding and less math.


chmod works like this - `chmod <triad>(operator)<permissions> filename`

* triad options are a(ll), u(ser), g(roup), o(ther) or empty (also means all). They can be combined as `go`, `ug`, `uo` or `ugo` (meaningless).
* operator options are `+` or `-`. As expected, addition adds permissions, substraction - removes.
* characters - `r`, `w`, `x`.


(Before that we restore the file back to its previous permissions `-rw-rw-r--`)

```bash
$ chmod +x afile
# or
$ chmod a+x afile
# CHANGES PERMISSIONS FOR EVERYONE (user, group ,other)!_!_1-1
# Number one mistake that everyone makes!
$ l afile
-rwxrw-r-- 1 syndbg syndbg 0 дек  4 00:39 afile
# Changing just user permissions
$ chmod u-x afile
```




That's it folks, I hope you found this useful.
