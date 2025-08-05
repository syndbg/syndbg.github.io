---
title: "Programming tips: Avoiding unbalanced parentheses errors"
date: 2014-02-09T11:21:29Z
draft: false
categories:
  - "tip"
  - "parentheses"
---

<p>It's very common for people to have unbalanced parenthesis after hours
of work and hundreds of lines of code.<br />
Also it's one of the more common questions in exams, find the
error/errors in the following code "...." and they give you way too many
parentheses.</p>
<p>One of the more popular solutions is to use a pencil and paper to follow
the code and see the mistake, but this takes resources and much more
time than it should.</p>
<p>The easier and better way is to count. Count what? Count the
parentheses.<br />
Any opening/left parenthesis is +1, any closing/right parenthesis is
-1. The result should always be 0 and never negative!<br />
So let's put this into examples.</p>
<div class="highlight"><pre><span class="p">((</span><span class="n">a</span> <span class="o">+</span> <span class="n">b</span><span class="p">)</span> <span class="o">*</span> <span class="n">t</span> <span class="o">/</span> <span class="mi">2</span> <span class="o">*</span> <span class="p">(</span><span class="mi">1</span> <span class="o">-</span> <span class="n">t</span><span class="p">)</span>
<span class="mi">1</span> <span class="mi">1</span>    <span class="o">-</span><span class="mi">1</span>           <span class="mi">1</span>    <span class="o">-</span><span class="mi">1</span>
</pre></div>


<p>And we notice that we have 1 remaining, which means the code is wrong.</p>
<p>Next example</p>
<div class="highlight"><pre><span class="p">(</span><span class="n">a</span> <span class="o">+</span> <span class="n">b</span><span class="p">)</span> <span class="o">*</span> <span class="n">t</span><span class="p">)</span> <span class="o">/</span> <span class="p">(</span><span class="mi">2</span> <span class="o">*</span> <span class="p">(</span><span class="mi">1</span> <span class="o">-</span> <span class="n">t</span><span class="p">)</span>
           <span class="err">↑</span>
<span class="o">+</span><span class="mi">1</span>   <span class="o">-</span><span class="mi">1</span>    <span class="o">-</span><span class="mi">1</span>
</pre></div>


<p>.... wait what? We get -1. If we get -1 the code is 100% wrong.</p>
<p>Good, now the time needed to track and correct parentheses is greatly
reduced!</p>