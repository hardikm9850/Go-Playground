package main

import "fmt"

// ------- Dependency Injection --------

type Notifier interface {
    Notify(message string)
}

// implement the interface with a struct
type EmailNotifier struct{}

func (e EmailNotifier) Notify(message string) {
    fmt.Println("Sending email with message:", message)
}

// accepts interface (polymorphism)
func SendNotification(n Notifier, message string) {
    n.Notify(message)
}

// ------- Mocking for testing --------

// Define an interface
type Database interface {
    Query(query string) string
}

// Real implementation
type RealDatabase struct{}

func (db RealDatabase) Query(query string) string {
    return "real result"
}

// Mock implementation
type MockDatabase struct{}

func (db MockDatabase) Query(query string) string {
    return "mock result"
}

func GetData(db Database) string {
    return db.Query("SELECT * FROM table")
}

// ------- Main --------

func main() {
    notifier := EmailNotifier{}
    SendNotification(notifier, "Hello, World!")

    // ----- mocking ------
    realDB := RealDatabase{}
    fmt.Println("Real DB:", GetData(realDB))

    mockDB := MockDatabase{}
    fmt.Println("Mock DB:", GetData(mockDB))
}
