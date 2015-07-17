---
layout:     post
title:      "Python: Decorators gotchas"
date:       2015-03-14 11:21:29
summary:
categories: python
---

## Intro

In Python, decorators are one of the most misunderstood features and one of the most blogged topic,
due to the their gotcha nature.


What's a gotcha:
> In programming, a gotcha is a feature of a system, a program or a programming language that works in the way it is
> documented but is counter-intuitive and almost invites mistakes because it is both enticingly easy to invoke and
> completely unexpected and/or unreasonable in its outcome.



**If you're not familiar with the idea behind decorators, head to [PEP-0318](https://www.python.org/dev/peps/pep-0318/).
It'll help understand and be able to write simple decorators.**


## However, some of the gotchas I've faced are:

**How do I pass args and kwargs to the decorator?**

The obvious approach (not working)

```python
def exposed(func, *attrs_to_expose):
    def _callable(*args, **kwargs):
        for attr in attrs_to_expose:
            print(getattr(func, attr))
        return func(*args, **kwargs)
    return _callable


@exposed('__name__', '__name__')
def decorated_fn(*args):
    ....

# Results into
>>> decorated_fn('1', '2', '3')
decorated_func
__main__
....
```

And the naive working non-obvious at first approach

```python
def exposed(*attrs_to_expose):
    def decorator(func):
        def _callable(*args, **kwargs):
            for attr in attrs_to_expose:
                print(getattr(func, attr))
            return func(*args, **kwargs)
        return _callable
    return decorator

# Results into
>>> decorated_fn('1', '2', '3')
totally_working
__main__
yep
```

**Same idea carries over for kwargs. BUT PLEASE, don't use `getattr` like that.** I mention the `functools.wraps` bug and how Django devs solved it.


**Why do we use functools.wraps?** Mainly legacy

To make debugging decorators way easier. It updates function attributes mentioned in `functools.WRAPPER_ASSIGNMENTS`, which at the time of writing this are `'__module__', '__name__', '__qualname__', '__doc__',
                       '__annotations__'`.


In code for earlier than Python 2.7:

```python
def exposed(func):
    def _callable(*args, **kwargs):
        print(func.__name__)
        print(func.__doc__)
        return func(*args, **kwargs)
    return _callable

# When we use it on:
@exposed
def totally_working(*args, **kwargs):
    """I lied"""
    raise Exception

# In Python versions earlier than 2.7
>>> totally_working('d', 3, 'c', 0, 'rators')
# nothing printed

# In Python versions 2.7.8 or and 3.3 and newer
>>> totally_working('d', 3, 'c', 0, 'rators')
totally_working
I lied
```

**Why do we implement our own `functools.wraps` solutions?** Again legacy

Nothing is without bugs and `functools.wraps` had a pretty long lasting bug that was reported http://bugs.python.org/issue3445 and fixed way later in Python3.3 and later.

Basically the `functools.WRAPPER_ASSIGNMENTS` at that time was aggressively trying to update a function's attributes, not taking into consideration the chances of them not existing.


The issue had Python developers think about ways to solve it. And from my limited experience, Django developers have the most elegant workaround. As seen https://github.com/django/django/blob/master/django/utils/decorators.py#L82


### In newer Python 2.7.x and 3.x the use `functools.wraps` is redundant
