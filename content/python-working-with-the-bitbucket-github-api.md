Date: 2014-08-14
Title: Python: Introduction to working with the Bitbucket/Github API
Tags:
Slug: python-working-with-the-bitbucket-github-api


<p>Nowdays Bitbucket and Github are the places to store your projects in an organized manner.</p>
<p>They have their pros and cons but both are irreplaceable. They offer RESTful Web APIs.</p>
<p>REST in fewest words posssible: you send a HTTP request to a URL, you get a HTTP response and data.</p>
<ol>
<li>
<p>Start with the amazing <a href="http://restbrowser.bitbucket.org/">Bitbucket REST API Browser</a>.
It's the best introduction tool. You get a very good representation what kind of request you're making,
where do you send it to and what response and data do you get.</p>
</li>
<li>
<p>After you feel comfortable with the use of a REST Web API, you must learn how to use it in Python.
2.1 You can use the amazing <a href="http://docs.python-requests.org/en/latest/">requests library</a> and just send requests to the URLs you already know.
2.2 You can use a library that wraps requests to URLs for you.</p>
</li>
<li>
<p>Decision making
I have experience with both and I can say that both have their use, depending on the complexity of your task.</p>
</li>
</ol>
<p>Finding a working and up-to-date Python library for Bitbucket and Github is hard and unrealiable.
You must trust the fellow developing the library had it all thought out, tested and documented but reality is quite different.</p>
<p>Most of the libraries are unstable and you rarely find something that works without a problems, but that's understandable and common.</p>
<p>You must make sacrifices and that's ok if you realize the risks or time isn't your concern.</p>
<p>Recently my fair share with the Github API was a task that involved a lot of data fetching. Repositories info, open/clossed issues, commits count and statistics.</p>
<p>Pretty much the whole range of Github API features. The Github devs did well with the RESTful Web API, but there wasn't an official Python Github library. Time was a concern since we had to have this feature working with our Django project in less than 5 days.</p>
<p>However first two days had to be spent in research and testing the libraries with small scripts. The best I could find was <a href="https://github.com/jacquev6/PyGithub/">PyGithub</a> which has a rough introduction,
but the best documentation you can wish for and the most Github API features implemented.</p>
<p>The reason for the use of PyGithub when even though time was a concern, is that we worked with a lot of Django models and I had to parse the data. The library already does it for us and I had more time to write on the actual problem.</p>
<h2>The testing</h2>
<p>... is hard. Testing with requests is the easiest it could get. However testing a library is a different thing. You have to mock to make it testable and I'm no good with mocking.</p>
<h2>The actual code</h2>
<p>Working with the Github API with requests</p>
<div class="highlight"><pre><span class="o">&gt;&gt;&gt;</span> <span class="n">r</span> <span class="o">=</span> <span class="n">requests</span><span class="o">.</span><span class="n">get</span><span class="p">(</span><span class="s">&#39;https://api.github.com/user&#39;</span><span class="p">,</span> <span class="n">auth</span><span class="o">=</span><span class="p">(</span><span class="s">&#39;user&#39;</span><span class="p">,</span> <span class="s">&#39;pass&#39;</span><span class="p">))</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">r</span><span class="o">.</span><span class="n">status_code</span>
<span class="mi">200</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">r</span><span class="o">.</span><span class="n">headers</span><span class="p">[</span><span class="s">&#39;content-type&#39;</span><span class="p">]</span>
<span class="s">&#39;application/json; charset=utf8&#39;</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">r</span><span class="o">.</span><span class="n">encoding</span>
<span class="s">&#39;utf-8&#39;</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">r</span><span class="o">.</span><span class="n">text</span>
<span class="s">u&#39;{&quot;type&quot;:&quot;User&quot;...&#39;</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">r</span><span class="o">.</span><span class="n">json</span><span class="p">()</span>
<span class="p">{</span><span class="s">u&#39;private_gists&#39;</span><span class="p">:</span> <span class="mi">419</span><span class="p">,</span> <span class="s">u&#39;total_private_repos&#39;</span><span class="p">:</span> <span class="mi">77</span><span class="p">,</span> <span class="o">...</span><span class="p">}</span>
</pre></div>


<p>The library equivalent</p>
<div class="highlight"><pre><span class="o">&gt;&gt;&gt;</span> <span class="kn">from</span> <span class="nn">github</span> <span class="kn">import</span> <span class="n">Github</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">gh</span> <span class="o">=</span> <span class="n">Github</span><span class="p">(</span><span class="s">&#39;user&#39;</span><span class="p">,</span> <span class="s">&#39;pass&#39;</span><span class="p">)</span> <span class="c"># an authenticated Github instance</span>
<span class="o">&lt;</span><span class="n">github</span><span class="o">.</span><span class="n">MainClass</span><span class="o">.</span><span class="n">Github</span> <span class="nb">object</span> <span class="n">at</span> <span class="mh">0x7f4945ee3410</span><span class="o">&gt;</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">usr</span> <span class="o">=</span> <span class="n">gh</span><span class="o">.</span><span class="n">get_user</span><span class="p">()</span> <span class="c"># returns an AuthenticatedUser</span>
<span class="o">&lt;</span><span class="n">github</span><span class="o">.</span><span class="n">AuthenticatedUser</span><span class="o">.</span><span class="n">AuthenticatedUser</span> <span class="nb">object</span> <span class="n">at</span> <span class="mh">0x7f4942f56190</span><span class="o">&gt;</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">usr</span><span class="o">.</span><span class="n">blog</span>
<span class="s">u&#39;syndbg.github.io&#39;</span>
<span class="o">&gt;&gt;&gt;</span> <span class="n">usr</span><span class="o">.</span><span class="n">created_at</span>
<span class="n">datetime</span><span class="o">.</span><span class="n">datetime</span><span class="p">(</span><span class="mi">2013</span><span class="p">,</span> <span class="mi">11</span><span class="p">,</span> <span class="mi">25</span><span class="p">,</span> <span class="mi">14</span><span class="p">,</span> <span class="mi">51</span><span class="p">,</span> <span class="mi">47</span><span class="p">)</span>
</pre></div>


<p>As you can see the JSON is automatically parsed for you and returned as a datetime object. Such small but important features save time.</p>
<p>You can build your own library and be confident to do so, it's a great practice!</p>
<h2>The project feature</h2>
<p>The custom Django to our project command - <a href="https://github.com/HackBulgaria/Odin/blob/master/students/management/commands/generate_certificates.py">here</a> and <a href="https://github.com/HackBulgaria/Odin/blob/master/students/management/commands/helpers/classes.py">the helpers classes</a>.</p>
