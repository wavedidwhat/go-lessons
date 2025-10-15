package types

import "fmt"

// ============================================
// BASIC TYPES IN GO
// ============================================

// Go has built-in types, but you can create custom types based on them
// This is called "type aliasing" or "type definition"

type myNumber int 
type myString string
type myFloat float64
type myBool bool

// Why create custom types?
// 1. Type safety - prevents mixing up similar data
// 2. Self-documenting code - makes intent clear
// 3. Can attach methods to custom types

// ============================================
// STRUCT TYPES
// ============================================

// Structs group related data together
// Similar to objects in other languages, but simpler

type myRole struct {
	Role string
}

// Structs can have multiple fields of different types
type accountBalance struct {
	amount   myFloat
	currency myString
}

// ============================================
// COLLECTION TYPES
// ============================================

// ARRAYS - Fixed size, cannot grow or shrink
type myArray [5]int

// SLICES - Dynamic size, can grow and shrink
type mySlice []string

// MAPS - Key-value pairs (like dictionaries/objects)
// Format: map[KeyType]ValueType
type myMap map[string]int

// ============================================
// FUNCTION TYPES
// ============================================

// Functions can be types too!
// This function takes two ints and returns an int
type myFunc func(int, int) int

// Example usage:
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// You can assign functions to variables of type myFunc
func demonstrateFunc() {
	var operation myFunc
	
	operation = add
	fmt.Println(operation(5, 3)) // 8
	
	operation = subtract
	fmt.Println(operation(5, 3)) // 2
}

// ============================================
// INTERFACE TYPES
// ============================================

// Interfaces define behavior (method signatures)
// Any type that implements these methods satisfies the interface
type myInterface interface {
	MyMethod() string
}

// ============================================
// STRUCT EMBEDDING (Composition)
// ============================================

// Go doesn't have inheritance, but has COMPOSITION
// You can embed structs inside other structs

// Create custom types for User fields
type UserID myNumber
type UserName myString
type UserBody myFloat
type UserActive myBool

type User struct {
	ID       UserID
	UserName UserName
	Body     UserBody
	Active   UserActive
	accountBalance  // Embedded struct - promotes its fields
}

// Method on User that uses embedded accountBalance
func (u User) BalanceString() string {
	// Can access embedded fields directly
	return fmt.Sprintf("%v %v", u.amount, u.currency)
}

// Method to display user info
func (u User) DisplayInfo() string {
	return fmt.Sprintf("User: %v (ID: %v) - Active: %v", 
		u.UserName, u.ID, u.Active)
}

// ============================================
// STRUCT EMBEDDING EXAMPLES
// ============================================

// Admin embeds User - gets all User fields and methods
type Admin struct {
	User  // Anonymous field - promotes User's fields
	Role  myRole
}

// Method specific to Admin
func (a Admin) DisplayRole() string {
	return fmt.Sprintf("Admin %v has role: %v", a.UserName, a.Role.Role)
}

// SuperAdmin embeds Admin - gets Admin AND User fields/methods
type SuperAdmin struct {
	Admin       // Gets both Admin and User fields/methods
	Permissions mySlice
}

// Method specific to SuperAdmin
func (sa SuperAdmin) HasPermission(perm string) bool {
	for _, p := range sa.Permissions {
		if p == perm {
			return true
		}
	}
	return false
}

// ============================================
// STRUCTS WITH SLICES
// ============================================

type Team struct {
	Name    string
	Members []User  // Slice of User structs
}

// Method to add a member to the team
func (t *Team) AddMember(user User) {
	t.Members = append(t.Members, user)
}

// Method to get team size
func (t Team) Size() int {
	return len(t.Members)
}

// ============================================
// STRUCTS WITH MAPS
// ============================================

type Company struct {
	Name      string
	Employees myMap  // map[string]int - employee names to IDs
}

// Method to add an employee
func (c *Company) AddEmployee(name string, id int) {
	if c.Employees == nil {
		c.Employees = make(myMap) // Initialize map if nil
	}
	c.Employees[name] = id
}

// Method to get employee ID
func (c Company) GetEmployeeID(name string) (int, bool) {
	id, exists := c.Employees[name]
	return id, exists
}

// ============================================
// USING myArray WITH INTERFACES
// ============================================

// Interface for processing data arrays
type DataProcessor interface {
	Process(data myArray) myArray
}

// Concrete implementation of DataProcessor
type ArrayDoubler struct{}

// Process doubles all values in the array
func (ad ArrayDoubler) Process(data myArray) myArray {
	var result myArray
	for i, val := range data {
		result[i] = val * 2
	}
	return result
}

// Another implementation
type ArraySquarer struct{}

func (as ArraySquarer) Process(data myArray) myArray {
	var result myArray
	for i, val := range data {
		result[i] = val * val
	}
	return result
}

// ============================================
// USING myInterface
// ============================================

// Let's make User implement myInterface
func (u User) MyMethod() string {
	return fmt.Sprintf("I am user %v", u.UserName)
}

