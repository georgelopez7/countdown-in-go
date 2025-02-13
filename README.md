# countdown-in-go

This is a high-performance implementation of the **Countdown Numbers Game** in **Go**. It randomly generates six numbers and a target number, then leverages **goroutines** to efficiently explore all possible arithmetic combinations to find solutions.

## Getting Started

#### 📂 Clone the Repository

```sh
git clone https://github.com/georgelopez7/countdown-in-go.git
cd countdown-in-go
```

#### 🏃🏻‍♂️ Run the Game

```sh
go run main.go
```

## How It Works

### 1️⃣ Generating Permutations with Heap's Algorithm

To explore all possible ways to arrange the six numbers, we use **Heap’s Algorithm**. This efficient method systematically generates **all permutations** by swapping elements in place. The key idea is:

- If there’s only **one element**, return it.
- Otherwise, recursively generate permutations for **n-1** elements.
- Swap elements based on an index to ensure all unique arrangements are considered.

This guarantees we explore every possible order of the numbers efficiently.

### 2️⃣ Exploring Solutions with Goroutines

Each permutation is passed to a **goroutine**, which explores possible arithmetic operations concurrently. We:

- Spawn multiple goroutines to **parallelize** the solution search.
- Use **channels** to collect valid solutions, which are later processed.
- This allows us to **maximize CPU utilization** and find solutions faster.

### 3️⃣ Recursively Solving with Arithmetic Operations

For each permutation, we recursively try all possible sequences of operations (`+`, `-`, `*`, `/`), ensuring:

- Every combination is explored **step by step**.
- Division is only performed if it results in an **integer**.
- Subtraction only occurs if the result is **non-negative**.
- If a valid solution matches the target, it is sent to the **solutions channel**.

## Demo

https://github.com/user-attachments/assets/29b02344-508a-4a50-8877-7fe5e1abaa27

