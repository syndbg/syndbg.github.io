---
title: "Publishing to PyPi Can Be A Nightmare"
date: 2015-04-05T11:21:29Z
draft: false
categories:
  - "python"
  - "pypi"
---

## Intro

Yesterday (04.04) I was having the most ridiculous issues with PyPi - from not being able to publish my package, to already publishing a version and not being able to update and at last - deleting my package from PyPi wouldn't actually delete it.


**This makes me ask why would you make PyPi so overengineered and quirky?**


## I want to publish a package

No reliable PyPi publishing tutorial. I have to rely on users' blogs to learn how to publish my package.

If it was my first PyPi package, I would've had to rely on unofficial tutorials, since the official ones are at best, **horrible**. Compare https://wiki.python.org/moin/CheeseShopTutorial to http://guides.rubygems.org/make-your-own-gem/ or https://docs.npmjs.com/getting-started/creating-node-modules .

It's obvious which is more concise and understandable. Not to mention way fewer steps needed.


## I want to delete a specific version

It's mostly just a visual thing. I published version `1.0.0` of my package on PyPi and I delete it from the Web UI. Does this means it's really deleted? No!

If I try to upload my package anew, I would get the message below, even though I deleted it from the PyPi's web UI.

```
Upload failed (400): This filename has previously been used, you should use a different version.
```


## I want to delete my project

Same problem as above. Deleting my project completely, still doesn't delete the versions from PyPi. It would hide them from users therefore rendering it impossible to install anything with `pip install my_package`, but the files would stil persist on PyPi servers.


## I just want my package to be reachable by people, damn it!

If you already published a version once and deleted it from the Web UI, chances are that you can never EVER again use that version number and you must settle with another one. Don't do my mistake and publish starting with `1.0.0` even if your package is ~100% test covered. It's not PyPi covered. Yeah, we have to cover that too.


The only workaround is when uploading to specify explicitly to upload as `zip`. Why it works? Cause by default it gets uploaded as `tar.gz`. So if you have already uploaded version `1.0.0`, it exists as `{package-name}-1.0.0.tar.gz` and you can upload it as `{package-name}-1.0.0.zip` if you write:

```
python setup.py sdist --formats=zip upload
```

However you should read a bit about `setup.py`'s `zip_safe` setting https://pythonhosted.org/setuptools/setuptools.html#setting-the-zip-safe-flag .


If you uploaded it as `zip` first, just do the reverse and upload with `formats=tar.gz` .


## Is there a future?

There are two sides of the coin, as seen in the articles:

* Alex Clark's http://blog.aclark.net/2011/01/31/in-defense-of-pypi/
* Peter Fein's http://i.wearpants.org/blog/elitism-and-the-importance-of-pypi/


I was willing to accept this situation if we were still in 2010-2012, but it's 2015 and PyPi isn't getting any better. Not at all, while JS, Ruby and the newer languages Rust and Go, all have better package publishing and management.


If Python wants to move forward and keep up with the other languages, being the best starter programming language isn't enough to keep people coding in Python. We need to fix the `Python 3 port` and `PyPi` problem as soon as possible.

Sincerely, a Pythonista.