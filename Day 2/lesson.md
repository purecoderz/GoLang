# The Beginner's Guide to Go `fmt`

As a professional Go developer, the `fmt` (Format) package is the tool I use to make my program **talk** to the user. We will break this down into **4 clear levels**.

---

## Level 1: Just get text on the screen
**The Tool:** `fmt.Println`

Think of this as "easy mode." You use this when you just want to dump information to the screen. It automatically puts spaces between things and hits "Enter" (new line) at the end for you.

**When to use it:** 90% of the time for simple messages.

```go
package main

import "fmt"

func main() {
    name := "Sam"
    age := 25

    // It puts a space between "Sam" and 25 automatically.
    // It also moves to the next line after printing.
    fmt.Println("User info:", name, age)
    
    fmt.Println("This is on a new line.")
}

```

**Output:**

```text
User info: Sam 25

```

### Level 2: Fill-in-the-blanks (`Printf`)

`Printf` stands for "Print Formatted." It works like a **Mad Libs** game. You write a sentence template, and use special codes (called **verbs**) where you want your variables to go.

* **Behavior:** It allows precise control.
* **New Line:** **NO.** You must add `\n` yourself if you want a new line.

**The Essential Verbs:**

* `%s`: For **Strings** (text).
* `%d`: For **Digits** (integers).
* `\n`: Represents a **New Line**.

```go
package main

import "fmt"

func main() {
    name := "Sam"
    age := 25

    // %s -> name
    // %d -> age
    // \n -> new line
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}

```

### Level 3: The Cheat Code (`%v`)

As a beginner, you might forget exactly which verb matches which variable type. Go gives you a "lazy" verb: **`%v`** (Value).

* `%v`: "I don't care what this is, just print the value."
* `%+v`: **Pro Tip.** Use this for structs to see field names (great for debugging).

```go
package main

import "fmt"

func main() {
    item := "Pizza"
    price := 20
    
    // Using %v for everything
    fmt.Printf("Item: %v, Price: %v\n", item, price)
}

```

### Level 4: Saving text without printing (`Sprintf`)

Sometimes you don't want to print to the console immediately. You want to save the text into a string variable (perhaps to email it later, or log it to a file).

* **S**printf = **String** Print Format.
* **Returns:** A string.

```go
package main

import "fmt"

func main() {
    name := "Sam"
    
    // This does NOT print to the screen.
    // It creates a string and saves it in 'msg'.
    msg := fmt.Sprintf("Welcome back, %s!", name)

    // We can print it later
    fmt.Println(msg)
}

```

---

## Part 2: Input (Listening to the User)

### The Basic Listener (`Scan`)

If you want to pause the program and wait for the user to type something, use `fmt.Scan`.

**Crucial Rule:** You must pass the **Address** of the variable using the `&` symbol. This allows `fmt` to go into memory and update that variable.

```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Println("Enter your name and age:")
    
    // The program pauses here until user types and hits Enter.
    // We use &name and &age so Scan can update them.
    fmt.Scan(&name, &age)

    fmt.Printf("Nice to meet you, %s. You are %d.\n", name, age)
}

```

---

## Part 3: Summary Cheat Sheet

| Function | Meaning | Does it add a new line? | Primary Use |
| --- | --- | --- | --- |
| `fmt.Println` | Print Line | **Yes** | simple output. |
| `fmt.Printf` | Print Format | **No** (Use `\n`) | Precise control (Mad Libs style). |
| `fmt.Sprintf` | String Format | N/A | **Saves** the formatted text into a variable. |
| `fmt.Scan` | Scan Input | N/A | Reads user input into a variable (requires `&`). |

### Common Verbs List

| Verb | Description | Example Input |
| --- | --- | --- |
| `%s` | String | "Hello" |
| `%d` | Decimal (Integer) | 10 |
| `%f` | Float (Decimal number) | 10.55 |
| `%v` | Value (Universal) | Any |
| `%+v` | Value with Field Names | `{Name:Sam Age:25}` |
| `%T` | Type of variable | `int`, `string` |

Learn to use the `fmt.Scan` function to collect user input and use it within your code.

**PAUSE 1. Update the code to add an exclamation mark**
Using what you have learnt in this lesson and previous, can you figure out how take user input and slot it in between 2 strings?

You are looking to print a sentence like this:

Hello Name!

e.g. Hello Angela!


**PAUSE 2. Solve it**

1. Create a greeting for your program.
2. Ask the user for the city that they grew up in and store it in a variable.
3. Ask the user for the name of a pet and store it in a variable.
4. Combine the name of their city and pet and show them their band name.
<div class="hint">
  You can use String Concatenation to combine strings with variables too!
  e.g. 

<code>
print("My name is ", name)
</code>
</div>

5. Make sure the input cursor shows on a new line:

<div class="hint">
  Think about how you used \n to add a new line to a string. Try putting it in some different places in your code until it does what you expect. Note, when you click into the output area you will be able to click on the end of the line as well as the new line. See the video solution for how it looks on my system.
</div>

