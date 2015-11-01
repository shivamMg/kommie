## kommie

#### A command-line productivity tool for <del>Linux</del> GNU/Linux :D

Every day you come across *commands* that you might need to save for future references. **kommie** helps you to store and manage these *commands*, interactively from the terminal. Commands are categorized, making them easier to manage.

### Install

**kommie** requires [Go](https://golang.org/dl/). Install it if you haven't already.

Install **kommie** with:

```
go get github.com/shivammg/kommie
```

Make sure you have `$GOPATH/bin` included in `$PATH`. If you don't, add this to your bash configuration file (.bashrc):

	PATH="$PATH:$GOPATH/bin"

### Manual

	kommie <category> <operation>

`<operation>` can be:

`<operation>` | Use
--- | ---
add | Add a command
del | Delete a command
mod | Modify a command
delcat | Delete a category
modcat | Modify a category

If `<category>` is specified, the operation is executed on it. If it is not specified, a list of stored categories is displayed. User is asked to enter a category and then the operation is executed on it.

Try to keep one-word category names, e.g. git, vim, emacs.

There is no add-category command, since a category is automatically created if it is new.

If both `<category>` and `<operation>` are not specified, a list of stored categories is displayed.

Besides the above operations, you can export the data file using:

	kommie export


### Licence

GNU General Public License v3.0

