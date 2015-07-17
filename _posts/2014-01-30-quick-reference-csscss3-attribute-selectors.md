---
layout:     post
title:      "Quick reference: CSS/CSS3 Attribute selectors"
date:       2014-01-30 11:21:29
summary:
categories: css3
---

<p><strong>The basic syntax</strong></p>
<p><code style="background-color:salmon;font-size:1.1rem;color:white;">selector[attribute]</code></p>
<p><strong>Exact match</strong> - finds the exact match of the value</p>
<p><code style="background-color:salmon;font-size:1.1rem;color:white;">selector[attribute=value]</code></p>
<p><strong>Partial match</strong> - finds a partial match of the whole word value.
Doesn't find
<code style="background-color:salmon;font-size:1.1rem;color:white;">valueExample</code>
or
<code style="background-color:salmon;font-size:1.1rem;color:white;">value-Example</code>
strings</p>
<p><code style="background-color:salmon;font-size:1.1rem;color:white;">selector[attribute~=value]</code></p>
<p><strong>Hyphen(dash) match</strong> - finds a match of the value value seperated by
dash. Doesn't matter if it's the first or last word in the string.</p>
<p><code style="background-color:salmon;font-size:1.1rem;color:white;">selector[attribute|=value]</code></p>
<p><strong>Start of value match</strong> - finds the value match in the start of the
string only. Includes
<code style="background-color:salmon;font-size:1.1rem;color:white;">valueExample</code>
and
<code style="background-color:salmon;font-size:1.1rem;color:white;">value-example</code>
strings</p>
<p><code style="background-color:salmon;font-size:1.1rem;color:white;">selector[attribute^=value]</code></p>
<p><strong>End of value match</strong> - finds a match of the value in the end of the
string only. Also includes
<code style="background-color:salmon;font-size:1.1rem;color:white;">valueExample</code>
and
<code style="background-color:salmon;font-size:1.1rem;color:white;">value-example</code>
strings</p>
<p><code style="background-color:salmon;font-size:1.1rem;color:white;">selector[attribute$=value]</code></p>
<p><strong>Any part of value match</strong> - same as partial match but includes
<code style="background-color:salmon;font-size:1.1rem;color:white;">valueExample</code>
and
<code style="background-color:salmon;font-size:1.1rem;color:white;">value-example</code>
strings</p>
<p><code style="background-color:salmon;font-size:1.1rem;color:white;">selector[attribute*=value]</code></p>
