# PGO - Effective Compilation Optimisations With Golang
Gophers Hannover, March 2024

---

## PGO

### What Is It?

* Profile Guided Optimisation
* Optimise **Hot Paths** based on CPU profiles
  * Includes dependencies and stdlib as well
* Techniques
  * Inlining
  * Devirtualization (We'll ignore this)
  * ... I think that's it for now
* Landed in Go 1.21 (Preview in 1.20)

### What is not?

* No Aggressive Optimisation of everything
* Can't magically improve unprofiled workloads
  > New code, using an old profile, won't be optimised.

---

## Why does PGO improve inlining?

### How does inlining in Go work?

* Inlining uses a budget
  * Each inlined function needs to stay within budget
* Inlining may be impossible due to
  * `go:` compiler directives
  * big functions (many lines or nodes)
  * ... reading the compiler code is kind of hard ...

## Inlining with PGO

* Increase the budget for functions
* Figure out which functions are worth increasing (Hot path)
* Doesn't change hard rules, such as `go:noinline` directives


---

## Should I use it?

### Pro

* Easy to use with single workload type applications
  * Merging profiles of multiple workloads is possible, but its work!
* Good for CPU intensive use cases
  * Parsing
  * String Manipulation
  * ...
* Profiles are reusable
* Can't degrade performance in comparison to baseline
* Can reduce cost due to improve CPU usage

### Con

* Profiles should be recaptured regularly for proper efficiency
  * Efficiency between releases can decrease if profiles aren't captured
    correctly.
* Increases compile times
  * Should be negligible
* Slightly bigger binaries
  * Should be negligible

---

## Comparison to C

* Demonstrative compilation using different optimisation levels
* A Look at the differences
  * Compile time
  * Binary size
  * Performance ... but we won't do that, it is a Go talk!

---

## Compiling in C - No Optimisations

Compiling the "Pawn Language" compiler with `-O0` flag:

```bash
# Prepare
git clone --depth 1 https://github.com/pawn-lang/compiler.git || true
cd compiler
rm -rf build && mkdir build && cd build
cmake ../source/compiler -DCMAKE_C_FLAGS="-m32 -O0" > /dev/null

# Compile
start=$(date +%s%N)
make > /dev/null
end=$(date +%s%N)
echo "Done! Elapsed time: $(($(($end-$start))/1000000)) ms"
echo "Binary Size: $(stat -c %s pawncc)"
```

Script output:

---

## Compiling in C - All Optimisations

Compiling the "Pawn Language" compiler with `-O3` flag:

```bash
# Prepare
git clone --depth 1 https://github.com/pawn-lang/compiler.git || true
cd compiler
rm -rf build && mkdir build && cd build
cmake ../source/compiler -DCMAKE_C_FLAGS="-m32 -O3" > /dev/null

# Compile
start=$(date +%s%N)
make > /dev/null
end=$(date +%s%N)
echo "Done! Elapsed time: $(($(($end-$start))/1000000)) ms"
echo "Binary Size: $(stat -c %s pawncc)"
```

Script output:

---

## PGO in Go - Live Demos

TO THE TERMINAL!

