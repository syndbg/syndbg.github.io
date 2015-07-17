---
layout:     post
title:      "Making the switch to Ubuntu. How to install PPAs"
date:       2014-02-20 11:21:29
summary:
categories: ubuntu ppa
---

<p>If you want instructions how to make a liveCD/liveUSB and how to install
Ubuntu head -
<a href="http://www.ubuntu.com/download/desktop/install-desktop-latest">here</a></p>
<h2>What is a PPA?</h2>
<p>Personal Package Archive, stuff that isn't included by default in the
OS, mostly developed by people not affiliated to the official developers
of Ubuntu (in our case).  INSTALL PPAs WITH CARE! You've been warned.</p>
<p>First of all open the Terminal (<em>CTRL + ALT + T</em>)</p>
<p><strong>Mostly everything you'll do in the future will involve around these
commands: </strong></p>
<h2>Order of installing a PPA:</h2>
<p>1) <em>sudo add-apt-repository</em></p>
<p>2) <em>sudo apt-get update</em></p>
<p>3) <em>sudo apt-get install</em></p>
<h2>What do they mean?</h2>
<p><strong>sudo</strong> - super user (root access). Everything power user-ish, like
installing, updating through the terminal is involved with sudo.</p>
<p><strong>apt</strong> - advanced packaging tool (it's actually an interface).
Everything with apt- prefix is a tool which uses apt. Why use apt?
Because it's the easiest way to install stuff on Linux, without scaring
yourself to death by the new OS if you're coming from Windows (like me)
or Mac</p>
<p><strong><em>sudo add-apt-repository </em></strong><em>- adds a repository to the OS, but
doesn't install it.</em></p>
<p><em>Example: sudo
add-apt-repository </em>http://ppa.launchpad.net/webupd8team/java/ubuntu</p>
<p><strong><em>sudo apt-get update</em></strong> - updates your software. Including stuff you
just installed<strong>
</strong></p>
<p><em>Example: sudo apt-get update</em></p>
<p><strong><em>sudo apt-get install</em></strong><strong> </strong>- yep, installs.<strong>
</strong></p>
<p><em>Example: sudo apt-get install</em> <em>oracle-java7-installer</em></p>
<p><strong> </strong></p>
<h2>A full step by step how to find and install a PPA</h2>
<p>​1) Finding. Good source for PPAs
is <a href="https://launchpad.net/~webupd8team">https://launchpad.net/~webupd8team</a> and
follow <a href="http://www.webupd8.org/">http://www.webupd8.org/</a></p>
<p>​2) Found something interesting, how to install it?</p>
<p><strong>If I want to install Java7 on my machine</strong></p>
<p>I go here <a href="https://launchpad.net/~webupd8team/+archive/java">https://launchpad.net/~webupd8team/+archive/java</a></p>
<p>Click the green <a href="https://launchpad.net/~webupd8team/+archive/java/+index#">Technical details about this
PPA</a></p>
<p>Copy only and only <a href="http://ppa.launchpad.net/webupd8team/java/ubuntu">http://ppa.launchpad.net/webupd8team/java/ubuntu</a></p>
<p>Open the Terminal (<em>CTRL+ALT+T</em>)</p>
<p>Write these lines:</p>
<p>1) <em>sudo
add-apt-repository <a href="http://ppa.launchpad.net/webupd8team/java/ubuntu">http://ppa.launchpad.net/webupd8team/java/ubuntu</a></em></p>
<p>​2) sudo apt-get update</p>
<p>​3) Now you wonder what can I install ? See
in <a href="http://ppa.launchpad.net/webupd8team/java/ubuntu">http://ppa.launchpad.net/webupd8team/java/ubuntu</a> below
the <strong>Overview of published packages</strong></p>
<p>In our case we install Java7, so we must see oracle-java7-installer</p>
<p>Then we write: <em>sudo apt-get install oracle-java7-installer</em></p>
<p>​4) Done! But anyway test with writing: <em>java</em></p>
<p>If no error shows up but documentation, it works!</p>
<p><a href="http://syndbg.wordpress.com/2014/02/20/ubuntu-ppas-you-should-have-installed-or-have-in-mind/">Ubuntu: PPAs you should have installed or have in mind</a></p>
