
---

# 🚦 Day 3: Conditionals & Logic

In programming, we often need the computer to make a decision. We do this using **Conditionals**.

### 1. Condition Check (`if`)

In Go, we use the `if` keyword.

* **No Parentheses:** We do not put `( )` around the condition.
* **Mandatory Braces:** We *must* use curly braces `{ }` to wrap the code that runs if true.

**Syntax:**

```go
if 5 > 2 {
    fmt.Println("This is true!")
}

```

### 2. What if the condition is false? (`else`)

The `else` keyword defines a block of code that runs if the check fails.

**⚠️ The Golden Rule of Go Logic:**
In Go, the `else` keyword **must** be on the same line as the closing brace `}` of the `if` statement. If you drop it to the next line, Go will crash.

**Correct Way:**

```go
canFly := false

if canFly {
    fmt.Println("I am flying!")
} else {  // <--- 'else' is cuddling the closing brace '}'
    fmt.Println("This is real life.")
}

```

**Wrong Way (Will not run):**

```go
if canFly {
    fmt.Println("I am flying!")
}         // Error here!
else {
    fmt.Println("This is real life.")
}

```

### 3. Blocks vs. Indentation

* **Go:** Uses **Curly Braces `{ }**` to decide what code belongs inside.

However, we **still indent** our code in Go to make it readable for humans.

**Example:**

```go
if 10 > 5 {
    // This line is inside the block because of the braces { }
    fmt.Println("Math works") 
}
// This line is outside the block
fmt.Println("Program continues...")

```

### 4. Comparison Operators

These symbols allow us to compare two values. They result in a `bool` (true or false).

| Operator | Meaning | Example |
| --- | --- | --- |
| `>` | Greater than | `5 > 2` (true) |
| `<` | Less than | `2 < 5` (true) |
| `>=` | Greater than or equal to | `age >= 18` |
| `<=` | Less than or equal to | `age <= 12` |
| `==` | **Equal to** (Note the double equals!) | `name == "Alice"` |
| `!=` | **Not equal to** | `exit != true` |


---

# ➗ The Modulo Operator (`%`)

The modulo operator is one of the most useful tools in programming. It divides two numbers and gives you the **remainder**.

* `6 % 2` is **0** (6 divides cleanly by 2, no remainder).
* `6 % 5` is **1** (5 goes into 6 once, with 1 left over).
* `6 % 4` is **2** (4 goes into 6 once, with 2 left over).

### ⏸️ PAUSE 1: Predict the Output

What do you think the code below will output?

```go
fmt.Println(10 % 3)

```

*(Answer: It will print **1**, because 3 goes into 10 three times (making 9), with **1** left over.)*

---

### ⏸️ PAUSE 2: The Odd or Even Checker

**Challenge:**
Write a program that asks the user for a number.

* If the number is **Even** (divisible by 2), print "Even".
* If the number is **Odd**, print "Odd".

<div class="hint">
<strong>💡 Hints for Go:</strong>




1. Create an <code>int</code> variable to hold the number.




2. Use <code>fmt.Scanln</code> to capture the user's input.




3. Use the modulo operator <code>%</code> to divide the number by 2.




4. If the remainder is <code>0</code>, it is Even. Otherwise, it is Odd.
</div>

---

### 🔐 Solution Key

```go
package main

import "fmt"

func main() {
    var number int

    fmt.Println("Enter a number:")
    fmt.Scanln(&number)

    // Check if the remainder when divided by 2 is 0
    if number % 2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}

```

### ⏸️ PAUSE: The Rollercoaster Check

**Task:** Write a program that asks the user for their height in cm.

1. If they are **over 120cm**, print "You can ride!"
2. **Else**, print "Sorry, you have to grow taller."


---

# 🔀 Module: Structural Logic Patterns

**Concept:** You can write as many if statements as you need, but **how** you structure them changes the logic completely.

