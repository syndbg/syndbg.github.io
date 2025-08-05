---
title: "Atom: Setting up for Python Development"
date: 2014-08-23T11:21:29Z
draft: false
categories:
  - "atom"
  - "python"
  - "development"
description: "Setting up Atom for Python Development"
---

<h2>1. Installation</h2>
<p>Taken from <a href="http://www.webupd8.org/2014/06/atom-text-editor-available-for-linux.html">http://www.webupd8.org/2014/06/atom-text-editor-available-for-linux.html</a></p>
<div class="highlight"><pre><span class="n">sudo</span> <span class="n">add</span><span class="o">-</span><span class="n">apt</span><span class="o">-</span><span class="n">repository</span> <span class="n">ppa</span><span class="o">:</span><span class="n">webupd8team</span><span class="o">/</span><span class="n">atom</span>
<span class="n">sudo</span> <span class="n">apt</span><span class="o">-</span><span class="n">get</span> <span class="n">update</span>
<span class="n">sudo</span> <span class="n">apt</span><span class="o">-</span><span class="n">get</span> <span class="n">install</span> <span class="n">atom</span>
</pre></div>


<h2>2. GitHub recommended first steps</h2>
<p>The following guides are a must-read:</p>
<ul>
<li><a href="https://github.com/atom/atom/blob/master/docs/getting-started.md">Getting Started</a></li>
<li><a href="https://github.com/atom/atom/blob/master/docs/customizing-atom.md">Customizing Atom</a></li>
<li><a href="https://github.com/atom/atom/blob/master/docs/debugging.md">Debugging</a></li>
</ul>
<h2>3. Packages for Python development</h2>
<p>By default Atom has Python language syntax highlighting and snippets, but lacks a linter.
I recommend <a href="https://bitbucket.org/tarek/flake8/src">flake8</a>, which is a wrapper for <a href="https://pypi.python.org/pypi/pyflakes">pyflakes</a>, <a href="https://pypi.python.org/pypi/pep8">pep8</a> and <a href="https://pypi.python.org/pypi/mccabe">Ned Batchelderâs McCabe script</a>.</p>
<h3>3.1. Install pypi package <a href="https://pypi.python.org/pypi/flake8">flake8</a></h3>
<p><em>Note</em>: have pip pre-installed</p>
<div class="highlight"><pre><span class="n">sudo</span> <span class="n">pip</span> <span class="n">install</span> <span class="n">flake8</span>
</pre></div>


<h3>3.2. Install Atom packages - <a href="https://atom.io/packages/linter">linter</a> and <a href="https://atom.io/packages/linter-flake8">linter-flake8</a></h3>
<p><code>linter</code> is a static code analysis tool and <code>linter-flake8</code> is an addon to <code>linter</code> and also an interface for <a href="https://bitbucket.org/tarek/flake8/src">flake8</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">linter</span>
<span class="n">apm</span> <span class="n">install</span> <span class="n">linter</span><span class="o">-</span><span class="n">flake8</span>
</pre></div>


<h2>4. Development settings</h2>
<p>These can be modified through the Settings panel, but it's faster if we just edit <code>~/.atom/config.cson</code></p>
<div class="highlight"><pre><span class="s">&#39;welcome&#39;</span><span class="o">:</span>
  <span class="s">&#39;showOnStartup&#39;</span><span class="o">:</span> <span class="kc">false</span>
<span class="s">&#39;metrics&#39;</span><span class="o">:</span> <span class="kc">false</span>
<span class="s">&#39;editor&#39;</span><span class="o">:</span>
  <span class="s">&#39;preferredLineLength&#39;</span><span class="o">:</span> <span class="mi">80</span>
  <span class="s">&#39;showInvisibles&#39;</span><span class="o">:</span> <span class="kc">true</span>
  <span class="s">&#39;tabLength&#39;</span><span class="o">:</span> <span class="mi">4</span>
  <span class="s">&#39;scrollSensitivity&#39;</span><span class="o">:</span> <span class="mi">60</span>
  <span class="s">&#39;showIndentGuide&#39;</span><span class="o">:</span> <span class="kc">true</span>
<span class="s">&#39;core&#39;</span><span class="o">:</span>
  <span class="s">&#39;disabledPackages&#39;</span><span class="o">:</span> <span class="p">[</span>
    <span class="s">&#39;welcome&#39;</span>
    <span class="s">&#39;styleguide&#39;</span>
    <span class="s">&#39;spell-check&#39;</span>
    <span class="s">&#39;metrics&#39;</span>
    <span class="s">&#39;feedback&#39;</span>
  <span class="p">]</span>
  <span class="s">&#39;ignoredNames&#39;</span><span class="o">:</span> <span class="p">[</span>
    <span class="s">&#39;*.pyc&#39;</span>
    <span class="s">&#39;*.*~&#39;</span>
  <span class="p">]</span>
