import datetime

# Custom configuration for the Sphinx documentation builder.
# All configuration specific to your project should be done in this file.
#
# The file is included in the common conf.py configuration file.
# You can modify any of the settings below or add any configuration that
# is not covered by the common conf.py file.
#
# For the full list of built-in configuration values, see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html
#
# If you're not familiar with Sphinx and don't want to use advanced
# features, it is sufficient to update the settings in the "Project
# information" section.

############################################################
### Project information
############################################################

# Product name
project = "Ubuntu on WSL"
author = "Canonical Ltd."

# The title you want to display for the documentation in the sidebar.
# You might want to include a version number here.
# To not display any title, set this option to an empty string.
html_title = project + " documentation"

# The default value uses CC-BY-SA as the license and the current year
# as the copyright year.
#
# If your documentation needs a different copyright license, use that
# instead of 'CC-BY-SA'. Also, if your documentation is included as
# part of the code repository of your project, it'll inherit the license
# of the code. So you'll need to specify that license here (instead of
# 'CC-BY-SA').
#
# For static works, it is common to provide the year of first publication.
# Another option is to give the first year and the current year
# for documentation that is often changed, e.g. 2022–2023 (note the en-dash).
#
# A way to check a GitHub repo's creation date is to obtain a classic GitHub
# token with 'repo' permissions here: https://github.com/settings/tokens
# Next, use 'curl' and 'jq' to extract the date from the GitHub API's output:
#
# curl -H 'Authorization: token <TOKEN>' \
#   -H 'Accept: application/vnd.github.v3.raw' \
#   https://api.github.com/repos/canonical/<REPO> | jq '.created_at'

# This adds an edit button to the top of each page
html_theme_options = {
    "source_repository": "https://github.com/canonical/ubuntu-pro-for-wsl/",
    "source_branch": "main",
    "source_directory": "docs/",
}

copyright = "%s CC-BY-SA, %s" % (datetime.date.today().year, author)

## Open Graph configuration - defines what is displayed as a link preview
## when linking to the documentation from another website (see https://ogp.me/)
# The URL where the documentation will be hosted (leave empty if you
# don't know yet)
# NOTE: If no ogp_* variable is defined (e.g. if you remove this section) the
# sphinxext.opengraph extension will be disabled.
ogp_site_url = "https://canonical-starter-pack.readthedocs-hosted.com/"
# The documentation website name (usually the same as the product name)
ogp_site_name = project
# The URL of an image or logo that is used in the preview
ogp_image = "https://assets.ubuntu.com/v1/253da317-image-document-ubuntudocs.svg"

# Update with the favicon for your product (default is the circle of friends)
html_favicon = "../msix/UbuntuProForWSL/Images/icon.ico"

# (Some settings must be part of the html_context dictionary, while others
#  are on root level. Don't move the settings.)
html_context = {
    # Change to the link to the website of your product (without "https://")
    # For example: "ubuntu.com/lxd" or "microcloud.is"
    # If there is no product website, edit the header template to remove the
    # link (see the readme for instructions).
    "product_page": "ubuntu.com/wsl",
    # Add your product tag (the orange part of your logo, will be used in the
    # header) to ".sphinx/_static" and change the path here (start with "_static")
    # (default is the circle of friends)
    "product_tag": "_static/tag.png",
    # Change to the discourse instance you want to be able to link to
    # using the :discourse: metadata at the top of a file
    # (use an empty value if you don't want to link)
    "discourse": "https://discourse.ubuntu.com/c/wsl/27",
    # Change to the Mattermost channel you want to link to
    # (use an empty value if you don't want to link)
    # 'mattermost': 'https://chat.canonical.com/canonical/channels/documentation',
    # Change to the GitHub info for your project
    "github_url": "https://github.com/canonical/ubuntu-pro-for-wsl",
    # Change to the Matrix channel you want to link to
    # (use an empty value if you don't want to link)
    "matrix": "https://matrix.to/#/#ubuntu-wsl:ubuntu.com",
    # Change to the branch for this version of the documentation
    "github_version": "main",
    # Change to the folder that contains the documentation
    # (usually "/" or "/docs/")
    "github_folder": "/docs/",
    # Change to an empty value if your GitHub repo doesn't have issues enabled.
    # This will disable the feedback button and the issue link in the footer.
    "github_issues": "enabled",
    # Controls the existence of Previous / Next buttons at the bottom of pages
    # Valid options: none, prev, next, both
    "sequential_nav": "none",
    # Controls if to display the contributors of a file or not
    "display_contributors": True,
    # Controls time frame for showing the contributors
    "display_contributors_since": "",
}

# If your project is on documentation.ubuntu.com, specify the project
# slug (for example, "lxd") here.
slug = ""

############################################################
### Redirects
############################################################

# Set up redirects (https://documatt.gitlab.io/sphinx-reredirects/usage.html)
# For example: 'explanation/old-name.html': '../how-to/prettify.html',
# You can also configure redirects in the Read the Docs project dashboard
# (see https://docs.readthedocs.io/en/stable/guides/redirects.html).
# NOTE: If this variable is not defined, set to None, or the dictionary is empty,
# the sphinx_reredirects extension will be disabled.

