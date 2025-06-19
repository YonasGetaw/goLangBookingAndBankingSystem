package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "cbe_bank_full/bank"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    storage, err := bank.NewStorage("db.json")
    if err != nil {
        fmt.Println("Error initializing storage:", err)
        return
    }

    for {
        fmt.Println("\n=== Commercial Bank of Ethiopia CLI ===")
        fmt.Println("1. Create Account")
        fmt.Println("2. Login")
        fmt.Println("3. Exit")
        fmt.Print("Select: ")
        choice := readLine(reader)

        switch choice {
        case "1":
            createAccount(reader, storage)
        case "2":
            loginFlow(reader, storage)
        case "3":
            storage.SaveAndExit()
            return
        default:
            fmt.Println("Invalid option")
        }
    }
}

func createAccount(reader *bufio.Reader, st *bank.Storage) {
    fmt.Print("Enter full name: ")
    name := readLine(reader)
    fmt.Print("Set 4-digit PIN: ")
    pin := readLine(reader)
    fmt.Print("Account type (1=Savings, 2=Current): ")
    t := readLine(reader)

    var atype bank.AccountType
    if t == "1" {
        atype = bank.Savings
    } else {
        atype = bank.Current
    }

    acc, err := st.AddAccount(name, pin, atype)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("✅ Account created! ID:", acc.ID)
}

func loginFlow(reader *bufio.Reader, st *bank.Storage) {
    fmt.Print("Enter Account ID: ")
    id := readLine(reader)
    fmt.Print("Enter PIN: ")
    pin := readLine(reader)

    acc := st.Authenticate(id, pin)
    if acc == nil {
        fmt.Println("Invalid credentials")
        return
    }

    bank.ApplyMonthlyInterest(acc, 0.5) // Apply interest for savings accounts

    for {
        fmt.Printf("\n--- Welcome %s | Balance: %.2f ETB ---\n", acc.Name, acc.Balance)
        fmt.Println("1. Deposit")
        fmt.Println("2. Withdraw")
        fmt.Println("3. Transfer")
        fmt.Println("4. Statement")
        fmt.Println("5. Logout")
        fmt.Print("Select: ")
        opt := readLine(reader)

        switch opt {
        case "1":
            fmt.Print("Amount: ")
            amt := readFloat(reader)
            if err := bank.Deposit(acc, amt); err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("💵 Deposited!")
            }
        case "2":
            fmt.Print("Amount: ")
            amt := readFloat(reader)
            if err := bank.Withdraw(acc, amt); err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("💸 Withdrawn!")
            }
        case "3":
            fmt.Print("Destination Account ID: ")
            destID := readLine(reader)
            dest := st.FindAccount(destID)
            if dest == nil {
                fmt.Println("❌ Destination not found")
                continue
            }
            fmt.Print("Amount: ")
            amt := readFloat(reader)
            if err := bank.Transfer(acc, dest, amt); err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("✅ Transferred!")
            }
        case "4":
            bank.PrintStatement(acc)
        case "5":
            st.SaveAndExit()
            return
        default:
            fmt.Println("Invalid option")
        }
    }
}

func readLine(r *bufio.Reader) string {
    input, _ := r.ReadString('\n')
    return strings.TrimSpace(input)
}

func readFloat(r *bufio.Reader) float64 {
    valStr := readLine(r)
    val, err := strconv.ParseFloat(valStr, 64)
    if err != nil {
        return 0
    }
    return val
}

// Save and exit
func (s *bank.Storage) SaveAndExit() {
    if err := s.Save(); err != nil {
        fmt.Println("Error saving data:", err)
    }
}
