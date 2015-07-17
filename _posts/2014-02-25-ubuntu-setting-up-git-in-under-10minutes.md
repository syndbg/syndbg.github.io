---
layout:     post
title:      "Ubuntu: Setting up git in under 10minutes"
date:       2014-02-25 11:21:29
summary:
categories: ubuntu git
---

<h2><a href="http://cli.learncodethehardway.org/book/">Become familiar with the Terminal beforehand.</a></h2>
<h2><a href="http://en.wikipedia.org/wiki/Git_(software)">What is git?</a></h2>
<h2><a href="http://en.wikipedia.org/wiki/GitHub">What is Github?</a></h2>
<h2>Setting it up:</h2>
<p>​1) Register at <a href="https://github.com/join‎">github</a> and login.<br />
2) Open the Terminal, install git<br />
<code>sudo apt-get install git sudo apt-get update</code><br />
3) Setup your git with your github username and email:  </p>
<p><code>git config --global user.name "Your Name" git config --global user.email "you@somewhere.com"</code></p>
<h2>What do you want to do?</h2>
<h3>Start my own repository and upload code as fast as possible:</h3>
<p>http://www.youtube.com/watch?v=ABYHBOE5oU4</p>
<p><a href="https://github.com/new">Create a new repository</a>  </p>
<p><code>mkdir YOURrepository cd YOURrepository git init gedit README.md git add README.md gedit Hello.py git add Hello.py git status git commit -m "init or whatever" git remote add origin https://github.com/yourUSERNAME/yourREPOSITORY.git git push -u origin master git status</code></p>
<h2>And if you still feel uncomfortable to use Terminal git, <a href="http://www.maketecheasier.com/6-useful-graphical-git-client-for-linux/">here are some graphical git clients for Linux</a></h2>