<span class="s">&#39;tree-view&#39;</span><span class="o">:</span>
  <span class="s">&#39;showOnRightSide&#39;</span><span class="o">:</span> <span class="kc">true</span>
<span class="s">&#39;linter-pep8&#39;</span><span class="o">:</span>
  <span class="s">&#39;ignoreErrorCodes&#39;</span><span class="o">:</span> <span class="p">[</span>
    <span class="s">&#39;W0232&#39;</span>
  <span class="p">]</span>
<span class="s">&#39;linter-flake8&#39;</span><span class="o">:</span>
  <span class="s">&#39;maxLineLength&#39;</span><span class="o">:</span> <span class="mi">80</span>
<span class="s">&#39;linter&#39;</span><span class="o">:</span>
  <span class="s">&#39;showStatusBarWhenCursorIsInErrorRange&#39;</span><span class="o">:</span> <span class="kc">true</span>
</pre></div>


<h2>5. Other useful development packages</h2>
<ul>
<li>
<p><a href="https://github.com/atom/sort-lines">Sort-lines</a> - sorts lines and it never gets tiring.</p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">sort</span><span class="o">-</span><span class="n">lines</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/mizchi/atom-git-grep">Git-grep</a> - performs a grep on your whole repo.</p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">git</span><span class="o">-</span><span class="n">grep</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/richrace/highlight-line">Highlight-line</a> - highlights the current line with some extra options as underlining with solid, dotted or solid line.</p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">highlight</span><span class="o">-</span><span class="n">line</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/smashwilson/merge-conflicts">Merge-conflicts</a> - a more visual way to resolve merge conflicts.</p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">merge</span><span class="o">-</span><span class="n">conflicts</span>
</pre></div>


</li>
</ul>
<h2>6. Nicer themes and syntax highlighting</h2>
<p>These are recommended by me, but many more can be found at <a href="https://atom.io/themes/">Atom.io themes</a></p>
<h3>6.1 UI themes</h3>
<ul>
<li>
<p><a href="https://github.com/jesseweed/seti-ui">Seti-UI, a dark colored UI theme</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">seti</span><span class="o">-</span><span class="n">ui</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/jesseweed/yeti-ui">Yeti-UI, a light colored UI theme</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">yeti</span><span class="o">-</span><span class="n">ui</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/dmackerman/atom-soda-dark-ui">Soda-dark-UI, a port of Sublime Text 3's Soda Dark UI theme</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">atom</span><span class="o">-</span><span class="n">soda</span><span class="o">-</span><span class="n">dark</span><span class="o">-</span><span class="n">ui</span>
</pre></div>


</li>
<li>
<p><a href="https://atom.io/themes/unity-ui">Unity-UI, a more native OSX experience</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">unity</span><span class="o">-</span><span class="n">ui</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/brentd/neutron-ui">Neutron-UI, an almost flat UI theme in shades of gunmetal and medium-contrast colors</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">neutron</span><span class="o">-</span><span class="n">ui</span>
</pre></div>


</li>
</ul>
<h3>6.2 Syntax highlighting themes</h3>
<ul>
<li>
<p><a href="https://github.com/kevinsawicki/monokai">Monokai</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">monokai</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/jesseweed/seti-syntax">Seti-syntax</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">seti</span><span class="o">-</span><span class="n">syntax</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/joaoafrmartins/seti-monokai">Seti-monokai</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">seti</span><span class="o">-</span><span class="n">monokai</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/jesseweed/yeti-syntax">Yeti-syntax</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">yeti</span><span class="o">-</span><span class="n">syntax</span>
</pre></div>


</li>
<li>
<p><a href="https://github.com/brentd/neutron-syntax">Neutron-syntax</a></p>
<div class="highlight"><pre><span class="n">apm</span> <span class="n">install</span> <span class="n">neutron</span><span class="o">-</span><span class="n">syntax</span>
</pre></div>


</li>
</ul>
<h2>7. Useful custom keybinds</h2>
<p>These are the custom keybinds I use. Edit ~/.atom/keymap.cson</p>
<div class="highlight"><pre><span class="s">&#39;.workspace&#39;</span><span class="o">:</span>
  <span class="s">&#39;ctrl-G&#39;</span><span class="o">:</span> <span class="s">&#39;git-grep:grep&#39;</span>

<span class="s">&#39;.body&#39;</span><span class="o">:</span>
  <span class="s">&#39;ctrl-O&#39;</span><span class="o">:</span> <span class="s">&#39;application:open-folder&#39;</span>
</pre></div>


<p>In the end, this is how my Atom looks like with Seti-ui and Seti-monokai
<a href="https://raw.githubusercontent.com/syndbg/syndbg.github.io/master/static/atom_configured.png"><img alt="Configured atom" src="https://raw.githubusercontent.com/syndbg/syndbg.github.io/master/static/atom_configured.png" title="Configured Atom" /></a></p>