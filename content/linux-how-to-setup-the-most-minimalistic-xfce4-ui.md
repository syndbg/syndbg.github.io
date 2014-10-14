Date: 2014-10-14
Title: Linux: How to setup the most minimalistic Xfce4 UI
Tags:
Slug: linux-how-to-setup-the-most-minimalistic-xfce4-ui



Most of us adore the minimalistic Unity. But we hate it for the memory hog it is.

So the only reasonable solution is to replace Unity with the rising in popularity, Xfce4.

Xfce is a lightweight desktop environment and that's all we need to know for this article. [(Read here)](http://www.xfce.org/).



## Xfce 4/latest installation:

* It comes preinstalled with [Mint Xfce](http://www.linuxmint.com/download.php) and [Xubuntu](http://xubuntu.org/).
* If you already have an Ubuntu you can't let go of, just `sudo apt-get install xubuntu-desktop`. It will install all the xfce4 candy and widgets you'll need.
* If you come from an other distro, please check the packages you have available and search for `xfce4*`.



## The problem:

Apps neatly hide their app bar and integrate their menu with Unity's top bar.
However this is not the same for other desktop environments.

Here's how Atom looks in Unity:
[![Unity Atom](http://i.imgur.com/h01xaFN.png)]


And here's how Atom looks in non-Unity (Cinnamon in this screenshot
[![Non-Unity Atom](http://i.imgur.com/0daTsbr.png)]



As you can see from miles away, we get a big fat app bar that takes away precious space.
Even more precious if we're limited on display size.




## The solution for the app bar/maximized windows problem:

1. The first solution is to compile from source [xfwm4-titleless](https://github.com/cedl38/xfwm4-titleless).
2. The second is to install `maximus`. `sudo apt-get install maximus`.


I suggest the second, since it's the easiest.


The final result is:

[![Xfce4 Atom](http://i.imgur.com/RmXbTfx.png)]



## Global Menu bar

Is still a persistent issue :(

I tried my luck with Arch's [gnome-globalmenu-xfce4 ](https://aur.archlinux.org/packages/gnome-globalmenu-xfce4/), compiled from source, but to no avail.

2.
