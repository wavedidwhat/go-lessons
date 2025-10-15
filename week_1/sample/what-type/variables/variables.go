package variables

// Declaring variables with var keyword
var (
	// Exported variable (starts with uppercase letter)
	Publicusername = "PublicUser"
	// Unexported variable (starts with lowercase letter)
	privateUsername = "PrivateUser"
)

var Name string = "wave"
// A variable declared without var keyword (short variable declaration)
// This is only allowed inside functions, not at package level
// Uncommenting the following line will cause a compilation error
// age := 24

// more idiomatic way to declare a variable without specifying the type
var Age = 24 // type inference, Go infers that Age is of type int

// You can also declare multiple variables in a single line
var (
	Country  = "Wonderland"
	City     = "Fictional City"
	Language = "Go"
)
// You can also declare a variable without an initial value
var Score int // defaults to 0 for int type

// You can declare a variable and assign a value later
var Height float64

func init() {
	Height = 5.9 // assigning value later
}
// Note: Variables that start with a lowercase letter are unexported and cannot be accessed from other packages.
// Variables that start with an uppercase letter are exported and can be accessed from other packages.
// For example, Publicusername can be accessed from other packages, but privateUsername cannot.
// Similarly, Name and Age can be accessed from other packages, but Score and Height can also be accessed since they are declared with var keyword.
// However, it's a good practice to keep variables unexported unless they need to be accessed from other packages.
// This encapsulation helps in maintaining the integrity of the data and prevents unintended modifications.
// In summary, use uppercase for exported variables and lowercase for unexported ones.



