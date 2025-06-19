# 🇪🇹 Commercial Bank of Ethiopia CLI & Ticket Booking System

Welcome to a Go-powered command-line project that simulates the operations of a **Commercial Bank of Ethiopia (CBE)** and a **Ticket Booking System**. This repository showcases core real-world applications built using clean, maintainable Go code — with file-based persistence and intuitive CLI interaction.

---

## 📌 Projects Included

### 🏦 1. Commercial Bank of Ethiopia – CLI Banking System

A full-featured banking application supporting:

- 🔐 Account creation with secure PIN login
- 💰 Deposit & Withdraw operations
- 🔄 Fund transfers between accounts
- 🧾 Transaction history & printed statements
- 📈 Monthly interest calculation (Savings)
- 🗃️ Local JSON file for persistent data

### 🎟️ 2. Ticket Booking System – CLI App

An easy-to-use CLI application for booking tickets, featuring:

- 📋 Ticket reservation with automatic ID generation
- 🔎 Availability checking & seat tracking
- ❌ Cancellation and updates
- 📄 Booking summaries with timestamps
- 💾 Lightweight JSON-based storage

---

## 🛠 Tech Stack

| Tech        | Description                          |
|-------------|--------------------------------------|
| **Go (Golang)** | Main programming language           |
| **UUID**    | Unique transaction & booking IDs     |
| **JSON File** | Data persistence (no DB required)   |
| **CLI**     | Text-based user interaction          |

---

## 🚀 Getting Started

### 📦 Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- `git` installed

### 📥 Clone the Repository

```bash
git clone https://github.com/yourusername/cbe-booking-system.git
cd cbe-booking-system