redirects = {
    # deprecated tutorials that will be reworked into new content
    "tutorials/dotnet-systemd": "https://github.com/ubuntu/WSL/blob/main/docs/tutorials/dotnet-systemd.md",
    "tutorials/interop": "https://github.com/ubuntu/WSL/blob/main/docs/tutorials/interop.md",
    # tutorials that have since been converted to howto guides
    "tutorials/cloud-init": "../../howto/cloud-init",
    "tutorials/data-science-and-engineering": "../../howto/data-science-and-engineering",
    "tutorials/gpu-cuda": "../../howto/gpu-cuda",
    "tutorials/vscode": "../../tutorials/develop-with-ubuntu-wsl",
    # change in diataxis names
    "guides/": "../../howto/",
    "guides/contributing": "../../howto/contributing",
    "guides/install-ubuntu-wsl2": "../../howto/install-ubuntu-wsl2",
    "guides/run-workflows-azure": "../../howto/run-workflows-azure",
    "explanations/": "../../explanation/",
    "explanations/ref-arch-explanation": "../../explanation/ref-arch-explanation",
    # improved url to account for merging of distro and UP4W app docs
    "tutorials/getting-started": "../../tutorials/getting-started-with-up4w",
    # account for old use of "tutorial/"
    "tutorial/": "../../tutorials/",
    "tutorial/getting-started": "../../tutorials/getting-started-with-up4w",
    "tutorial/deployment": "../../tutorials/deployment",
    # ... even for URLs that never existed
    "tutorial/getting-started-with-up4w": "../../tutorials/getting-started-with-up4w",
    "tutorial/develop-with-ubuntu-wsl": "../../tutorials/develop-with-ubuntu-wsl",
    # old UP4W explainer is redundant so redirect to arch explanation
    "explanations/up4w": "../../explanation/ref-arch-explanation",
    # deprecated feature so point users to relevant doc
    "guides/autoinstall": "../../howto/cloud-init",
}

############################################################
### Link checker exceptions
############################################################

# Links to ignore when checking links
linkcheck_ignore = [
    "http://127.0.0.1:8000",
    # Linkcheck does not have access to the repo
    "https://github.com/canonical/ubuntu-pro-for-wsl/*",
    # This page redirects to SSO login:
    "https://ubuntu.com/pro/dashboard",
    # Only users logged in to MS Store with their account registered for beta can access this link
    "https://apps.microsoft.com/detail/9PD1WZNBDXKZ",
]

# Pages on which to ignore anchors
# (This list will be appended to linkcheck_anchors_ignore_for_url)
custom_linkcheck_anchors_ignore_for_url = []

############################################################
### Additions to default configuration
############################################################

## The following settings are appended to the default configuration.
## Use them to extend the default functionality.

# Remove this variable to disable the MyST parser extensions.
custom_myst_extensions = []

# Add custom Sphinx extensions as needed.
# This array contains recommended extensions that should be used.
# NOTE: The following extensions are handled automatically and do
# not need to be added here: myst_parser, sphinx_copybutton, sphinx_design,
# sphinx_reredirects, sphinxcontrib.jquery, sphinxext.opengraph
custom_extensions = [
    "sphinx_tabs.tabs",
    "canonical.youtube-links",
    "canonical.related-links",
    "canonical.custom-rst-roles",
    "canonical.terminal-output",
    "notfound.extension",
]

# Add custom required Python modules that must be added to the
# .sphinx/requirements.txt file.
# NOTE: The following modules are handled automatically and do not need to be
# added here: canonical-sphinx-extensions, furo, linkify-it-py, myst-parser,
# pyspelling, sphinx, sphinx-autobuild, sphinx-copybutton, sphinx-design,
# sphinx-notfound-page, sphinx-reredirects, sphinx-tabs, sphinxcontrib-jquery,
# sphinxext-opengraph
custom_required_modules = []

# Add files or directories that should be excluded from processing.
custom_excludes = [
    "doc-cheat-sheet*",
    "diagrams/readme.md",
]

# Add CSS files (located in .sphinx/_static/)
custom_html_css_files = []

# Add JavaScript files (located in .sphinx/_static/)
custom_html_js_files = []

## The following settings override the default configuration.

# Specify a reST string that is included at the end of each file.
# If commented out, use the default (which pulls the reuse/links.txt
# file into each reST file).
# custom_rst_epilog = ''

# By default, the documentation includes a feedback button at the top.
# You can disable it by setting the following configuration to True.
disable_feedback_button = False

# Add tags that you want to use for conditional inclusion of text
# (https://www.sphinx-doc.org/en/master/usage/restructuredtext/directives.html#tags)
custom_tags = []

# If you are using the :manpage: role, set this variable to the URL for the version
# that you want to link to:
# manpages_url = "https://manpages.ubuntu.com/manpages/noble/en/man{section}/{page}.{section}.html"

############################################################
### Additional configuration
############################################################

## Add any configuration that is not covered by the common conf.py file.

# Define a :center: role that can be used to center the content of table cells.
rst_prolog = """
.. role:: center
   :class: align-center
"""

# Define a selector that only adds copy buttons to code blocks without the class `no-copy`
copybutton_selector = "div:not(.no-copy) > div.highlight > pre"

# Define prompts to be excluded from copying when a copy button is used
copybutton_prompt_text = r"^.*?[\$>]\s+"
copybutton_prompt_is_regexp = True
