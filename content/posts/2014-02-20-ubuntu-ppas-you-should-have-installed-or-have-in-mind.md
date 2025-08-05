---
title: "Ubuntu: PPAs you should have installed or have in mind"
date: 2014-02-20T11:21:29Z
draft: false
categories:
  - "ubuntu"
  - "ppa"
---

<h2><a href="http://syndbg.wordpress.com/2014/02/20/making-the-switch-to-ubuntu-how-to-install-ppas/">How to install PPAs</a></h2>
<h2>Faster updates and installs: apt-fast</h2>
<h3>What is apt-fast:</h3>
<p>A script that uses command line download accelerators. It drastically
improves updating and installing</p>
<h3>How to install:</h3>
<p><code>sudo add-apt-repository ppa:apt-fast/stable sudo apt-get update sudo apt-get install apt-fast</code></p>
<p>In future you use apt-fast everywhere you would've typed apt-get</p>
<h3>Example</h3>
<p><code>sudo apt-fast update sudo apt-fast upgrade sudo apt-fast install sudo apt-fast autoremove</code></p>
<h2>Manage PPAs easily: y-ppa-manager</h2>
<h3>What is y-ppa-manager:</h3>
<p>A tool for easy management of PPAs. You can remove/purge, update and
edit source, as in the screenshot.</p>
<p><a href="http://syndbg.files.wordpress.com/2014/02/y-ppa-manager.png"><img alt="y-ppa-manager" src="http://syndbg.files.wordpress.com/2014/02/y-ppa-manager.png?w=300" /></a></p>
<h3>How to install:</h3>
<p><code>sudo add-apt-repository ppa:webupd8team/y-ppa-manager sudo apt-get update sudo apt-get install y-ppa-manager</code></p>
<h2>Extend battery life: tlp</h2>
<h3>What is tlp:</h3>
<p>A command line tool that automatically applies tweaks to reduce your
battery usage</p>
<h3>How to install:</h3>
<p><code>sudo add-apt-repository ppa:linrunner/tlp sudo apt-get update sudo apt-get install tlp tlp-rdw</code></p>
<p>It will automatically start at system start-up after you restart for the
first time. However if you want to start it as fast as possible:</p>
<p>\&lt;code<em>sudo tlp start</em></p>
<h2>Java7: oracle-java7-installer or openjdk-7-jdk</h2>
<h3>What is the difference between JRE and JDK:</h3>
<p><strong>JRE (Java Runtime Environment)</strong> installs the Java Virtual Machine and
enables you to use any java program on your machine</p>
<p><strong>JDK (Java Development Kit)</strong> does everything JRE does and also
installs the development tools to compile, debug and document your java
programs.</p>
<p>Installing JDK from the get-go is better, so I'll show you how to
install JDK in case you decide to develop java programs in the future.</p>
<h3><a href="http://stackoverflow.com/a/1977354">What is the difference between openjdk and oracle jdk?</a></h3>
<h3>How to install Oracle Java7 JDK:</h3>
<p><code>sudo add-apt-repository ppa:webupd8team/java sudo apt-get update sudo apt-get install oracle-java7-installer</code></p>
<h3>How to install OpenJDK Java7 JDK:</h3>
<p><code>sudo apt-get install openjdk-7-jdk</code></p>
<h3>How to test if it's installed correctly:</h3>
<p><strong>Try running:</strong></p>
<p><code>java -version</code></p>
<p><strong>It should return something similar to:</strong>  </p>
<p><code>java version "1.7.0_51" OpenJDK Runtime Environment (IcedTea 2.4.4) (7u51-2.4.4-0ubuntu0.13.10.1) OpenJDK 64-Bit Server VM (build 24.45-b08, mixed mode)</code></p>
<h2>Development with Sublime Text: sublime-text-2 or sublime-text-3</h2>
<h3>Difference between ST2 and ST3:</h3>
<p>Some plugins don't have support for ST3 yet. If you're just starting to
use Sublime Text you won't care about the support of plugins you don't
know, so go for ST3.<br />
But if you're already a ST2 user, wondering if you can switch to ST3
-> check with <a href="http://www.caniswitchtosublimetext3.com/">Can I switch to Sublime Text
3?</a></p>
<h3>How to install ST2 anyway</h3>
<p><code>sudo add-apt-repository ppa:webupd8team/sublime-text-2 sudo apt-get update sudo apt-get install sublime-text-installer</code></p>
<h3>How to install ST3:</h3>
<p><code>sudo add-apt-repository ppa:webupd8team/sublime-text-3 sudo apt-get update sudo apt-get install sublime-text-installer</code></p>
<h2>Stop hurting your eyes, use Redshift/f.lux: gtk-redshift or fluxgui</h2>
<h3>What do both programs do?</h3>
<p>They adjust the color temperature of the screen to adapt to the day
time. If you're an early rise or a night owl (like me), you should stop
blinding yourself and use Redshift or f.lux</p>
<h3>What to install</h3>
<p>I suggest Redshift because f.lux isn't supporting Ubuntu 13.10 and
probably won't support future versions (if not updated by the developer)</p>
<h3>How to install</h3>
<p><code>sudo add-apt-repository ppa:jonls/redshift-ppa sudo apt-get update sudo apt-get install gtk-redshift</code></p>
<p>I suggest gtk-redshift over redshift. gtk-redshift has a toggle tray so
you can toggle it on/off and quit it at any time</p>
<h3>How to use it</h3>
<p><a href="http://www.latlong.net/">Get your latitude and longitude</a><br />
<strong>Example command:</strong>  </p>
<p><code>gtk-redshift -l YOUR_LATITUDE:YOUR_LONGITUDE -t DAY_TEMP:NIGHT_TEMP -g GAMMA -v</code><br />
<strong>You can use my settings replaced with your latitude/longitude:</strong><br />
<code>gtk-redshift -l 42.697839:23.321670 -t 5700:3600 -g 0.8 -v</code></p>
<h3>That's all. If you have any suggestions, comment and I'll add them</h3>