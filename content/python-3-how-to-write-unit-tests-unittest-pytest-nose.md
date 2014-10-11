Date: 2014-04-20
Title: Python 3 How to: Write Unit tests (PyUnit, pytest, nose)
Tags:
Slug: python-3-how-to-write-unit-tests-unittest-pytest-nose


<h3>In the following blog post I'll cover the following frameworks for unit testing.</h3>
<p>​1) <a href="http://pyunit.sourceforge.net/" title="PyUnit ">PyUnit</a>, also known as
unittest in the Python modules.
2) <a href="http://pytest.org/latest/" title="Pytest">Pytest</a> - a mature
full-featured Python testing tool that helps you write better programs.
3) <a href="http://nose.readthedocs.org" title="Nose">Nose</a> - a nicer testing for
Python? It's more like a unit test loader since it supports all the
above for loading.</p>
<h3>How to get them?</h3>
<p>​1) PyUnit - already installed for you. Just import unittest in your
unit tests file</p>
<p>​2) Pytest - devs guide
-> <a href="http://pytest.org/latest/getting-started.html" title="Pytest install guide">here</a></p>
<p>​3) Nose - devs guide ->
<a href="https://nose.readthedocs.org/en/latest/" title="Nose install guide">here</a></p>
<h3>Get to know TDD:</h3>
<p>TDD - Test-Driven Development or also known as <em>Red, Green and
Refactor. </em>Can be explained simple in a image.</p>
<p><a href="http://thejackalofjavascript.com/wp-content/uploads/2014/02/tdd_flow.gif"><img alt="TDD flow" src="http://thejackalofjavascript.com/wp-content/uploads/2014/02/tdd_flow.gif" /></a></p>
<h3>How to start writing tests with PyUnit:</h3>
<p>​1) Import unittest</p>
<p>​2) Import the code you want to test</p>
<p>​3) Create a class that inherits  from unittest.TestCase</p>
<p>​4) Create methods for the class you created that MUST start with test</p>
<p>​5) Run the tests with the python3 interpreter in the console or your
IDE</p>
<p><strong><a href="https://docs.python.org/3.4/library/unittest.html#assert-methods" title="PyUnit methods">Useful methods to
know</a> (framework's
API)</strong></p>
<p><strong><a href="https://github.com/syndbg/Mini-projects/blob/master/Python/Introduction%20to%20unit%20testing%20frameworks/pyunit_bankaccount_simplified.py" title="PyUnit source code">Unit test source
code</a></strong></p>
<p><strong>And when run in the terminal</strong></p>
<p><a href="http://i.imgur.com/kggP1vj.png"><img alt="PyUnit terminal showcase" src="http://i.imgur.com/kggP1vj.png" /></a></p>
<h3>How to start writing tests with Pytest:</h3>
<p>​1) Install pytest</p>
<p>​2) Import the  code you want to test</p>
<p>​3) Write test functions, it's optional to create a class with methods.</p>
<p>​4) Run the tests with <em>py.test</em> in the console</p>
<p>For Pytest all you need to know is how Python 3's <em>assert</em> <em>works.
 </em><a href="http://pytest.org/latest/assert.html" title="Pytest methods">Read a little more about Pytest and
assert</a></p>
<p><strong><a href="https://github.com/syndbg/Mini-projects/blob/master/Python/Introduction%20to%20unit%20testing%20frameworks/pytest_bankaccount_simplified.py" title="PyTest source code">Unit test source code<strong> (framework's
API)</strong></a></strong></p>
<p><strong>And when run in the terminal</strong></p>
<p><a href="http://i.imgur.com/Cpfc0Ed.png"><img alt="Pytest terminal showcase" src="http://i.imgur.com/Cpfc0Ed.png" /></a></p>
<h3>How to start writing tests with Nose:</h3>
<p>​1) Install nose</p>
<p>​2) Import the code you want to test</p>
<p>​3) Import nose.tools</p>
<p>​4) Create a class that starts wtih Test</p>
<p>​5) Create tests (methods) that start with test</p>
<p>​6) Run the test with <em>nosetests</em> in the console</p>
<p><a href="http://nose.readthedocs.org/en/latest/testing_tools.html#module-nose.tools" title="Nose methods"><strong>Useful methods  to
know</strong></a> .<strong> (framework's
API)</strong></p>
<p><strong>Most of the time you'll use:</strong></p>
<p><strong>*  <em>ok_</em> which is a short-hand for *assert</strong>*</p>
<p><strong>*  eq_ which is (a, b) and tests if a == b</strong></p>
<p><strong><a href="https://github.com/syndbg/Mini-projects/blob/master/Python/Introduction%20to%20unit%20testing%20frameworks/nose_bankaccount_simplified.py" title="Nose source code">Unit test source
code</a></strong></p>
<p><strong>And when run in the terminal</strong></p>
<p><a href="http://i.imgur.com/wCaBrhm.png"><img alt="Nose showwcase" src="http://i.imgur.com/wCaBrhm.png" /></a></p>
