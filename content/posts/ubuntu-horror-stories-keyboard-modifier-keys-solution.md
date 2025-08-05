---
title: "Ubuntu bug solution: Keyboard's modifier keys aren't working"
date: 2014-02-20T11:21:29Z
draft: false
categories:
  - "ubuntu"
---

<h2>What are modifier keys?</h2>
<p><strong>Alt</strong>, <strong>Control</strong>, <strong>Shift</strong> and<strong>Windows key</strong> are modifier keys.
They're mostly used in combination with other keys.</p>
<p>Examples: Ctrl+Alt+Delete, Alt+Tab and others...</p>
<h2>The story:</h2>
<p>As an owner of <a href="https://www.google.bg/search?q=zowie+celeritas&amp;oq=zowie+celeritas&amp;aqs=chrome.0.69i59.1229j0j4&amp;sourceid=chrome&amp;ie=UTF-8">Zowie
Celeritas</a> I
understood the hard way that my keyboard isn't really supported by the
Linux kernel.</p>
<p>Any modifier keys combinations don't work.</p>
<p>The problem is <strong>not Ubuntu-specific, </strong>I tried Mint and Debian -
problem still persists.</p>
<p>This is a common problem for other keyboards too, so the solution may or
may not work for you.</p>
<h2>The solution:</h2>
<p><strong>If your keyboard is PS/2</strong> -> you shouldn't have a problem to begin
with. But you can still try with an adapter from PS/2 to USB, to see if
your keyboard works.</p>
<p><strong>If your keyboard is USB</strong> -> use a USB to PS/2 adapter.
<a href="http://www.mzoori.com/images/product/verybigs/ce91fffc544bdce5c10ba61e1565f0a9.jpg">This</a>
has worked for many</p>
<p><strong>If your keyboard is USB but you don't have a PS/2 port</strong> <strong>(Zowie
Celeritas and other mechanical keyboards in combination with a
laptop)</strong> -> If you don't have a USB to PS2 adapter buy the above one,
else use the one provided with your keyboard. Afterwards buy <a href="http://www.usbwholesale.com/images/ps2%20usb%20c.jpg">this PS/2
to USB adapter</a>or
a similar one. It must be a real adapter and convert the signal.  The
one I provided worked for me and many others.</p>
<p><strong>Good luck and I hope you resolve this problem!</strong></p>