Date: 2014-11-27
Title: Python: Modules vs Packages
Tags:
Slug: python-modules-vs-packages



## Intro

Coming from different language, the definition of a module and package may get mixed up. To confirm/clear up confusion, I wrote this article for myself and others who tend to mix modules and packages.



## What's a module?

In dynamic languages (ex. Python, Ruby, JS) a module is a single file's contents - classes, functions, variables.


Let's look at Django's config file for example (https://docs.djangoproject.com/en/1.7/ref/settings/)

The Django framework utilizes every one of them, some are defaulted, some must be explicitly defined. But point is they're a module that is read by other programs.

In a Django asset, you can access it through django.conf - `from django.conf import settings`.

But if you wanted to create your own module, in JS or Ruby you may need to explicitly define a module, but in Python you just write code.


Example: filename `foobar.py`
```python

foobar = 'foobar'

def foo():
    ...


def bar():
    ...
```

And this is a module. No need to explicitly write module anywhere.

If you wanted to use it somewhere, it'll just work as shown below.

```python
import foobar
# or
from foobar import foo


foobar.bar
foo()
```



## What's a package?

A package is the extended module idea. Instead of havaing only one module, you can have dozens. We already know that modules are just files, so a package is a bunch of files. However there's a certain boiler plate to this.

If you want to create a package, you need to create an `__init__.py` file in the same directory. The name is strict!


So if we have our own `accounts` Django package, it'll look like this:
```
accounts
├── admin.py # module
├── authentication.py # module
├── __init__.py # boilerplate for the package
├── models.py # module
├── urls.py # module
└── views.py # module
```


All of the modules/files can be accessed in other Django assets as:

```python
import accounts # basic
from accounts import admin # specific
from accounts.models import User # most specific

# 1
backend = accounts.authentication.AuthenticationBackend()
# 2
account_admin = admin.AccountAdmin()
# 3
user = User(name='Josh')
```

But it can be a bit confusing.

```python
In [1]: import accounts
In [2]: accounts
Out[2]: <module 'accounts' from '/home/syndbg/tdd-app/accounts/__init__.py'>
# This is a package, but it says module, you might say.
# However when we bring help(accounts)
In [3]: help(accounts)
    Help on package accounts:

    NAME
        accounts

    PACKAGE CONTENTS
        admin
        authentication
        models
        urls
        views

    FILE
        /home/syndbg/tdd-app/accounts/__init__.py

# So it just read __init__.py and said it was a module.
```

Packages can contain other packages.

If we want to keep everything tidy and use SoC (seperation of concerns), we would like to have our own folder for tests. And it'll look like this.

```
├── tests
│   ├── __init__.py
│   ├── test_authentication.py
│   ├── test_models.py
│   └── test_views.py

```

^ The above is another Python package. And that package is located in, as you might guess, the accounts package.

The whole picture:

```
├── admin.py
├── authentication.py
├── __init__.py
├── models.py
├── static
│   ├── accounts.js
│   └── tests
│       └── tests.html
├── tests
│   ├── __init__.py
│   ├── test_authentication.py
│   ├── test_models.py
│   └── test_views.py
├── urls.py
└── views.py
```

## Verdict

Nowadays modules and packages are an important pattern and feature to many languages that want to keep developers happy.

Java 9 will support modules. Python, Ruby, JS - already do. The list of languages will get bigger and bigger. Start using modules and packages! :)
