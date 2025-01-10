package example



// Service defines the interface for the business logic.
type Service interface {
    Add(a, b int) int
    Multiply(a, b int) int
}
