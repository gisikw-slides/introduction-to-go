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
<% uv :lang => "c" do %>
_(__,___,____,_____){___/__<=_____?_(__,___+_____,____,_____):!(___%__)?_(__,___+_____,___%__,_____):___%__==___/
__&&!____?(printf("%d\t",___/__),_(__,___+_____,____,_____)):(___%__>_____&&___%__<___/__)?_(__,___+_____,
____+!(___/__%(___%__)),_____):___<__*__?_(__,___+_____,____,_____):0;}main(void){_(100, 0, 0, 1);}
<% end %>

# Specifications

- Compiled
- Imperative, structured
- Concurrent
- Strongly typed (explicit or inferred)