// Admin also implements myInterface (inherits from User)
// But can override it
func (a Admin) MyMethod() string {
	return fmt.Sprintf("I am admin %v with role %v", a.UserName, a.Role.Role)
}

// Function that accepts any type implementing myInterface
func PrintMyMethod(m myInterface) {
	fmt.Println(m.MyMethod())
}

// ============================================
// USING myFunc
// ============================================

type Calculator struct {
	Operation myFunc
}

func (c Calculator) Calculate(a, b int) int {
	return c.Operation(a, b)
}

// ============================================
// PRACTICAL EXAMPLES
// ============================================

func DemonstrateAllTypes() {
	// 1. Basic custom types
	var num myNumber = 42
	var str myString = "hello"
	var fl myFloat = 3.14
	var b myBool = true
	
	fmt.Printf("Number: %v, String: %v, Float: %v, Bool: %v\n", num, str, fl, b)
	
	// 2. Create a User
	user := User{
		ID:       UserID(1),
		UserName: UserName("Alice"),
		Body:     UserBody(65.5),
		Active:   UserActive(true),
		accountBalance: accountBalance{
			amount:   myFloat(1000.50),
			currency: myString("USD"),
		},
	}
	
	fmt.Println(user.DisplayInfo())
	fmt.Println(user.BalanceString())
	
	// 3. Create an Admin (has User fields too!)
	admin := Admin{
		User: user,
		Role: myRole{Role: "Manager"},
	}
	
	fmt.Println(admin.DisplayRole())
	// Can access User fields directly
	fmt.Println("Admin username:", admin.UserName)
	
	// 4. Create a SuperAdmin
	superAdmin := SuperAdmin{
		Admin:       admin,
		Permissions: mySlice{"read", "write", "delete", "admin"},
	}
	
	fmt.Println("Has 'admin' permission:", superAdmin.HasPermission("admin"))
	fmt.Println("Has 'execute' permission:", superAdmin.HasPermission("execute"))
	
	// 5. Working with Team (slice of structs)
	team := Team{
		Name: "Engineering",
	}
	team.AddMember(user)
	team.AddMember(User{
		ID:       UserID(2),
		UserName: UserName("Bob"),
		Active:   UserActive(true),
	})
	fmt.Printf("Team '%s' has %d members\n", team.Name, team.Size())
	
	// 6. Working with Company (map)
	company := Company{
		Name: "TechCorp",
	}
	company.AddEmployee("Alice", 1)
	company.AddEmployee("Bob", 2)
	company.AddEmployee("Charlie", 3)
	
	if id, exists := company.GetEmployeeID("Alice"); exists {
		fmt.Printf("Alice's ID: %d\n", id)
	}
	
	// 7. Using myArray with DataProcessor interface
	data := myArray{1, 2, 3, 4, 5}
	
	var processor DataProcessor
	processor = ArrayDoubler{}
	doubled := processor.Process(data)
	fmt.Println("Doubled:", doubled)
	
	processor = ArraySquarer{}
	squared := processor.Process(data)
	fmt.Println("Squared:", squared)
	
	// 8. Using myInterface
	PrintMyMethod(user)
	PrintMyMethod(admin)
	
	// 9. Using myFunc
	calc := Calculator{Operation: add}
	fmt.Println("5 + 3 =", calc.Calculate(5, 3))
	
	calc.Operation = subtract
	fmt.Println("5 - 3 =", calc.Calculate(5, 3))
	
	// Can even use anonymous function
	calc.Operation = func(a, b int) int {
		return a * b
	}
	fmt.Println("5 * 3 =", calc.Calculate(5, 3))
}

// ============================================
// KEY TAKEAWAYS
// ============================================

/*
!note Type System Summary

1. **Basic Types**: int, string, float64, bool
   - Can create custom types: `type UserID int`

2. **Structs**: Group related data
   - Like lightweight objects
   - Can embed other structs (composition)

3. **Collections**:
   - Arrays: Fixed size `[5]int`
   - Slices: Dynamic size `[]string`
   - Maps: Key-value pairs `map[string]int`

4. **Functions as Types**: `type myFunc func(int, int) int`
   - Functions are first-class citizens
   - Can be passed around like values

5. **Interfaces**: Define behavior
   - Any type that implements the methods satisfies the interface
   - Enables polymorphism

6. **Composition over Inheritance**:
   - Embed structs instead of inheriting
   - More flexible and explicit
*/

/*
!warning Common Gotchas

1. **Maps must be initialized**: 
```go
   var m map[string]int  // nil map, will panic on write!
   m = make(map[string]int)  // Now safe to use
```

2. **Pointer receivers for mutation**:
```go
   func (t *Team) AddMember(u User) {  // Pointer receiver
       t.Members = append(t.Members, u)  // Modifies original
   }
```

3. **Embedded field access**:
```go
   admin.UserName  // Access User field directly
   admin.User.UserName  // Also valid but verbose
```

4. **Zero values**:
   - Structs: all fields set to their zero values
   - Slices/maps: nil (must initialize before use)
*/