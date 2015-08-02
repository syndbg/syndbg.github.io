---
layout:     post
title:      "Review: The Well Grounded Rubyist, 2nd edition"
date:       2015-08-01 11:21:29
summary:    Review of Manning and David A. Black's The Well Grounded Rubyist, 2nd edition book
categories: review ruby book
---

![The Well Grounded Rubyist, 2nd edition cover](http://ecx.images-amazon.com/images/I/51qtp4NzjRL._SX397_BO1,204,203,200_.jpg "The Well Grounded Rubyist, 2nd edition cover")

## Links

* Manning - http://www.manning.com/black3/
* Amazon - http://www.amazon.com/The-Well-Grounded-Rubyist-David-Black/dp/1617291692
* Goodreads - https://www.goodreads.com/book/show/3892688-the-well-grounded-rubyist


## My thoughts

The Well-Grounded Rubyist, 2nd edition is a great read for every aspiring Ruby developer. Not flawless, but a book that may teach you a great deal of how Ruby works.


The book is a great fountain of knowledge so to say. Reading and understanding everything may take a good while, so I've added a `Must-read content` section with content that is *mostly* not your ordinary `learn Ruby tutorial information`.

I really liked the `This chapter covers` intro pages at the start of each chapter. It prepares the reader well what to expect and possibly research before heading into the chapter. Likewise, the summary at the end of each chapter does the opposite a favor.

Also due to the flaws I found in some sections, I suggest reading the `Flaws` section if you really care about them. I think you should since they're pointing out some wrongs in the info provided.

And of course a `What to improve` section is addressing the `Flaws`.

Before proceeding further if you just care about the rating I would give the book - it's **4/5**. The flaws and what to improve sections would've made it a `5` for me. The writing style is good, but not as clear and informative as I would like it to be all the time.


## Must-read content

**Chapter 1. Bootstrapping your Ruby literacy**:

I would say the whole chapter is a must read for new Ruby, but I would place special importance on:

`1.2.3. The site_ruby (RbConfig::CONFIG[sitedir]) and vendor_ruby (RbConf- fig::CONFIG[vendordir]) directories` - Useful knowledge on where Ruby is located and how it loads/requires files. This applies for `rvm` and `rbenv` users since `site_ruby` and `vendor_ruby` are part of Ruby.

`1.3.3. “Require”-ing a feature` - explains the difference between `load` and `require`.

**Chapter 2. Objects, methods, and local variables**:

`2.3.1. Identifying objects uniquely with the object_id method` -
shows how *almost* everything in Ruby is an object as said in [Seeing Everything as an Object at ruby-lang.org](https://www.ruby-lang.org/en/about/). Technically not everything can be an Object in Ruby - statements, methods, operators and blocks aren't. Read `Flaws` section of the post!

`2.3.3. Sending messages to objects with the send method` - sending messages is a core concept in Ruby. Beyond the basics this section has a must-read good code style subsection `Using __send__ or public_send instead of send`. Safety with `send` is required and this subsection delivers good advice!

**Chapter 3. Organizing objects with classes**

`3.4.2. Summary of attr_* methods` - really neat and concise `attr_*` method and equivalent method with a body code comparison.

`3.6 (the whole subsection)` - really well explains the *blurry* definition of class methods in Ruby and compares them to instance methods.

**Chapter 4. Modules and program organization**:

`4.1.2 Mixing a module into a class - Syntax of require/load vs. syntax of include` - I doubt that anyone expects including a module to be anything different than loading from the `memory`, but anyway it's mentioned and explained here.

`4.2. Modules, classes, and method lookup` - the whole section and its subsections about the process of method lookup when mixing in modules are core Ruby knowledge. The author does a good job at explaining it even though I have some mentions about it in `Flaws`. Also `4.2.2 Defining the same method more than once` and `4.2.3 How prepend works` explains the quirks of `prepend`, include order and including the same module twice and what actually would happen.

`4.4 Class/module design and naming` - the most important part of the chapter, when to use a module or a class.

**Chapter 5. The default object (self), scope, and visibility**:

`5.1.2 The top-level self object` - just good to know why you can call Object methods in global scope and they would work. You can also define instance variables and instance_variables in the global scope and it would work. Knowing that this *would* work and is actually related to the global scope would save you troubles when debugging.

`5.1.3 Self inside class, module and method definitions` - as title says, it's immensely useful.

`5.2.2 Local scope` - expected block scope and standard OO scoping, but the code graph makes it really easier to understand for newcomers. A good chapter!

`5.2.4 Scope and resolution of constants` - as title says and also the paragraph about how to refer to the STL `String` if you're in a class scope where you have a class of your own `String`, to refer to the STL one you say `::String`. Which isn't exclussive for referring to STL classes, but for starting the search from the top-level down. And of course this would hit the STL library first.

`5.2.5 Class variables syntax, scope and visibility` - really well demistifies some of the weird parts of class variables.

`5.3 Deploying method access-rules` - as title says. The `Private and singleton are different` section is handy knowledge and something not so well-known beforehand.

`6.3 Iterators and code blocks` - Ruby's interesting take on iterators and code blocks is a mandatory read. The whole section is key to understanding code blocks.


## Flaws

However through the books I've found some code examples and sections that need additional explanation.

`2.3.1. Identifying objects uniquely with the object_id method`:

It's not mentioned in the section, but `object_ids` can be reused for optimization as mentioned in http://ruby-doc.org/core-2.2.2/Object.html#method-i-object_id . In fact it's written that "every object in Ruby has an unique ID number associated with it". As seen in the docs, that's not always correct and may lead to confusion.


One of them:

```ruby
c = Class.new do
  def say_hello
    puts "Hello!"
  end
end
```

You are supposed to create a class that can call `say_hello`. However `c` is a class, not an instance.
Therefore the author's c#say_hello won't do any good, since you need to instantiate the class. first and then call `say_hello` on the instance.


----

In section `3.6.5 Class methods vs instance methods`, subsection `A note on method notation`,
it's useful to note that `Ticket::most_expensive` and the double colon notation is actually valid and used in Ruby code.

And also there are code graph examples that aren't inconsistent with the text.

To mention one - the subsection in `4.2.1 Illustrating the basics of method lookup`.
The text describes a situation where `D` is a subclass of `C` which has a mixin `M`. `D` does not have a mixin.
However in the code graph class `D` has a mixin `M` and `C` doesn't. Therefore the code graph and text only lead to confusion even though what they explain is correct.
 The way it's presented makes people double-check it for 2minutes before making sense of it.

`4.2.3 How prepend works` - code graph that shows how prepend works reuses old code graph with a slight change of prepend in one of the examples. Nevertheless the code example given in the previous page and the code graph use different classes. It should be consistent and either the code example should be changed to match the obfuscated naming `D`, `C` and `M` or change the code graph to match the previous page's
code example.

`4.2.4 The rules of method lookup summarized` - includes the recently introduced `prepend` to the previous method lookup flow and shows the whole flow of method lookup.

`Chapter 4` - What I would've liked to read in the chapter is the difference between `include` and `extend`. It's mandatory to know. [StackOverflow to the rescue](http://stackoverflow.com/questions/156362/what-is-the-difference-between-include-and-extend-in-ruby). But one would expect to be mentioned here.


`5.1.3 Self inside class, module and method definitions, subsection "Using self instead of hard-coded class names"` - ok this may be an exaggeration my problem with this is that it's forced. You normally wouldn't need a section like this to be even mentioned if you didn't provide code examples with the bad practice of defining class methods using `ClassName.method_name` instead of the better style of just used `self.method_name`.

`Chapter 5` - this may be a stretch but I haven't found any mention of the difference between using `self.instance_var` or `@instance_var`. [StackOverflow to the rescue](http://stackoverflow.com/questions/1693243/instance-variable-self-vs). If you're familiar with Python property descriptors, you'll find a reason to use the `self.instance_var` way.

----

## What to improve

My two top suggestions to the book are:

* better define the idea you try to prove. Either by rephrasing or explaining in more sentences. The code examples with arrows and numbers are great. The author definitely does good explanations most of the time!
* have all the book code examples in the provided book source code. Not all of the code examples are in the book source code.
