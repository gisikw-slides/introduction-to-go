author: Kevin W. Gisi
company: Twin Cities Code Camp 8&mdash;April 10, 2010
title: Introduction to Go
subtitle:
footer: <a href='http://www.kevingisi.com'>Blog</a> | <a href='http://speakerrate.com/talks/2116-introduction-to-go'>SpeakerRate</a> | <a href="http://github.com/gisikw-slides/introduction-to-go">Source Code</a>
subfooter: Copyright &copy; 2010 by Kevin W. Gisi
slides-url: http://gisikw-slides.github.com/introduction-to-go/
code-engine: uv
code-theme: all_hallows_eve
code-line-numbers: false

# Summary

Go is a brand-spanking-new systems language that Google released in November, 2009. Every wonder how awesome C would be if it was garbage-collected, concurrent, and didn't take a few weeks to compile? Wake up; it's here! We'll take a look at this new language that steals some of the dynamic flexibility of Python and Ruby, the performance of C, and a compile time that you'll miss if you blink.

# Hello, world
<%= code 'code/hello_world.go', :lang => "go" %>

# Why Go?

- It's a systems language
- It's fun, like dynamic languages

# We Already Have a Systems Language!
Like C
<%= code 'code/c.c', :lang => "c" %>

# We Already Have Fun Languages!
<%= code 'code/ruby.rb' %>

# It Runs on Linux and OS X!
<img src='images/linux.jpg'/>
<img src='images/snow_leopard.jpg'/>

# And Also...
<img src='images/vmware.jpg'/>
<img src='images/virtualbox.png'/>

# Specifications

- Compiled
- Imperative, structured
- Concurrent
- Strongly typed (explicit or inferred)

# Variables &amp; Types
- int, float
- int8, int32, float64
- uint, ufloat
- string

# Variables: Pointers and Arrays
Pointers
- [TODO]

Arrays
<%= code 'code/variables.go#arrays', :lang => 'go' %>

# Variables: Slices and Maps 
Slices

Maps
<%= code 'code/variables.go#maps', :lang => 'go' %>

# Variable Declaration

<%= code 'code/variables.go#declarations', :lang => "go" %>

# Go &#8800; C
- Semicolons optional (implied)
- Curly braces MUST start on the same line
- No parentheses in <code>if</code>s and <code>for</code>s
- Garbage collected
- Arrays aren't pointers

# Methods
- Pass by value
- Multiple return values

# What? Multiple Return Values?!

# Concurrency
<img src='images/sheldon.png' />

# Goroutines
- NOT threads
- Independent code
- Communication over shared memory

# Channels

# Threading

# Networking

# Interfaces
