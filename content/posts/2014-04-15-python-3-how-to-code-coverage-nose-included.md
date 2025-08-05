---
title: "Python 3 How to: Code coverage (nose way included)"
date: 2014-04-15T11:21:29Z
draft: false
categories:
  - "python"
  - "testing"
  - "coverage"
---

<p><strong>First of all, you must download and install
<a href="https://pypi.python.org/pypi/coverage/3.7.1">coverage.py</a> with pip.</strong></p>
<h3 style="text-align:left;">The steps are:</h3>
<p>​1) First have written tests for the Python program you want to test
2) Run code coverage on the unit tests you've written.
3) Export the report to html for better readability and more details.</p>
<h4 style="text-align:left;">and in code:</h4>
<p><code>coverage run test_something.py</code>
<code>coverage report -m</code>
<code>coverage html</code></p>
<p><strong>The code coverage can now be found in a directory htmlconv/ which is
in the same directory where you run the coverage commands. Open
htmlconv/index.html and see the detailed html version of the code
coverage.</strong></p>
<p><em>Note: the -m argument is to show a tab for the missing code coverage</em></p>
<h3>And how it looks in action:</h3>
<p><a href="http://i.imgur.com/l1JKolA.png" title="Code Coverage in action picture"><img alt="Code coverage showcase" src="http://i.imgur.com/l1JKolA.png" /></a></p>
<h2>The nose way</h2>
<p>​1) You must have nose installed. (pip install nose)
2) Again, you must have written unit tests.
3) Run the code coverage on the unit tests. Done! All in one command :)</p>
<h4 style="text-align:left;">and in code:</h4>
<p><code>nosetests --with-coverage --cover-html test_something.py</code></p>
<p><strong>The code coverage can now be found in a directory cover/ which is in
the same directory where you run the coverage commands. Open
cover/index.html and see the detailed html version of the code
coverage.</strong>
<em>Note: Nose code coverage doesn't include the unit test file in the
code coverage. To include it add <code>--cover-tests</code> to the command you
execute.</em></p>
<p><a href="http://i.imgur.com/074oiIV.png" title="Nose code Coverage in action picture"><img alt="Nose code coverage showcase" src="http://i.imgur.com/074oiIV.png" /></a></p>