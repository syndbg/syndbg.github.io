Date: 2014-03-13
Title: Linux: How to customize your terminal
Tags:
Slug: linux-how-to-customize-your-terminal


<p>In this short tutorial, I'll tell you how to modify your bash into
something more useful for you.</p>
<p>You can customize your terminal's colors, background, title, cursor and
bash prompt.</p>
<h3><strong>To customize your colors:</strong></h3>
<p>​1) Open the Terminal Menu</p>
<p>​2) Navigate to <strong>Edit Profiles</strong></p>
<p>3 ) Create a new profile and <strong>Edit</strong> it</p>
<p>​4) After you've done editing it, make sure you set <strong>Profile used when
launching a new terminal</strong> to your new profile's name.<strong><br />
</strong></p>
<h3>To customize your bash prompt (PS1):</h3>
<p>We'll use the great <a href="http://bashrcgenerator.com/">bashrcgenerator.com</a></p>
<p>​1) Open bashgenerator and make your dream PS1</p>
<p>​2) After you're done, open the terminal's settings. Write <strong>sudo gedit
\~/.bashrc</strong></p>
<p>​3) First of all save your current <strong>.bashrc</strong> to <strong>.bashrc_backup
, </strong>just in case things don't work out from the first time.</p>
<p>​4) Find PS1 in .bashrc and replace it with <strong><em>*Your generated .bashrc
PS1 and additional functions</em></strong>* from bashgenerator.</p>
<p>​5) Close any terminal instances</p>
<p>​6) Open a new terminal and rejoice.</p>
<p>If things didn't work from the first time, replace the non-working
<strong>.bashrc</strong> with the <strong>.bashrc_backup</strong> and try again!</p>
<p>If you just want ready to use, awesome PS1, and the colors are Tango:<br />
<a href="http://syndbg.files.wordpress.com/2014/03/screenshot-from-2014-03-13-132054.png"><img alt="Screenshot from 2014-03-13
13:20:54" src="http://syndbg.files.wordpress.com/2014/03/screenshot-from-2014-03-13-132054.png?w=300" /></a></p>
<p><code>PS1="\[\e[00;33m\]\u\[\e[0m\]\[\e[00;37m\] \[\e[0m\]\[\e[00;31m\]@\[\e[0m\]\[\e[00;37m\] \[\e[0m\]\[\e[00;33m\]\w\[\e[0m\]\[\e[00;37m\]\n\[\e[0m\]\[\e[00;31m\]&gt;\[\e[0m\]"</code></p>
<h3><strong>Extra round: Show current git branch</strong></h3>
<p>Add an extra function in the .bashrc file  </p>
<p><code># show git branch parse_git_branch() {   git branch 2&gt; /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/(\1)/' }</code></p>
<p>And somewhere in the PS1 add the following: <code>\\$(parse_git_branch)</code></p>
<p>Complete PS1:  </p>
<p><code>PS1="\[\e[00;33m\]\u\[\e[0m\]\[\e[00;37m\] \[\e[0m\]\[\e[00;31m\]@\[\e[0m\]\[\e[00;37m\] \[\e[0m\]\[\e[00;33m\]\w\[\e[0m\]\[\e[00;37m\] \$(parse_git_branch) \n\[\e[0m\]\[\e[00;31m\]&gt;\[\e[0m\]"</code></p>
