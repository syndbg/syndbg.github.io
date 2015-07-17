---
layout:     post
title:      "Learning Django"
date:       2015-03-26 11:21:29
summary:    Learning Django resources
categories: django python
---

Django is the go-to batteries included web framework for Pythonista.

However starting to create web apps with Django is a bit rough,
if you don't have the comfort of someone mentoring you.

From my experience as a Django developer for 1 year, 3-4 projects and close to 3 versions of the framework,
I will share the materials that helped me start with Django without
prior web and MVC pattern experience.



First of all I am a book guy, I like books. I think books are the best tutorials since they're better structured than blog tutorials who cover only small portions of the greater picture.



## My recommendations are:


**You're not confident you know Python?**. Kein problem. Months ago I was amazed by [Bill Lubanovic's Introducing Python](http://shop.oreilly.com/product/0636920028659.do) and how clear, concise and practical the book is. The book I wish I had when I started writing Python. You'll learn to express yourself in Python code faster than you can speak in [Parseltongue](http://harrypotter.wikia.com/wiki/Parseltongue)!


**If you have no web or MVC pattern experience start with Flask.** Yep, Flask the microframework. You need to understand how a really basic web app works and how to achieve things in a non-pattern enforced way.

Flask gives you the freedom to mimic MVC or structure your web app the way it makes sense to you. Either way, it's a web microframework that doesn't overwhelm you at first.

Go for Miguel Grinberg's amazing [Flask Web Development Book](http://flaskbook.com/). It's the best web development starter book you can wish for!


**But if you won't settle with anything different than Django or have web experience or experience with MVC**,

Go for the playful [Tango with Django tutorial](http://www.tangowithdjango.com/). It's updated for Django 1.7 and touches upon everything you'll need to know (at least as basics) if you want to create a Django web app from start to finish.


**Ok you have the basics down, you look to improve and master Django**.

Here are the latest and greatest (in my eyes) books that you should read - [Test-Driven Development with Python](http://shop.oreilly.com/product/0636920029533.do). The best book I've read on TDD and the book that taught me the most on how to develop and ship quality software. Author Harry J.W Percival is a TDD beast and will have you learn how to test every tiny bit of Django. Doing unit and integration tests, you'll learn the most about the framework. Also great benefits are learning how to deploy, add a bit of JS fanciness and automate the deployment process of your web app. Definitely one of the best books I've read about Python.


After you woke up in your new self as TDD(iddy), take a look at the tasty [Two Scoops of Django](http://twoscoopspress.org/), available for 1.5/1.6 at the time of writing. But don't let that stop you, it's experts Daniel Greenfeld and Audrey Roy's advice for Django developers. Even being 1-2-3 versions behind won't change the fact that it's still a MVC framework that functions in the same way but with new features. You might even learn more by seeing how experts dealt with specific problems in earlier versions and how (maybe/hopefully) Django project core devs made that easier/better.


**You got the JS hype and you need to combine it with Django without causing a nuclear reaction**.

We all have been there. But since November 2014 we have a book that teaches you all the things you need to know to make JS and Django behave under 1 roof.
The book is [Lightweight Django](http://shop.oreilly.com/product/0636920032502.do). Buuuut no AngularJS? I got you, check out thinkerster.io's [Building Web Applications with Django and AngularJS](https://thinkster.io/django-angularjs-tutorial/) and their other AngularJS oriented tutorials for Rails and Firebase.
You'll quickly realize that adding a JS frontend to your web app involves the same abstract steps in most web frameworks.


**Useful resources**:

![request response cycle](http://rnevius.github.io/django_request_response_cycle.png "request response cycle")


**Notes**:

* At the time of writing this, the latest stable release of Django is 1.7.
* In Django Model-View-Controller translates to Model-Template-View.
