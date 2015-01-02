newsreader
==========

# $B$O$8$a$K(B
Go $B8@8l$N3+H/$KIT47$l$J$N$G!":nK!$,$h$/J,$+$C$F$$$^$;$s!#(B

# $BL\E*(B
NHK $B$N%K%e!<%9$r<hF@$7$F!"(BOpen JTalk $B$GFI$_>e$2$k%9%/%j%W%H$G$9!#(B

# $B;vA0$K%$%s%9%H!<%k$7$F$*$/$b$N(B
* Go $B8@8l(B
* Open JTalk
* aplay

$B%F%9%H$O(B Debian wheezy $B$G9T$$$^$7$?!#(B

# $B;H$$J}(B
## $B4D6-@_Dj(B
    $ mkdir -p newsreader				# $B<B9T%U%!%$%k$NCV$->l=j(B
    $ cd newsreader
    $ go build src/reader.go
    $ go build src/input_nhknews.go
    $ mkdir -p source					# $BFI$_>e$2$k%F%-%9%H$NCV$->l=j(B

## $B<B9T(B
    $ ./input_nhknews					# $B5-;v$r%F%-%9%H%U%!%$%k2=(B
    $ ./reader							# $BFI$_>e$2(B
	$ rm source/*						# $B%F%-%9%H$r:o=|(B

# $B@bL@(B
$B%*%W%7%g%s$O(B -h $B$G3NG'$7$F$/$@$5$$!#(B

## input_nhknews.go
NHK $B$+$i%K%e!<%9$r<hF@$7$F!"5-;v$4$H$K%F%-%9%H%U%!%$%k$r=PNO$7$^$9!#%G%U%)%k%H$G$O(B source/ $B$K=PNO$7$^$9!#D9J8$r(B Open JTalk $B$KEO$9$HF0:n$,IT0BDj$K$J$k$N$G!"5-;v$4$H$KJ,3d$7$F$$$^$9!#(B

## reader.go
$B;XDj%G%#%l%/%H%j$K$"$k%F%-%9%H%U%!%$%k$rFI$_>e$2$^$9!#%G%U%)%k%H$G$O(B source/ $B$N%F%-%9%H%U%!%$%k$,BP>]$G!"<j$G=q$$$?%F%-%9%H$G$b$+$^$$$^$;$s!#(B


# $B$5$$$4$K(B
$B:#2s$O%K%e!<%9$r<hF@$9$k$h$&$K$7$?$1$l$I!"(BTL $B$r<hF@$7$?$j!"(BJenkins $B$N%(%i!<$r<hF@$7$?$j$9$k$H<BMQ@-$,>e$,$k$H;W$$$^$9!#(B


