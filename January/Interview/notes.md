MONTH 1 — Go Fundamentals & Testing

Interview Level: Junior → Strong Junior

Focus: writing correct, testable, production-ready Go code.

You must master

Variables, functions, structs

Errors as values

Zero values

Unit testing (testing package)

1. Tell me about yourself

I’m a backend Go developer focused on building reliable services, writing clean unit-tested code, and understanding performance and concurrency. I enjoy turning simple programs into scalable systems.

2. Why does Go use error returns instead of exceptions?

Go uses explicit error returns to make control flow clear and predictable. Errors aren’t hidden or thrown implicitly, which reduces crashes, improves testability, and makes production systems easier to reason about.

3. What are zero values and why are they useful?

Every Go variable has a safe default value:

int → 0

string → ""

bool → false

pointers/slices/maps → nil

Zero values prevent undefined state, reduce bugs, and simplify code.

4. How do you test edge cases?

I use table-driven tests to cover:

normal cases

boundary values

invalid inputs

This improves coverage and keeps tests clean and scalable.

5. How do you write a unit test in Go?

Using the testing package in _test.go files:

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}


Run with:

go test

What this month proves

By the end of Month 1, you can show that you:

understand Go fundamentals

handle errors intentionally

write clean, testable code

think about edge cases, not just happy paths