## Demo:
[Try it out first here](https://appbrewery.github.io/python-day1-demo/)



---
## Math And Logic
You cannot build software without doing math and logic. In Go, "Operators" are the symbols that tell the compiler to perform specific mathematical or logical manipulations.

Here is the breakdown of the essential operators your students need to know today.

---

## ➗ 1. Arithmetic Operators (Math)

Go handles math strictly. You generally can only perform math on variables of the **same type** (e.g., `int` + `int`).

| Operator | Name | Example | Result |
| --- | --- | --- | --- |
| `+` | Addition | `10 + 5` | `15` |
| `-` | Subtraction | `10 - 2` | `8` |
| `*` | Multiplication | `4 * 3` | `12` |
| `/` | Division | `10 / 2` | `5` |
| `%` | Modulo (Remainder) | `10 % 3` | `1` |

### ⚠️ The "Integer Division" Trap

This is the most common bug for beginners.
If you divide two **integers**, Go will **drop the decimal**. It does not round up; it simply cuts off the tail.

* `7 / 2` = `3` (Not 3.5)
* `7.0 / 2.0` = `3.5` (If using floats)

---

## 🔄 2. The Increment Operators

In loops (which we will learn next), we often need to count up by one. Go makes this easy, but unlike other languages (like C++ or JavaScript), these are **statements**, not expressions. You can't put them inside a print statement.

* **`x++`** : Adds 1 to x. (Same as `x = x + 1`)
* **`x--`** : Subtracts 1 from x.

**Valid:**

```go
count := 0
count++ 

```

**Invalid (Error):**

```go
fmt.Println(count++) // ERROR! Go doesn't allow this.

```

---

## 🎯 3. Assignment Operators

Programmers are lazy. We like shortcuts.

| Operator | Long Form | Meaning |
| --- | --- | --- |
| `+=` | `x = x + 5` | Add 5 to x and save it. |
| `-=` | `x = x - 5` | Subtract 5 from x and save it. |
| `*=` | `x = x * 5` | Multiply x by 5 and save it. |

---

## 🧠 4. Logical Operators (True/False)

We use these to combine multiple checks inside an `if` statement.

| Operator | Name | Logic |
| --- | --- | --- |
| `&&` | **AND** | Returns true only if **BOTH** sides are true. |
| `||` | **OR** | Returns true if **AT LEAST ONE** side is true. |
| `!` | **NOT** | Flips the value (`!true` becomes `false`). |

---

### 🛠️ Quick Exercise: The Leap Year Checker

## Basic Operators

Learn to use the basic mathematical operators, +, -, *, /, // and **

## PEMDAS
Parentheses, Exponents, Multiplication/Division, Addition/Subtraction

### PAUSE 1. What is the output of this code? 
`print(3 * 3 + 3 / 3 - 3)`

### PAUSE 2. Change the code so it outputs 3?
`print(3 * 3 + 3 / 3 - 3)`



---

# 🧱 Lesson: Types, Errors

In Go, types are strict. If you try to mix them up (like adding a number to a word), the program won't even start. It will give you a **Compile Time Error**.

### 1. Type Mismatch Errors

In dynamic languages, you might get a `TypeError` while the program is running. In Go, the compiler stops you immediately.

**Example of a Breakdown:**

```go
fmt.Println(len(12345)) 
// ❌ COMPILER ERROR: invalid argument 12345 (type int) for len

```

The `len()` function expects a **String** (or Array/Slice). It does not know how to count the digits of an **Integer** directly.

#### ⏸️ PAUSE 1: Fix the Code

**Task:** rewrite the code so `len()` accepts the value. We need to make that number look like a string.

```go
// Fix this line so it prints "5"
fmt.Println(len(12345)) 

```

*(Hint: Wrap the number in quotes `""`)*

---

### 2. Type Checking (`%T`)

How do we know what type a variable is? In Go, we don't have a `type()` function. Instead, we use the `fmt` package with a special "Verb": **`%T`**.

**Syntax:**

```go
fmt.Printf("Type: %T \n", "Hello") // Output: Type: string

```

#### ⏸️ PAUSE 2: The Type Inspector

**Task:** Write a program that prints the type for 4 different values using `fmt.Printf` and `%T`.

1. A String (`"Go"`)
2. An Integer (`42`)
3. A Float (`3.14`)
4. A Boolean (`true`)

**Expected Output:**
`string int float64 bool`

---

# 💰 Day 2 Capstone Project: The Tip Calculator

**Goal:** Build a program that calculates exactly how much everyone needs to pay when splitting a bill at a restaurant.

### The Scenario
You are out with friends. The bill arrives. Everyone is looking at you to figure out the math.
You need to calculate the **split amount** per person, including the **tip**.

### 📝 The Requirements

1.  **Greeting:** Welcome the user to the calculator.
2.  **Input 1:** Ask for the **Total Bill** (e.g., `$150.00`).
3.  **Input 2:** Ask for the **Tip Percentage** (e.g., `12` for 12%).
4.  **Input 3:** Ask for the **Number of People** splitting the bill (e.g., `5`).
5.  **The Logic:** Calculate the share per person using this formula:
    `Result = (Bill / People) * (1 + (Tip / 100))`
6.  **Output:** Print the final amount per person, **rounded to 2 decimal places**.

### 💡 Example Interaction
```text
Welcome to the Tip Calculator!
What was the total bill? 
> 150.00
How much tip would you like to give? (10, 12, or 15)
> 12
How many people to split the bill?
> 5

Each person should pay: $33.60
