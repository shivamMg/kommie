## kommie

#### A command-line productivity tool for <del>Linux</del> GNU/Linux :D

Store and manage categorized commands for future references.


### Install

**kommie** requires [Go](https://golang.org/dl/). Install it if you haven't already.

If you do, download and install **kommie** using:

```
go get github.com/shivammg/kommie
```

Make sure to have `$GOPATH/bin` in your `$PATH`. If you don't, add this to your bash configuration file (.bashrc):

	PATH="$PATH:$GOPATH/bin"


### Manual

A valid command is of the form:

```
kommie <category> <operation>
```

`<operation>` can be one of the following:

	add     Add command to <category>
	del     Delete command from <category>
	mod     Modify command in <category>
	delcat  Delete <category>
	modcat  Modify <category>

 - If `<operation>` is not specified, commands inside `<category>` are displayed.
 - If `<category>` is not specified, all categories are displayed and user is asked for input. The corresponding `<operation>` is then performed.
 - If both are not specified, all categories are displayed.

Besides these, there are two more commands:

	kommie export

to export stored json data.

	kommie help

to display help.

> **Tip:** Use small lowercase names for categories, ex. *vim*, *git*, *vocab*.

I have included an example data file. You can copy it to your home directory and view it in the program.

### Licence

[GNU General Public License v3.0](https://github.com/shivamMg/kommie/blob/master/LICENSE)

