copper-cli
==========

copper-cli is a small command line tool that uses the [Copper] template engine for Go to
render text using a template.

**Usage:** `copper-cli <template.txt >output.txt -data KEY=VALUE -data KEY_2=VALUE_2 ...`

Example
-------

**`template.html`**

```
<html>
  <head>
    <title><% html(title) %></title>
  </head>
  <body>
    <p>
      Hello, <% html(who) %>!
    </p>
  </body>
</html>
```

**Command line: `copper-cli <template.html >output.html -data title=Example -data who=world`**

**`output.html`**

```
<html>
  <head>
    <title>Example</title>
  </head>
  <body>
    <p>
      Hello, world!
    </p>
  </body>
</html>
```



[Copper]: https://github.com/blizzy78/copper
