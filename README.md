crror
=====

Cat says error messages for various compilers


How to install
==============

1. install golang
   
   Ubuntu:
   ```
   $ sudo apt-get install golang
   ```

2. clone
   ```
   $ git clone https://github.com/TakedaHiromasa/crror-cat-edition.git ~/crror-cat-edition
   ```

3. setup $PATH and $GOPATH
   ```
   $ export PATH=$PATH:~/saycat
   $ export GOPATH=~/golang
   ```

you may need to have the above 2 export lines in your .bashrc


How to use
==========

feed output from a compiler/linker commands

    $ saycat make
 
    
or

    $ saycat gcc test.c


feed compiler/linker output from a log file

    $ make 2>&1 > build.log
    $ saycat build.log


