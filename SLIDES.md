# PGO - Effective Compilation Optimisations With Golang
Gophers Hannover, March 2024

---

## Compiling in C

* Demonstrative compilation using different optimisation levels
* A Look at the differences
  * Compile time
  * Binary size
  * Performance ... but we won't do that

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
stat -c %s pawncc
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
stat -c %s pawncc
```

Script output:

---