There are three distinct patterns. Understanding the difference is the key to writing bug-free code.

---

## 1. The "Mutually Exclusive" Chain
**Syntax:** `if` ... `else if` ... `else`

Use this when **only ONE** thing should happen. As soon as the computer finds a `true` condition, it executes that block and **skips** the rest.

* **Note:** Go uses `else if` (two words).

```go
score := 85

if score >= 90 {
    fmt.Println("Grade A") // If this runs...
} else if score >= 80 {
    fmt.Println("Grade B") // ...this is skipped.
} else {
    fmt.Println("Grade C") // ...and this is skipped.
}

```

---

## 2. The "Tunnel" (Nested Ifs)

**Syntax:** An `if` inside another `if`.

Use this when the inner condition **depends** on the outer condition. You only get to check the second condition if the first one was true.

```go
if hasTicket {
    fmt.Println("Ticket accepted.")
    
    // We only check for a gun IF they have a ticket
    if hasGun == false {
        fmt.Println("Security passed. You may enter.")
    } else {
        fmt.Println("Security Alert! Gun detected.")
    }
}

```

---

## 3. The "Checklist" (Multiple Independent Ifs)

**Syntax:** Multiple `if` statements with no `else` connecting them.

Use this when **multiple things** can happen at the same time. These checks are completely unrelated. The computer will check **every single one**, regardless of the previous results.

```go
// A user can be ALL of these things at once
if isUser {
    fmt.Println("Loading Profile...")
}

if isAdmin {
    fmt.Println("Loading Admin Panel...")
}

if isSubscribed {
    fmt.Println("Loading Premium Content...")
}

```

---

# 🎢 Final Challenge: The Rollercoaster CLI

**Objective:** Translate a complex logical flow into Go.

You are building a digital ticket booth for a theme park. The pricing rules are specific and depend on multiple factors (Height, Age, and Add-ons).

### 📝 The Rules (Logic Flow)

1.  **The Gatekeeper (Nested Logic):**
    * First, ask for the user's **Height** (cm).
    * If they are **under 120cm**, they cannot ride. Print a message and stop.
    * If they are **120cm or taller**, they can proceed to the ticket booth.

2.  **Ticket Pricing (Mutually Exclusive Chain):**
    * Once inside, ask for their **Age**.
    * **Child (Under 12):** $5
    * **Youth (12 to 18 inclusive):** $7
    * **Adult (Over 18):** $12
    * *Store this base price in a `bill` variable.*

3.  **The Upsell (Independent Logic):**
    * Ask: "Do you want a photo taken? (Y/N)"
    * If they answer **"Y"**, add **$3** to their bill.
    * (Ignore this step if they said "N").

4.  **The Receipt:**
    * Print the final calculated bill (e.g., "Your final bill is $15").


---

### 🏁 Expected Output

**Scenario 1 (Too short):**
```text
Welcome to the Rollercoaster!
What is your height in cm?
> 110
Sorry, you have to grow taller before you can ride.
```


---

## ⏸️ PAUSE: The Pizza Order Challenge

**Objective:** Write a program that calculates the final bill for a pizza order. This requires you to mix **Exclusive** logic (Size) with **Independent** logic (Toppings).

**The Rules:**

1. **Size:** A pizza can only be one size.
* Small (S): $15
* Medium (M): $20
* Large (L): $25


2. **Pepperoni:** If they want pepperoni, add **$2** (regardless of size).
3. **Cheese:** If they want extra cheese, add **$1**.

**Task:**

1. Ask for size. Set the base `bill` using an `if/else if/else` chain.
2. Ask about pepperoni. If true, add to the bill.
3. Ask about cheese. If true, add to the bill.
4. Print the final bill.

**Checklist:**

* [ ] Logic for Size (Mutually Exclusive)
* [ ] Logic for Pepperoni (Independent)
* [ ] Logic for Cheese (Independent)
* [ ] Print Final Bill

---
