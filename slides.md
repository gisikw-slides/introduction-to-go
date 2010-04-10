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

# Fun Languages are Slow at Runtime
<img src='images/benchmarks.png'/>

# Fast Languages are Slow to Compile
<img src='images/xkcd_compiling.png'/>

# Alternative: Go
<s>Go is slow at runtime AND at compile time!</s>

- 1/2&mdash;1/3 LOC for C++ benchmarks
- 1&mdash;1/16 of the speed (libraries?)

Priorities:
- Compile time
- Performance
- Readability

# It Runs on Linux 
<img src='images/linux.jpg'/>

# It Runs on OS X!
<img src='images/snow_leopard.jpg'/>

# It Runs on Virtual Machines!
<img src='images/vmware.jpg'/>

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
- Use them for reference
- DON'T manipulate them!

Arrays
<%= code 'code/variables.go#arrays', :lang => 'go' %>

# Variables: Slices and Maps 
Slices
- "Pointers" for arrays
- Contains pointers to each object in a range of the array
- Used for passing array values by reference

Maps
<%= code 'code/variables.go#maps', :lang => 'go' %>

# Variable Declaration

<%= code 'code/variables.go#declarations', :lang => "go" %>

# Variable Allocation
new()
- allocates heap space
- zero-initializes the space
- returns the address
- used for ints, floats, structs

make()
- allocates heap space
- creates the object (and underlying data structure)
- returns the value
- used for maps, channels, slices

# Go &#8800; C
- Semicolons optional (implied)
- Curly braces MUST start on the same line
- No parentheses in <code>if</code>s and <code>for</code>s
- Garbage collected
- Arrays aren't pointers

# Methods
- CamelCase - public
- camelCase - package-level
- Pass by value
- Multiple return values

# What? Multiple Return Values?!
<%= code 'code/methods.go#multiple_return', :lang => 'go' %>
How do we access them?
<%= code 'code/methods.go#call_gimme_five', :lang => 'go' %>

# Why Multiple Return Values?
<%= code 'code/methods.go#why_multiple', :lang => 'go' %>

# Named Results
<%= code 'code/methods.go#named_results', :lang => 'go' %>

# Concurrency
<img src='images/sheldon.png' />

# Goroutines
- NOT threads
- Independent code
- Communication over shared memory

# Threads
- Exist
- Span across multiple cores
- Go load-balances them
- Don't worry about it

# Goroutine Example
<%= code 'code/concurrency.go#goroutines', :lang => 'go' %>
- No access to spawned goroutines
- No thread.join equivalent

# Channels
- Like Unix pipes
- Communicate across goroutines
- Optionally blocking/non-blocking

# Channel Example
<%= code 'code/concurrency.go#channels', :lang => 'go' %>

# Interfaces
