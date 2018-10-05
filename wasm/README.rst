======================================
Frontend Programming in Go WebAssembly
======================================

This directory contains examples for Go WebAssembly_.
For GopherJS_ examples, visit `root directory`_.

Development Environment:

  - `Ubuntu 18.04`_
  - Go_ (Go 1.11 or later for Go WebAssembly)


Install
+++++++

.. code-block:: bash

  $ GOARCH=wasm GOOS=js go get -u github.com/siongui/godom/wasm

Examples:

  - Hello World (`demo <https://siongui.github.io/frontend-programming-in-go/wasm/001-hello-world/demo/>`_)
  - querySelector (`demo <https://siongui.github.io/frontend-programming-in-go/wasm/002-querySelector/demo/>`_)
  - querySelectorAll (`demo <https://siongui.github.io/frontend-programming-in-go/wasm/002-querySelectorAll/demo/>`_)

.. _Ubuntu 18.04: http://releases.ubuntu.com/18.04/
.. _Go: https://golang.org/dl/
.. _GopherJS: http://www.gopherjs.org/
.. _WebAssembly: https://duckduckgo.com/?q=webassembly
.. _root directory: https://github.com/siongui/frontend-programming-in-go
