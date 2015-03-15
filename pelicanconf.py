# Theme-specific settings
SITEURL = 'http://blog.syndbg.com'
SITENAME = 'Anton Antonov\'s mindspace'
DOMAIN = 'blog.syndbg.com'
BIO_TEXT = 'Open-source developer and an eternal apprentice.'
FOOTER_TEXT = 'Powered by <a href="http://getpelican.com">Pelican</a> and <a href="http://pages.github.com">GitHub&nbsp;Pages</a>.'

SITE_AUTHOR = 'Anton Antonov'
TWITTER_USERNAME = '@syndbg'
INDEX_DESCRIPTION = 'Open-source developer and an eternal apprentice.'

SIDEBAR_LINKS = [
    '<a href="/about/">About</a>',
    '<a href="/books/">Books</a>',
    '<a href="/cheatsheet/">Cheatsheet</a>',
]

ICONS_PATH = 'images/icons'

GOOGLE_FONTS = [
    "Raleway:400,600",
    "Droid Sans",
]

SOCIAL_ICONS = [
    ('mailto:anton.synd.antonov@gmail.com', 'Email (anton.synd.antonov@gmail.com)', 'fa-envelope'),
    ('http://twitter.com/syndbg', 'Twitter', 'fa-twitter'),
    ('http://github.com/syndbg', 'GitHub', 'fa-github'),
    ('https://www.linkedin.com/in/syndbg', 'Linkedin', 'fa-linkedin'),
]

# Pelican settings
RELATIVE_URLS = True
DEFAULT_LANG = u'en'
TIMEZONE = 'Europe/Sofia'
DEFAULT_DATE = 'fs'
DEFAULT_DATE_FORMAT = '%B %d, %Y'
DEFAULT_PAGINATION = False

THEME = 'pneumatic'

ARTICLE_URL = '{date:%Y}/{date:%m}/{slug}/'
ARTICLE_SAVE_AS = ARTICLE_URL + 'index.html'

PAGE_URL = '{slug}/'
PAGE_SAVE_AS = PAGE_URL + 'index.html'

ARCHIVES_SAVE_AS = 'archive/index.html'
YEAR_ARCHIVE_SAVE_AS = '{date:%Y}/index.html'
MONTH_ARCHIVE_SAVE_AS = '{date:%Y}/{date:%m}/index.html'

DISQUS_SITENAME = 'syndbg'
DISQUS_SHORTNAME = 'syndbg'

# Disable authors, categories, tags, and category pages
DIRECT_TEMPLATES = ['index', 'archives']
CATEGORY_SAVE_AS = ''

# Disables Atom feed generation
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None

TYPOGRIFY = True
MD_EXTENSIONS = ['admonition', 'codehilite(linenums=True)', 'extra']

CACHE_CONTENT = False
DELETE_OUTPUT_DIRECTORY = True
OUTPUT_PATH = 'output'
PATH = 'content'

STATIC_PATHS = ['images', 'uploads', 'extra']
IGNORE_FILES = ['.DS_Store']

extras = ['CNAME', 'favicon.ico', 'keybase.txt', 'robots.txt']
EXTRA_PATH_METADATA = {'extra/%s' % file: {'path': file} for file in extras}

PLUGIN_PATHS = ['plugins']
PLUGINS = ['assets', 'neighbors', 'share_post']
ASSET_SOURCE_PATHS = ['static']
ASSET_CONFIG = [
    ('cache', False),
    ('manifest', False),
    ('url_expire', False),
    ('versions', False),
]
