author: Kevin W. Gisi
company: Twin Cities Code Camp 8&mdash;April 10, 2010
title: Introduction to Go
subtitle:
footer: <a href='http://www.kevingisi.com'>Blog</a> | <a href='http://speakerrate.com/talks/2116-introduction-to-go'>SpeakerRate</a>
subfooter: Copyright &copy; 2010 by Kevin W. Gisi
slides-url: http://gisikw-slides.github.com/introduction-to-go/
code-theme: eiffel
code-line-numbers: false

# Summary

Go is a brand-spanking-new systems language that Google released in November, 2009. Every wonder how awesome C would be if it was garbage-collected, concurrent, and didn't take a few weeks to compile? Wake up; it's here! We'll take a look at this new language that steals some of the dynamic flexibility of Python and Ruby, the performance of C, and a compile time that you'll miss if you blink.

# Why Do We Need a Systems Language?
<<<<<<< HEAD:slides.md
<% uv do %>>
#include <stdio.h>
main(t,_,a)char *a;{return!0<t?t<3?main(-79,-13,a+main(-87,1-_,
main(-86,0,a+1)+a)):1,t<_?main(t+1,_,a):3,main(-94,-27+t,a)&&t==2?_<13?
main(2,_+1,"%s %d %d\n"):9:16:t<0?t<-72?main(_,t,
"@n'+,#'/*{}w+/w#cdnr/+,{}r/*de}+,/*{*+,/w{%+,/w#q#n+,/#{l,+,/n{n+,/+#n+,/#\
;#q#n+,/+k#;*+,/'r :'d*'3,}{w+K w'K:'+}e#';dq#'l \
q#'+d'K#!/+k#;q#'r}eKK#}w'r}eKK{nl]'/#;#q#n'){)#}w'){){nl]'/+#n';d}rw' i;# \
){nl]!/n{n#'; r{#w'r nc{nl]'/#{l,+'K {rw' iK{;[{nl]'/w#q#n'wk nw' \
iwk{KK{nl]!/w{%'l##w#' i; :{nl]'/*{q#'ld;r'}{nlwb!/*de}'c \
;;{nl'-{}rw]'/+,}##'*}#nc,',#nw]'/+kd'+e}+;#'rdq#w! nr'/ ') }+}{rl#'{n' ')# \
}'+}##(!!/")
:t<-50?_==*a?putchar(31[a]):main(-65,_,a+1):main((*a=='/')+t,_,a+1)
  :0<t?main(2,2,"%s"):*a=='/'||main(0,main(-61,*a,
  "!ek;dc i@bK'(q)-[w]*%n+r3#l,{}:\nuwloca-O;m .vpbks,fxntdCeghiry"),a+1);}
=======
<% uv do %>
_(__,___,____,_____){___/__<=_____?_(__,___+_____,____,_____):!(___%__)?_(__,___+_____,___%__,_____):___%__==___/
__&&!____?(printf("%d\t",___/__),_(__,___+_____,____,_____)):(___%__>_____&&___%__<___/__)?_(__,___+_____,
____+!(___/__%(___%__)),_____):___<__*__?_(__,___+_____,____,_____):0;}main(void){_(100, 0, 0, 1);}
>>>>>>> master:slides.md
<% end %>

# Specifications

- Compiled
- Imperative, structured
- Concurrent
- Strongly typed (explicit or inferred)
