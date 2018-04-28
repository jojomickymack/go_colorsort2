## SDL2 gfx colorsorting app example 2

this is a refactored version of an earlier project that reduces code repetition by using reflection and the 
sort interface.

this requires go-sdl2 and was created with GOPATH set to the root of the project

navigate to project root

``
export GOPATH=$(pwd)
``

cd into the bin directory

``
export GOBIN=$(pwd)
``

now install source for sdl2 and gfx (requires gcc compiler and sdl2/gfx libs/headers, see [installation 
instructions](https://github.com/veandco/go-sdl2)).


```
go get github.com/veandco/go-sdl2/sdl
go get github.com/veandco/go-sdl2/gfx    
```

with those dependencies installed, you should be able to compile the project with

``
go install colorsort
``
