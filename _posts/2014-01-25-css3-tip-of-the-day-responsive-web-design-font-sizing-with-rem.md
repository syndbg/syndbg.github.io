---
layout:     post
title:      "CSS3 tip of the day: Responsive web design - Font-sizing with rem"
date:       2014-01-25 11:21:29
summary:    CSS3 rem font-sizing
categories: css3 font sizing
---

<p>To create good responsive web design you must use proportionally-sized
text units like <em>em </em>and<em> % (percents).</em></p>
<p>However, CSS3 introduces a new proportionately-sized text
unit <strong><em>rem</em></strong>.</p>
<p><strong><em>rem</em> </strong>- stands for <em>root em</em>.</p>
<p>1em = 1 x current font size. If current font-size is 15px. 1em = 15px
and 2em = 30px and so on.</p>
<p>But! <em>rem </em>is always calculated relative to the font size of the
\&lt;html> element. Not the font size of the containing element.  3rem is
always 3 times the font size, wherever you apply it.</p>
<p><strong>Which browsers support this? </strong>Chrome, Firefox3.6+, Safari 5+, IE9+,
Opera 11.10+ (but not 11.10)</p>
<p><strong>The old common em calculation problem:</strong>It always gets complicated to
calculate the resulting font sizing in pixels after applying em. Sure
it's easy to use 1 or 2em but what happens when you need more accurate
sizing like 1.64, 0.67 ... ?</p>
<p>Well, people have thought of a solution back in the days of <em>em</em>.
Fortunately, the same solution can be applied to <em>rem.</em></p>
<p>Since you should be using <em>rem</em> in the future, always make sure that
1<em>rem</em> is equal to 10px.</p>
<p><strong>How? </strong></p>
<div class="highlight"><pre>    <span class="nt">html</span> <span class="p">{</span> <span class="k">font-size</span><span class="o">:</span> <span class="m">62</span><span class="o">.</span><span class="m">5</span><span class="o">%</span><span class="p">;</span> <span class="p">}</span> <span class="c">/* 10px */</span>

    <span class="nt">p</span> <span class="p">{</span> <span class="k">font-size</span><span class="o">:</span> <span class="m">3</span><span class="o">.</span><span class="m">2</span><span class="n">rem</span><span class="p">;</span> <span class="p">}</span> <span class="c">/* 32px */</span>

    <span class="nt">li</span> <span class="p">{</span> <span class="k">font-size</span><span class="o">:</span> <span class="m">1</span><span class="o">.</span><span class="m">6</span><span class="n">rem</span><span class="p">;</span> <span class="p">}</span> <span class="c">/* 16px */</span>
</pre></div>
